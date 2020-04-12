// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package calendar

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type EventMessage struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EventMessage) Reset()         { *m = EventMessage{} }
func (m *EventMessage) String() string { return proto.CompactTextString(m) }
func (*EventMessage) ProtoMessage()    {}
func (*EventMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

func (m *EventMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventMessage.Unmarshal(m, b)
}
func (m *EventMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventMessage.Marshal(b, m, deterministic)
}
func (m *EventMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMessage.Merge(m, src)
}
func (m *EventMessage) XXX_Size() int {
	return xxx_messageInfo_EventMessage.Size(m)
}
func (m *EventMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EventMessage proto.InternalMessageInfo

func (m *EventMessage) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EventMessage) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *EventMessage) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EventMessage) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func init() {
	proto.RegisterType((*EventMessage)(nil), "calendar.EventMessage")
}

func init() {
	proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e)
}

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x8f, 0xbd, 0x6e, 0xc3, 0x20,
	0x14, 0x85, 0xe5, 0x9f, 0x56, 0x2d, 0x54, 0x1d, 0x50, 0x55, 0x21, 0x2f, 0xb5, 0x3a, 0x75, 0xc2,
	0x92, 0x3b, 0x77, 0xec, 0xe8, 0xc5, 0xc9, 0x0b, 0x60, 0x73, 0x63, 0x21, 0xd9, 0xc6, 0x82, 0x1b,
	0x3f, 0x42, 0x9e, 0x3b, 0x04, 0x07, 0xc9, 0x4b, 0x36, 0xf8, 0xce, 0x41, 0x7c, 0x87, 0x50, 0x58,
	0x61, 0x46, 0xb1, 0x58, 0x83, 0x86, 0xbd, 0xf4, 0x72, 0x84, 0x59, 0x49, 0x5b, 0x7c, 0x0d, 0xc6,
	0x0c, 0x23, 0x54, 0x81, 0x77, 0xe7, 0x53, 0x85, 0x7a, 0x02, 0x87, 0x72, 0x5a, 0xb6, 0xea, 0xf7,
	0x25, 0x21, 0x6f, 0xff, 0xb7, 0xa7, 0x0d, 0x38, 0x27, 0x07, 0x60, 0xef, 0x24, 0xd5, 0x8a, 0x27,
	0x65, 0xf2, 0x93, 0xb5, 0xfe, 0xc4, 0x3e, 0xc8, 0x13, 0x6a, 0x1c, 0x81, 0xa7, 0x1e, 0xbd, 0xb6,
	0xdb, 0x85, 0x95, 0x84, 0x2a, 0x70, 0xbd, 0xd5, 0x0b, 0x6a, 0x33, 0xf3, 0x2c, 0x64, 0x7b, 0xc4,
	0x04, 0xc9, 0x95, 0x44, 0xe0, 0xb9, 0x8f, 0x68, 0x5d, 0x88, 0x4d, 0x44, 0x44, 0x11, 0x71, 0x8c,
	0x22, 0x6d, 0xe8, 0xd5, 0xcd, 0xdd, 0xe3, 0x00, 0x76, 0xd5, 0x3d, 0xb0, 0x3f, 0x42, 0x9d, 0xdf,
	0x10, 0xb5, 0x3e, 0x45, 0xdc, 0x24, 0xf6, 0xba, 0xc5, 0x03, 0xde, 0x3d, 0x87, 0x8f, 0x7e, 0xaf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x03, 0x7e, 0xb9, 0x01, 0x18, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventServiceClient interface {
	SendMessage(ctx context.Context, in *EventMessage, opts ...grpc.CallOption) (*EventMessage, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) SendMessage(ctx context.Context, in *EventMessage, opts ...grpc.CallOption) (*EventMessage, error) {
	out := new(EventMessage)
	err := c.cc.Invoke(ctx, "/calendar.EventService/sendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
type EventServiceServer interface {
	SendMessage(context.Context, *EventMessage) (*EventMessage, error)
}

// UnimplementedEventServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (*UnimplementedEventServiceServer) SendMessage(ctx context.Context, req *EventMessage) (*EventMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.EventService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).SendMessage(ctx, req.(*EventMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calendar.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendMessage",
			Handler:    _EventService_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event.proto",
}
