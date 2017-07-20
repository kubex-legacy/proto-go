// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

/*
Package users is a generated protocol buffer package.

It is generated from these files:
	users.proto

It has these top-level messages:
	RetrieveRequest
	UserResponse
	UsersResponse
*/
package users

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RetrieveRequest struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *RetrieveRequest) Reset()                    { *m = RetrieveRequest{} }
func (m *RetrieveRequest) String() string            { return proto.CompactTextString(m) }
func (*RetrieveRequest) ProtoMessage()               {}
func (*RetrieveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RetrieveRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type UserResponse struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
}

func (m *UserResponse) Reset()                    { *m = UserResponse{} }
func (m *UserResponse) String() string            { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()               {}
func (*UserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserResponse) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type UsersResponse struct {
	Users map[string]*UserResponse `protobuf:"bytes,1,rep,name=users" json:"users,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *UsersResponse) Reset()                    { *m = UsersResponse{} }
func (m *UsersResponse) String() string            { return proto.CompactTextString(m) }
func (*UsersResponse) ProtoMessage()               {}
func (*UsersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UsersResponse) GetUsers() map[string]*UserResponse {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*RetrieveRequest)(nil), "users.RetrieveRequest")
	proto.RegisterType((*UserResponse)(nil), "users.UserResponse")
	proto.RegisterType((*UsersResponse)(nil), "users.UsersResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Users service

type UsersClient interface {
	Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*UsersResponse, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := grpc.Invoke(ctx, "/users.Users/Retrieve", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Users service

type UsersServer interface {
	Retrieve(context.Context, *RetrieveRequest) (*UsersResponse, error)
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Retrieve(ctx, req.(*RetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "users.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Retrieve",
			Handler:    _Users_Retrieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}

func init() { proto.RegisterFile("users.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xcd, 0x4b, 0xfb, 0x40,
	0x10, 0x25, 0xe9, 0x2f, 0x3f, 0x9a, 0x89, 0x5f, 0xac, 0x45, 0x42, 0x45, 0x5a, 0xe3, 0xa5, 0x1e,
	0x4c, 0x20, 0x22, 0x14, 0x6f, 0x0a, 0x1e, 0xf5, 0xb0, 0xe0, 0xc5, 0x8b, 0x24, 0x66, 0xac, 0x4b,
	0xf3, 0xe5, 0xee, 0xa6, 0xd8, 0x3f, 0xc4, 0xff, 0x57, 0x76, 0xb6, 0xf1, 0x0b, 0x6f, 0xfb, 0xde,
	0x9b, 0x79, 0xf3, 0x66, 0x16, 0x82, 0x4e, 0xa1, 0x54, 0x71, 0x2b, 0x1b, 0xdd, 0x30, 0x8f, 0x40,
	0x74, 0x02, 0xbb, 0x1c, 0xb5, 0x14, 0xb8, 0x42, 0x8e, 0xaf, 0x1d, 0x2a, 0xcd, 0xf6, 0x60, 0x20,
	0x0a, 0x15, 0x3a, 0xd3, 0xc1, 0xcc, 0xe7, 0xe6, 0x19, 0xb5, 0xb0, 0x75, 0xaf, 0x50, 0x72, 0x54,
	0x6d, 0x53, 0x2b, 0x64, 0x3b, 0xe0, 0x8a, 0x22, 0x74, 0xa6, 0xce, 0xcc, 0xe7, 0xae, 0x28, 0xd8,
	0x08, 0x3c, 0xac, 0x32, 0x51, 0x86, 0x2e, 0x51, 0x16, 0xb0, 0x23, 0x80, 0x67, 0x21, 0x95, 0x7e,
	0xac, 0xb3, 0x0a, 0xc3, 0x01, 0x49, 0x3e, 0x31, 0x77, 0x59, 0x85, 0xec, 0x10, 0xfc, 0x32, 0xeb,
	0xd5, 0x7f, 0xa4, 0x0e, 0x0d, 0x61, 0xc4, 0xe8, 0xdd, 0x81, 0x6d, 0x33, 0x52, 0x7d, 0xce, 0xbc,
	0x00, 0x9b, 0x98, 0x72, 0x05, 0xe9, 0x24, 0xb6, 0xcb, 0xfc, 0x28, 0xb2, 0xe8, 0xa6, 0xd6, 0x72,
	0xcd, 0x6d, 0xf5, 0xf8, 0x16, 0xe0, 0x8b, 0x34, 0xab, 0x2d, 0x71, 0xbd, 0x49, 0x6e, 0x9e, 0xec,
	0x14, 0xbc, 0x55, 0x56, 0x76, 0x48, 0xd1, 0x83, 0x74, 0xff, 0x9b, 0x6d, 0xef, 0xca, 0x6d, 0xc5,
	0xa5, 0x3b, 0x77, 0xd2, 0x2b, 0xf0, 0xc8, 0x8e, 0xcd, 0x61, 0xd8, 0xdf, 0x8d, 0x1d, 0x6c, 0x9a,
	0x7e, 0x1d, 0x72, 0x3c, 0xfa, 0x2b, 0xe3, 0xf5, 0xf1, 0xc3, 0x64, 0x21, 0xf4, 0x4b, 0x97, 0xc7,
	0x4f, 0x4d, 0x95, 0x2c, 0xbb, 0x1c, 0xdf, 0x12, 0xfa, 0x92, 0xb3, 0x45, 0x93, 0x50, 0x43, 0xfe,
	0x9f, 0xf0, 0xf9, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xfc, 0xb7, 0x0a, 0xb1, 0x01, 0x00,
	0x00,
}