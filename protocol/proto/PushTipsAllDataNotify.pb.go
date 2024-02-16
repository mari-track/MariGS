// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: PushTipsAllDataNotify.proto

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

type PushTipsAllDataNotify_CmdId int32

const (
	PushTipsAllDataNotify_NONE             PushTipsAllDataNotify_CmdId = 0
	PushTipsAllDataNotify_CMD_ID           PushTipsAllDataNotify_CmdId = 2221
	PushTipsAllDataNotify_ENET_CHANNEL_ID  PushTipsAllDataNotify_CmdId = 0
	PushTipsAllDataNotify_ENET_IS_RELIABLE PushTipsAllDataNotify_CmdId = 1
)

// Enum value maps for PushTipsAllDataNotify_CmdId.
var (
	PushTipsAllDataNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		2221: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	PushTipsAllDataNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2221,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x PushTipsAllDataNotify_CmdId) Enum() *PushTipsAllDataNotify_CmdId {
	p := new(PushTipsAllDataNotify_CmdId)
	*p = x
	return p
}

func (x PushTipsAllDataNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PushTipsAllDataNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PushTipsAllDataNotify_proto_enumTypes[0].Descriptor()
}

func (PushTipsAllDataNotify_CmdId) Type() protoreflect.EnumType {
	return &file_PushTipsAllDataNotify_proto_enumTypes[0]
}

func (x PushTipsAllDataNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PushTipsAllDataNotify_CmdId.Descriptor instead.
func (PushTipsAllDataNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PushTipsAllDataNotify_proto_rawDescGZIP(), []int{0, 0}
}

type PushTipsAllDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PushTipsList []*PushTipsData `protobuf:"bytes,1,rep,name=push_tips_list,json=pushTipsList,proto3" json:"push_tips_list,omitempty"`
}

func (x *PushTipsAllDataNotify) Reset() {
	*x = PushTipsAllDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PushTipsAllDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushTipsAllDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushTipsAllDataNotify) ProtoMessage() {}

func (x *PushTipsAllDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PushTipsAllDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushTipsAllDataNotify.ProtoReflect.Descriptor instead.
func (*PushTipsAllDataNotify) Descriptor() ([]byte, []int) {
	return file_PushTipsAllDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PushTipsAllDataNotify) GetPushTipsList() []*PushTipsData {
	if x != nil {
		return x.PushTipsList
	}
	return nil
}

var File_PushTipsAllDataNotify_proto protoreflect.FileDescriptor

var file_PushTipsAllDataNotify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x50, 0x75, 0x73, 0x68, 0x54, 0x69, 0x70, 0x73, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74,
	0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x50, 0x75, 0x73, 0x68, 0x54, 0x69, 0x70, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x15, 0x50, 0x75, 0x73,
	0x68, 0x54, 0x69, 0x70, 0x73, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x39, 0x0a, 0x0e, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x70, 0x73, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x54, 0x69, 0x70, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x0c, 0x70, 0x75, 0x73, 0x68, 0x54, 0x69, 0x70, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d, 0x0a,
	0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xad, 0x11, 0x12, 0x13, 0x0a,
	0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45,
	0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PushTipsAllDataNotify_proto_rawDescOnce sync.Once
	file_PushTipsAllDataNotify_proto_rawDescData = file_PushTipsAllDataNotify_proto_rawDesc
)

func file_PushTipsAllDataNotify_proto_rawDescGZIP() []byte {
	file_PushTipsAllDataNotify_proto_rawDescOnce.Do(func() {
		file_PushTipsAllDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PushTipsAllDataNotify_proto_rawDescData)
	})
	return file_PushTipsAllDataNotify_proto_rawDescData
}

var file_PushTipsAllDataNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PushTipsAllDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PushTipsAllDataNotify_proto_goTypes = []interface{}{
	(PushTipsAllDataNotify_CmdId)(0), // 0: proto.PushTipsAllDataNotify.CmdId
	(*PushTipsAllDataNotify)(nil),    // 1: proto.PushTipsAllDataNotify
	(*PushTipsData)(nil),             // 2: proto.PushTipsData
}
var file_PushTipsAllDataNotify_proto_depIdxs = []int32{
	2, // 0: proto.PushTipsAllDataNotify.push_tips_list:type_name -> proto.PushTipsData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PushTipsAllDataNotify_proto_init() }
func file_PushTipsAllDataNotify_proto_init() {
	if File_PushTipsAllDataNotify_proto != nil {
		return
	}
	file_PushTipsData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PushTipsAllDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushTipsAllDataNotify); i {
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
			RawDescriptor: file_PushTipsAllDataNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PushTipsAllDataNotify_proto_goTypes,
		DependencyIndexes: file_PushTipsAllDataNotify_proto_depIdxs,
		EnumInfos:         file_PushTipsAllDataNotify_proto_enumTypes,
		MessageInfos:      file_PushTipsAllDataNotify_proto_msgTypes,
	}.Build()
	File_PushTipsAllDataNotify_proto = out.File
	file_PushTipsAllDataNotify_proto_rawDesc = nil
	file_PushTipsAllDataNotify_proto_goTypes = nil
	file_PushTipsAllDataNotify_proto_depIdxs = nil
}
