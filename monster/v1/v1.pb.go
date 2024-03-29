// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/v1.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type UserRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_39a831d85a350775, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Success              bool     `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39a831d85a350775, []int{1}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type PatientRequest struct {
	Id                   int32    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PatientRequest) Reset()         { *m = PatientRequest{} }
func (m *PatientRequest) String() string { return proto.CompactTextString(m) }
func (*PatientRequest) ProtoMessage()    {}
func (*PatientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_39a831d85a350775, []int{2}
}

func (m *PatientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatientRequest.Unmarshal(m, b)
}
func (m *PatientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatientRequest.Marshal(b, m, deterministic)
}
func (m *PatientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatientRequest.Merge(m, src)
}
func (m *PatientRequest) XXX_Size() int {
	return xxx_messageInfo_PatientRequest.Size(m)
}
func (m *PatientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PatientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PatientRequest proto.InternalMessageInfo

func (m *PatientRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PatientRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PatientResponse struct {
	Id                   int32    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PatientResponse) Reset()         { *m = PatientResponse{} }
func (m *PatientResponse) String() string { return proto.CompactTextString(m) }
func (*PatientResponse) ProtoMessage()    {}
func (*PatientResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39a831d85a350775, []int{3}
}

func (m *PatientResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatientResponse.Unmarshal(m, b)
}
func (m *PatientResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatientResponse.Marshal(b, m, deterministic)
}
func (m *PatientResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatientResponse.Merge(m, src)
}
func (m *PatientResponse) XXX_Size() int {
	return xxx_messageInfo_PatientResponse.Size(m)
}
func (m *PatientResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PatientResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PatientResponse proto.InternalMessageInfo

func (m *PatientResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PatientResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "v1.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "v1.UserResponse")
	proto.RegisterType((*PatientRequest)(nil), "v1.PatientRequest")
	proto.RegisterType((*PatientResponse)(nil), "v1.PatientResponse")
}

func init() { proto.RegisterFile("v1/v1.proto", fileDescriptor_39a831d85a350775) }

var fileDescriptor_39a831d85a350775 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4d, 0x4f, 0x84, 0x30,
	0x10, 0x86, 0x6d, 0x5d, 0x56, 0x9d, 0x35, 0xbb, 0x66, 0xbc, 0x34, 0x9c, 0x48, 0x4f, 0x9c, 0xd0,
	0xfa, 0x91, 0xe8, 0xd9, 0x83, 0x57, 0xd3, 0xc4, 0x1f, 0x80, 0x30, 0x07, 0x0e, 0x02, 0x32, 0xa5,
	0xbf, 0xdf, 0x00, 0x22, 0x34, 0xdc, 0x3a, 0x6f, 0xfa, 0xf4, 0x79, 0x3b, 0x70, 0xf0, 0xe6, 0xce,
	0x9b, 0xac, 0xed, 0x1a, 0xd7, 0xa0, 0xf4, 0x46, 0x1b, 0x38, 0x7c, 0x32, 0x75, 0x96, 0x7e, 0x7a,
	0x62, 0x87, 0x47, 0x90, 0x55, 0xa9, 0x44, 0x22, 0xd2, 0xc8, 0xca, 0xaa, 0x44, 0x84, 0x5d, 0x9d,
	0x7f, 0x93, 0x92, 0x89, 0x48, 0xaf, 0xec, 0x78, 0xd6, 0x2f, 0x70, 0x3d, 0x21, 0xdc, 0x36, 0x35,
	0xd3, 0x86, 0x51, 0x70, 0xc1, 0x7d, 0x51, 0x10, 0xb3, 0x3a, 0x4f, 0x44, 0x7a, 0x69, 0xe7, 0x51,
	0x3f, 0xc1, 0xf1, 0x23, 0x77, 0x15, 0xd5, 0x2e, 0xf4, 0xed, 0x36, 0xbe, 0x68, 0xe5, 0x7b, 0x86,
	0xd3, 0x3f, 0x15, 0x28, 0xb7, 0xd8, 0x7e, 0xc1, 0x1e, 0x7a, 0x88, 0x86, 0x9a, 0x8c, 0x06, 0xe0,
	0xad, 0xa3, 0xdc, 0xd1, 0x30, 0xe2, 0x29, 0xf3, 0x26, 0x5b, 0x7d, 0x39, 0xbe, 0x59, 0x82, 0xe9,
	0x75, 0x7d, 0x86, 0xaf, 0x00, 0xef, 0xe4, 0xfe, 0xac, 0x88, 0xc3, 0x8d, 0xb0, 0x78, 0x7c, 0x1b,
	0x64, 0x33, 0x78, 0x2f, 0xbe, 0xf6, 0xe3, 0x6e, 0x1f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x18,
	0x0a, 0x38, 0x4d, 0x6a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetPatient(ctx context.Context, in *PatientRequest, opts ...grpc.CallOption) (Users_GetPatientClient, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/v1.Users/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetPatient(ctx context.Context, in *PatientRequest, opts ...grpc.CallOption) (Users_GetPatientClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Users_serviceDesc.Streams[0], "/v1.Users/GetPatient", opts...)
	if err != nil {
		return nil, err
	}
	x := &usersGetPatientClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Users_GetPatientClient interface {
	Recv() (*PatientResponse, error)
	grpc.ClientStream
}

type usersGetPatientClient struct {
	grpc.ClientStream
}

func (x *usersGetPatientClient) Recv() (*PatientResponse, error) {
	m := new(PatientResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	CreateUser(context.Context, *UserRequest) (*UserResponse, error)
	GetPatient(*PatientRequest, Users_GetPatientServer) error
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Users/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CreateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetPatient_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PatientRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UsersServer).GetPatient(m, &usersGetPatientServer{stream})
}

type Users_GetPatientServer interface {
	Send(*PatientResponse) error
	grpc.ServerStream
}

type usersGetPatientServer struct {
	grpc.ServerStream
}

func (x *usersGetPatientServer) Send(m *PatientResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Users_CreateUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPatient",
			Handler:       _Users_GetPatient_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1/v1.proto",
}
