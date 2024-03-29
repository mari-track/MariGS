// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: WatcherEventNotify.proto

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

type WatcherEventNotify_CmdId int32

const (
	WatcherEventNotify_NONE             WatcherEventNotify_CmdId = 0
	WatcherEventNotify_CMD_ID           WatcherEventNotify_CmdId = 2203
	WatcherEventNotify_ENET_CHANNEL_ID  WatcherEventNotify_CmdId = 0
	WatcherEventNotify_ENET_IS_RELIABLE WatcherEventNotify_CmdId = 1
	WatcherEventNotify_IS_ALLOW_CLIENT  WatcherEventNotify_CmdId = 1
)

// Enum value maps for WatcherEventNotify_CmdId.
var (
	WatcherEventNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		2203: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	WatcherEventNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2203,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x WatcherEventNotify_CmdId) Enum() *WatcherEventNotify_CmdId {
	p := new(WatcherEventNotify_CmdId)
	*p = x
	return p
}

func (x WatcherEventNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WatcherEventNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_WatcherEventNotify_proto_enumTypes[0].Descriptor()
}

func (WatcherEventNotify_CmdId) Type() protoreflect.EnumType {
	return &file_WatcherEventNotify_proto_enumTypes[0]
}

func (x WatcherEventNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WatcherEventNotify_CmdId.Descriptor instead.
func (WatcherEventNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_WatcherEventNotify_proto_rawDescGZIP(), []int{0, 0}
}

type WatcherEventNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WatcherId   uint32 `protobuf:"varint,1,opt,name=watcher_id,json=watcherId,proto3" json:"watcher_id,omitempty"`
	AddProgress uint32 `protobuf:"varint,2,opt,name=add_progress,json=addProgress,proto3" json:"add_progress,omitempty"`
}

func (x *WatcherEventNotify) Reset() {
	*x = WatcherEventNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WatcherEventNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatcherEventNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatcherEventNotify) ProtoMessage() {}

func (x *WatcherEventNotify) ProtoReflect() protoreflect.Message {
	mi := &file_WatcherEventNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatcherEventNotify.ProtoReflect.Descriptor instead.
func (*WatcherEventNotify) Descriptor() ([]byte, []int) {
	return file_WatcherEventNotify_proto_rawDescGZIP(), []int{0}
}

func (x *WatcherEventNotify) GetWatcherId() uint32 {
	if x != nil {
		return x.WatcherId
	}
	return 0
}

func (x *WatcherEventNotify) GetAddProgress() uint32 {
	if x != nil {
		return x.AddProgress
	}
	return 0
}

var File_WatcherEventNotify_proto protoreflect.FileDescriptor

var file_WatcherEventNotify_proto_rawDesc = []byte{
	0x0a, 0x18, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x12, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44,
	0x10, 0x9b, 0x11, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e,
	0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WatcherEventNotify_proto_rawDescOnce sync.Once
	file_WatcherEventNotify_proto_rawDescData = file_WatcherEventNotify_proto_rawDesc
)

func file_WatcherEventNotify_proto_rawDescGZIP() []byte {
	file_WatcherEventNotify_proto_rawDescOnce.Do(func() {
		file_WatcherEventNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_WatcherEventNotify_proto_rawDescData)
	})
	return file_WatcherEventNotify_proto_rawDescData
}

var file_WatcherEventNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_WatcherEventNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WatcherEventNotify_proto_goTypes = []interface{}{
	(WatcherEventNotify_CmdId)(0), // 0: WatcherEventNotify.CmdId
	(*WatcherEventNotify)(nil),    // 1: WatcherEventNotify
}
var file_WatcherEventNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_WatcherEventNotify_proto_init() }
func file_WatcherEventNotify_proto_init() {
	if File_WatcherEventNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_WatcherEventNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatcherEventNotify); i {
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
			RawDescriptor: file_WatcherEventNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WatcherEventNotify_proto_goTypes,
		DependencyIndexes: file_WatcherEventNotify_proto_depIdxs,
		EnumInfos:         file_WatcherEventNotify_proto_enumTypes,
		MessageInfos:      file_WatcherEventNotify_proto_msgTypes,
	}.Build()
	File_WatcherEventNotify_proto = out.File
	file_WatcherEventNotify_proto_rawDesc = nil
	file_WatcherEventNotify_proto_goTypes = nil
	file_WatcherEventNotify_proto_depIdxs = nil
}
