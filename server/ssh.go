package server

import (
	"io"
	"net"
	"sync"

	"github.com/gliderlabs/ssh"
	"github.com/oklog/run"
	log "github.com/sirupsen/logrus"
	gossh "golang.org/x/crypto/ssh"
)

const (
	forwardedStreamlocalChannelType     = "forwarded-streamlocal@openssh.com"
	streamlocalForwardChannelType       = "streamlocal-forward@openssh.com"
	cancelStreamlocalForwardChannelType = "cancel-streamlocal-forward@openssh.com"
)

type streamlocalChannelForwardMsg struct {
	SocketPath string
}

type forwardedStreamlocalPayload struct {
	SocketPath string
	Reserved0  string
}

func newStreamlocalForwardHandler(sessionDialListener SessionDialListener, logger log.FieldLogger) *streamlocalForwardHandler {
	return &streamlocalForwardHandler{
		sessionDialListener: sessionDialListener,
		forwards:            make(map[string]net.Listener),
		logger:              logger,
	}
}

type streamlocalForwardHandler struct {
	sessionDialListener SessionDialListener
	forwards            map[string]net.Listener
	logger              log.FieldLogger
	sync.Mutex
}

func (h *streamlocalForwardHandler) listen(ctx ssh.Context, ln net.Listener, sessionID string, logger log.FieldLogger) {
	conn := ctx.Value(ssh.ContextKeyConn).(*gossh.ServerConn)

	for {
		c, err := ln.Accept()
		if err != nil {
			break
		}
		payload := gossh.Marshal(&forwardedStreamlocalPayload{
			SocketPath: sessionID,
		})
		go func() {
			ch, reqs, err := conn.OpenChannel(forwardedStreamlocalChannelType, payload)
			if err != nil {
				logger.WithError(err).Info("error opening channel")
				c.Close()
				return
			}

			closeAll := func() {
				ch.Close()
				c.Close()
			}

			var g run.Group
			{
				g.Add(func() error {
					gossh.DiscardRequests(reqs)
					return nil
				}, func(err error) {
					closeAll()
				})
			}
			{
				g.Add(func() error {
					_, err := io.Copy(ch, c)
					return err
				}, func(err error) {
					closeAll()
				})
			}
			{
				g.Add(func() error {
					_, err := io.Copy(c, ch)
					return err
				}, func(err error) {
					closeAll()
				})
			}

			go func() { _ = g.Run() }()
		}()
	}
}

func (h *streamlocalForwardHandler) Handler(ctx ssh.Context, srv *ssh.Server, req *gossh.Request) (bool, []byte) {
	switch req.Type {
	case streamlocalForwardChannelType:
		var reqPayload streamlocalChannelForwardMsg
		if err := gossh.Unmarshal(req.Payload, &reqPayload); err != nil {
			h.logger.WithError(err).Info("error parsing streamlocal payload")
			return false, []byte(err.Error())
		}

		if srv.ReversePortForwardingCallback == nil || !srv.ReversePortForwardingCallback(ctx, reqPayload.SocketPath, 0) {
			return false, []byte("port forwarding is disabled")
		}

		sessionID := reqPayload.SocketPath
		logger := h.logger.WithFields(log.Fields{"session-id": sessionID})
		ln, err := h.sessionDialListener.Listen(sessionID)
		if err != nil {
			logger.WithError(err).Info("error listening socketing")
			return false, []byte(err.Error())
		}

		h.trackListener(sessionID, ln)

		var g run.Group
		{
			g.Add(func() error {
				<-ctx.Done()
				return ctx.Err()
			}, func(err error) {
				h.closeListener(sessionID)
			})
		}
		{
			g.Add(func() error {
				h.listen(ctx, ln, sessionID, logger)
				return nil
			}, func(err error) {
				h.closeListener(sessionID)
			})
		}

		go func() { _ = g.Run() }()

		return true, nil
	case cancelStreamlocalForwardChannelType:
		var reqPayload streamlocalChannelForwardMsg
		if err := gossh.Unmarshal(req.Payload, &reqPayload); err != nil {
			h.logger.WithError(err).Info("error parsing steamlocal payload")
			return false, []byte(err.Error())
		}

		h.closeListener(reqPayload.SocketPath)
		return true, nil

	default:
		return false, nil
	}
}

func (h *streamlocalForwardHandler) trackListener(sessionID string, ln net.Listener) {
	h.Lock()
	defer h.Unlock()
	h.forwards[sessionID] = ln
}

func (h *streamlocalForwardHandler) closeListener(sessionID string) {
	h.Lock()
	defer h.Unlock()

	ln, ok := h.forwards[sessionID]
	if ok {
		ln.Close()
	}

	delete(h.forwards, sessionID)
}
