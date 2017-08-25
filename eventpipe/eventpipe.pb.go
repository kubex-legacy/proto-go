// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eventpipe.proto

/*
Package eventpipe is a generated protocol buffer package.

It is generated from these files:
	eventpipe.proto

It has these top-level messages:
	Event
	CreateEventResponse
*/
package eventpipe

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

type Event struct {
	DistributorId string                     `protobuf:"bytes,1,opt,name=distributor_id,json=distributorId" json:"distributor_id,omitempty"`
	ProjectId     string                     `protobuf:"bytes,2,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	VendorId      string                     `protobuf:"bytes,3,opt,name=vendor_id,json=vendorId" json:"vendor_id,omitempty"`
	AppId         string                     `protobuf:"bytes,4,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	EventId       string                     `protobuf:"bytes,5,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	Occurred      *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=occurred" json:"occurred,omitempty"`
	Attributes    []*Event_Attribute         `protobuf:"bytes,7,rep,name=attributes" json:"attributes,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetDistributorId() string {
	if m != nil {
		return m.DistributorId
	}
	return ""
}

func (m *Event) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *Event) GetVendorId() string {
	if m != nil {
		return m.VendorId
	}
	return ""
}

func (m *Event) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *Event) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Event) GetOccurred() *google_protobuf.Timestamp {
	if m != nil {
		return m.Occurred
	}
	return nil
}

func (m *Event) GetAttributes() []*Event_Attribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type Event_Attribute struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Event_Attribute) Reset()                    { *m = Event_Attribute{} }
func (m *Event_Attribute) String() string            { return proto.CompactTextString(m) }
func (*Event_Attribute) ProtoMessage()               {}
func (*Event_Attribute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Event_Attribute) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event_Attribute) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CreateEventResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *CreateEventResponse) Reset()                    { *m = CreateEventResponse{} }
func (m *CreateEventResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateEventResponse) ProtoMessage()               {}
func (*CreateEventResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateEventResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Event)(nil), "eventpipe.Event")
	proto.RegisterType((*Event_Attribute)(nil), "eventpipe.Event.Attribute")
	proto.RegisterType((*CreateEventResponse)(nil), "eventpipe.CreateEventResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EventPipe service

type EventPipeClient interface {
	Write(ctx context.Context, in *Event, opts ...grpc.CallOption) (*CreateEventResponse, error)
}

type eventPipeClient struct {
	cc *grpc.ClientConn
}

func NewEventPipeClient(cc *grpc.ClientConn) EventPipeClient {
	return &eventPipeClient{cc}
}

func (c *eventPipeClient) Write(ctx context.Context, in *Event, opts ...grpc.CallOption) (*CreateEventResponse, error) {
	out := new(CreateEventResponse)
	err := grpc.Invoke(ctx, "/eventpipe.EventPipe/Write", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EventPipe service

type EventPipeServer interface {
	Write(context.Context, *Event) (*CreateEventResponse, error)
}

func RegisterEventPipeServer(s *grpc.Server, srv EventPipeServer) {
	s.RegisterService(&_EventPipe_serviceDesc, srv)
}

func _EventPipe_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventPipeServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventpipe.EventPipe/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventPipeServer).Write(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventPipe_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eventpipe.EventPipe",
	HandlerType: (*EventPipeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _EventPipe_Write_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eventpipe.proto",
}

func init() { proto.RegisterFile("eventpipe.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xd1, 0x4b, 0xc3, 0x30,
	0x10, 0xc6, 0x59, 0x67, 0xb7, 0xf6, 0xc4, 0x29, 0x51, 0xa1, 0x56, 0xd4, 0x31, 0x19, 0xec, 0xc5,
	0x16, 0x27, 0x08, 0xfa, 0xa6, 0xa2, 0xd0, 0x37, 0x29, 0x82, 0xe0, 0x8b, 0xb4, 0xcd, 0x59, 0xa3,
	0xdb, 0x12, 0x92, 0x74, 0xf8, 0xff, 0xf9, 0x8f, 0x49, 0x93, 0xb5, 0x0e, 0xf1, 0xf1, 0xbe, 0xef,
	0x77, 0x97, 0xfb, 0x2e, 0xb0, 0x8d, 0x4b, 0x5c, 0x68, 0xc1, 0x04, 0x46, 0x42, 0x72, 0xcd, 0x89,
	0xdf, 0x0a, 0xe1, 0x49, 0xc9, 0x79, 0x39, 0xc3, 0xd8, 0x18, 0x79, 0xf5, 0x16, 0x6b, 0x36, 0x47,
	0xa5, 0xb3, 0xb9, 0xb0, 0xec, 0xe8, 0xdb, 0x01, 0xf7, 0xbe, 0xc6, 0xc9, 0x18, 0x06, 0x94, 0x29,
	0x2d, 0x59, 0x5e, 0x69, 0x2e, 0x5f, 0x19, 0x0d, 0x3a, 0xc3, 0xce, 0xc4, 0x4f, 0xb7, 0xd6, 0xd4,
	0x84, 0x92, 0x23, 0x00, 0x21, 0xf9, 0x07, 0x16, 0xba, 0x46, 0x1c, 0x83, 0xf8, 0x2b, 0x25, 0xa1,
	0xe4, 0x10, 0xfc, 0x25, 0x2e, 0xa8, 0x1d, 0xd0, 0x35, 0xae, 0x67, 0x85, 0x84, 0x92, 0x7d, 0xe8,
	0x65, 0x42, 0xd4, 0xce, 0x86, 0x71, 0xdc, 0x4c, 0x88, 0x84, 0x92, 0x03, 0xf0, 0xcc, 0xc6, 0xb5,
	0xe1, 0x1a, 0xa3, 0x6f, 0xea, 0x84, 0x92, 0x4b, 0xf0, 0x78, 0x51, 0x54, 0x52, 0x22, 0x0d, 0x7a,
	0xc3, 0xce, 0x64, 0x73, 0x1a, 0x46, 0x36, 0x52, 0xd4, 0x44, 0x8a, 0x9e, 0x9a, 0x48, 0x69, 0xcb,
	0x92, 0x6b, 0x80, 0x4c, 0xdb, 0xad, 0x51, 0x05, 0xfd, 0x61, 0xd7, 0x74, 0xfe, 0x1e, 0xca, 0x44,
	0x8e, 0x6e, 0x1a, 0x24, 0x5d, 0xa3, 0xc3, 0x73, 0xf0, 0x5b, 0x83, 0x0c, 0xc0, 0x69, 0x2f, 0xe1,
	0x30, 0x4a, 0xf6, 0xc0, 0x5d, 0x66, 0xb3, 0x0a, 0x57, 0xc9, 0x6d, 0x31, 0x8a, 0x61, 0xf7, 0x4e,
	0x62, 0xa6, 0xd1, 0xcc, 0x4d, 0x51, 0x09, 0xbe, 0x50, 0x48, 0x02, 0xe8, 0xab, 0xaa, 0x28, 0x50,
	0x29, 0x33, 0xc1, 0x4b, 0x9b, 0x72, 0xfa, 0x00, 0xbe, 0x41, 0x1f, 0x99, 0x40, 0x72, 0x05, 0xee,
	0xb3, 0x64, 0x1a, 0xc9, 0xce, 0xdf, 0x0d, 0xc3, 0xe3, 0x35, 0xe5, 0x9f, 0x17, 0x6e, 0xc7, 0x2f,
	0xa7, 0x25, 0xd3, 0xef, 0x55, 0x1e, 0x15, 0x7c, 0x1e, 0x7f, 0x56, 0x39, 0x7e, 0xd9, 0xbf, 0x3e,
	0x2b, 0x79, 0xdc, 0xb6, 0xe6, 0x3d, 0xa3, 0x5d, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x83, 0xfb,
	0x2a, 0x59, 0x2b, 0x02, 0x00, 0x00,
}