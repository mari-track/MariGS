// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: WatcherEventTypeNotify.proto

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

type WatcherEventTypeNotify_CmdId int32

const (
	WatcherEventTypeNotify_NONE             WatcherEventTypeNotify_CmdId = 0
	WatcherEventTypeNotify_CMD_ID           WatcherEventTypeNotify_CmdId = 2204
	WatcherEventTypeNotify_ENET_CHANNEL_ID  WatcherEventTypeNotify_CmdId = 0
	WatcherEventTypeNotify_ENET_IS_RELIABLE WatcherEventTypeNotify_CmdId = 1
	WatcherEventTypeNotify_IS_ALLOW_CLIENT  WatcherEventTypeNotify_CmdId = 1
)

// Enum value maps for WatcherEventTypeNotify_CmdId.
var (
	WatcherEventTypeNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		2204: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	WatcherEventTypeNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2204,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x WatcherEventTypeNotify_CmdId) Enum() *WatcherEventTypeNotify_CmdId {
	p := new(WatcherEventTypeNotify_CmdId)
	*p = x
	return p
}

func (x WatcherEventTypeNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WatcherEventTypeNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_WatcherEventTypeNotify_proto_enumTypes[0].Descriptor()
}

func (WatcherEventTypeNotify_CmdId) Type() protoreflect.EnumType {
	return &file_WatcherEventTypeNotify_proto_enumTypes[0]
}

func (x WatcherEventTypeNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WatcherEventTypeNotify_CmdId.Descriptor instead.
func (WatcherEventTypeNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_WatcherEventTypeNotify_proto_rawDescGZIP(), []int{0, 0}
}

type WatcherEventTypeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WatcherTriggerType uint32   `protobuf:"varint,1,opt,name=watcher_trigger_type,json=watcherTriggerType,proto3" json:"watcher_trigger_type,omitempty"`
	ParamList          []uint32 `protobuf:"varint,2,rep,packed,name=param_list,json=paramList,proto3" json:"param_list,omitempty"`
	AddProgress        uint32   `protobuf:"varint,3,opt,name=add_progress,json=addProgress,proto3" json:"add_progress,omitempty"`
}

func (x *WatcherEventTypeNotify) Reset() {
	*x = WatcherEventTypeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WatcherEventTypeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatcherEventTypeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatcherEventTypeNotify) ProtoMessage() {}

func (x *WatcherEventTypeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_WatcherEventTypeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatcherEventTypeNotify.ProtoReflect.Descriptor instead.
func (*WatcherEventTypeNotify) Descriptor() ([]byte, []int) {
	return file_WatcherEventTypeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *WatcherEventTypeNotify) GetWatcherTriggerType() uint32 {
	if x != nil {
		return x.WatcherTriggerType
	}
	return 0
}

func (x *WatcherEventTypeNotify) GetParamList() []uint32 {
	if x != nil {
		return x.ParamList
	}
	return nil
}

func (x *WatcherEventTypeNotify) GetAddProgress() uint32 {
	if x != nil {
		return x.AddProgress
	}
	return 0
}

var File_WatcherEventTypeNotify_proto protoreflect.FileDescriptor

var file_WatcherEventTypeNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0,
	0x01, 0x0a, 0x16, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x30, 0x0a, 0x14, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x64,
	0x64, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x61, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x62, 0x0a,
	0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x9c, 0x11, 0x12, 0x13, 0x0a,
	0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45,
	0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41,
	0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WatcherEventTypeNotify_proto_rawDescOnce sync.Once
	file_WatcherEventTypeNotify_proto_rawDescData = file_WatcherEventTypeNotify_proto_rawDesc
)

func file_WatcherEventTypeNotify_proto_rawDescGZIP() []byte {
	file_WatcherEventTypeNotify_proto_rawDescOnce.Do(func() {
		file_WatcherEventTypeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_WatcherEventTypeNotify_proto_rawDescData)
	})
	return file_WatcherEventTypeNotify_proto_rawDescData
}

var file_WatcherEventTypeNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_WatcherEventTypeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WatcherEventTypeNotify_proto_goTypes = []interface{}{
	(WatcherEventTypeNotify_CmdId)(0), // 0: WatcherEventTypeNotify.CmdId
	(*WatcherEventTypeNotify)(nil),    // 1: WatcherEventTypeNotify
}
var file_WatcherEventTypeNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_WatcherEventTypeNotify_proto_init() }
func file_WatcherEventTypeNotify_proto_init() {
	if File_WatcherEventTypeNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_WatcherEventTypeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatcherEventTypeNotify); i {
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
			RawDescriptor: file_WatcherEventTypeNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WatcherEventTypeNotify_proto_goTypes,
		DependencyIndexes: file_WatcherEventTypeNotify_proto_depIdxs,
		EnumInfos:         file_WatcherEventTypeNotify_proto_enumTypes,
		MessageInfos:      file_WatcherEventTypeNotify_proto_msgTypes,
	}.Build()
	File_WatcherEventTypeNotify_proto = out.File
	file_WatcherEventTypeNotify_proto_rawDesc = nil
	file_WatcherEventTypeNotify_proto_goTypes = nil
	file_WatcherEventTypeNotify_proto_depIdxs = nil
}
