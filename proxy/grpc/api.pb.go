// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package pbrpc is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	Empty
	Pong
	Err
	ContainerID
	LogOpts
	Container
	Containers
	Io
	WindowSize
	ExecOptions
*/
package pbrpc

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

type Empty struct {
	Auth string `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Empty) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

type Pong struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Pong) Reset()                    { *m = Pong{} }
func (m *Pong) String() string            { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()               {}
func (*Pong) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Pong) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Err struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *Err) Reset()                    { *m = Err{} }
func (m *Err) String() string            { return proto.CompactTextString(m) }
func (*Err) ProtoMessage()               {}
func (*Err) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Err) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type ContainerID struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Auth string `protobuf:"bytes,2,opt,name=auth" json:"auth,omitempty"`
}

func (m *ContainerID) Reset()                    { *m = ContainerID{} }
func (m *ContainerID) String() string            { return proto.CompactTextString(m) }
func (*ContainerID) ProtoMessage()               {}
func (*ContainerID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ContainerID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ContainerID) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

type LogOpts struct {
	C      *ContainerID `protobuf:"bytes,1,opt,name=c" json:"c,omitempty"`
	Follow bool         `protobuf:"varint,2,opt,name=follow" json:"follow,omitempty"`
	Tail   string       `protobuf:"bytes,3,opt,name=tail" json:"tail,omitempty"`
}

func (m *LogOpts) Reset()                    { *m = LogOpts{} }
func (m *LogOpts) String() string            { return proto.CompactTextString(m) }
func (*LogOpts) ProtoMessage()               {}
func (*LogOpts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LogOpts) GetC() *ContainerID {
	if m != nil {
		return m.C
	}
	return nil
}

func (m *LogOpts) GetFollow() bool {
	if m != nil {
		return m.Follow
	}
	return false
}

func (m *LogOpts) GetTail() string {
	if m != nil {
		return m.Tail
	}
	return ""
}

// Container instance
type Container struct {
	Id            string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name          string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Image         string   `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
	Command       string   `protobuf:"bytes,4,opt,name=command" json:"command,omitempty"`
	State         string   `protobuf:"bytes,5,opt,name=state" json:"state,omitempty"`
	Status        string   `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
	Ips           []string `protobuf:"bytes,7,rep,name=ips" json:"ips,omitempty"`
	Shell         string   `protobuf:"bytes,8,opt,name=shell" json:"shell,omitempty"`
	PodName       string   `protobuf:"bytes,9,opt,name=pod_name,json=podName" json:"pod_name,omitempty"`
	ContainerName string   `protobuf:"bytes,10,opt,name=container_name,json=containerName" json:"container_name,omitempty"`
	Namespace     string   `protobuf:"bytes,11,opt,name=namespace" json:"namespace,omitempty"`
	RunningNode   string   `protobuf:"bytes,12,opt,name=running_node,json=runningNode" json:"running_node,omitempty"`
	LocServer     string   `protobuf:"bytes,13,opt,name=loc_server,json=locServer" json:"loc_server,omitempty"`
	ExecCMD       string   `protobuf:"bytes,14,opt,name=execCMD" json:"execCMD,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Container) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Container) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Container) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Container) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Container) GetIps() []string {
	if m != nil {
		return m.Ips
	}
	return nil
}

func (m *Container) GetShell() string {
	if m != nil {
		return m.Shell
	}
	return ""
}

func (m *Container) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *Container) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *Container) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Container) GetRunningNode() string {
	if m != nil {
		return m.RunningNode
	}
	return ""
}

func (m *Container) GetLocServer() string {
	if m != nil {
		return m.LocServer
	}
	return ""
}

func (m *Container) GetExecCMD() string {
	if m != nil {
		return m.ExecCMD
	}
	return ""
}

type Containers struct {
	Cs []*Container `protobuf:"bytes,1,rep,name=cs" json:"cs,omitempty"`
}

func (m *Containers) Reset()                    { *m = Containers{} }
func (m *Containers) String() string            { return proto.CompactTextString(m) }
func (*Containers) ProtoMessage()               {}
func (*Containers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Containers) GetCs() []*Container {
	if m != nil {
		return m.Cs
	}
	return nil
}

type Io struct {
	In  []byte `protobuf:"bytes,1,opt,name=in,proto3" json:"in,omitempty"`
	Out []byte `protobuf:"bytes,2,opt,name=out,proto3" json:"out,omitempty"`
}

