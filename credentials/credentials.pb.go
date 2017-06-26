// Code generated by protoc-gen-go.
// source: credentials.proto
// DO NOT EDIT!

/*
Package credentials is a generated protocol buffer package.

Storage of applications

It is generated from these files:
	credentials.proto

It has these top-level messages:
	Credential
	RetrieveRequest
	StoreRequest
	StoreResponse
	CredentialsResponse
*/
package credentials

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

type Credential struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Value     string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	ProjectId string `protobuf:"bytes,3,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	UserId    string `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	VendorId  string `protobuf:"bytes,5,opt,name=vendor_id,json=vendorId" json:"vendor_id,omitempty"`
	AppId     string `protobuf:"bytes,6,opt,name=app_id,json=appId" json:"app_id,omitempty"`
}

func (m *Credential) Reset()                    { *m = Credential{} }
func (m *Credential) String() string            { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()               {}
func (*Credential) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Credential) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Credential) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Credential) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *Credential) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Credential) GetVendorId() string {
	if m != nil {
		return m.VendorId
	}
	return ""
}

func (m *Credential) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type RetrieveRequest struct {
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	VendorId  string `protobuf:"bytes,3,opt,name=vendor_id,json=vendorId" json:"vendor_id,omitempty"`
	AppId     string `protobuf:"bytes,4,opt,name=app_id,json=appId" json:"app_id,omitempty"`
}

func (m *RetrieveRequest) Reset()                    { *m = RetrieveRequest{} }
func (m *RetrieveRequest) String() string            { return proto.CompactTextString(m) }
func (*RetrieveRequest) ProtoMessage()               {}
func (*RetrieveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RetrieveRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *RetrieveRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *RetrieveRequest) GetVendorId() string {
	if m != nil {
		return m.VendorId
	}
	return ""
}

