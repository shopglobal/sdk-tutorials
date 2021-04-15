// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nameservice/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// this line is used by starport scaffolding # proto/tx/message
type MsgCreateWhois struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price int32  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (m *MsgCreateWhois) Reset()         { *m = MsgCreateWhois{} }
func (m *MsgCreateWhois) String() string { return proto.CompactTextString(m) }
func (*MsgCreateWhois) ProtoMessage()    {}
func (*MsgCreateWhois) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4dec9b515ec245, []int{0}
}
func (m *MsgCreateWhois) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateWhois) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateWhois.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateWhois) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateWhois.Merge(m, src)
}
func (m *MsgCreateWhois) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateWhois) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateWhois.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateWhois proto.InternalMessageInfo

func (m *MsgCreateWhois) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgCreateWhois) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgCreateWhois) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type MsgCreateWhoisResponse struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgCreateWhoisResponse) Reset()         { *m = MsgCreateWhoisResponse{} }
func (m *MsgCreateWhoisResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateWhoisResponse) ProtoMessage()    {}
func (*MsgCreateWhoisResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4dec9b515ec245, []int{1}
}
func (m *MsgCreateWhoisResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateWhoisResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateWhoisResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateWhoisResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateWhoisResponse.Merge(m, src)
}
func (m *MsgCreateWhoisResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateWhoisResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateWhoisResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateWhoisResponse proto.InternalMessageInfo

func (m *MsgCreateWhoisResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MsgUpdateWhois struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Id    uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Price int32  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (m *MsgUpdateWhois) Reset()         { *m = MsgUpdateWhois{} }
func (m *MsgUpdateWhois) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateWhois) ProtoMessage()    {}
func (*MsgUpdateWhois) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4dec9b515ec245, []int{2}
}
func (m *MsgUpdateWhois) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateWhois) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateWhois.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateWhois) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateWhois.Merge(m, src)
}
func (m *MsgUpdateWhois) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateWhois) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateWhois.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateWhois proto.InternalMessageInfo

func (m *MsgUpdateWhois) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgUpdateWhois) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MsgUpdateWhois) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgUpdateWhois) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type MsgUpdateWhoisResponse struct {
}

func (m *MsgUpdateWhoisResponse) Reset()         { *m = MsgUpdateWhoisResponse{} }
func (m *MsgUpdateWhoisResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateWhoisResponse) ProtoMessage()    {}
func (*MsgUpdateWhoisResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4dec9b515ec245, []int{3}
}
func (m *MsgUpdateWhoisResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateWhoisResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateWhoisResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateWhoisResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateWhoisResponse.Merge(m, src)
}
func (m *MsgUpdateWhoisResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateWhoisResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateWhoisResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateWhoisResponse proto.InternalMessageInfo

type MsgDeleteWhois struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Id    uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgDeleteWhois) Reset()         { *m = MsgDeleteWhois{} }
func (m *MsgDeleteWhois) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteWhois) ProtoMessage()    {}
func (*MsgDeleteWhois) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4dec9b515ec245, []int{4}
}
func (m *MsgDeleteWhois) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteWhois) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteWhois.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteWhois) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteWhois.Merge(m, src)
}
func (m *MsgDeleteWhois) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteWhois) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteWhois.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteWhois proto.InternalMessageInfo

func (m *MsgDeleteWhois) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgDeleteWhois) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MsgDeleteWhoisResponse struct {
}

func (m *MsgDeleteWhoisResponse) Reset()         { *m = MsgDeleteWhoisResponse{} }
func (m *MsgDeleteWhoisResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteWhoisResponse) ProtoMessage()    {}
func (*MsgDeleteWhoisResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4dec9b515ec245, []int{5}
}
func (m *MsgDeleteWhoisResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteWhoisResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteWhoisResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteWhoisResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteWhoisResponse.Merge(m, src)
}
func (m *MsgDeleteWhoisResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteWhoisResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteWhoisResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteWhoisResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateWhois)(nil), "username.nameservice.nameservice.MsgCreateWhois")
	proto.RegisterType((*MsgCreateWhoisResponse)(nil), "username.nameservice.nameservice.MsgCreateWhoisResponse")
	proto.RegisterType((*MsgUpdateWhois)(nil), "username.nameservice.nameservice.MsgUpdateWhois")
	proto.RegisterType((*MsgUpdateWhoisResponse)(nil), "username.nameservice.nameservice.MsgUpdateWhoisResponse")
	proto.RegisterType((*MsgDeleteWhois)(nil), "username.nameservice.nameservice.MsgDeleteWhois")
	proto.RegisterType((*MsgDeleteWhoisResponse)(nil), "username.nameservice.nameservice.MsgDeleteWhoisResponse")
}

