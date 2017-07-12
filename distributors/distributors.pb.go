// Code generated by protoc-gen-go. DO NOT EDIT.
// source: distributors.proto

/*
Package distributors is a generated protocol buffer package.

It is generated from these files:
	distributors.proto

It has these top-level messages:
	RetrieveRequest
	DistributorResponse
*/
package distributors

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
	Domain string `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
}

func (m *RetrieveRequest) Reset()                    { *m = RetrieveRequest{} }
func (m *RetrieveRequest) String() string            { return proto.CompactTextString(m) }
func (*RetrieveRequest) ProtoMessage()               {}
func (*RetrieveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RetrieveRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

type DistributorResponse struct {
	Id           string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Domain       string `protobuf:"bytes,4,opt,name=domain" json:"domain,omitempty"`
	PrimaryColor string `protobuf:"bytes,5,opt,name=primaryColor" json:"primaryColor,omitempty"`
	AccentColor  string `protobuf:"bytes,6,opt,name=accentColor" json:"accentColor,omitempty"`
}

func (m *DistributorResponse) Reset()                    { *m = DistributorResponse{} }
func (m *DistributorResponse) String() string            { return proto.CompactTextString(m) }
func (*DistributorResponse) ProtoMessage()               {}
func (*DistributorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DistributorResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DistributorResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DistributorResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DistributorResponse) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *DistributorResponse) GetPrimaryColor() string {
	if m != nil {
		return m.PrimaryColor
	}
	return ""
}

func (m *DistributorResponse) GetAccentColor() string {
	if m != nil {
		return m.AccentColor
	}
	return ""
}

func init() {
	proto.RegisterType((*RetrieveRequest)(nil), "distributors.RetrieveRequest")
	proto.RegisterType((*DistributorResponse)(nil), "distributors.DistributorResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Distributors service

type DistributorsClient interface {
	Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*DistributorResponse, error)
}

type distributorsClient struct {
	cc *grpc.ClientConn
}

func NewDistributorsClient(cc *grpc.ClientConn) DistributorsClient {
	return &distributorsClient{cc}
}

func (c *distributorsClient) Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*DistributorResponse, error) {
	out := new(DistributorResponse)
	err := grpc.Invoke(ctx, "/distributors.Distributors/Retrieve", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Distributors service

type DistributorsServer interface {
	Retrieve(context.Context, *RetrieveRequest) (*DistributorResponse, error)
}

func RegisterDistributorsServer(s *grpc.Server, srv DistributorsServer) {
	s.RegisterService(&_Distributors_serviceDesc, srv)
}

func _Distributors_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorsServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distributors.Distributors/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorsServer).Retrieve(ctx, req.(*RetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Distributors_serviceDesc = grpc.ServiceDesc{
	ServiceName: "distributors.Distributors",
	HandlerType: (*DistributorsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Retrieve",
			Handler:    _Distributors_Retrieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "distributors.proto",
}

func init() { proto.RegisterFile("distributors.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0xc7, 0x69, 0x5d, 0x8b, 0x8e, 0x45, 0x61, 0x04, 0x29, 0x82, 0xb0, 0xf6, 0x20, 0xeb, 0xc1,
	0x16, 0xf4, 0x0d, 0xd4, 0xa3, 0xa7, 0x1e, 0xc5, 0x4b, 0xdb, 0x0c, 0xeb, 0xa0, 0xcd, 0xd4, 0x7c,
	0x88, 0x3e, 0x9a, 0x6f, 0x27, 0x66, 0x2b, 0xa6, 0xb2, 0xb7, 0xe4, 0xf7, 0x9f, 0x8f, 0xe4, 0x07,
	0xa8, 0xd8, 0x3a, 0xc3, 0x9d, 0x77, 0x62, 0x6c, 0x35, 0x1a, 0x71, 0x82, 0x79, 0xcc, 0xca, 0x4b,
	0x38, 0x6a, 0xc8, 0x19, 0xa6, 0x77, 0x6a, 0xe8, 0xcd, 0x93, 0x75, 0x78, 0x02, 0x99, 0x92, 0xa1,
	0x65, 0x5d, 0x24, 0xcb, 0x64, 0xb5, 0xdf, 0x4c, 0xb7, 0xf2, 0x2b, 0x81, 0xe3, 0xfb, 0xbf, 0xde,
	0x86, 0xec, 0x28, 0xda, 0x12, 0x1e, 0x42, 0xca, 0x6a, 0xaa, 0x4d, 0x59, 0x21, 0xc2, 0x42, 0xb7,
	0x03, 0x15, 0x69, 0x20, 0xe1, 0x8c, 0x4b, 0x38, 0x50, 0x64, 0x7b, 0xc3, 0xa3, 0x63, 0xd1, 0xc5,
	0x4e, 0x88, 0x62, 0x14, 0x6d, 0x5d, 0xc4, 0x5b, 0xb1, 0x84, 0x7c, 0x34, 0x3c, 0xb4, 0xe6, 0xf3,
	0x4e, 0x5e, 0xc5, 0x14, 0xbb, 0x21, 0x9d, 0xb1, 0x9f, 0xe9, 0x6d, 0xdf, 0x93, 0x76, 0x9b, 0x92,
	0x6c, 0x33, 0x3d, 0x42, 0xd7, 0x4f, 0x90, 0x47, 0x4f, 0xb7, 0xf8, 0x00, 0x7b, 0xbf, 0xdf, 0xc6,
	0xb3, 0x6a, 0x66, 0xe9, 0x9f, 0x8e, 0xd3, 0xf3, 0x79, 0xbc, 0xc5, 0xc0, 0xed, 0xea, 0xf1, 0x62,
	0xcd, 0xee, 0xd9, 0x77, 0x55, 0x2f, 0x43, 0xfd, 0xe2, 0x3b, 0xfa, 0xa8, 0x83, 0xec, 0xab, 0xb5,
	0xd4, 0x71, 0x77, 0x97, 0x05, 0x7c, 0xf3, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x72, 0x94, 0x9f, 0x7b,
	0x99, 0x01, 0x00, 0x00,
}