// Code generated by protoc-gen-go.
// source: imperium.proto
// DO NOT EDIT!

/*
Package imperium is a generated protocol buffer package.

It is generated from these files:
	imperium.proto

It has these top-level messages:
	CertificateRequest
	CertificateResponse
*/
package imperium

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type CertificateRequest struct {
	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
}

func (m *CertificateRequest) Reset()                    { *m = CertificateRequest{} }
func (m *CertificateRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateRequest) ProtoMessage()               {}
func (*CertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CertificateRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type CertificateResponse struct {
	Certificate     string                     `protobuf:"bytes,1,opt,name=certificate" json:"certificate,omitempty"`
	PrivateKey      string                     `protobuf:"bytes,2,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	RootCertificate string                     `protobuf:"bytes,3,opt,name=root_certificate,json=rootCertificate" json:"root_certificate,omitempty"`
	Hostname        string                     `protobuf:"bytes,4,opt,name=hostname" json:"hostname,omitempty"`
	ExpiryTime      *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=expiry_time,json=expiryTime" json:"expiry_time,omitempty"`
}

func (m *CertificateResponse) Reset()                    { *m = CertificateResponse{} }
func (m *CertificateResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateResponse) ProtoMessage()               {}
func (*CertificateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CertificateResponse) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

func (m *CertificateResponse) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *CertificateResponse) GetRootCertificate() string {
	if m != nil {
		return m.RootCertificate
	}
	return ""
}

func (m *CertificateResponse) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *CertificateResponse) GetExpiryTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpiryTime
	}
	return nil
}

func init() {
	proto.RegisterType((*CertificateRequest)(nil), "imperium.CertificateRequest")
	proto.RegisterType((*CertificateResponse)(nil), "imperium.CertificateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Imperium service

type ImperiumClient interface {
	Request(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
}

type imperiumClient struct {
	cc *grpc.ClientConn
}

func NewImperiumClient(cc *grpc.ClientConn) ImperiumClient {
	return &imperiumClient{cc}
}

func (c *imperiumClient) Request(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := grpc.Invoke(ctx, "/imperium.Imperium/Request", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Imperium service

type ImperiumServer interface {
	Request(context.Context, *CertificateRequest) (*CertificateResponse, error)
}

func RegisterImperiumServer(s *grpc.Server, srv ImperiumServer) {
	s.RegisterService(&_Imperium_serviceDesc, srv)
}

func _Imperium_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImperiumServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imperium.Imperium/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImperiumServer).Request(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Imperium_serviceDesc = grpc.ServiceDesc{
	ServiceName: "imperium.Imperium",
	HandlerType: (*ImperiumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _Imperium_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "imperium.proto",
}

func init() { proto.RegisterFile("imperium.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0x80, 0xa9, 0xba, 0x59, 0x5f, 0x41, 0x25, 0x22, 0x94, 0xa2, 0xac, 0x14, 0x0f, 0x13, 0x31,
	0x85, 0x79, 0xf4, 0xa6, 0x20, 0x0c, 0x6f, 0xc5, 0x93, 0x97, 0x92, 0x76, 0x6f, 0x5d, 0x98, 0x59,
	0x62, 0x9a, 0xca, 0xfa, 0xbf, 0xfa, 0xc7, 0x48, 0x7f, 0x64, 0x56, 0x64, 0xc7, 0xf7, 0xf5, 0x7b,
	0x0d, 0xf9, 0x02, 0xa7, 0x5c, 0x28, 0xd4, 0xbc, 0x12, 0x54, 0x69, 0x69, 0x24, 0x71, 0xed, 0x1c,
	0x4c, 0x0a, 0x29, 0x8b, 0x0f, 0x8c, 0x5b, 0x9e, 0x55, 0xcb, 0xd8, 0x70, 0x81, 0xa5, 0x61, 0x42,
	0x75, 0x6a, 0x74, 0x07, 0xe4, 0x19, 0xb5, 0xe1, 0x4b, 0x9e, 0x33, 0x83, 0x09, 0x7e, 0x56, 0x58,
	0x1a, 0x72, 0x09, 0x63, 0xa6, 0x54, 0xca, 0x17, 0xbe, 0x13, 0x3a, 0xd3, 0x93, 0x64, 0xc4, 0x94,
	0x9a, 0x2f, 0xa2, 0x6f, 0x07, 0x2e, 0xfe, 0xd8, 0xa5, 0x92, 0x9b, 0x12, 0x49, 0x08, 0x5e, 0xfe,
	0x8b, 0xfb, 0x9d, 0x21, 0x22, 0x13, 0xf0, 0x94, 0xe6, 0x5f, 0xcc, 0x60, 0xba, 0xc6, 0xda, 0x3f,
	0x68, 0x0d, 0xe8, 0xd1, 0x2b, 0xd6, 0xe4, 0x16, 0xce, 0xb5, 0x94, 0x26, 0x1d, 0xfe, 0xe7, 0xb0,
	0xb5, 0xce, 0x1a, 0x3e, 0x38, 0x95, 0x04, 0xe0, 0xae, 0x64, 0x69, 0x36, 0x4c, 0xa0, 0x7f, 0xd4,
	0x2a, 0xbb, 0x99, 0x3c, 0x82, 0x87, 0x5b, 0xc5, 0x75, 0x9d, 0x36, 0x17, 0xf5, 0x47, 0xa1, 0x33,
	0xf5, 0x66, 0x01, 0xed, 0x2a, 0x50, 0x5b, 0x81, 0xbe, 0xd9, 0x0a, 0x09, 0x74, 0x7a, 0x03, 0x66,
	0x09, 0xb8, 0xf3, 0x3e, 0x1c, 0x79, 0x81, 0x63, 0x1b, 0xe3, 0x8a, 0xee, 0xf2, 0xfe, 0x4f, 0x15,
	0x5c, 0xef, 0xf9, 0xda, 0xa5, 0x79, 0xba, 0x79, 0x8f, 0x0a, 0x6e, 0x56, 0x55, 0x46, 0x73, 0x29,
	0xe2, 0x75, 0x95, 0xe1, 0xb6, 0x7b, 0x8c, 0xfb, 0x42, 0xc6, 0x76, 0x33, 0x1b, 0xb7, 0xe8, 0xe1,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x23, 0xe1, 0xa3, 0xff, 0xc9, 0x01, 0x00, 0x00,
}
