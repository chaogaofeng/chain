// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chain/permission/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be400f032c1f7d4c, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be400f032c1f7d4c, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetRoleAccountRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryGetRoleAccountRequest) Reset()         { *m = QueryGetRoleAccountRequest{} }
func (m *QueryGetRoleAccountRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetRoleAccountRequest) ProtoMessage()    {}
func (*QueryGetRoleAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be400f032c1f7d4c, []int{2}
}
func (m *QueryGetRoleAccountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetRoleAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetRoleAccountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetRoleAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetRoleAccountRequest.Merge(m, src)
}
func (m *QueryGetRoleAccountRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetRoleAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetRoleAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetRoleAccountRequest proto.InternalMessageInfo

func (m *QueryGetRoleAccountRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type QueryGetRoleAccountResponse struct {
	RoleAccount RoleAccount `protobuf:"bytes,1,opt,name=roleAccount,proto3" json:"roleAccount"`
}

func (m *QueryGetRoleAccountResponse) Reset()         { *m = QueryGetRoleAccountResponse{} }
func (m *QueryGetRoleAccountResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetRoleAccountResponse) ProtoMessage()    {}
func (*QueryGetRoleAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be400f032c1f7d4c, []int{3}
}
func (m *QueryGetRoleAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetRoleAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetRoleAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetRoleAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetRoleAccountResponse.Merge(m, src)
}
func (m *QueryGetRoleAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetRoleAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetRoleAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetRoleAccountResponse proto.InternalMessageInfo

func (m *QueryGetRoleAccountResponse) GetRoleAccount() RoleAccount {
	if m != nil {
		return m.RoleAccount
	}
	return RoleAccount{}
}

type QueryAllRoleAccountRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllRoleAccountRequest) Reset()         { *m = QueryAllRoleAccountRequest{} }
func (m *QueryAllRoleAccountRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllRoleAccountRequest) ProtoMessage()    {}
func (*QueryAllRoleAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be400f032c1f7d4c, []int{4}
}
func (m *QueryAllRoleAccountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllRoleAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllRoleAccountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllRoleAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllRoleAccountRequest.Merge(m, src)
}
func (m *QueryAllRoleAccountRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllRoleAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllRoleAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllRoleAccountRequest proto.InternalMessageInfo

func (m *QueryAllRoleAccountRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllRoleAccountResponse struct {
	RoleAccount []RoleAccount       `protobuf:"bytes,1,rep,name=roleAccount,proto3" json:"roleAccount"`
	Pagination  *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllRoleAccountResponse) Reset()         { *m = QueryAllRoleAccountResponse{} }
func (m *QueryAllRoleAccountResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllRoleAccountResponse) ProtoMessage()    {}
func (*QueryAllRoleAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be400f032c1f7d4c, []int{5}
}
func (m *QueryAllRoleAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllRoleAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllRoleAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllRoleAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllRoleAccountResponse.Merge(m, src)
}
func (m *QueryAllRoleAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllRoleAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllRoleAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllRoleAccountResponse proto.InternalMessageInfo

func (m *QueryAllRoleAccountResponse) GetRoleAccount() []RoleAccount {
	if m != nil {
		return m.RoleAccount
	}
	return nil
}

func (m *QueryAllRoleAccountResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "glodnet.chain.permission.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "glodnet.chain.permission.QueryParamsResponse")
	proto.RegisterType((*QueryGetRoleAccountRequest)(nil), "glodnet.chain.permission.QueryGetRoleAccountRequest")
	proto.RegisterType((*QueryGetRoleAccountResponse)(nil), "glodnet.chain.permission.QueryGetRoleAccountResponse")
	proto.RegisterType((*QueryAllRoleAccountRequest)(nil), "glodnet.chain.permission.QueryAllRoleAccountRequest")
	proto.RegisterType((*QueryAllRoleAccountResponse)(nil), "glodnet.chain.permission.QueryAllRoleAccountResponse")
}

func init() { proto.RegisterFile("chain/permission/query.proto", fileDescriptor_be400f032c1f7d4c) }

var fileDescriptor_be400f032c1f7d4c = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4f, 0x6b, 0x13, 0x41,
	0x18, 0xc6, 0xb3, 0xb5, 0x46, 0x9c, 0x80, 0x87, 0xb1, 0x87, 0xb0, 0xd6, 0x35, 0x8c, 0x58, 0x83,
	0x7f, 0x66, 0x6c, 0xfc, 0x83, 0x27, 0x21, 0x05, 0xed, 0x49, 0xa8, 0x0b, 0x5e, 0xbc, 0xc8, 0x64,
	0x33, 0x6c, 0x17, 0x26, 0xf3, 0x6e, 0x77, 0x26, 0x62, 0x11, 0x41, 0xfc, 0x02, 0x0a, 0x7e, 0x0e,
	0x6f, 0xfa, 0x1d, 0x7a, 0x2c, 0x78, 0xf1, 0x24, 0x92, 0xf8, 0x41, 0x24, 0x33, 0x53, 0xb2, 0x6d,
	0x76, 0x49, 0xc4, 0x5b, 0x32, 0xf3, 0x3c, 0xef, 0xf3, 0x7b, 0xf7, 0x7d, 0x77, 0xd1, 0x66, 0xb2,
	0xcf, 0x33, 0xc5, 0x72, 0x51, 0x8c, 0x32, 0xad, 0x33, 0x50, 0xec, 0x60, 0x2c, 0x8a, 0x43, 0x9a,
	0x17, 0x60, 0x00, 0xb7, 0x53, 0x09, 0x43, 0x25, 0x0c, 0xb5, 0x2a, 0x3a, 0x57, 0x85, 0x1b, 0x29,
	0xa4, 0x60, 0x45, 0x6c, 0xf6, 0xcb, 0xe9, 0xc3, 0xcd, 0x14, 0x20, 0x95, 0x82, 0xf1, 0x3c, 0x63,
	0x5c, 0x29, 0x30, 0xdc, 0x64, 0xa0, 0xb4, 0xbf, 0xbd, 0x95, 0x80, 0x1e, 0x81, 0x66, 0x03, 0xae,
	0x85, 0x8b, 0x61, 0x6f, 0xb6, 0x07, 0xc2, 0xf0, 0x6d, 0x96, 0xf3, 0x34, 0x53, 0x56, 0xec, 0xb5,
	0x57, 0x17, 0xb8, 0x72, 0x5e, 0xf0, 0xd1, 0x49, 0xa9, 0xeb, 0x0b, 0xd7, 0x05, 0x48, 0xf1, 0x9a,
	0x27, 0x09, 0x8c, 0x95, 0x71, 0x22, 0xb2, 0x81, 0xf0, 0x8b, 0x59, 0xca, 0x9e, 0x75, 0xc6, 0xe2,
	0x60, 0x2c, 0xb4, 0x21, 0x2f, 0xd1, 0xe5, 0x53, 0xa7, 0x3a, 0x07, 0xa5, 0x05, 0x7e, 0x82, 0x9a,
	0x2e, 0xa1, 0x1d, 0x74, 0x82, 0x6e, 0xab, 0xd7, 0xa1, 0x75, 0xbd, 0x53, 0xe7, 0xdc, 0x59, 0x3f,
	0xfa, 0x75, 0xad, 0x11, 0x7b, 0x17, 0x79, 0x84, 0x42, 0x5b, 0x76, 0x57, 0x98, 0x18, 0xa4, 0xe8,
	0x3b, 0x12, 0x1f, 0x8a, 0xdb, 0xe8, 0x02, 0x1f, 0x0e, 0x0b, 0xa1, 0x5d, 0xf9, 0x8b, 0xf1, 0xc9,
	0x5f, 0x22, 0xd1, 0x95, 0x4a, 0x9f, 0xc7, 0x7a, 0x8e, 0x5a, 0xc5, 0xfc, 0xd8, 0xb3, 0xdd, 0xa8,
	0x67, 0x2b, 0xd5, 0xf0, 0x80, 0x65, 0x3f, 0x19, 0x7a, 0xca, 0xbe, 0x94, 0x15, 0x94, 0xcf, 0x10,
	0x9a, 0x0f, 0xc2, 0x67, 0x6d, 0x51, 0x37, 0x35, 0x3a, 0x9b, 0x1a, 0x75, 0xcb, 0xe1, 0xa7, 0x46,
	0xf7, 0x78, 0x2a, 0xbc, 0x37, 0x2e, 0x39, 0xc9, 0xf7, 0xc0, 0x37, 0x75, 0x36, 0xa6, 0xae, 0xa9,
	0x73, 0xff, 0xd3, 0x14, 0xde, 0x3d, 0x85, 0xbd, 0x66, 0xb1, 0x6f, 0x2e, 0xc5, 0x76, 0x2c, 0x65,
	0xee, 0xde, 0x87, 0x75, 0x74, 0xde, 0x72, 0xe3, 0x4f, 0x01, 0x6a, 0xba, 0x31, 0xe3, 0x3b, 0xf5,
	0x5c, 0x8b, 0xdb, 0x15, 0xde, 0x5d, 0x51, 0xed, 0xd2, 0x49, 0xf7, 0xe3, 0x8f, 0x3f, 0x5f, 0xd6,
	0x08, 0xee, 0x30, 0x6f, 0x63, 0x35, 0x7b, 0x8f, 0xbf, 0x05, 0xa8, 0x55, 0x7a, 0x0e, 0xf8, 0xc1,
	0x92, 0xa0, 0xca, 0x3d, 0x0c, 0x1f, 0xfe, 0xa3, 0xcb, 0x63, 0x3e, 0xb6, 0x98, 0x3d, 0x7c, 0xaf,
	0x1e, 0xb3, 0xfc, 0xfe, 0xb1, 0x77, 0x7e, 0xbb, 0xdf, 0xe3, 0xaf, 0x01, 0xba, 0x54, 0xaa, 0xd8,
	0x97, 0x72, 0x29, 0x79, 0xe5, 0x6e, 0x2e, 0x25, 0xaf, 0x5e, 0x35, 0x42, 0x2d, 0x79, 0x17, 0x6f,
	0xad, 0x46, 0xbe, 0xf3, 0xf4, 0x68, 0x12, 0x05, 0xc7, 0x93, 0x28, 0xf8, 0x3d, 0x89, 0x82, 0xcf,
	0xd3, 0xa8, 0x71, 0x3c, 0x8d, 0x1a, 0x3f, 0xa7, 0x51, 0xe3, 0xd5, 0xed, 0x34, 0x33, 0xfb, 0xe3,
	0x01, 0x4d, 0x60, 0x74, 0xa6, 0xd6, 0xdb, 0x72, 0x35, 0x73, 0x98, 0x0b, 0x3d, 0x68, 0xda, 0x2f,
	0xd0, 0xfd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x80, 0xba, 0xc6, 0x5f, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a RoleAccount by index.
	RoleAccount(ctx context.Context, in *QueryGetRoleAccountRequest, opts ...grpc.CallOption) (*QueryGetRoleAccountResponse, error)
	// Queries a list of RoleAccount items.
	RoleAccountAll(ctx context.Context, in *QueryAllRoleAccountRequest, opts ...grpc.CallOption) (*QueryAllRoleAccountResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/glodnet.chain.permission.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RoleAccount(ctx context.Context, in *QueryGetRoleAccountRequest, opts ...grpc.CallOption) (*QueryGetRoleAccountResponse, error) {
	out := new(QueryGetRoleAccountResponse)
	err := c.cc.Invoke(ctx, "/glodnet.chain.permission.Query/RoleAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RoleAccountAll(ctx context.Context, in *QueryAllRoleAccountRequest, opts ...grpc.CallOption) (*QueryAllRoleAccountResponse, error) {
	out := new(QueryAllRoleAccountResponse)
	err := c.cc.Invoke(ctx, "/glodnet.chain.permission.Query/RoleAccountAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a RoleAccount by index.
	RoleAccount(context.Context, *QueryGetRoleAccountRequest) (*QueryGetRoleAccountResponse, error)
	// Queries a list of RoleAccount items.
	RoleAccountAll(context.Context, *QueryAllRoleAccountRequest) (*QueryAllRoleAccountResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) RoleAccount(ctx context.Context, req *QueryGetRoleAccountRequest) (*QueryGetRoleAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleAccount not implemented")
}
func (*UnimplementedQueryServer) RoleAccountAll(ctx context.Context, req *QueryAllRoleAccountRequest) (*QueryAllRoleAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleAccountAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glodnet.chain.permission.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RoleAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetRoleAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RoleAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glodnet.chain.permission.Query/RoleAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RoleAccount(ctx, req.(*QueryGetRoleAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RoleAccountAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllRoleAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RoleAccountAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glodnet.chain.permission.Query/RoleAccountAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RoleAccountAll(ctx, req.(*QueryAllRoleAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "glodnet.chain.permission.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "RoleAccount",
			Handler:    _Query_RoleAccount_Handler,
		},
		{
			MethodName: "RoleAccountAll",
			Handler:    _Query_RoleAccountAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chain/permission/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetRoleAccountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetRoleAccountRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetRoleAccountRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetRoleAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetRoleAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetRoleAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.RoleAccount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllRoleAccountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllRoleAccountRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllRoleAccountRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllRoleAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllRoleAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllRoleAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.RoleAccount) > 0 {
		for iNdEx := len(m.RoleAccount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RoleAccount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetRoleAccountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetRoleAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.RoleAccount.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllRoleAccountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllRoleAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RoleAccount) > 0 {
		for _, e := range m.RoleAccount {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetRoleAccountRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetRoleAccountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetRoleAccountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetRoleAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetRoleAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetRoleAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RoleAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllRoleAccountRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllRoleAccountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllRoleAccountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllRoleAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllRoleAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllRoleAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoleAccount = append(m.RoleAccount, RoleAccount{})
			if err := m.RoleAccount[len(m.RoleAccount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
