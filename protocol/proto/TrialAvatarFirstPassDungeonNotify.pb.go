// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: TrialAvatarFirstPassDungeonNotify.proto

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

type TrialAvatarFirstPassDungeonNotify_CmdId int32

const (
	TrialAvatarFirstPassDungeonNotify_NONE             TrialAvatarFirstPassDungeonNotify_CmdId = 0
	TrialAvatarFirstPassDungeonNotify_CMD_ID           TrialAvatarFirstPassDungeonNotify_CmdId = 2047
	TrialAvatarFirstPassDungeonNotify_ENET_CHANNEL_ID  TrialAvatarFirstPassDungeonNotify_CmdId = 0
	TrialAvatarFirstPassDungeonNotify_ENET_IS_RELIABLE TrialAvatarFirstPassDungeonNotify_CmdId = 1
)

// Enum value maps for TrialAvatarFirstPassDungeonNotify_CmdId.
var (
	TrialAvatarFirstPassDungeonNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		2047: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	TrialAvatarFirstPassDungeonNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2047,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x TrialAvatarFirstPassDungeonNotify_CmdId) Enum() *TrialAvatarFirstPassDungeonNotify_CmdId {
	p := new(TrialAvatarFirstPassDungeonNotify_CmdId)
	*p = x
	return p
}

func (x TrialAvatarFirstPassDungeonNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrialAvatarFirstPassDungeonNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_TrialAvatarFirstPassDungeonNotify_proto_enumTypes[0].Descriptor()
}

func (TrialAvatarFirstPassDungeonNotify_CmdId) Type() protoreflect.EnumType {
	return &file_TrialAvatarFirstPassDungeonNotify_proto_enumTypes[0]
}

func (x TrialAvatarFirstPassDungeonNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrialAvatarFirstPassDungeonNotify_CmdId.Descriptor instead.
func (TrialAvatarFirstPassDungeonNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_TrialAvatarFirstPassDungeonNotify_proto_rawDescGZIP(), []int{0, 0}
}

type TrialAvatarFirstPassDungeonNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrialAvatarIndexId uint32 `protobuf:"varint,1,opt,name=trial_avatar_index_id,json=trialAvatarIndexId,proto3" json:"trial_avatar_index_id,omitempty"`
}

func (x *TrialAvatarFirstPassDungeonNotify) Reset() {
	*x = TrialAvatarFirstPassDungeonNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TrialAvatarFirstPassDungeonNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrialAvatarFirstPassDungeonNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrialAvatarFirstPassDungeonNotify) ProtoMessage() {}

func (x *TrialAvatarFirstPassDungeonNotify) ProtoReflect() protoreflect.Message {
	mi := &file_TrialAvatarFirstPassDungeonNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrialAvatarFirstPassDungeonNotify.ProtoReflect.Descriptor instead.
func (*TrialAvatarFirstPassDungeonNotify) Descriptor() ([]byte, []int) {
	return file_TrialAvatarFirstPassDungeonNotify_proto_rawDescGZIP(), []int{0}
}

func (x *TrialAvatarFirstPassDungeonNotify) GetTrialAvatarIndexId() uint32 {
	if x != nil {
		return x.TrialAvatarIndexId
	}
	return 0
}

var File_TrialAvatarFirstPassDungeonNotify_proto protoreflect.FileDescriptor

var file_TrialAvatarFirstPassDungeonNotify_proto_rawDesc = []byte{
	0x0a, 0x27, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa5, 0x01, 0x0a, 0x21, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x31, 0x0a, 0x15, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64,
	0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06,
	0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xff, 0x0f, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45,
	0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TrialAvatarFirstPassDungeonNotify_proto_rawDescOnce sync.Once
	file_TrialAvatarFirstPassDungeonNotify_proto_rawDescData = file_TrialAvatarFirstPassDungeonNotify_proto_rawDesc
)

func file_TrialAvatarFirstPassDungeonNotify_proto_rawDescGZIP() []byte {
	file_TrialAvatarFirstPassDungeonNotify_proto_rawDescOnce.Do(func() {
		file_TrialAvatarFirstPassDungeonNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_TrialAvatarFirstPassDungeonNotify_proto_rawDescData)
	})
	return file_TrialAvatarFirstPassDungeonNotify_proto_rawDescData
}

var file_TrialAvatarFirstPassDungeonNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_TrialAvatarFirstPassDungeonNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TrialAvatarFirstPassDungeonNotify_proto_goTypes = []interface{}{
	(TrialAvatarFirstPassDungeonNotify_CmdId)(0), // 0: proto.TrialAvatarFirstPassDungeonNotify.CmdId
	(*TrialAvatarFirstPassDungeonNotify)(nil),    // 1: proto.TrialAvatarFirstPassDungeonNotify
}
var file_TrialAvatarFirstPassDungeonNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TrialAvatarFirstPassDungeonNotify_proto_init() }
func file_TrialAvatarFirstPassDungeonNotify_proto_init() {
	if File_TrialAvatarFirstPassDungeonNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TrialAvatarFirstPassDungeonNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrialAvatarFirstPassDungeonNotify); i {
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
			RawDescriptor: file_TrialAvatarFirstPassDungeonNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TrialAvatarFirstPassDungeonNotify_proto_goTypes,
		DependencyIndexes: file_TrialAvatarFirstPassDungeonNotify_proto_depIdxs,
		EnumInfos:         file_TrialAvatarFirstPassDungeonNotify_proto_enumTypes,
		MessageInfos:      file_TrialAvatarFirstPassDungeonNotify_proto_msgTypes,
	}.Build()
	File_TrialAvatarFirstPassDungeonNotify_proto = out.File
	file_TrialAvatarFirstPassDungeonNotify_proto_rawDesc = nil
	file_TrialAvatarFirstPassDungeonNotify_proto_goTypes = nil
	file_TrialAvatarFirstPassDungeonNotify_proto_depIdxs = nil
}
