// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consensus/sim/proto/sim.proto

package sim

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

// The request message containing the playerID and the value of the message being fed into the consensus engine
type Request struct {
	PlayerID             int32    `protobuf:"varint,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
	InternalMsgType      uint32   `protobuf:"varint,2,opt,name=internalMsgType,proto3" json:"internalMsgType,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_sim_4f723b2e18ed9d42, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetPlayerID() int32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *Request) GetInternalMsgType() uint32 {
	if m != nil {
		return m.InternalMsgType
	}
	return 0
}

func (m *Request) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// The request message telling the server to initialize the necessary parameters
type InitRequest struct {
	NHonest              int32    `protobuf:"varint,1,opt,name=nHonest,proto3" json:"nHonest,omitempty"`
	NFS                  int32    `protobuf:"varint,2,opt,name=nFS,proto3" json:"nFS,omitempty"`
	NBF                  int32    `protobuf:"varint,3,opt,name=nBF,proto3" json:"nBF,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitRequest) Reset()         { *m = InitRequest{} }
func (m *InitRequest) String() string { return proto.CompactTextString(m) }
func (*InitRequest) ProtoMessage()    {}
func (*InitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sim_4f723b2e18ed9d42, []int{1}
}
func (m *InitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitRequest.Unmarshal(m, b)
}
func (m *InitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitRequest.Marshal(b, m, deterministic)
}
func (dst *InitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitRequest.Merge(dst, src)
}
func (m *InitRequest) XXX_Size() int {
	return xxx_messageInfo_InitRequest.Size(m)
}
func (m *InitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitRequest proto.InternalMessageInfo

func (m *InitRequest) GetNHonest() int32 {
	if m != nil {
		return m.NHonest
	}
	return 0
}

func (m *InitRequest) GetNFS() int32 {
	if m != nil {
		return m.NFS
	}
	return 0
}

func (m *InitRequest) GetNBF() int32 {
	if m != nil {
		return m.NBF
	}
	return 0
}

// The response message returning the output of the consensus engine
type Reply struct {
	MessageType          int32    `protobuf:"varint,1,opt,name=messageType,proto3" json:"messageType,omitempty"`
	InternalMsgType      uint32   `protobuf:"varint,2,opt,name=internalMsgType,proto3" json:"internalMsgType,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_sim_4f723b2e18ed9d42, []int{2}
}
func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (dst *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(dst, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetMessageType() int32 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

func (m *Reply) GetInternalMsgType() uint32 {
	if m != nil {
		return m.InternalMsgType
	}
	return 0
}

func (m *Reply) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Proposal struct {
	PlayerID             int32    `protobuf:"varint,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
	InternalMsgType      uint32   `protobuf:"varint,2,opt,name=internalMsgType,proto3" json:"internalMsgType,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_sim_4f723b2e18ed9d42, []int{3}
}
func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proposal.Unmarshal(m, b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
}
func (dst *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(dst, src)
}
func (m *Proposal) XXX_Size() int {
	return xxx_messageInfo_Proposal.Size(m)
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

func (m *Proposal) GetPlayerID() int32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *Proposal) GetInternalMsgType() uint32 {
	if m != nil {
		return m.InternalMsgType
	}
	return 0
}

func (m *Proposal) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// an empty message
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_sim_4f723b2e18ed9d42, []int{4}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Request)(nil), "sim.Request")
	proto.RegisterType((*InitRequest)(nil), "sim.InitRequest")
	proto.RegisterType((*Reply)(nil), "sim.Reply")
	proto.RegisterType((*Proposal)(nil), "sim.Proposal")
	proto.RegisterType((*Empty)(nil), "sim.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SimulatorClient is the client API for Simulator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SimulatorClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (Simulator_PingClient, error)
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (Simulator_InitClient, error)
	Exit(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type simulatorClient struct {
	cc *grpc.ClientConn
}

func NewSimulatorClient(cc *grpc.ClientConn) SimulatorClient {
	return &simulatorClient{cc}
}

func (c *simulatorClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (Simulator_PingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Simulator_serviceDesc.Streams[0], "/sim.Simulator/Ping", opts...)
	if err != nil {
		return nil, err
	}
	x := &simulatorPingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Simulator_PingClient interface {
	Recv() (*Reply, error)
	grpc.ClientStream
}

type simulatorPingClient struct {
	grpc.ClientStream
}

func (x *simulatorPingClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *simulatorClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (Simulator_InitClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Simulator_serviceDesc.Streams[1], "/sim.Simulator/Init", opts...)
	if err != nil {
		return nil, err
	}
	x := &simulatorInitClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Simulator_InitClient interface {
	Recv() (*Proposal, error)
	grpc.ClientStream
}

type simulatorInitClient struct {
	grpc.ClientStream
}

func (x *simulatorInitClient) Recv() (*Proposal, error) {
	m := new(Proposal)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *simulatorClient) Exit(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/sim.Simulator/Exit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimulatorServer is the server API for Simulator service.
type SimulatorServer interface {
	Ping(*Request, Simulator_PingServer) error
	Init(*InitRequest, Simulator_InitServer) error
	Exit(context.Context, *Empty) (*Empty, error)
}

func RegisterSimulatorServer(s *grpc.Server, srv SimulatorServer) {
	s.RegisterService(&_Simulator_serviceDesc, srv)
}

func _Simulator_Ping_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SimulatorServer).Ping(m, &simulatorPingServer{stream})
}

type Simulator_PingServer interface {
	Send(*Reply) error
	grpc.ServerStream
}

type simulatorPingServer struct {
	grpc.ServerStream
}

func (x *simulatorPingServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

func _Simulator_Init_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InitRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SimulatorServer).Init(m, &simulatorInitServer{stream})
}

type Simulator_InitServer interface {
	Send(*Proposal) error
	grpc.ServerStream
}

type simulatorInitServer struct {
	grpc.ServerStream
}

func (x *simulatorInitServer) Send(m *Proposal) error {
	return x.ServerStream.SendMsg(m)
}

func _Simulator_Exit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimulatorServer).Exit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sim.Simulator/Exit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimulatorServer).Exit(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Simulator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sim.Simulator",
	HandlerType: (*SimulatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exit",
			Handler:    _Simulator_Exit_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Ping",
			Handler:       _Simulator_Ping_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Init",
			Handler:       _Simulator_Init_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "consensus/sim/proto/sim.proto",
}

func init() { proto.RegisterFile("consensus/sim/proto/sim.proto", fileDescriptor_sim_4f723b2e18ed9d42) }

var fileDescriptor_sim_4f723b2e18ed9d42 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0x17, 0xb7, 0xba, 0xed, 0xce, 0xe1, 0x08, 0x3e, 0x94, 0x82, 0x50, 0x8a, 0x0f, 0x05,
	0x61, 0x13, 0xfd, 0x07, 0xc3, 0x15, 0x87, 0x08, 0x23, 0xf3, 0x0f, 0x44, 0xb9, 0x96, 0x40, 0x9a,
	0xc4, 0x26, 0x15, 0xfb, 0xe6, 0x4f, 0x97, 0xa6, 0xad, 0x14, 0x5f, 0xc5, 0xb7, 0x73, 0x3e, 0xc2,
	0xb9, 0x37, 0x27, 0x81, 0xcb, 0x57, 0xad, 0x2c, 0x2a, 0x5b, 0xd9, 0x8d, 0x15, 0xc5, 0xc6, 0x94,
	0xda, 0xe9, 0x46, 0xad, 0xbd, 0xa2, 0x63, 0x2b, 0x8a, 0x04, 0x61, 0xca, 0xf0, 0xbd, 0x42, 0xeb,
	0x68, 0x04, 0x33, 0x23, 0x79, 0x8d, 0xe5, 0xfe, 0x3e, 0x24, 0x31, 0x49, 0x03, 0xf6, 0xe3, 0x69,
	0x0a, 0xe7, 0x42, 0x39, 0x2c, 0x15, 0x97, 0x4f, 0x36, 0x7f, 0xae, 0x0d, 0x86, 0x27, 0x31, 0x49,
	0x97, 0xec, 0x37, 0xa6, 0x17, 0x10, 0x7c, 0x70, 0x59, 0x61, 0x38, 0x8e, 0x49, 0x3a, 0x67, 0xad,
	0x49, 0x1e, 0x61, 0xb1, 0x57, 0xc2, 0xf5, 0xa3, 0x42, 0x98, 0xaa, 0x07, 0xad, 0xd0, 0xba, 0x6e,
	0x52, 0x6f, 0xe9, 0x0a, 0xc6, 0x2a, 0x3b, 0xfa, 0xf0, 0x80, 0x35, 0xd2, 0x93, 0x6d, 0xe6, 0xe3,
	0x1a, 0xb2, 0xcd, 0x12, 0x01, 0x01, 0x43, 0x23, 0x6b, 0x1a, 0xc3, 0xa2, 0x40, 0x6b, 0x79, 0x8e,
	0x7e, 0xa3, 0x36, 0x6a, 0x88, 0xfe, 0xbc, 0xf7, 0x1b, 0xcc, 0x0e, 0xa5, 0x36, 0xda, 0x72, 0xf9,
	0xaf, 0xfd, 0x4c, 0x21, 0xd8, 0x15, 0xc6, 0xd5, 0xb7, 0x5f, 0x04, 0xe6, 0x47, 0x51, 0x54, 0x92,
	0x3b, 0x5d, 0xd2, 0x2b, 0x98, 0x1c, 0x84, 0xca, 0xe9, 0xd9, 0xba, 0x79, 0xb6, 0xae, 0xbd, 0x08,
	0x3a, 0x67, 0x64, 0x9d, 0x8c, 0x6e, 0x08, 0xbd, 0x86, 0x49, 0x53, 0x2e, 0x5d, 0x79, 0x3e, 0xe8,
	0x39, 0x5a, 0x7a, 0xd2, 0xdf, 0xc0, 0x1f, 0x8e, 0x61, 0xb2, 0xfb, 0x14, 0x8e, 0xb6, 0x21, 0x7e,
	0x68, 0x34, 0xd0, 0xc9, 0xe8, 0xe5, 0xd4, 0x7f, 0x8f, 0xbb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x77, 0x0f, 0xa4, 0xf5, 0x3f, 0x02, 0x00, 0x00,
}
