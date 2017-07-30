// Code generated by protoc-gen-go.
// source: nebulapb.proto
// DO NOT EDIT!

/*
Package nebulapb is a generated protocol buffer package.

It is generated from these files:
	nebulapb.proto

It has these top-level messages:
	Empty
	StreamRequest
	EntryStreamResponse
	QuitStreamRequest
	ServerEntry
	ServerStatus
	GetServerEntryRequest
	GetServerEntryResponse
	AddServerEntryRequest
	AddServerEntryResponse
	RemoveServerEntryRequest
	RemoveServerEntryResponse
	PushStatusRequest
	PushStatusResponse
	FetchStatusRequest
	FetchStatusResponse
*/
package nebulapb

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

//
// STREAM
//
type StreamType int32

const (
	StreamType_QUIT     StreamType = 0
	StreamType_CONNECT  StreamType = 1
	StreamType_RESTORED StreamType = 2
	StreamType_DISPATCH StreamType = 3
	StreamType_SYNC     StreamType = 4
)

var StreamType_name = map[int32]string{
	0: "QUIT",
	1: "CONNECT",
	2: "RESTORED",
	3: "DISPATCH",
	4: "SYNC",
}
var StreamType_value = map[string]int32{
	"QUIT":     0,
	"CONNECT":  1,
	"RESTORED": 2,
	"DISPATCH": 3,
	"SYNC":     4,
}