func (m *Io) Reset()                    { *m = Io{} }
func (m *Io) String() string            { return proto.CompactTextString(m) }
func (*Io) ProtoMessage()               {}
func (*Io) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Io) GetIn() []byte {
	if m != nil {
		return m.In
	}
	return nil
}

func (m *Io) GetOut() []byte {
	if m != nil {
		return m.Out
	}
	return nil
}

type WindowSize struct {
	Height int32 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
	Width  int32 `protobuf:"varint,2,opt,name=width" json:"width,omitempty"`
}

func (m *WindowSize) Reset()                    { *m = WindowSize{} }
func (m *WindowSize) String() string            { return proto.CompactTextString(m) }
func (*WindowSize) ProtoMessage()               {}
func (*WindowSize) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *WindowSize) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *WindowSize) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

type ExecOptions struct {
	Cmd  *Io         `protobuf:"bytes,1,opt,name=cmd" json:"cmd,omitempty"`
	C    *Container  `protobuf:"bytes,2,opt,name=c" json:"c,omitempty"`
	Err  string      `protobuf:"bytes,3,opt,name=err" json:"err,omitempty"`
	Auth string      `protobuf:"bytes,4,opt,name=auth" json:"auth,omitempty"`
	Ws   *WindowSize `protobuf:"bytes,5,opt,name=ws" json:"ws,omitempty"`
}

func (m *ExecOptions) Reset()                    { *m = ExecOptions{} }
func (m *ExecOptions) String() string            { return proto.CompactTextString(m) }
func (*ExecOptions) ProtoMessage()               {}
func (*ExecOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ExecOptions) GetCmd() *Io {
	if m != nil {
		return m.Cmd
	}
	return nil
}

func (m *ExecOptions) GetC() *Container {
	if m != nil {
		return m.C
	}
	return nil
}

func (m *ExecOptions) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *ExecOptions) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

func (m *ExecOptions) GetWs() *WindowSize {
	if m != nil {
		return m.Ws
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "pbrpc.empty")
	proto.RegisterType((*Pong)(nil), "pbrpc.pong")
	proto.RegisterType((*Err)(nil), "pbrpc.err")
	proto.RegisterType((*ContainerID)(nil), "pbrpc.ContainerID")
	proto.RegisterType((*LogOpts)(nil), "pbrpc.logOpts")
	proto.RegisterType((*Container)(nil), "pbrpc.Container")
	proto.RegisterType((*Containers)(nil), "pbrpc.Containers")
	proto.RegisterType((*Io)(nil), "pbrpc.io")
	proto.RegisterType((*WindowSize)(nil), "pbrpc.windowSize")
	proto.RegisterType((*ExecOptions)(nil), "pbrpc.execOptions")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ContainerServer service

type ContainerServerClient interface {
	GetInfo(ctx context.Context, in *ContainerID, opts ...grpc.CallOption) (*Container, error)
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Containers, error)
	Start(ctx context.Context, in *ContainerID, opts ...grpc.CallOption) (*Err, error)
	Stop(ctx context.Context, in *ContainerID, opts ...grpc.CallOption) (*Err, error)
	Restart(ctx context.Context, in *ContainerID, opts ...grpc.CallOption) (*Err, error)
	Exec(ctx context.Context, opts ...grpc.CallOption) (ContainerServer_ExecClient, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error)
	Logs(ctx context.Context, in *LogOpts, opts ...grpc.CallOption) (ContainerServer_LogsClient, error)
}

type containerServerClient struct {
	cc *grpc.ClientConn
}

func NewContainerServerClient(cc *grpc.ClientConn) ContainerServerClient {
	return &containerServerClient{cc}
}

