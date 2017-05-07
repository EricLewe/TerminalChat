// Code generated by protoc-gen-go.
// source: Wat.proto
// DO NOT EDIT!

/*
Package wat is a generated protocol buffer package.

It is generated from these files:
	Wat.proto

It has these top-level messages:
	LoginRequest
	RegisterRequest
	LoginReply
	Request
	ChatMessageReply
	ConversationRequest
	ConversationReply
*/
package wat

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

type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginReply struct {
	Username        string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	MessageOfTheDay string `protobuf:"bytes,2,opt,name=messageOfTheDay" json:"messageOfTheDay,omitempty"`
}

func (m *LoginReply) Reset()                    { *m = LoginReply{} }
func (m *LoginReply) String() string            { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()               {}
func (*LoginReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LoginReply) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginReply) GetMessageOfTheDay() string {
	if m != nil {
		return m.MessageOfTheDay
	}
	return ""
}

type Request struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Request) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type ChatMessageReply struct {
	ConversationId int32                      `protobuf:"varint,1,opt,name=conversationId" json:"conversationId,omitempty"`
	Content        string                     `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
	Sent           *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=sent" json:"sent,omitempty"`
	SentByUser     string                     `protobuf:"bytes,4,opt,name=sentByUser" json:"sentByUser,omitempty"`
}

func (m *ChatMessageReply) Reset()                    { *m = ChatMessageReply{} }
func (m *ChatMessageReply) String() string            { return proto.CompactTextString(m) }
func (*ChatMessageReply) ProtoMessage()               {}
func (*ChatMessageReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ChatMessageReply) GetConversationId() int32 {
	if m != nil {
		return m.ConversationId
	}
	return 0
}

func (m *ChatMessageReply) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ChatMessageReply) GetSent() *google_protobuf.Timestamp {
	if m != nil {
		return m.Sent
	}
	return nil
}

func (m *ChatMessageReply) GetSentByUser() string {
	if m != nil {
		return m.SentByUser
	}
	return ""
}

type ConversationRequest struct {
	Id      int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Request *Request `protobuf:"bytes,2,opt,name=request" json:"request,omitempty"`
}

func (m *ConversationRequest) Reset()                    { *m = ConversationRequest{} }
func (m *ConversationRequest) String() string            { return proto.CompactTextString(m) }
func (*ConversationRequest) ProtoMessage()               {}
func (*ConversationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ConversationRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ConversationRequest) GetRequest() *Request {
	if m != nil {
		return m.Request
	}
	return nil
}

type ConversationReply struct {
	Id            int32                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	FromDate      *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=fromDate" json:"fromDate,omitempty"`
	Name          string                     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	LatestMessage *ChatMessageReply          `protobuf:"bytes,4,opt,name=latestMessage" json:"latestMessage,omitempty"`
}

func (m *ConversationReply) Reset()                    { *m = ConversationReply{} }
func (m *ConversationReply) String() string            { return proto.CompactTextString(m) }
func (*ConversationReply) ProtoMessage()               {}
func (*ConversationReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ConversationReply) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ConversationReply) GetFromDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.FromDate
	}
	return nil
}