func (x StreamType) String() string {
	return proto.EnumName(StreamType_name, int32(x))
}
func (StreamType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StreamRequest struct {
	Name string     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type StreamType `protobuf:"varint,2,opt,name=type,enum=nebulapb.StreamType" json:"type,omitempty"`
}

func (m *StreamRequest) Reset()                    { *m = StreamRequest{} }
func (m *StreamRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()               {}
func (*StreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StreamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StreamRequest) GetType() StreamType {
	if m != nil {
		return m.Type
	}
	return StreamType_QUIT
}

type EntryStreamResponse struct {
	Type   StreamType `protobuf:"varint,1,opt,name=type,enum=nebulapb.StreamType" json:"type,omitempty"`
	Target string     `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
}

func (m *EntryStreamResponse) Reset()                    { *m = EntryStreamResponse{} }
func (m *EntryStreamResponse) String() string            { return proto.CompactTextString(m) }
func (*EntryStreamResponse) ProtoMessage()               {}
func (*EntryStreamResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EntryStreamResponse) GetType() StreamType {
	if m != nil {
		return m.Type
	}
	return StreamType_QUIT
}

func (m *EntryStreamResponse) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

type QuitStreamRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *QuitStreamRequest) Reset()                    { *m = QuitStreamRequest{} }
func (m *QuitStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*QuitStreamRequest) ProtoMessage()               {}
func (*QuitStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *QuitStreamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// --
// Server Entry
// --
type ServerEntry struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=displayName" json:"displayName,omitempty"`
	Address     string `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	Port        int32  `protobuf:"varint,4,opt,name=port" json:"port,omitempty"`
	Motd        string `protobuf:"bytes,5,opt,name=motd" json:"motd,omitempty"`
}

func (m *ServerEntry) Reset()                    { *m = ServerEntry{} }
func (m *ServerEntry) String() string            { return proto.CompactTextString(m) }
func (*ServerEntry) ProtoMessage()               {}
func (*ServerEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ServerEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServerEntry) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ServerEntry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ServerEntry) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ServerEntry) GetMotd() string {
	if m != nil {
		return m.Motd
	}
	return ""
}

type ServerStatus struct {
	Online        string `protobuf:"bytes,1,opt,name=online" json:"online,omitempty"`
	OnlinePlayers int32  `protobuf:"varint,2,opt,name=onlinePlayers" json:"onlinePlayers,omitempty"`
	MaxPlayers    int32  `protobuf:"varint,3,opt,name=maxPlayers" json:"maxPlayers,omitempty"`
}

func (m *ServerStatus) Reset()                    { *m = ServerStatus{} }
func (m *ServerStatus) String() string            { return proto.CompactTextString(m) }
func (*ServerStatus) ProtoMessage()               {}
func (*ServerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ServerStatus) GetOnline() string {
	if m != nil {
		return m.Online
	}
	return ""
}

func (m *ServerStatus) GetOnlinePlayers() int32 {
	if m != nil {
		return m.OnlinePlayers
	}
	return 0
}

func (m *ServerStatus) GetMaxPlayers() int32 {
	if m != nil {
		return m.MaxPlayers
	}
	return 0
}

//
// ServerEntry
//
type GetServerEntryRequest struct {
}

func (m *GetServerEntryRequest) Reset()                    { *m = GetServerEntryRequest{} }
func (m *GetServerEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetServerEntryRequest) ProtoMessage()               {}
func (*GetServerEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type GetServerEntryResponse struct {
	Entry []*ServerEntry `protobuf:"bytes,1,rep,name=entry" json:"entry,omitempty"`
}

func (m *GetServerEntryResponse) Reset()                    { *m = GetServerEntryResponse{} }
func (m *GetServerEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*GetServerEntryResponse) ProtoMessage()               {}
func (*GetServerEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetServerEntryResponse) GetEntry() []*ServerEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type AddServerEntryRequest struct {
	Entry *ServerEntry `protobuf:"bytes,1,opt,name=entry" json:"entry,omitempty"`
}

func (m *AddServerEntryRequest) Reset()                    { *m = AddServerEntryRequest{} }
func (m *AddServerEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*AddServerEntryRequest) ProtoMessage()               {}
func (*AddServerEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AddServerEntryRequest) GetEntry() *ServerEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type AddServerEntryResponse struct {
}

func (m *AddServerEntryResponse) Reset()                    { *m = AddServerEntryResponse{} }
func (m *AddServerEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*AddServerEntryResponse) ProtoMessage()               {}
func (*AddServerEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type RemoveServerEntryRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *RemoveServerEntryRequest) Reset()                    { *m = RemoveServerEntryRequest{} }
func (m *RemoveServerEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveServerEntryRequest) ProtoMessage()               {}
func (*RemoveServerEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RemoveServerEntryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RemoveServerEntryResponse struct {
}

func (m *RemoveServerEntryResponse) Reset()                    { *m = RemoveServerEntryResponse{} }
func (m *RemoveServerEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveServerEntryResponse) ProtoMessage()               {}
func (*RemoveServerEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

// PushStatus
type PushStatusRequest struct {
	Entry []*PushStatusRequest_StatusEntry `protobuf:"bytes,1,rep,name=entry" json:"entry,omitempty"`
}

func (m *PushStatusRequest) Reset()                    { *m = PushStatusRequest{} }
func (m *PushStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*PushStatusRequest) ProtoMessage()               {}
func (*PushStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *PushStatusRequest) GetEntry() []*PushStatusRequest_StatusEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type PushStatusRequest_StatusEntry struct {
	ServerName string        `protobuf:"bytes,1,opt,name=serverName" json:"serverName,omitempty"`
	Status     *ServerStatus `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *PushStatusRequest_StatusEntry) Reset()         { *m = PushStatusRequest_StatusEntry{} }
func (m *PushStatusRequest_StatusEntry) String() string { return proto.CompactTextString(m) }
func (*PushStatusRequest_StatusEntry) ProtoMessage()    {}
func (*PushStatusRequest_StatusEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{12, 0}
}

func (m *PushStatusRequest_StatusEntry) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *PushStatusRequest_StatusEntry) GetStatus() *ServerStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type PushStatusResponse struct {
}

func (m *PushStatusResponse) Reset()                    { *m = PushStatusResponse{} }
func (m *PushStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*PushStatusResponse) ProtoMessage()               {}
func (*PushStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

// FetchStatus
type FetchStatusRequest struct {
}

func (m *FetchStatusRequest) Reset()                    { *m = FetchStatusRequest{} }
func (m *FetchStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchStatusRequest) ProtoMessage()               {}
func (*FetchStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type FetchStatusResponse struct {
	Entry []*ServerEntry `protobuf:"bytes,1,rep,name=entry" json:"entry,omitempty"`
}

func (m *FetchStatusResponse) Reset()                    { *m = FetchStatusResponse{} }
func (m *FetchStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchStatusResponse) ProtoMessage()               {}
func (*FetchStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *FetchStatusResponse) GetEntry() []*ServerEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "nebulapb.Empty")
	proto.RegisterType((*StreamRequest)(nil), "nebulapb.StreamRequest")
	proto.RegisterType((*EntryStreamResponse)(nil), "nebulapb.EntryStreamResponse")
	proto.RegisterType((*QuitStreamRequest)(nil), "nebulapb.QuitStreamRequest")
	proto.RegisterType((*ServerEntry)(nil), "nebulapb.ServerEntry")
	proto.RegisterType((*ServerStatus)(nil), "nebulapb.ServerStatus")
	proto.RegisterType((*GetServerEntryRequest)(nil), "nebulapb.GetServerEntryRequest")
	proto.RegisterType((*GetServerEntryResponse)(nil), "nebulapb.GetServerEntryResponse")
	proto.RegisterType((*AddServerEntryRequest)(nil), "nebulapb.AddServerEntryRequest")
	proto.RegisterType((*AddServerEntryResponse)(nil), "nebulapb.AddServerEntryResponse")
	proto.RegisterType((*RemoveServerEntryRequest)(nil), "nebulapb.RemoveServerEntryRequest")
	proto.RegisterType((*RemoveServerEntryResponse)(nil), "nebulapb.RemoveServerEntryResponse")
	proto.RegisterType((*PushStatusRequest)(nil), "nebulapb.PushStatusRequest")
	proto.RegisterType((*PushStatusRequest_StatusEntry)(nil), "nebulapb.PushStatusRequest.StatusEntry")
	proto.RegisterType((*PushStatusResponse)(nil), "nebulapb.PushStatusResponse")
	proto.RegisterType((*FetchStatusRequest)(nil), "nebulapb.FetchStatusRequest")
	proto.RegisterType((*FetchStatusResponse)(nil), "nebulapb.FetchStatusResponse")
	proto.RegisterEnum("nebulapb.StreamType", StreamType_name, StreamType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Nebula service

type NebulaClient interface {
	// -
	// MISC
	// -
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// -
	// STREAM
	// -
	// API -> Bungee (Notify Server Status)
	EntryStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Nebula_EntryStreamClient, error)
	QuitStream(ctx context.Context, in *QuitStreamRequest, opts ...grpc.CallOption) (*Empty, error)
	// -
	// ServerEntry
	// -
	// API -> Bungee (ServerEntry)
	GetServerEntry(ctx context.Context, in *GetServerEntryRequest, opts ...grpc.CallOption) (*GetServerEntryResponse, error)
	// API <- App
	AddServerEntry(ctx context.Context, in *AddServerEntryRequest, opts ...grpc.CallOption) (*AddServerEntryResponse, error)
	// API <- App
	RemoveServerEntry(ctx context.Context, in *RemoveServerEntryRequest, opts ...grpc.CallOption) (*RemoveServerEntryResponse, error)
	// -
	// ServerStatus
	// -
	// Bungee -> API (Server Status)
	PushStatus(ctx context.Context, in *PushStatusRequest, opts ...grpc.CallOption) (*PushStatusResponse, error)
	// Spigot <- API (Server Status)
	FetchStatus(ctx context.Context, in *FetchStatusRequest, opts ...grpc.CallOption) (*FetchStatusResponse, error)
}

type nebulaClient struct {
	cc *grpc.ClientConn
}

func NewNebulaClient(cc *grpc.ClientConn) NebulaClient {
	return &nebulaClient{cc}
}

func (c *nebulaClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/nebulapb.Nebula/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nebulaClient) EntryStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Nebula_EntryStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Nebula_serviceDesc.Streams[0], c.cc, "/nebulapb.Nebula/EntryStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &nebulaEntryStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Nebula_EntryStreamClient interface {
	Recv() (*EntryStreamResponse, error)
	grpc.ClientStream
}

type nebulaEntryStreamClient struct {
	grpc.ClientStream
}

func (x *nebulaEntryStreamClient) Recv() (*EntryStreamResponse, error) {
	m := new(EntryStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nebulaClient) QuitStream(ctx context.Context, in *QuitStreamRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/nebulapb.Nebula/QuitStream", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nebulaClient) GetServerEntry(ctx context.Context, in *GetServerEntryRequest, opts ...grpc.CallOption) (*GetServerEntryResponse, error) {
	out := new(GetServerEntryResponse)
	err := grpc.Invoke(ctx, "/nebulapb.Nebula/GetServerEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nebulaClient) AddServerEntry(ctx context.Context, in *AddServerEntryRequest, opts ...grpc.CallOption) (*AddServerEntryResponse, error) {
	out := new(AddServerEntryResponse)
	err := grpc.Invoke(ctx, "/nebulapb.Nebula/AddServerEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nebulaClient) RemoveServerEntry(ctx context.Context, in *RemoveServerEntryRequest, opts ...grpc.CallOption) (*RemoveServerEntryResponse, error) {
	out := new(RemoveServerEntryResponse)
	err := grpc.Invoke(ctx, "/nebulapb.Nebula/RemoveServerEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nebulaClient) PushStatus(ctx context.Context, in *PushStatusRequest, opts ...grpc.CallOption) (*PushStatusResponse, error) {
	out := new(PushStatusResponse)
	err := grpc.Invoke(ctx, "/nebulapb.Nebula/PushStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nebulaClient) FetchStatus(ctx context.Context, in *FetchStatusRequest, opts ...grpc.CallOption) (*FetchStatusResponse, error) {
	out := new(FetchStatusResponse)
	err := grpc.Invoke(ctx, "/nebulapb.Nebula/FetchStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Nebula service

type NebulaServer interface {
	// -
	// MISC
	// -
	Ping(context.Context, *Empty) (*Empty, error)
	// -
	// STREAM
	// -
	// API -> Bungee (Notify Server Status)
	EntryStream(*StreamRequest, Nebula_EntryStreamServer) error
	QuitStream(context.Context, *QuitStreamRequest) (*Empty, error)
	// -
	// ServerEntry
	// -
	// API -> Bungee (ServerEntry)
	GetServerEntry(context.Context, *GetServerEntryRequest) (*GetServerEntryResponse, error)
	// API <- App
	AddServerEntry(context.Context, *AddServerEntryRequest) (*AddServerEntryResponse, error)
	// API <- App
	RemoveServerEntry(context.Context, *RemoveServerEntryRequest) (*RemoveServerEntryResponse, error)
	// -
	// ServerStatus
	// -
	// Bungee -> API (Server Status)
	PushStatus(context.Context, *PushStatusRequest) (*PushStatusResponse, error)
	// Spigot <- API (Server Status)
	FetchStatus(context.Context, *FetchStatusRequest) (*FetchStatusResponse, error)
}

func RegisterNebulaServer(s *grpc.Server, srv NebulaServer) {
	s.RegisterService(&_Nebula_serviceDesc, srv)
}

func _Nebula_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NebulaServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebulapb.Nebula/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NebulaServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nebula_EntryStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NebulaServer).EntryStream(m, &nebulaEntryStreamServer{stream})
}

type Nebula_EntryStreamServer interface {
	Send(*EntryStreamResponse) error
	grpc.ServerStream
}

type nebulaEntryStreamServer struct {
	grpc.ServerStream
}

func (x *nebulaEntryStreamServer) Send(m *EntryStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Nebula_QuitStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuitStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NebulaServer).QuitStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebulapb.Nebula/QuitStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NebulaServer).QuitStream(ctx, req.(*QuitStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nebula_GetServerEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NebulaServer).GetServerEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebulapb.Nebula/GetServerEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NebulaServer).GetServerEntry(ctx, req.(*GetServerEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nebula_AddServerEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddServerEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NebulaServer).AddServerEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebulapb.Nebula/AddServerEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NebulaServer).AddServerEntry(ctx, req.(*AddServerEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nebula_RemoveServerEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveServerEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NebulaServer).RemoveServerEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebulapb.Nebula/RemoveServerEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NebulaServer).RemoveServerEntry(ctx, req.(*RemoveServerEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nebula_PushStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NebulaServer).PushStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebulapb.Nebula/PushStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NebulaServer).PushStatus(ctx, req.(*PushStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nebula_FetchStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NebulaServer).FetchStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebulapb.Nebula/FetchStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NebulaServer).FetchStatus(ctx, req.(*FetchStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Nebula_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nebulapb.Nebula",
	HandlerType: (*NebulaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Nebula_Ping_Handler,
		},
		{
			MethodName: "QuitStream",
			Handler:    _Nebula_QuitStream_Handler,
		},
		{
			MethodName: "GetServerEntry",
			Handler:    _Nebula_GetServerEntry_Handler,
		},
		{
			MethodName: "AddServerEntry",
			Handler:    _Nebula_AddServerEntry_Handler,
		},
		{
			MethodName: "RemoveServerEntry",
			Handler:    _Nebula_RemoveServerEntry_Handler,
		},
		{
			MethodName: "PushStatus",
			Handler:    _Nebula_PushStatus_Handler,
		},
		{
			MethodName: "FetchStatus",
			Handler:    _Nebula_FetchStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EntryStream",
			Handler:       _Nebula_EntryStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "nebulapb.proto",
}

func init() { proto.RegisterFile("nebulapb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 670 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xdb, 0x4e, 0xdb, 0x40,
	0x10, 0x8d, 0x89, 0x13, 0x60, 0x0c, 0x34, 0x0c, 0x10, 0xdc, 0xd0, 0x4b, 0xb4, 0xad, 0x44, 0x44,
	0x25, 0xab, 0x4a, 0x5f, 0xdb, 0x07, 0x08, 0x69, 0x1b, 0xa9, 0x0d, 0xc1, 0x09, 0xaa, 0xfa, 0xd0,
	0x4a, 0x06, 0xaf, 0xc0, 0x52, 0x7c, 0xa9, 0x77, 0x83, 0xea, 0x0f, 0xe8, 0x07, 0xf5, 0x9f, 0xfa,
	0x21, 0x95, 0xd7, 0x36, 0x5e, 0xc7, 0xa6, 0x48, 0x7d, 0xf3, 0xce, 0x9e, 0x39, 0x73, 0x66, 0xe7,
	0x62, 0xd8, 0xf2, 0xe8, 0xe5, 0x62, 0x6e, 0x05, 0x97, 0x46, 0x10, 0xfa, 0xdc, 0xc7, 0xb5, 0xec,
	0x4c, 0x56, 0xa1, 0x31, 0x74, 0x03, 0x1e, 0x91, 0xcf, 0xb0, 0x39, 0xe5, 0x21, 0xb5, 0x5c, 0x93,
	0xfe, 0x58, 0x50, 0xc6, 0x11, 0x41, 0xf5, 0x2c, 0x97, 0xea, 0x4a, 0x57, 0xe9, 0xad, 0x9b, 0xe2,
	0x1b, 0x7b, 0xa0, 0xf2, 0x28, 0xa0, 0xfa, 0x4a, 0x57, 0xe9, 0x6d, 0xf5, 0x77, 0x8d, 0x3b, 0xda,
	0xc4, 0x75, 0x16, 0x05, 0xd4, 0x14, 0x08, 0xf2, 0x05, 0x76, 0x86, 0x1e, 0x0f, 0xa3, 0x8c, 0x93,
	0x05, 0xbe, 0xc7, 0x72, 0x02, 0xe5, 0x21, 0x02, 0x6c, 0x43, 0x93, 0x5b, 0xe1, 0x35, 0xe5, 0x22,
	0xd8, 0xba, 0x99, 0x9e, 0xc8, 0x21, 0x6c, 0x9f, 0x2f, 0x1c, 0xfe, 0xa0, 0x56, 0xf2, 0x4b, 0x01,
	0x6d, 0x4a, 0xc3, 0x5b, 0x1a, 0x0a, 0x21, 0x95, 0xf9, 0x74, 0x41, 0xb3, 0x1d, 0x16, 0xcc, 0xad,
	0x68, 0x1c, 0x5f, 0x25, 0x91, 0x64, 0x13, 0xea, 0xb0, 0x6a, 0xd9, 0x76, 0x48, 0x19, 0xd3, 0xeb,
	0xe2, 0x36, 0x3b, 0xc6, 0x7c, 0x81, 0x1f, 0x72, 0x5d, 0xed, 0x2a, 0xbd, 0x86, 0x29, 0xbe, 0x63,
	0x9b, 0xeb, 0x73, 0x5b, 0x6f, 0x24, 0x31, 0xe2, 0x6f, 0x32, 0x87, 0x8d, 0x44, 0xc6, 0x94, 0x5b,
	0x7c, 0xc1, 0xe2, 0xc4, 0x7c, 0x6f, 0xee, 0x78, 0x99, 0x92, 0xf4, 0x84, 0x2f, 0x61, 0x33, 0xf9,
	0x9a, 0xcc, 0xad, 0x88, 0x86, 0x4c, 0xa8, 0x69, 0x98, 0x45, 0x23, 0x3e, 0x03, 0x70, 0xad, 0x9f,
	0x19, 0xa4, 0x2e, 0x20, 0x92, 0x85, 0xec, 0xc3, 0xde, 0x07, 0xca, 0xa5, 0xbc, 0xd3, 0x27, 0x22,
	0x43, 0x68, 0x2f, 0x5f, 0xa4, 0x35, 0x79, 0x05, 0x0d, 0x1a, 0x1b, 0x74, 0xa5, 0x5b, 0xef, 0x69,
	0xfd, 0x3d, 0xa9, 0x28, 0x12, 0x3a, 0xc1, 0x90, 0x53, 0xd8, 0x3b, 0xb6, 0xed, 0x32, 0xbf, 0xcc,
	0xa2, 0x3c, 0xc8, 0xa2, 0x43, 0x7b, 0x99, 0x25, 0x11, 0x43, 0x0c, 0xd0, 0x4d, 0xea, 0xfa, 0xb7,
	0xb4, 0x22, 0x44, 0x55, 0x95, 0x0f, 0xe0, 0x71, 0x05, 0x3e, 0x25, 0xfb, 0xad, 0xc0, 0xf6, 0x64,
	0xc1, 0x6e, 0x92, 0x97, 0xcf, 0x68, 0xde, 0x15, 0xf3, 0x3d, 0xcc, 0x95, 0x96, 0xb0, 0x46, 0x72,
	0x92, 0xb5, 0x77, 0xbe, 0x81, 0x26, 0x59, 0xe3, 0x82, 0x30, 0x11, 0x7a, 0x9c, 0x4b, 0x93, 0x2c,
	0x68, 0x40, 0x93, 0x09, 0xb8, 0xa8, 0xa7, 0xd6, 0x6f, 0x2f, 0x3f, 0x4c, 0x1a, 0x30, 0x45, 0x91,
	0x5d, 0x40, 0x59, 0x46, 0x9a, 0xc9, 0x2e, 0xe0, 0x7b, 0xca, 0xaf, 0x8a, 0xea, 0xc8, 0x09, 0xec,
	0x14, 0xac, 0xff, 0x51, 0xd0, 0xa3, 0x11, 0x40, 0x3e, 0x7b, 0xb8, 0x06, 0xea, 0xf9, 0xc5, 0x68,
	0xd6, 0xaa, 0xa1, 0x06, 0xab, 0x83, 0xb3, 0xf1, 0x78, 0x38, 0x98, 0xb5, 0x14, 0xdc, 0x80, 0x35,
	0x73, 0x38, 0x9d, 0x9d, 0x99, 0xc3, 0xd3, 0xd6, 0x4a, 0x7c, 0x3a, 0x1d, 0x4d, 0x27, 0xc7, 0xb3,
	0xc1, 0xc7, 0x56, 0x3d, 0x76, 0x99, 0x7e, 0x1d, 0x0f, 0x5a, 0x6a, 0xff, 0x8f, 0x0a, 0xcd, 0xb1,
	0x08, 0x85, 0x47, 0xa0, 0x4e, 0x1c, 0xef, 0x1a, 0x1f, 0xe5, 0xb1, 0xc5, 0x9a, 0xe9, 0x2c, 0x1b,
	0x48, 0x0d, 0x47, 0xa0, 0x49, 0xab, 0x02, 0xf7, 0x97, 0x97, 0x42, 0x9a, 0x6d, 0xe7, 0xa9, 0xe4,
	0x5a, 0x5e, 0x2d, 0xa4, 0xf6, 0x5a, 0xc1, 0xb7, 0x00, 0xf9, 0x72, 0xc0, 0x83, 0xdc, 0xa1, 0xb4,
	0x32, 0xaa, 0x84, 0x5c, 0xc0, 0x56, 0x71, 0x44, 0xf0, 0x79, 0x0e, 0xaa, 0x9c, 0xaa, 0x4e, 0xf7,
	0x7e, 0x40, 0x26, 0x2b, 0xa6, 0x2d, 0x36, 0xbb, 0x4c, 0x5b, 0x39, 0x4c, 0x32, 0xed, 0x3d, 0x73,
	0x52, 0xc3, 0xef, 0xb0, 0x5d, 0xea, 0x7c, 0x24, 0xb9, 0xe3, 0x7d, 0x63, 0xd4, 0x79, 0xf1, 0x4f,
	0xcc, 0x1d, 0xff, 0x08, 0x20, 0x6f, 0x44, 0xf9, 0x2d, 0x4b, 0x53, 0xd2, 0x79, 0x52, 0x7d, 0x79,
	0x47, 0xf5, 0x09, 0x34, 0xa9, 0x4f, 0x51, 0x82, 0x97, 0x9b, 0x5a, 0x2e, 0x73, 0x45, 0x73, 0x93,
	0xda, 0x49, 0x0f, 0x74, 0x8f, 0x72, 0x83, 0x45, 0xde, 0xd5, 0x0d, 0xbf, 0x71, 0x2c, 0xc3, 0x0a,
	0x9c, 0xd4, 0xe7, 0x64, 0x23, 0xe9, 0xbf, 0x49, 0xfc, 0x97, 0x63, 0x97, 0x4d, 0xf1, 0xb7, 0x7b,
	0xf3, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x60, 0x5b, 0x8d, 0xff, 0x06, 0x00, 0x00,
}