func init() { proto.RegisterFile("nameservice/tx.proto", fileDescriptor_df4dec9b515ec245) }

var fileDescriptor_df4dec9b515ec245 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x41, 0x6b, 0x83, 0x30,
	0x1c, 0xc5, 0x1b, 0x6b, 0x07, 0x4b, 0xa1, 0x87, 0x50, 0x86, 0xec, 0x10, 0xc4, 0x93, 0x27, 0x1d,
	0x2b, 0x8c, 0x9d, 0xb7, 0x5d, 0x85, 0x21, 0x8c, 0xc1, 0x4e, 0x6b, 0xf5, 0x8f, 0x0d, 0xac, 0x46,
	0x12, 0xbb, 0xd5, 0x6f, 0xb1, 0x8f, 0xb5, 0x63, 0x8f, 0x3b, 0x0e, 0xfd, 0x0e, 0x3b, 0x0f, 0xe3,
	0x2c, 0x11, 0x84, 0xda, 0x5b, 0xfe, 0xc6, 0xf7, 0x7e, 0x8f, 0x17, 0xfe, 0x78, 0x9e, 0x2e, 0x37,
	0x20, 0x41, 0xbc, 0xb3, 0x08, 0xfc, 0x7c, 0xe7, 0x65, 0x82, 0xe7, 0x9c, 0xd8, 0x5b, 0x09, 0xa2,
	0xbe, 0xf1, 0xb4, 0x6b, 0xfd, 0xec, 0x3c, 0xe2, 0x59, 0x20, 0x93, 0x7b, 0x01, 0xcb, 0x1c, 0x9e,
	0xd7, 0x9c, 0x49, 0x32, 0xc7, 0x13, 0xfe, 0x91, 0x82, 0xb0, 0x90, 0x8d, 0xdc, 0xf3, 0xb0, 0x19,
	0x08, 0xc1, 0x66, 0x2d, 0xb3, 0x0c, 0xf5, 0x51, 0x9d, 0xeb, 0x3f, 0x33, 0xc1, 0x22, 0xb0, 0xc6,
	0x36, 0x72, 0x27, 0x61, 0x33, 0x38, 0x2e, 0xbe, 0xe8, 0x3a, 0x86, 0x20, 0x33, 0x9e, 0x4a, 0x20,
	0x33, 0x6c, 0xb0, 0x58, 0xd9, 0x9a, 0xa1, 0xc1, 0x62, 0xe7, 0x55, 0xb1, 0x9f, 0xb2, 0xf8, 0x08,
	0xbb, 0xd1, 0x19, 0xad, 0xee, 0x90, 0x65, 0xdc, 0x97, 0xc5, 0xd4, 0xb3, 0x58, 0x2a, 0x8b, 0x46,
	0x68, 0xb3, 0x38, 0x37, 0x8a, 0xfd, 0x00, 0x6f, 0x70, 0x12, 0xfb, 0xdf, 0x51, 0xd3, 0xb5, 0x8e,
	0xd7, 0xbf, 0x06, 0x1e, 0x07, 0x32, 0x21, 0x05, 0x9e, 0xea, 0x75, 0x5e, 0x79, 0xc7, 0xde, 0xc0,
	0xeb, 0xd6, 0x75, 0x79, 0x7b, 0xaa, 0xe2, 0x50, 0x70, 0x81, 0xa7, 0x7a, 0x9b, 0xc3, 0xd0, 0x9a,
	0x62, 0x20, 0xba, 0xa7, 0xcf, 0x1a, 0xad, 0x97, 0x39, 0x0c, 0xad, 0x29, 0x06, 0xa2, 0x7b, 0x8a,
	0xbf, 0x0b, 0xbe, 0x4a, 0x8a, 0xf6, 0x25, 0x45, 0x3f, 0x25, 0x45, 0x9f, 0x15, 0x1d, 0xed, 0x2b,
	0x3a, 0xfa, 0xae, 0xe8, 0xe8, 0x65, 0x91, 0xb0, 0x7c, 0xbd, 0x5d, 0x79, 0x11, 0xdf, 0xf8, 0xad,
	0xbb, 0xaf, 0x2f, 0xca, 0xae, 0x33, 0xe5, 0x45, 0x06, 0x72, 0x75, 0xa6, 0x56, 0x67, 0xf1, 0x17,
	0x00, 0x00, 0xff, 0xff, 0x79, 0x48, 0x12, 0x6b, 0x52, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	CreateWhois(ctx context.Context, in *MsgCreateWhois, opts ...grpc.CallOption) (*MsgCreateWhoisResponse, error)
	UpdateWhois(ctx context.Context, in *MsgUpdateWhois, opts ...grpc.CallOption) (*MsgUpdateWhoisResponse, error)
	DeleteWhois(ctx context.Context, in *MsgDeleteWhois, opts ...grpc.CallOption) (*MsgDeleteWhoisResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateWhois(ctx context.Context, in *MsgCreateWhois, opts ...grpc.CallOption) (*MsgCreateWhoisResponse, error) {
	out := new(MsgCreateWhoisResponse)
	err := c.cc.Invoke(ctx, "/username.nameservice.nameservice.Msg/CreateWhois", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateWhois(ctx context.Context, in *MsgUpdateWhois, opts ...grpc.CallOption) (*MsgUpdateWhoisResponse, error) {
	out := new(MsgUpdateWhoisResponse)
	err := c.cc.Invoke(ctx, "/username.nameservice.nameservice.Msg/UpdateWhois", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteWhois(ctx context.Context, in *MsgDeleteWhois, opts ...grpc.CallOption) (*MsgDeleteWhoisResponse, error) {
	out := new(MsgDeleteWhoisResponse)
	err := c.cc.Invoke(ctx, "/username.nameservice.nameservice.Msg/DeleteWhois", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	CreateWhois(context.Context, *MsgCreateWhois) (*MsgCreateWhoisResponse, error)
	UpdateWhois(context.Context, *MsgUpdateWhois) (*MsgUpdateWhoisResponse, error)
	DeleteWhois(context.Context, *MsgDeleteWhois) (*MsgDeleteWhoisResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateWhois(ctx context.Context, req *MsgCreateWhois) (*MsgCreateWhoisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWhois not implemented")
}
func (*UnimplementedMsgServer) UpdateWhois(ctx context.Context, req *MsgUpdateWhois) (*MsgUpdateWhoisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWhois not implemented")
}
func (*UnimplementedMsgServer) DeleteWhois(ctx context.Context, req *MsgDeleteWhois) (*MsgDeleteWhoisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWhois not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateWhois_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateWhois)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateWhois(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/username.nameservice.nameservice.Msg/CreateWhois",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateWhois(ctx, req.(*MsgCreateWhois))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateWhois_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateWhois)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateWhois(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/username.nameservice.nameservice.Msg/UpdateWhois",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateWhois(ctx, req.(*MsgUpdateWhois))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteWhois_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteWhois)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteWhois(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/username.nameservice.nameservice.Msg/DeleteWhois",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteWhois(ctx, req.(*MsgDeleteWhois))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "username.nameservice.nameservice.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWhois",
			Handler:    _Msg_CreateWhois_Handler,
		},
		{
			MethodName: "UpdateWhois",
			Handler:    _Msg_UpdateWhois_Handler,
		},
		{
			MethodName: "DeleteWhois",
			Handler:    _Msg_DeleteWhois_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nameservice/tx.proto",
}

func (m *MsgCreateWhois) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateWhois) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateWhois) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Price != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Price))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateWhoisResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateWhoisResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateWhoisResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateWhois) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateWhois) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateWhois) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Price != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Price))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateWhoisResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateWhoisResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateWhoisResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteWhois) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteWhois) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteWhois) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteWhoisResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteWhoisResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteWhoisResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateWhois) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Price != 0 {
		n += 1 + sovTx(uint64(m.Price))
	}
	return n
}

func (m *MsgCreateWhoisResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func (m *MsgUpdateWhois) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Price != 0 {
		n += 1 + sovTx(uint64(m.Price))
	}
	return n
}

func (m *MsgUpdateWhoisResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteWhois) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func (m *MsgDeleteWhoisResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateWhois) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateWhois: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateWhois: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			m.Price = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Price |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateWhoisResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateWhoisResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateWhoisResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateWhois) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateWhois: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateWhois: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			m.Price = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Price |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateWhoisResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateWhoisResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateWhoisResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeleteWhois) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeleteWhois: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteWhois: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeleteWhoisResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeleteWhoisResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteWhoisResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)