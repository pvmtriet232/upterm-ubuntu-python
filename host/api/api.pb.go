// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Identifier_Type int32

const (
	Identifier_HOST   Identifier_Type = 0
	Identifier_CLIENT Identifier_Type = 1
)

var Identifier_Type_name = map[int32]string{
	0: "HOST",
	1: "CLIENT",
}

var Identifier_Type_value = map[string]int32{
	"HOST":   0,
	"CLIENT": 1,
}

func (x Identifier_Type) String() string {
	return proto.EnumName(Identifier_Type_name, int32(x))
}

func (Identifier_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2, 0}
}

type GetSessionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSessionRequest) Reset()         { *m = GetSessionRequest{} }
func (m *GetSessionRequest) String() string { return proto.CompactTextString(m) }
func (*GetSessionRequest) ProtoMessage()    {}
func (*GetSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *GetSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionRequest.Unmarshal(m, b)
}
func (m *GetSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionRequest.Marshal(b, m, deterministic)
}
func (m *GetSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionRequest.Merge(m, src)
}
func (m *GetSessionRequest) XXX_Size() int {
	return xxx_messageInfo_GetSessionRequest.Size(m)
}
func (m *GetSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionRequest proto.InternalMessageInfo

type GetSessionResponse struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Command              []string `protobuf:"bytes,2,rep,name=command,proto3" json:"command,omitempty"`
	ForceCommand         []string `protobuf:"bytes,3,rep,name=force_command,json=forceCommand,proto3" json:"force_command,omitempty"`
	Host                 string   `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	NodeAddr             string   `protobuf:"bytes,5,opt,name=node_addr,json=nodeAddr,proto3" json:"node_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSessionResponse) Reset()         { *m = GetSessionResponse{} }
func (m *GetSessionResponse) String() string { return proto.CompactTextString(m) }
func (*GetSessionResponse) ProtoMessage()    {}
func (*GetSessionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *GetSessionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionResponse.Unmarshal(m, b)
}
func (m *GetSessionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionResponse.Marshal(b, m, deterministic)
}
func (m *GetSessionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionResponse.Merge(m, src)
}
func (m *GetSessionResponse) XXX_Size() int {
	return xxx_messageInfo_GetSessionResponse.Size(m)
}
func (m *GetSessionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionResponse proto.InternalMessageInfo

func (m *GetSessionResponse) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *GetSessionResponse) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *GetSessionResponse) GetForceCommand() []string {
	if m != nil {
		return m.ForceCommand
	}
	return nil
}

func (m *GetSessionResponse) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *GetSessionResponse) GetNodeAddr() string {
	if m != nil {
		return m.NodeAddr
	}
	return ""
}