func (c *containerServerClient) GetInfo(ctx context.Context, in *ContainerID, opts ...grpc.CallOption) (*Container, error) {
	out := new(Container)
	err := grpc.Invoke(ctx, "/pbrpc.containerServer/GetInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServerClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Containers, error) {
	out := new(Containers)
	err := grpc.Invoke(ctx, "/pbrpc.containerServer/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServerClient) Start(ctx context.Context, in *ContainerID, opts ...grpc.CallOption) (*Err, error) {
	out := new(Err)
	err := grpc.Invoke(ctx, "/pbrpc.containerServer/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServerClient) Stop(ctx context.Context, in *ContainerID, opts ...grpc.CallOption) (*Err, error) {
	out := new(Err)
	err := grpc.Invoke(ctx, "/pbrpc.containerServer/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServerClient) Restart(ctx context.Context, in *ContainerID, opts ...grpc.CallOption) (*Err, error) {
	out := new(Err)
	err := grpc.Invoke(ctx, "/pbrpc.containerServer/Restart", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServerClient) Exec(ctx context.Context, opts ...grpc.CallOption) (ContainerServer_ExecClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ContainerServer_serviceDesc.Streams[0], c.cc, "/pbrpc.containerServer/Exec", opts...)
	if err != nil {
		return nil, err
	}
	x := &containerServerExecClient{stream}
	return x, nil
}

type ContainerServer_ExecClient interface {
	Send(*ExecOptions) error
	Recv() (*ExecOptions, error)
	grpc.ClientStream
}

type containerServerExecClient struct {
	grpc.ClientStream
}

func (x *containerServerExecClient) Send(m *ExecOptions) error {
	return x.ClientStream.SendMsg(m)
}

func (x *containerServerExecClient) Recv() (*ExecOptions, error) {
	m := new(ExecOptions)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *containerServerClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := grpc.Invoke(ctx, "/pbrpc.containerServer/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServerClient) Logs(ctx context.Context, in *LogOpts, opts ...grpc.CallOption) (ContainerServer_LogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ContainerServer_serviceDesc.Streams[1], c.cc, "/pbrpc.containerServer/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &containerServerLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ContainerServer_LogsClient interface {
	Recv() (*Io, error)
	grpc.ClientStream
}

type containerServerLogsClient struct {
	grpc.ClientStream
}

func (x *containerServerLogsClient) Recv() (*Io, error) {
	m := new(Io)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ContainerServer service

type ContainerServerServer interface {
	GetInfo(context.Context, *ContainerID) (*Container, error)
	List(context.Context, *Empty) (*Containers, error)
	Start(context.Context, *ContainerID) (*Err, error)
	Stop(context.Context, *ContainerID) (*Err, error)
	Restart(context.Context, *ContainerID) (*Err, error)
	Exec(ContainerServer_ExecServer) error
	Ping(context.Context, *Empty) (*Pong, error)
	Logs(*LogOpts, ContainerServer_LogsServer) error
}

func RegisterContainerServerServer(s *grpc.Server, srv ContainerServerServer) {
	s.RegisterService(&_ContainerServer_serviceDesc, srv)
}

func _ContainerServer_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServerServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbrpc.containerServer/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServerServer).GetInfo(ctx, req.(*ContainerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerServer_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbrpc.containerServer/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServerServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerServer_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbrpc.containerServer/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServerServer).Start(ctx, req.(*ContainerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerServer_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbrpc.containerServer/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServerServer).Stop(ctx, req.(*ContainerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerServer_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServerServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbrpc.containerServer/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServerServer).Restart(ctx, req.(*ContainerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerServer_Exec_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ContainerServerServer).Exec(&containerServerExecServer{stream})
}

type ContainerServer_ExecServer interface {
	Send(*ExecOptions) error
	Recv() (*ExecOptions, error)
	grpc.ServerStream
}

type containerServerExecServer struct {
	grpc.ServerStream
}

func (x *containerServerExecServer) Send(m *ExecOptions) error {
	return x.ServerStream.SendMsg(m)
}

func (x *containerServerExecServer) Recv() (*ExecOptions, error) {
	m := new(ExecOptions)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ContainerServer_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbrpc.containerServer/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServerServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerServer_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogOpts)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContainerServerServer).Logs(m, &containerServerLogsServer{stream})
}

type ContainerServer_LogsServer interface {
	Send(*Io) error
	grpc.ServerStream
}

type containerServerLogsServer struct {
	grpc.ServerStream
}

func (x *containerServerLogsServer) Send(m *Io) error {
	return x.ServerStream.SendMsg(m)
}

var _ContainerServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbrpc.containerServer",
	HandlerType: (*ContainerServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _ContainerServer_GetInfo_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ContainerServer_List_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _ContainerServer_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _ContainerServer_Stop_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _ContainerServer_Restart_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _ContainerServer_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Exec",
			Handler:       _ContainerServer_Exec_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Logs",
			Handler:       _ContainerServer_Logs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x6e, 0xdb, 0x38,
	0x10, 0xb5, 0x64, 0x29, 0x8e, 0x47, 0x8e, 0x37, 0x21, 0x16, 0xbb, 0x5c, 0x67, 0x5b, 0x38, 0x2a,
	0x52, 0xb8, 0x28, 0x60, 0x24, 0x6e, 0x4f, 0xbd, 0x26, 0x45, 0x11, 0x20, 0x4d, 0x0a, 0xf9, 0xd0,
	0x63, 0xa0, 0x50, 0x8c, 0x4c, 0x40, 0x22, 0x09, 0x92, 0xae, 0xd3, 0xfe, 0x46, 0xbf, 0xb4, 0xe8,
	0x0f, 0x14, 0xa4, 0x68, 0x39, 0x48, 0x7c, 0xc8, 0x6d, 0xde, 0xcc, 0xe3, 0xe3, 0xd3, 0xcc, 0x88,
	0xd0, 0xcf, 0x25, 0x9b, 0x4a, 0x25, 0x8c, 0x40, 0xb1, 0xbc, 0x55, 0x92, 0xa4, 0x87, 0x10, 0xd3,
	0x5a, 0x9a, 0xef, 0x08, 0x41, 0x94, 0x2f, 0xcd, 0x02, 0x07, 0xe3, 0x60, 0xd2, 0xcf, 0x5c, 0x9c,
	0x62, 0x88, 0xa4, 0xe0, 0x25, 0xda, 0x87, 0x6e, 0xad, 0x4b, 0x5f, 0xb2, 0x61, 0xfa, 0x2f, 0x74,
	0xa9, 0x52, 0xb6, 0x40, 0x95, 0x5a, 0x17, 0xa8, 0x52, 0xe9, 0x29, 0x24, 0x67, 0x82, 0x9b, 0x9c,
	0x71, 0xaa, 0x2e, 0xce, 0xd1, 0x10, 0x42, 0x56, 0xf8, 0x7a, 0xc8, 0x8a, 0xf6, 0x96, 0xf0, 0xc1,
	0x2d, 0x5f, 0xa1, 0x57, 0x89, 0xf2, 0x5a, 0x1a, 0x8d, 0xc6, 0x10, 0x10, 0xc7, 0x4e, 0x66, 0x68,
	0xea, 0x0c, 0x4e, 0x1f, 0xa8, 0x65, 0x01, 0x41, 0xff, 0xc0, 0xce, 0x9d, 0xa8, 0x2a, 0xb1, 0x72,
	0x12, 0xbb, 0x99, 0x47, 0x56, 0xd8, 0xe4, 0xac, 0xc2, 0xdd, 0x46, 0xd8, 0xc6, 0xe9, 0xaf, 0x10,
	0xfa, 0xed, 0xf1, 0x6d, 0x56, 0x78, 0x5e, 0xd3, 0xb5, 0x15, 0x1b, 0xa3, 0xbf, 0x21, 0x66, 0x75,
	0x5e, 0x52, 0x2f, 0xd3, 0x00, 0x84, 0xa1, 0x47, 0x44, 0x5d, 0xe7, 0xbc, 0xc0, 0x91, 0xcb, 0xaf,
	0xa1, 0xe5, 0x6b, 0x93, 0x1b, 0x8a, 0xe3, 0x86, 0xef, 0x80, 0xf5, 0x68, 0x83, 0xa5, 0xc6, 0x3b,
	0x2e, 0xed, 0x91, 0xed, 0x16, 0x93, 0x1a, 0xf7, 0xc6, 0x5d, 0xdb, 0x2d, 0x26, 0xb5, 0x3b, 0xbf,
	0xa0, 0x55, 0x85, 0x77, 0xfd, 0x79, 0x0b, 0xd0, 0x7f, 0xb0, 0x2b, 0x45, 0x71, 0xe3, 0xdc, 0xf5,
	0x9b, 0x0b, 0xa5, 0x28, 0xae, 0xac, 0xc1, 0x63, 0x18, 0x92, 0xf5, 0x17, 0x35, 0x04, 0x70, 0x84,
	0xbd, 0x36, 0xeb, 0x68, 0xff, 0x43, 0xdf, 0x16, 0xb5, 0xcc, 0x09, 0xc5, 0x89, 0x63, 0x6c, 0x12,
	0xe8, 0x08, 0x06, 0x6a, 0xc9, 0x39, 0xe3, 0xe5, 0x0d, 0x17, 0x05, 0xc5, 0x03, 0x47, 0x48, 0x7c,
	0xee, 0x4a, 0x14, 0x14, 0xbd, 0x00, 0xa8, 0x04, 0xb9, 0xd1, 0x54, 0x7d, 0xa3, 0x0a, 0xef, 0x35,
	0x0a, 0x95, 0x20, 0x73, 0x97, 0xb0, 0x1d, 0xa1, 0xf7, 0x94, 0x9c, 0x7d, 0x3e, 0xc7, 0xc3, 0xc6,
	0xa0, 0x87, 0xe9, 0x14, 0xa0, 0x6d, 0xb9, 0x9d, 0x67, 0x48, 0x34, 0x0e, 0xc6, 0xdd, 0x49, 0x32,
	0xdb, 0x7f, 0x3c, 0xd0, 0x2c, 0x24, 0x3a, 0x7d, 0x0d, 0x21, 0x13, 0x6e, 0x36, 0xdc, 0xcd, 0x66,
	0x90, 0x85, 0x8c, 0xdb, 0x4e, 0x89, 0xa5, 0x71, 0xa3, 0x19, 0x64, 0x36, 0x4c, 0x3f, 0x00, 0xac,
	0x18, 0x2f, 0xc4, 0x6a, 0xce, 0x7e, 0xb8, 0x0e, 0x2f, 0x28, 0x2b, 0x17, 0xc6, 0x9d, 0x89, 0x33,
	0x8f, 0x6c, 0x3f, 0x57, 0xac, 0xf0, 0xfb, 0x15, 0x67, 0x0d, 0x48, 0x7f, 0x06, 0x90, 0x58, 0x7f,
	0xd7, 0xd2, 0x30, 0xc1, 0x35, 0x3a, 0x84, 0x2e, 0xa9, 0x0b, 0xbf, 0x67, 0x7d, 0x6f, 0x8b, 0x89,
	0xcc, 0x66, 0xd1, 0x4b, 0xbb, 0x82, 0xa1, 0x2b, 0x3d, 0x75, 0x1c, 0x90, 0xf5, 0xca, 0x77, 0xdb,
	0x95, 0x6f, 0x77, 0x3a, 0xda, 0xec, 0x34, 0x3a, 0x82, 0x70, 0xa5, 0xdd, 0x56, 0x24, 0xb3, 0x03,
	0x2f, 0xb3, 0xf1, 0x9f, 0x85, 0x2b, 0x3d, 0xfb, 0x1d, 0xc2, 0x5f, 0xed, 0xd4, 0x7c, 0x5f, 0x4f,
	0xa1, 0xf7, 0x89, 0x9a, 0x0b, 0x7e, 0x27, 0xd0, 0x96, 0xfd, 0x1f, 0x3d, 0x31, 0x94, 0x76, 0xd0,
	0x1b, 0x88, 0x2e, 0x99, 0x36, 0x68, 0xe0, 0x6b, 0xee, 0x6f, 0x1e, 0x1d, 0x3c, 0x66, 0x6a, 0x47,
	0x8d, 0xe7, 0x26, 0x57, 0x66, 0xab, 0x36, 0xac, 0xcf, 0x2b, 0xab, 0x3a, 0x81, 0x68, 0x6e, 0x84,
	0x7c, 0x06, 0xf3, 0x2d, 0xf4, 0x32, 0xaa, 0x9f, 0x29, 0xfb, 0x1e, 0xa2, 0x8f, 0xf7, 0x94, 0xb4,
	0xcc, 0x07, 0x53, 0x19, 0x6d, 0xc9, 0xa5, 0x9d, 0x49, 0x70, 0x12, 0xa0, 0x57, 0x10, 0x7d, 0x61,
	0xbc, 0x7c, 0xf4, 0x89, 0x89, 0x47, 0xf6, 0x85, 0x4a, 0x3b, 0xe8, 0x18, 0xa2, 0x4b, 0x51, 0x6a,
	0x34, 0xf4, 0x69, 0xff, 0xa4, 0x8c, 0x36, 0xf3, 0x4d, 0x3b, 0x27, 0xc1, 0xed, 0x8e, 0x7b, 0xfd,
	0xde, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x20, 0x90, 0x7a, 0x2e, 0x0a, 0x05, 0x00, 0x00,
}
