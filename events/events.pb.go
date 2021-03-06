// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
//      protoc-gen-go v1.24.0-devel
//      protoc        v3.11.2
// source: events.proto

package events

import (
        context "context"
        proto "github.com/golang/protobuf/proto"
        grpc "google.golang.org/grpc"
        codes "google.golang.org/grpc/codes"
        status "google.golang.org/grpc/status"
        protoreflect "google.golang.org/protobuf/reflect/protoreflect"
        protoimpl "google.golang.org/protobuf/runtime/protoimpl"
        reflect "reflect"
        sync "sync"
)

const (
        // Verify that this generated code is sufficiently up-to-date.
        _ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
        // Verify that runtime/protoimpl is sufficiently up-to-date.
        _ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Message struct {
        state         protoimpl.MessageState
        sizeCache     protoimpl.SizeCache
        unknownFields protoimpl.UnknownFields

        Event string `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
        Data  string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Message) Reset() {
        *x = Message{}
        if protoimpl.UnsafeEnabled {
                mi := &file_events_proto_msgTypes[0]
                ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
                ms.StoreMessageInfo(mi)
        }
}

func (x *Message) String() string {
        return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
        mi := &file_events_proto_msgTypes[0]
        if protoimpl.UnsafeEnabled && x != nil {
                ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
                if ms.LoadMessageInfo() == nil {
                        ms.StoreMessageInfo(mi)
                }
                return ms
        }
        return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
        return file_events_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetEvent() string {
        if x != nil {
                return x.Event
        }
        return ""
}

func (x *Message) GetData() string {
        if x != nil {
                return x.Data
        }
        return ""
}

var File_events_proto protoreflect.FileDescriptor

var file_events_proto_rawDesc = []byte{
        0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
        0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x33, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
        0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
        0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
        0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x3b, 0x0a, 0x06, 0x53,
        0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x31, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
        0x12, 0x0f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
        0x65, 0x1a, 0x0f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
        0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
        file_events_proto_rawDescOnce sync.Once
        file_events_proto_rawDescData = file_events_proto_rawDesc
)

func file_events_proto_rawDescGZIP() []byte {
        file_events_proto_rawDescOnce.Do(func() {
                file_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_proto_rawDescData)
        })
        return file_events_proto_rawDescData
}

var file_events_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_events_proto_goTypes = []interface{}{
        (*Message)(nil), // 0: events.Message
}
var file_events_proto_depIdxs = []int32{
        0, // 0: events.Stream.Connect:input_type -> events.Message
        0, // 1: events.Stream.Connect:output_type -> events.Message
        1, // [1:2] is the sub-list for method output_type
        0, // [0:1] is the sub-list for method input_type
        0, // [0:0] is the sub-list for extension type_name
        0, // [0:0] is the sub-list for extension extendee
        0, // [0:0] is the sub-list for field type_name
}

func init() { file_events_proto_init() }
func file_events_proto_init() {
        if File_events_proto != nil {
                return
        }
        if !protoimpl.UnsafeEnabled {
                file_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
                        switch v := v.(*Message); i {
                        case 0:
                                return &v.state
                        case 1:
                                return &v.sizeCache
                        case 2:
                                return &v.unknownFields
                        default:
                                return nil
                        }
                }
        }
        type x struct{}
        out := protoimpl.TypeBuilder{
                File: protoimpl.DescBuilder{
                        GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
                        RawDescriptor: file_events_proto_rawDesc,
                        NumEnums:      0,
                        NumMessages:   1,
                        NumExtensions: 0,
                        NumServices:   1,
                },
                GoTypes:           file_events_proto_goTypes,
                DependencyIndexes: file_events_proto_depIdxs,
                MessageInfos:      file_events_proto_msgTypes,
        }.Build()
        File_events_proto = out.File
        file_events_proto_rawDesc = nil
        file_events_proto_goTypes = nil
        file_events_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamClient is the client API for Stream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamClient interface {
        Connect(ctx context.Context, opts ...grpc.CallOption) (Stream_ConnectClient, error)
}

type streamClient struct {
        cc grpc.ClientConnInterface
}

func NewStreamClient(cc grpc.ClientConnInterface) StreamClient {
        return &streamClient{cc}
}

func (c *streamClient) Connect(ctx context.Context, opts ...grpc.CallOption) (Stream_ConnectClient, error) {
        stream, err := c.cc.NewStream(ctx, &_Stream_serviceDesc.Streams[0], "/events.Stream/Connect", opts...)
        if err != nil {
                return nil, err
        }
        x := &streamConnectClient{stream}
        return x, nil
}

type Stream_ConnectClient interface {
        Send(*Message) error
        Recv() (*Message, error)
        grpc.ClientStream
}

type streamConnectClient struct {
        grpc.ClientStream
}

func (x *streamConnectClient) Send(m *Message) error {
        return x.ClientStream.SendMsg(m)
}

func (x *streamConnectClient) Recv() (*Message, error) {
        m := new(Message)
        if err := x.ClientStream.RecvMsg(m); err != nil {
                return nil, err
        }
        return m, nil
}

// StreamServer is the server API for Stream service.
type StreamServer interface {
        Connect(Stream_ConnectServer) error
}

// UnimplementedStreamServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServer struct {
}

func (*UnimplementedStreamServer) Connect(Stream_ConnectServer) error {
        return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}

func RegisterStreamServer(s *grpc.Server, srv StreamServer) {
        s.RegisterService(&_Stream_serviceDesc, srv)
}

func _Stream_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
        return srv.(StreamServer).Connect(&streamConnectServer{stream})
}

type Stream_ConnectServer interface {
        Send(*Message) error
        Recv() (*Message, error)
        grpc.ServerStream
}

type streamConnectServer struct {
        grpc.ServerStream
}

func (x *streamConnectServer) Send(m *Message) error {
        return x.ServerStream.SendMsg(m)
}

func (x *streamConnectServer) Recv() (*Message, error) {
        m := new(Message)
        if err := x.ServerStream.RecvMsg(m); err != nil {
                return nil, err
        }
        return m, nil
}

var _Stream_serviceDesc = grpc.ServiceDesc{
        ServiceName: "events.Stream",
        HandlerType: (*StreamServer)(nil),
        Methods:     []grpc.MethodDesc{},
        Streams: []grpc.StreamDesc{
                {
                        StreamName:    "Connect",
                        Handler:       _Stream_Connect_Handler,
                        ServerStreams: true,
                        ClientStreams: true,
                },
        },
        Metadata: "events.proto",
}