type Identifier struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 Identifier_Type `protobuf:"varint,2,opt,name=type,proto3,enum=api.Identifier_Type" json:"type,omitempty"`
	NodeAddr             string          `protobuf:"bytes,3,opt,name=node_addr,json=nodeAddr,proto3" json:"node_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Identifier) Reset()         { *m = Identifier{} }
func (m *Identifier) String() string { return proto.CompactTextString(m) }
func (*Identifier) ProtoMessage()    {}
func (*Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identifier.Unmarshal(m, b)
}
func (m *Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identifier.Marshal(b, m, deterministic)
}
func (m *Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identifier.Merge(m, src)
}
func (m *Identifier) XXX_Size() int {
	return xxx_messageInfo_Identifier.Size(m)
}
func (m *Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_Identifier proto.InternalMessageInfo

func (m *Identifier) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Identifier) GetType() Identifier_Type {
	if m != nil {
		return m.Type
	}
	return Identifier_HOST
}

func (m *Identifier) GetNodeAddr() string {
	if m != nil {
		return m.NodeAddr
	}
	return ""
}

func init() {
	proto.RegisterEnum("api.Identifier_Type", Identifier_Type_name, Identifier_Type_value)
	proto.RegisterType((*GetSessionRequest)(nil), "api.GetSessionRequest")
	proto.RegisterType((*GetSessionResponse)(nil), "api.GetSessionResponse")
	proto.RegisterType((*Identifier)(nil), "api.Identifier")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xdf, 0x4a, 0x32, 0x41,
	0x18, 0xc6, 0xbf, 0xfd, 0xf3, 0x99, 0xfb, 0x62, 0x62, 0x6f, 0x51, 0x8b, 0x19, 0xc8, 0x76, 0xb2,
	0x47, 0x0a, 0x76, 0x05, 0x22, 0x51, 0x42, 0x24, 0xac, 0x9e, 0x2f, 0x93, 0xf3, 0x6a, 0x03, 0x39,
	0x33, 0xed, 0x4c, 0x81, 0x87, 0x75, 0x0b, 0x5d, 0x42, 0x97, 0xd4, 0x2d, 0x74, 0x21, 0xe1, 0xb8,
	0x66, 0xe5, 0xd9, 0xee, 0xef, 0x79, 0x98, 0x67, 0xf8, 0x0d, 0x44, 0x4c, 0x8b, 0x8e, 0x2e, 0x94,
	0x55, 0x18, 0x30, 0x2d, 0x9a, 0xad, 0xb9, 0x52, 0xf3, 0x07, 0xea, 0x32, 0x2d, 0xba, 0x4c, 0x4a,
	0x65, 0x99, 0x15, 0x4a, 0x9a, 0x75, 0x25, 0x39, 0x84, 0x83, 0x2b, 0xb2, 0x63, 0x32, 0x46, 0x28,
	0x99, 0xd1, 0xe3, 0x13, 0x19, 0x9b, 0xbc, 0x7b, 0x80, 0x3f, 0xa9, 0xd1, 0x4a, 0x1a, 0xc2, 0x33,
	0x00, 0xb3, 0x46, 0xb9, 0xe0, 0xb1, 0xd7, 0xf6, 0xd2, 0x28, 0x8b, 0x4a, 0x32, 0xe4, 0x18, 0xc3,
	0xde, 0x54, 0x2d, 0x16, 0x4c, 0xf2, 0xd8, 0x6f, 0x07, 0x69, 0x94, 0x6d, 0x7e, 0xf1, 0x1c, 0xf6,
	0x67, 0xaa, 0x98, 0x52, 0xbe, 0xc9, 0x03, 0x97, 0xd7, 0x1c, 0x1c, 0x94, 0x25, 0x84, 0xf0, 0x5e,
	0x19, 0x1b, 0x87, 0xee, 0x5c, 0xf7, 0x8d, 0xa7, 0x10, 0x49, 0xc5, 0x29, 0x67, 0x9c, 0x17, 0xf1,
	0x7f, 0x17, 0x54, 0x57, 0xa0, 0xcf, 0x79, 0x91, 0xbc, 0x78, 0x00, 0x43, 0x4e, 0xd2, 0x8a, 0x99,
	0xa0, 0x02, 0xeb, 0xe0, 0x7f, 0xdf, 0xca, 0x17, 0x1c, 0x53, 0x08, 0xed, 0x52, 0x53, 0xec, 0xb7,
	0xbd, 0xb4, 0xde, 0x3b, 0xea, 0xac, 0xb4, 0x6c, 0xeb, 0x9d, 0xc9, 0x52, 0x53, 0xe6, 0x1a, 0xbf,
	0x57, 0x82, 0x3f, 0x2b, 0x2d, 0x08, 0x57, 0x55, 0xac, 0x42, 0x78, 0x3d, 0x1a, 0x4f, 0x1a, 0xff,
	0x10, 0xa0, 0x32, 0xb8, 0x19, 0x5e, 0xde, 0x4e, 0x1a, 0x5e, 0x2f, 0x87, 0x5a, 0x9f, 0x2f, 0x84,
	0x1c, 0x53, 0xf1, 0x2c, 0xa6, 0x84, 0x23, 0x80, 0xad, 0x38, 0x3c, 0x76, 0xa3, 0x3b, 0x7e, 0x9b,
	0x27, 0x3b, 0x7c, 0x6d, 0x38, 0x69, 0xbc, 0x7e, 0x7c, 0xbe, 0xf9, 0x80, 0xd5, 0x6e, 0xa9, 0xf5,
	0xae, 0xe2, 0x9e, 0xe9, 0xe2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x18, 0x01, 0xcc, 0xd6, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminServiceClient interface {
	GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*GetSessionResponse, error)
}

type adminServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdminServiceClient(cc *grpc.ClientConn) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*GetSessionResponse, error) {
	out := new(GetSessionResponse)
	err := c.cc.Invoke(ctx, "/api.AdminService/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
type AdminServiceServer interface {
	GetSession(context.Context, *GetSessionRequest) (*GetSessionResponse, error)
}

// UnimplementedAdminServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (*UnimplementedAdminServiceServer) GetSession(ctx context.Context, req *GetSessionRequest) (*GetSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}

func RegisterAdminServiceServer(s *grpc.Server, srv AdminServiceServer) {
	s.RegisterService(&_AdminService_serviceDesc, srv)
}

func _AdminService_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AdminService/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetSession(ctx, req.(*GetSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSession",
			Handler:    _AdminService_GetSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
