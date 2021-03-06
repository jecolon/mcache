// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mcache.proto

package mcache

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Status int32

const (
	Status_OK        Status = 0
	Status_NOT_FOUND Status = 1
	Status_STORE_ERR Status = 2
)

var Status_name = map[int32]string{
	0: "OK",
	1: "NOT_FOUND",
	2: "STORE_ERR",
}

var Status_value = map[string]int32{
	"OK":        0,
	"NOT_FOUND": 1,
	"STORE_ERR": 2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a91e42676df706b3, []int{0}
}

type Entry struct {
	Stat                 Status   `protobuf:"varint,2,opt,name=stat,proto3,enum=mcache.Status" json:"stat,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Ctype                string   `protobuf:"bytes,4,opt,name=ctype,proto3" json:"ctype,omitempty"`
	Mtime                int64    `protobuf:"varint,5,opt,name=mtime,proto3" json:"mtime,omitempty"`
	Content              []byte   `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Size                 int64    `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91e42676df706b3, []int{0}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetStat() Status {
	if m != nil {
		return m.Stat
	}
	return Status_OK
}

func (m *Entry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Entry) GetCtype() string {
	if m != nil {
		return m.Ctype
	}
	return ""
}

func (m *Entry) GetMtime() int64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func (m *Entry) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Entry) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func init() {
	proto.RegisterEnum("mcache.Status", Status_name, Status_value)
	proto.RegisterType((*Entry)(nil), "mcache.Entry")
}

func init() { proto.RegisterFile("mcache.proto", fileDescriptor_a91e42676df706b3) }

var fileDescriptor_a91e42676df706b3 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x97, 0xfe, 0xc9, 0xd8, 0xcb, 0x36, 0xca, 0x8b, 0x87, 0xe0, 0xa9, 0x14, 0x84, 0xe2,
	0xa1, 0x87, 0xf9, 0x11, 0x5c, 0xf5, 0x20, 0x2c, 0x12, 0xe7, 0x79, 0xcc, 0xf0, 0x82, 0x43, 0xdb,
	0x8e, 0xf5, 0xdd, 0xa1, 0xe2, 0xf7, 0xf0, 0xeb, 0x4a, 0x12, 0x7b, 0xf0, 0xa4, 0xb7, 0xe7, 0xf7,
	0xe4, 0x07, 0xc9, 0x13, 0x98, 0x37, 0x76, 0x6f, 0x5f, 0xa9, 0x3a, 0x9e, 0x3a, 0xee, 0x50, 0x06,
	0x2a, 0xbe, 0x04, 0xa4, 0x75, 0xcb, 0xa7, 0x01, 0x0b, 0x48, 0x7a, 0xde, 0xb3, 0x8a, 0x72, 0x51,
	0x2e, 0x57, 0xcb, 0xea, 0x47, 0x77, 0xdd, 0xb9, 0x37, 0xfe, 0x0c, 0x33, 0x88, 0xdf, 0x68, 0x50,
	0x71, 0x2e, 0xca, 0x99, 0x71, 0x11, 0x2f, 0x20, 0xb5, 0x3c, 0x1c, 0x49, 0x25, 0xbe, 0x0b, 0xe0,
	0xda, 0x86, 0x0f, 0x0d, 0xa9, 0x34, 0x17, 0x65, 0x6c, 0x02, 0xa0, 0x82, 0xa9, 0xed, 0x5a, 0xa6,
	0x96, 0x95, 0xcc, 0x45, 0x39, 0x37, 0x23, 0x22, 0x42, 0xd2, 0x1f, 0x3e, 0x48, 0x4d, 0xbd, 0xee,
	0xf3, 0x75, 0x05, 0x32, 0xdc, 0x8d, 0x12, 0x22, 0xfd, 0x90, 0x4d, 0x70, 0x01, 0xb3, 0x8d, 0xde,
	0xee, 0xee, 0xf4, 0xf3, 0x66, 0x9d, 0x09, 0x87, 0x4f, 0x5b, 0x6d, 0xea, 0x5d, 0x6d, 0x4c, 0x16,
	0xad, 0x3e, 0x21, 0xbd, 0x75, 0x2f, 0xc6, 0x2b, 0x88, 0xef, 0x89, 0x71, 0x31, 0x2e, 0xf0, 0xf3,
	0x2e, 0x7f, 0x63, 0x31, 0x71, 0xda, 0xe3, 0xf9, 0x5f, 0xda, 0x9a, 0xde, 0xff, 0xd2, 0x5e, 0xa4,
	0xff, 0xd6, 0x9b, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x52, 0x96, 0x1b, 0x9a, 0x66, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CacheClient is the client API for Cache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CacheClient interface {
	Get(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*Entry, error)
	Put(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*Entry, error)
	Del(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*Entry, error)
}

type cacheClient struct {
	cc *grpc.ClientConn
}

func NewCacheClient(cc *grpc.ClientConn) CacheClient {
	return &cacheClient{cc}
}

func (c *cacheClient) Get(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := c.cc.Invoke(ctx, "/mcache.Cache/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) Put(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := c.cc.Invoke(ctx, "/mcache.Cache/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) Del(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := c.cc.Invoke(ctx, "/mcache.Cache/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServer is the server API for Cache service.
type CacheServer interface {
	Get(context.Context, *Entry) (*Entry, error)
	Put(context.Context, *Entry) (*Entry, error)
	Del(context.Context, *Entry) (*Entry, error)
}

// UnimplementedCacheServer can be embedded to have forward compatible implementations.
type UnimplementedCacheServer struct {
}

func (*UnimplementedCacheServer) Get(ctx context.Context, req *Entry) (*Entry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCacheServer) Put(ctx context.Context, req *Entry) (*Entry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedCacheServer) Del(ctx context.Context, req *Entry) (*Entry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}

func RegisterCacheServer(s *grpc.Server, srv CacheServer) {
	s.RegisterService(&_Cache_serviceDesc, srv)
}

func _Cache_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcache.Cache/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Get(ctx, req.(*Entry))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcache.Cache/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Put(ctx, req.(*Entry))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcache.Cache/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Del(ctx, req.(*Entry))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mcache.Cache",
	HandlerType: (*CacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Cache_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Cache_Put_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _Cache_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mcache.proto",
}