func (m *ConversationReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ConversationReply) GetLatestMessage() *ChatMessageReply {
	if m != nil {
		return m.LatestMessage
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "wat.LoginRequest")
	proto.RegisterType((*RegisterRequest)(nil), "wat.RegisterRequest")
	proto.RegisterType((*LoginReply)(nil), "wat.LoginReply")
	proto.RegisterType((*Request)(nil), "wat.Request")
	proto.RegisterType((*ChatMessageReply)(nil), "wat.ChatMessageReply")
	proto.RegisterType((*ConversationRequest)(nil), "wat.ConversationRequest")
	proto.RegisterType((*ConversationReply)(nil), "wat.ConversationReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Chat service

type ChatClient interface {
	VerifyLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	SendMessage(ctx context.Context, in *ChatMessageReply, opts ...grpc.CallOption) (*Request, error)
	RouteConversation(ctx context.Context, in *Request, opts ...grpc.CallOption) (Chat_RouteConversationClient, error)
	RouteChat(ctx context.Context, in *ConversationRequest, opts ...grpc.CallOption) (Chat_RouteChatClient, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) VerifyLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/wat.Chat/verifyLogin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) SendMessage(ctx context.Context, in *ChatMessageReply, opts ...grpc.CallOption) (*Request, error) {
	out := new(Request)
	err := grpc.Invoke(ctx, "/wat.Chat/sendMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) RouteConversation(ctx context.Context, in *Request, opts ...grpc.CallOption) (Chat_RouteConversationClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Chat_serviceDesc.Streams[0], c.cc, "/wat.Chat/RouteConversation", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatRouteConversationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chat_RouteConversationClient interface {
	Recv() (*ConversationReply, error)
	grpc.ClientStream
}

type chatRouteConversationClient struct {
	grpc.ClientStream
}

func (x *chatRouteConversationClient) Recv() (*ConversationReply, error) {
	m := new(ConversationReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) RouteChat(ctx context.Context, in *ConversationRequest, opts ...grpc.CallOption) (Chat_RouteChatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Chat_serviceDesc.Streams[1], c.cc, "/wat.Chat/RouteChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatRouteChatClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chat_RouteChatClient interface {
	Recv() (*ChatMessageReply, error)
	grpc.ClientStream
}

type chatRouteChatClient struct {
	grpc.ClientStream
}

func (x *chatRouteChatClient) Recv() (*ChatMessageReply, error) {
	m := new(ChatMessageReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Chat service

type ChatServer interface {
	VerifyLogin(context.Context, *LoginRequest) (*LoginReply, error)
	SendMessage(context.Context, *ChatMessageReply) (*Request, error)
	RouteConversation(*Request, Chat_RouteConversationServer) error
	RouteChat(*ConversationRequest, Chat_RouteChatServer) error
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_VerifyLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).VerifyLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wat.Chat/VerifyLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).VerifyLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMessageReply)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wat.Chat/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SendMessage(ctx, req.(*ChatMessageReply))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_RouteConversation_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServer).RouteConversation(m, &chatRouteConversationServer{stream})
}

type Chat_RouteConversationServer interface {
	Send(*ConversationReply) error
	grpc.ServerStream
}

type chatRouteConversationServer struct {
	grpc.ServerStream
}

func (x *chatRouteConversationServer) Send(m *ConversationReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Chat_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConversationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServer).RouteChat(m, &chatRouteChatServer{stream})
}

type Chat_RouteChatServer interface {
	Send(*ChatMessageReply) error
	grpc.ServerStream
}

type chatRouteChatServer struct {
	grpc.ServerStream
}

func (x *chatRouteChatServer) Send(m *ChatMessageReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "verifyLogin",
			Handler:    _Chat_VerifyLogin_Handler,
		},
		{
			MethodName: "sendMessage",
			Handler:    _Chat_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RouteConversation",
			Handler:       _Chat_RouteConversation_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RouteChat",
			Handler:       _Chat_RouteChat_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "Wat.proto",
}

func init() { proto.RegisterFile("Wat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xd1, 0x6a, 0xdb, 0x30,
	0x14, 0x8d, 0x93, 0x6c, 0x69, 0x6e, 0xba, 0x66, 0xd1, 0xd8, 0x30, 0x1e, 0x6c, 0xc5, 0xb0, 0x92,
	0x27, 0xb5, 0xa4, 0x63, 0x2f, 0x63, 0x30, 0xd2, 0x32, 0x56, 0x58, 0x59, 0x11, 0x1d, 0x7d, 0x56,
	0xe3, 0x1b, 0xc7, 0x60, 0x5b, 0x9e, 0xa4, 0x34, 0xf3, 0xff, 0xec, 0x71, 0x5f, 0xb7, 0x2f, 0x18,
	0x92, 0xec, 0xe2, 0x98, 0x50, 0x06, 0x7d, 0xb2, 0x74, 0xee, 0xb9, 0xc7, 0xe7, 0xde, 0x23, 0x18,
	0xde, 0x70, 0x4d, 0x0b, 0x29, 0xb4, 0x20, 0xbd, 0x0d, 0xd7, 0xc1, 0xdb, 0x58, 0x88, 0x38, 0xc5,
	0x63, 0x0b, 0xdd, 0xae, 0x97, 0xc7, 0x3a, 0xc9, 0x50, 0x69, 0x9e, 0x15, 0x8e, 0x15, 0x7e, 0x81,
	0xfd, 0x6f, 0x22, 0x4e, 0x72, 0x86, 0x3f, 0xd7, 0xa8, 0x34, 0x09, 0x60, 0x6f, 0xad, 0x50, 0xe6,
	0x3c, 0x43, 0xdf, 0x3b, 0xf4, 0xa6, 0x43, 0x76, 0x7f, 0x37, 0xb5, 0x82, 0x2b, 0xb5, 0x11, 0x32,
	0xf2, 0xbb, 0xae, 0x56, 0xdf, 0xc3, 0x0b, 0x18, 0x33, 0x8c, 0x13, 0xa5, 0x51, 0x3e, 0x56, 0x8a,
	0x01, 0x54, 0x96, 0x8a, 0xb4, 0x7c, 0x50, 0x65, 0x0a, 0xe3, 0x0c, 0x95, 0xe2, 0x31, 0x7e, 0x5f,
	0x5e, 0xaf, 0xf0, 0x9c, 0x97, 0x95, 0x58, 0x1b, 0x0e, 0xdf, 0xc1, 0xe0, 0x3f, 0x6c, 0x85, 0xbf,
	0x3d, 0x78, 0x7e, 0xb6, 0xe2, 0xfa, 0xd2, 0xb5, 0x3b, 0x07, 0x47, 0x70, 0xb0, 0x10, 0xf9, 0x1d,
	0x4a, 0xc5, 0x75, 0x22, 0xf2, 0x8b, 0xc8, 0xb6, 0x3d, 0x61, 0x2d, 0x94, 0xf8, 0x30, 0x58, 0x88,
	0x5c, 0x63, 0xae, 0x2b, 0x17, 0xf5, 0x95, 0x50, 0xe8, 0x2b, 0x03, 0xf7, 0x0e, 0xbd, 0xe9, 0x68,
	0x16, 0x50, 0x17, 0x0a, 0xad, 0x43, 0xa1, 0xd7, 0x75, 0x28, 0xcc, 0xf2, 0xc8, 0x1b, 0x00, 0xf3,
	0x9d, 0x97, 0x3f, 0x14, 0x4a, 0xbf, 0x6f, 0xc5, 0x1a, 0x48, 0x78, 0x09, 0x2f, 0xce, 0x1a, 0xff,
	0xae, 0x27, 0x3b, 0x80, 0x6e, 0x52, 0x9b, 0xeb, 0x26, 0x11, 0x39, 0x82, 0x81, 0x74, 0x25, 0x6b,
	0x68, 0x34, 0xdb, 0xa7, 0x1b, 0xae, 0x69, 0x45, 0x67, 0x75, 0x31, 0xfc, 0xe3, 0xc1, 0x64, 0x5b,
	0xcf, 0x8c, 0xdd, 0x56, 0xfb, 0x00, 0x7b, 0x4b, 0x29, 0xb2, 0x73, 0xae, 0xb1, 0x92, 0x7b, 0x68,
	0x90, 0x7b, 0x2e, 0x21, 0xd0, 0xb7, 0xbb, 0xee, 0xd9, 0x31, 0xec, 0x99, 0x7c, 0x84, 0x67, 0x29,
	0xd7, 0xa8, 0xea, 0x45, 0xdb, 0x19, 0x47, 0xb3, 0x97, 0xd6, 0x5f, 0x3b, 0x00, 0xb6, 0xcd, 0x9d,
	0xfd, 0xf5, 0xa0, 0x6f, 0x38, 0xe4, 0x14, 0x46, 0x77, 0x28, 0x93, 0x65, 0x69, 0x9f, 0x0b, 0x99,
	0xd8, 0xee, 0xe6, 0x6b, 0x0e, 0xc6, 0x4d, 0xa8, 0x48, 0xcb, 0xb0, 0x43, 0xde, 0xc3, 0x48, 0x61,
	0x1e, 0x55, 0x62, 0x64, 0xf7, 0x2f, 0x83, 0xad, 0x4d, 0x85, 0x1d, 0xf2, 0x09, 0x26, 0x4c, 0xac,
	0x35, 0x36, 0xd7, 0x44, 0xb6, 0x48, 0xc1, 0x2b, 0xa7, 0xd4, 0xde, 0x63, 0xd8, 0x39, 0xf1, 0xc8,
	0x67, 0x18, 0xba, 0x76, 0x63, 0xdb, 0xdf, 0x41, 0x74, 0x12, 0xbb, 0xcd, 0x18, 0x85, 0xf9, 0x09,
	0xbc, 0x4e, 0x04, 0x8d, 0x65, 0xb1, 0xa0, 0xf8, 0x8b, 0x67, 0x45, 0x8a, 0x8a, 0xae, 0x30, 0x4d,
	0xc5, 0x46, 0xc8, 0x34, 0x9a, 0x8f, 0xbf, 0x9a, 0xf3, 0x8d, 0x39, 0x5f, 0x99, 0x30, 0xae, 0xbc,
	0xdb, 0xa7, 0x36, 0x95, 0xd3, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x15, 0xbd, 0x96, 0xac, 0x13,
	0x04, 0x00, 0x00,
}
