// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: ServerTimeNotify.proto

package proto

import (
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

type ServerTimeNotify_CmdId int32

const (
	ServerTimeNotify_NONE             ServerTimeNotify_CmdId = 0
	ServerTimeNotify_CMD_ID           ServerTimeNotify_CmdId = 10
	ServerTimeNotify_ENET_CHANNEL_ID  ServerTimeNotify_CmdId = 1
	ServerTimeNotify_ENET_IS_RELIABLE ServerTimeNotify_CmdId = 1
)

// Enum value maps for ServerTimeNotify_CmdId.
var (
	ServerTimeNotify_CmdId_name = map[int32]string{
		0:  "NONE",
		10: "CMD_ID",
		1:  "ENET_CHANNEL_ID",
		// Duplicate value: 1: "ENET_IS_RELIABLE",
	}
	ServerTimeNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           10,
		"ENET_CHANNEL_ID":  1,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x ServerTimeNotify_CmdId) Enum() *ServerTimeNotify_CmdId {
	p := new(ServerTimeNotify_CmdId)
	*p = x
	return p
}

func (x ServerTimeNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServerTimeNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_ServerTimeNotify_proto_enumTypes[0].Descriptor()
}

func (ServerTimeNotify_CmdId) Type() protoreflect.EnumType {
	return &file_ServerTimeNotify_proto_enumTypes[0]
}

func (x ServerTimeNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServerTimeNotify_CmdId.Descriptor instead.
func (ServerTimeNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_ServerTimeNotify_proto_rawDescGZIP(), []int{0, 0}
}

type ServerTimeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerTime uint64 `protobuf:"varint,1,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`
}

func (x *ServerTimeNotify) Reset() {
	*x = ServerTimeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ServerTimeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerTimeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerTimeNotify) ProtoMessage() {}

func (x *ServerTimeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ServerTimeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerTimeNotify.ProtoReflect.Descriptor instead.
func (*ServerTimeNotify) Descriptor() ([]byte, []int) {
	return file_ServerTimeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ServerTimeNotify) GetServerTime() uint64 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

var File_ServerTimeNotify_proto protoreflect.FileDescriptor

var file_ServerTimeNotify_proto_rawDesc = []byte{
	0x0a, 0x16, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4c,
	0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x0a, 0x12, 0x13, 0x0a,
	0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44,
	0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45,
	0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ServerTimeNotify_proto_rawDescOnce sync.Once
	file_ServerTimeNotify_proto_rawDescData = file_ServerTimeNotify_proto_rawDesc
)

func file_ServerTimeNotify_proto_rawDescGZIP() []byte {
	file_ServerTimeNotify_proto_rawDescOnce.Do(func() {
		file_ServerTimeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ServerTimeNotify_proto_rawDescData)
	})
	return file_ServerTimeNotify_proto_rawDescData
}

var file_ServerTimeNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ServerTimeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ServerTimeNotify_proto_goTypes = []interface{}{
	(ServerTimeNotify_CmdId)(0), // 0: ServerTimeNotify.CmdId
	(*ServerTimeNotify)(nil),    // 1: ServerTimeNotify
}
var file_ServerTimeNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ServerTimeNotify_proto_init() }
func file_ServerTimeNotify_proto_init() {
	if File_ServerTimeNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ServerTimeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerTimeNotify); i {
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
			RawDescriptor: file_ServerTimeNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ServerTimeNotify_proto_goTypes,
		DependencyIndexes: file_ServerTimeNotify_proto_depIdxs,
		EnumInfos:         file_ServerTimeNotify_proto_enumTypes,
		MessageInfos:      file_ServerTimeNotify_proto_msgTypes,
	}.Build()
	File_ServerTimeNotify_proto = out.File
	file_ServerTimeNotify_proto_rawDesc = nil
	file_ServerTimeNotify_proto_goTypes = nil
	file_ServerTimeNotify_proto_depIdxs = nil
}
