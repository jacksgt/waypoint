// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logs.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

type Logs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Logs) Reset()         { *m = Logs{} }
func (m *Logs) String() string { return proto.CompactTextString(m) }
func (*Logs) ProtoMessage()    {}
func (*Logs) Descriptor() ([]byte, []int) {
	return fileDescriptor_logs_355d40aed9ccd76d, []int{0}
}
func (m *Logs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Logs.Unmarshal(m, b)
}
func (m *Logs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Logs.Marshal(b, m, deterministic)
}
func (dst *Logs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Logs.Merge(dst, src)
}
func (m *Logs) XXX_Size() int {
	return xxx_messageInfo_Logs.Size(m)
}
func (m *Logs) XXX_DiscardUnknown() {
	xxx_messageInfo_Logs.DiscardUnknown(m)
}

var xxx_messageInfo_Logs proto.InternalMessageInfo

type Logs_Resp struct {
	// stream_id is the stream ID to connect to to get access to the
	// LogViewer service.
	StreamId             uint32   `protobuf:"varint,1,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Logs_Resp) Reset()         { *m = Logs_Resp{} }
func (m *Logs_Resp) String() string { return proto.CompactTextString(m) }
func (*Logs_Resp) ProtoMessage()    {}
func (*Logs_Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_logs_355d40aed9ccd76d, []int{0, 0}
}
func (m *Logs_Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Logs_Resp.Unmarshal(m, b)
}
func (m *Logs_Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Logs_Resp.Marshal(b, m, deterministic)
}
func (dst *Logs_Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Logs_Resp.Merge(dst, src)
}
func (m *Logs_Resp) XXX_Size() int {
	return xxx_messageInfo_Logs_Resp.Size(m)
}
func (m *Logs_Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Logs_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Logs_Resp proto.InternalMessageInfo

func (m *Logs_Resp) GetStreamId() uint32 {
	if m != nil {
		return m.StreamId
	}
	return 0
}

type Logs_NextBatchResp struct {
	Events               []*Logs_Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Logs_NextBatchResp) Reset()         { *m = Logs_NextBatchResp{} }
func (m *Logs_NextBatchResp) String() string { return proto.CompactTextString(m) }
func (*Logs_NextBatchResp) ProtoMessage()    {}
func (*Logs_NextBatchResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_logs_355d40aed9ccd76d, []int{0, 1}
}
func (m *Logs_NextBatchResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Logs_NextBatchResp.Unmarshal(m, b)
}
func (m *Logs_NextBatchResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Logs_NextBatchResp.Marshal(b, m, deterministic)
}
func (dst *Logs_NextBatchResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Logs_NextBatchResp.Merge(dst, src)
}
func (m *Logs_NextBatchResp) XXX_Size() int {
	return xxx_messageInfo_Logs_NextBatchResp.Size(m)
}
func (m *Logs_NextBatchResp) XXX_DiscardUnknown() {
	xxx_messageInfo_Logs_NextBatchResp.DiscardUnknown(m)
}

var xxx_messageInfo_Logs_NextBatchResp proto.InternalMessageInfo

func (m *Logs_NextBatchResp) GetEvents() []*Logs_Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type Logs_Event struct {
	Partition            string               `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Contents             string               `protobuf:"bytes,3,opt,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Logs_Event) Reset()         { *m = Logs_Event{} }
func (m *Logs_Event) String() string { return proto.CompactTextString(m) }
func (*Logs_Event) ProtoMessage()    {}
func (*Logs_Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_logs_355d40aed9ccd76d, []int{0, 2}
}
func (m *Logs_Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Logs_Event.Unmarshal(m, b)
}
func (m *Logs_Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Logs_Event.Marshal(b, m, deterministic)
}
func (dst *Logs_Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Logs_Event.Merge(dst, src)
}
func (m *Logs_Event) XXX_Size() int {
	return xxx_messageInfo_Logs_Event.Size(m)
}
func (m *Logs_Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Logs_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Logs_Event proto.InternalMessageInfo

func (m *Logs_Event) GetPartition() string {
	if m != nil {
		return m.Partition
	}
	return ""
}

func (m *Logs_Event) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Logs_Event) GetContents() string {
	if m != nil {
		return m.Contents
	}
	return ""
}

func init() {
	proto.RegisterType((*Logs)(nil), "proto.Logs")
	proto.RegisterType((*Logs_Resp)(nil), "proto.Logs.Resp")
	proto.RegisterType((*Logs_NextBatchResp)(nil), "proto.Logs.NextBatchResp")
	proto.RegisterType((*Logs_Event)(nil), "proto.Logs.Event")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogPlatformClient is the client API for LogPlatform service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogPlatformClient interface {
	LogsSpec(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FuncSpec, error)
	Logs(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*Logs_Resp, error)
}

type logPlatformClient struct {
	cc *grpc.ClientConn
}

func NewLogPlatformClient(cc *grpc.ClientConn) LogPlatformClient {
	return &logPlatformClient{cc}
}

func (c *logPlatformClient) LogsSpec(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FuncSpec, error) {
	out := new(FuncSpec)
	err := c.cc.Invoke(ctx, "/proto.LogPlatform/LogsSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logPlatformClient) Logs(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*Logs_Resp, error) {
	out := new(Logs_Resp)
	err := c.cc.Invoke(ctx, "/proto.LogPlatform/Logs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogPlatformServer is the server API for LogPlatform service.
type LogPlatformServer interface {
	LogsSpec(context.Context, *empty.Empty) (*FuncSpec, error)
	Logs(context.Context, *FuncSpec_Args) (*Logs_Resp, error)
}

func RegisterLogPlatformServer(s *grpc.Server, srv LogPlatformServer) {
	s.RegisterService(&_LogPlatform_serviceDesc, srv)
}

func _LogPlatform_LogsSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogPlatformServer).LogsSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LogPlatform/LogsSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogPlatformServer).LogsSpec(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogPlatform_Logs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuncSpec_Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogPlatformServer).Logs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LogPlatform/Logs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogPlatformServer).Logs(ctx, req.(*FuncSpec_Args))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogPlatform_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LogPlatform",
	HandlerType: (*LogPlatformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogsSpec",
			Handler:    _LogPlatform_LogsSpec_Handler,
		},
		{
			MethodName: "Logs",
			Handler:    _LogPlatform_Logs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logs.proto",
}

// LogViewerClient is the client API for LogViewer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogViewerClient interface {
	NextLogBatch(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Logs_NextBatchResp, error)
}

type logViewerClient struct {
	cc *grpc.ClientConn
}

func NewLogViewerClient(cc *grpc.ClientConn) LogViewerClient {
	return &logViewerClient{cc}
}

func (c *logViewerClient) NextLogBatch(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Logs_NextBatchResp, error) {
	out := new(Logs_NextBatchResp)
	err := c.cc.Invoke(ctx, "/proto.LogViewer/NextLogBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogViewerServer is the server API for LogViewer service.
type LogViewerServer interface {
	NextLogBatch(context.Context, *empty.Empty) (*Logs_NextBatchResp, error)
}

func RegisterLogViewerServer(s *grpc.Server, srv LogViewerServer) {
	s.RegisterService(&_LogViewer_serviceDesc, srv)
}

func _LogViewer_NextLogBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogViewerServer).NextLogBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LogViewer/NextLogBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogViewerServer).NextLogBatch(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogViewer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LogViewer",
	HandlerType: (*LogViewerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NextLogBatch",
			Handler:    _LogViewer_NextLogBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logs.proto",
}

func init() { proto.RegisterFile("logs.proto", fileDescriptor_logs_355d40aed9ccd76d) }

var fileDescriptor_logs_355d40aed9ccd76d = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x5d, 0x4b, 0xfb, 0x30,
	0x14, 0xc6, 0xe9, 0x7f, 0x2f, 0xac, 0x67, 0x1b, 0x7f, 0x0d, 0x22, 0x33, 0x13, 0x1c, 0x7a, 0x33,
	0x6f, 0x32, 0xe8, 0x6e, 0xc4, 0xbb, 0x09, 0x13, 0x84, 0x32, 0xa4, 0x8a, 0xb7, 0xd2, 0x75, 0x59,
	0x2c, 0xb4, 0x4d, 0x48, 0xce, 0x7c, 0xc1, 0xaf, 0xeb, 0x07, 0x91, 0x26, 0xdd, 0x8b, 0x13, 0xaf,
	0xc2, 0x39, 0xcf, 0x8f, 0x27, 0xe7, 0x79, 0x00, 0x32, 0x29, 0x0c, 0x53, 0x5a, 0xa2, 0x24, 0x0d,
	0xfb, 0xd0, 0xbe, 0x90, 0x52, 0x64, 0x7c, 0x64, 0xa7, 0xf9, 0x6a, 0x39, 0xe2, 0xb9, 0xc2, 0x0f,
	0xc7, 0xd0, 0xb3, 0x7d, 0x11, 0xd3, 0x9c, 0x1b, 0x8c, 0x73, 0x55, 0x01, 0x1d, 0x95, 0xad, 0x44,
	0x5a, 0xb8, 0xe9, 0xfc, 0xcb, 0x83, 0x7a, 0x28, 0x85, 0xa1, 0x17, 0x50, 0x8f, 0xb8, 0x51, 0xa4,
	0x0f, 0xbe, 0x41, 0xcd, 0xe3, 0xfc, 0x39, 0x5d, 0xf4, 0xbc, 0x81, 0x37, 0xec, 0x46, 0x2d, 0xb7,
	0xb8, 0x5b, 0xd0, 0x6b, 0xe8, 0xce, 0xf8, 0x3b, 0xde, 0xc4, 0x98, 0xbc, 0x58, 0xfa, 0x12, 0x9a,
	0xfc, 0x95, 0x17, 0x68, 0x7a, 0xde, 0xa0, 0x36, 0x6c, 0x07, 0x87, 0xce, 0x96, 0x95, 0x96, 0x6c,
	0x5a, 0x2a, 0x51, 0x05, 0xd0, 0x4f, 0x68, 0xd8, 0x05, 0x39, 0x05, 0x5f, 0xc5, 0x1a, 0x53, 0x4c,
	0x65, 0x61, 0x7f, 0xf0, 0xa3, 0xed, 0x82, 0x5c, 0x81, 0xbf, 0xb9, 0xb8, 0xf7, 0x6f, 0xe0, 0x0d,
	0xdb, 0x01, 0x65, 0x2e, 0x13, 0x5b, 0x67, 0x62, 0x8f, 0x6b, 0x22, 0xda, 0xc2, 0x84, 0x42, 0x2b,
	0x91, 0x05, 0xda, 0x6b, 0x6a, 0xd6, 0x76, 0x33, 0x07, 0x1a, 0xda, 0xa1, 0x14, 0xf7, 0x59, 0x8c,
	0x4b, 0xa9, 0x73, 0x32, 0x86, 0x56, 0x79, 0xe1, 0x83, 0xe2, 0x09, 0x39, 0xfe, 0xe5, 0x3e, 0x2d,
	0xeb, 0xa4, 0xff, 0xab, 0x28, 0xb7, 0xab, 0x22, 0xb1, 0x20, 0x73, 0x4d, 0x91, 0xa3, 0x3d, 0x81,
	0x4d, 0xb4, 0x30, 0xf4, 0x60, 0x37, 0x79, 0xd9, 0x4d, 0x30, 0x03, 0x3f, 0x94, 0xe2, 0x29, 0xe5,
	0x6f, 0x5c, 0x93, 0x09, 0x74, 0xca, 0xe6, 0x42, 0x29, 0x6c, 0x79, 0x7f, 0xfe, 0x7a, 0xb2, 0x6b,
	0xf3, 0xa3, 0xeb, 0x79, 0xd3, 0x2a, 0xe3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x82, 0xfa, 0x8e,
	0x0f, 0x12, 0x02, 0x00, 0x00,
}