func (m *RetrieveRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type StoreRequest struct {
	ProjectId   string        `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	UserId      string        `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	VendorId    string        `protobuf:"bytes,3,opt,name=vendor_id,json=vendorId" json:"vendor_id,omitempty"`
	AppId       string        `protobuf:"bytes,4,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	Credentials []*Credential `protobuf:"bytes,5,rep,name=credentials" json:"credentials,omitempty"`
}

func (m *StoreRequest) Reset()                    { *m = StoreRequest{} }
func (m *StoreRequest) String() string            { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()               {}
func (*StoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StoreRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *StoreRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *StoreRequest) GetVendorId() string {
	if m != nil {
		return m.VendorId
	}
	return ""
}

func (m *StoreRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *StoreRequest) GetCredentials() []*Credential {
	if m != nil {
		return m.Credentials
	}
	return nil
}

type StoreResponse struct {
	Stored bool `protobuf:"varint,1,opt,name=stored" json:"stored,omitempty"`
}

func (m *StoreResponse) Reset()                    { *m = StoreResponse{} }
func (m *StoreResponse) String() string            { return proto.CompactTextString(m) }
func (*StoreResponse) ProtoMessage()               {}
func (*StoreResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StoreResponse) GetStored() bool {
	if m != nil {
		return m.Stored
	}
	return false
}

type CredentialsResponse struct {
	Credentials []*Credential `protobuf:"bytes,1,rep,name=credentials" json:"credentials,omitempty"`
}

func (m *CredentialsResponse) Reset()                    { *m = CredentialsResponse{} }
func (m *CredentialsResponse) String() string            { return proto.CompactTextString(m) }
func (*CredentialsResponse) ProtoMessage()               {}
func (*CredentialsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CredentialsResponse) GetCredentials() []*Credential {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func init() {
	proto.RegisterType((*Credential)(nil), "credentials.Credential")
	proto.RegisterType((*RetrieveRequest)(nil), "credentials.RetrieveRequest")
	proto.RegisterType((*StoreRequest)(nil), "credentials.StoreRequest")
	proto.RegisterType((*StoreResponse)(nil), "credentials.StoreResponse")
	proto.RegisterType((*CredentialsResponse)(nil), "credentials.CredentialsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Credentials service

type CredentialsClient interface {
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error)
	Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*CredentialsResponse, error)
}

type credentialsClient struct {
	cc *grpc.ClientConn
}

func NewCredentialsClient(cc *grpc.ClientConn) CredentialsClient {
	return &credentialsClient{cc}
}

func (c *credentialsClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := grpc.Invoke(ctx, "/credentials.Credentials/Store", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialsClient) Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*CredentialsResponse, error) {
	out := new(CredentialsResponse)
	err := grpc.Invoke(ctx, "/credentials.Credentials/Retrieve", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Credentials service

type CredentialsServer interface {
	Store(context.Context, *StoreRequest) (*StoreResponse, error)
	Retrieve(context.Context, *RetrieveRequest) (*CredentialsResponse, error)
}

func RegisterCredentialsServer(s *grpc.Server, srv CredentialsServer) {
	s.RegisterService(&_Credentials_serviceDesc, srv)
}

func _Credentials_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credentials.Credentials/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Credentials_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credentials.Credentials/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServer).Retrieve(ctx, req.(*RetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Credentials_serviceDesc = grpc.ServiceDesc{
	ServiceName: "credentials.Credentials",
	HandlerType: (*CredentialsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _Credentials_Store_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _Credentials_Retrieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "credentials.proto",
}

func init() { proto.RegisterFile("credentials.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x93, 0xcd, 0x4e, 0xc2, 0x40,
	0x14, 0x85, 0x33, 0x85, 0x56, 0xb8, 0xf8, 0x13, 0xaf, 0x3f, 0x54, 0xd4, 0x84, 0x90, 0x18, 0xd8,
	0x08, 0x09, 0xae, 0xdc, 0xb8, 0xd0, 0x55, 0x5d, 0x99, 0xba, 0x73, 0x63, 0x0a, 0x73, 0x83, 0x55,
	0x64, 0xc6, 0xce, 0x94, 0xb8, 0xf1, 0x41, 0xdc, 0xf8, 0x16, 0xbe, 0x9f, 0xe9, 0xb4, 0xc2, 0x40,
	0xc4, 0xb8, 0x73, 0x79, 0xcf, 0x99, 0x99, 0xef, 0xb4, 0x67, 0x06, 0xb6, 0x87, 0x09, 0x71, 0x9a,
	0xe8, 0x38, 0x1a, 0xab, 0xae, 0x4c, 0x84, 0x16, 0x58, 0xb3, 0xa4, 0xd6, 0x07, 0x03, 0xb8, 0x9a,
	0xcd, 0xb8, 0x09, 0x4e, 0xcc, 0x7d, 0xd6, 0x64, 0x9d, 0x6a, 0xe8, 0xc4, 0x1c, 0x77, 0xc1, 0x9d,
	0x46, 0xe3, 0x94, 0x7c, 0xc7, 0x48, 0xf9, 0x80, 0xc7, 0x00, 0x32, 0x11, 0x8f, 0x34, 0xd4, 0xf7,
	0x31, 0xf7, 0x4b, 0xc6, 0xaa, 0x16, 0x4a, 0xc0, 0xb1, 0x0e, 0x6b, 0xa9, 0xa2, 0x24, 0xf3, 0xca,
	0xc6, 0xf3, 0xb2, 0x31, 0xe0, 0x78, 0x08, 0xd5, 0x29, 0x4d, 0xb8, 0x30, 0x96, 0x6b, 0xac, 0x4a,
	0x2e, 0x04, 0x1c, 0xf7, 0xc0, 0x8b, 0xa4, 0xcc, 0x1c, 0x2f, 0x67, 0x45, 0x52, 0x06, 0xbc, 0xf5,
	0x06, 0x5b, 0x21, 0xe9, 0x24, 0xa6, 0x29, 0x85, 0xf4, 0x92, 0x92, 0xd2, 0x4b, 0x78, 0xf6, 0x0b,
	0xde, 0x59, 0x8d, 0x2f, 0xad, 0xc4, 0x97, 0x6d, 0xfc, 0x27, 0x83, 0xf5, 0x5b, 0x2d, 0x92, 0xff,
	0x80, 0xe3, 0x39, 0xd8, 0x5d, 0xf9, 0x6e, 0xb3, 0xd4, 0xa9, 0xf5, 0xeb, 0x5d, 0xbb, 0xd2, 0x79,
	0x77, 0xe1, 0x42, 0xaf, 0x6d, 0xd8, 0x28, 0x62, 0x2b, 0x29, 0x26, 0x8a, 0x70, 0x1f, 0x3c, 0x95,
	0x09, 0x79, 0xe6, 0x4a, 0x58, 0x4c, 0xad, 0x1b, 0xd8, 0x99, 0x9f, 0xa1, 0x66, 0xcb, 0x97, 0xd0,
	0xec, 0xef, 0xe8, 0xfe, 0x3b, 0x83, 0x9a, 0x75, 0x24, 0x5e, 0x80, 0x6b, 0xa2, 0xe0, 0xc1, 0xc2,
	0x76, 0xfb, 0xaf, 0x36, 0x1a, 0x3f, 0x59, 0x45, 0x94, 0x6b, 0xa8, 0x7c, 0xdf, 0x00, 0x3c, 0x5a,
	0x58, 0xb7, 0x74, 0x31, 0x1a, 0xcd, 0x15, 0xf9, 0x66, 0x9f, 0x75, 0xd9, 0xbe, 0x3b, 0x19, 0xc5,
	0xfa, 0x21, 0x1d, 0x74, 0x87, 0xe2, 0xb9, 0xf7, 0x94, 0x0e, 0xe8, 0xb5, 0x67, 0x5e, 0xc5, 0xe9,
	0x48, 0xf4, 0xac, 0xcd, 0x03, 0xcf, 0xa8, 0x67, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x75,
	0x12, 0x38, 0x40, 0x03, 0x00, 0x00,
}