// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: AvatarEquipAffixStartNotify.proto

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

type AvatarEquipAffixStartNotify_CmdId int32

const (
	AvatarEquipAffixStartNotify_NONE             AvatarEquipAffixStartNotify_CmdId = 0
	AvatarEquipAffixStartNotify_CMD_ID           AvatarEquipAffixStartNotify_CmdId = 1747
	AvatarEquipAffixStartNotify_ENET_CHANNEL_ID  AvatarEquipAffixStartNotify_CmdId = 0
	AvatarEquipAffixStartNotify_ENET_IS_RELIABLE AvatarEquipAffixStartNotify_CmdId = 1
)

// Enum value maps for AvatarEquipAffixStartNotify_CmdId.
var (
	AvatarEquipAffixStartNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		1747: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	AvatarEquipAffixStartNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           1747,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x AvatarEquipAffixStartNotify_CmdId) Enum() *AvatarEquipAffixStartNotify_CmdId {
	p := new(AvatarEquipAffixStartNotify_CmdId)
	*p = x
	return p
}

func (x AvatarEquipAffixStartNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AvatarEquipAffixStartNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_AvatarEquipAffixStartNotify_proto_enumTypes[0].Descriptor()
}

func (AvatarEquipAffixStartNotify_CmdId) Type() protoreflect.EnumType {
	return &file_AvatarEquipAffixStartNotify_proto_enumTypes[0]
}

func (x AvatarEquipAffixStartNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AvatarEquipAffixStartNotify_CmdId.Descriptor instead.
func (AvatarEquipAffixStartNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_AvatarEquipAffixStartNotify_proto_rawDescGZIP(), []int{0, 0}
}

type AvatarEquipAffixStartNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarGuid     uint64                `protobuf:"varint,1,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
	EquipAffixInfo *AvatarEquipAffixInfo `protobuf:"bytes,2,opt,name=equip_affix_info,json=equipAffixInfo,proto3" json:"equip_affix_info,omitempty"`
}

func (x *AvatarEquipAffixStartNotify) Reset() {
	*x = AvatarEquipAffixStartNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarEquipAffixStartNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarEquipAffixStartNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarEquipAffixStartNotify) ProtoMessage() {}

func (x *AvatarEquipAffixStartNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarEquipAffixStartNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarEquipAffixStartNotify.ProtoReflect.Descriptor instead.
func (*AvatarEquipAffixStartNotify) Descriptor() ([]byte, []int) {
	return file_AvatarEquipAffixStartNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarEquipAffixStartNotify) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

func (x *AvatarEquipAffixStartNotify) GetEquipAffixInfo() *AvatarEquipAffixInfo {
	if x != nil {
		return x.EquipAffixInfo
	}
	return nil
}

var File_AvatarEquipAffixStartNotify_proto protoreflect.FileDescriptor

var file_AvatarEquipAffixStartNotify_proto_rawDesc = []byte{
	0x0a, 0x21, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x71, 0x75, 0x69, 0x70, 0x41, 0x66, 0x66,
	0x69, 0x78, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x41, 0x66, 0x66, 0x69, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xce, 0x01, 0x0a, 0x1b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x71, 0x75, 0x69, 0x70, 0x41,
	0x66, 0x66, 0x69, 0x78, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x47, 0x75, 0x69, 0x64,
	0x12, 0x3f, 0x0a, 0x10, 0x65, 0x71, 0x75, 0x69, 0x70, 0x5f, 0x61, 0x66, 0x66, 0x69, 0x78, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x45, 0x71, 0x75, 0x69, 0x70, 0x41, 0x66, 0x66, 0x69, 0x78, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x41, 0x66, 0x66, 0x69, 0x78, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xd3,
	0x0d, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45,
	0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49,
	0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarEquipAffixStartNotify_proto_rawDescOnce sync.Once
	file_AvatarEquipAffixStartNotify_proto_rawDescData = file_AvatarEquipAffixStartNotify_proto_rawDesc
)

func file_AvatarEquipAffixStartNotify_proto_rawDescGZIP() []byte {
	file_AvatarEquipAffixStartNotify_proto_rawDescOnce.Do(func() {
		file_AvatarEquipAffixStartNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarEquipAffixStartNotify_proto_rawDescData)
	})
	return file_AvatarEquipAffixStartNotify_proto_rawDescData
}

var file_AvatarEquipAffixStartNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_AvatarEquipAffixStartNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarEquipAffixStartNotify_proto_goTypes = []interface{}{
	(AvatarEquipAffixStartNotify_CmdId)(0), // 0: AvatarEquipAffixStartNotify.CmdId
	(*AvatarEquipAffixStartNotify)(nil),    // 1: AvatarEquipAffixStartNotify
	(*AvatarEquipAffixInfo)(nil),           // 2: AvatarEquipAffixInfo
}
var file_AvatarEquipAffixStartNotify_proto_depIdxs = []int32{
	2, // 0: AvatarEquipAffixStartNotify.equip_affix_info:type_name -> AvatarEquipAffixInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AvatarEquipAffixStartNotify_proto_init() }
func file_AvatarEquipAffixStartNotify_proto_init() {
	if File_AvatarEquipAffixStartNotify_proto != nil {
		return
	}
	file_AvatarEquipAffixInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AvatarEquipAffixStartNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarEquipAffixStartNotify); i {
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
			RawDescriptor: file_AvatarEquipAffixStartNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarEquipAffixStartNotify_proto_goTypes,
		DependencyIndexes: file_AvatarEquipAffixStartNotify_proto_depIdxs,
		EnumInfos:         file_AvatarEquipAffixStartNotify_proto_enumTypes,
		MessageInfos:      file_AvatarEquipAffixStartNotify_proto_msgTypes,
	}.Build()
	File_AvatarEquipAffixStartNotify_proto = out.File
	file_AvatarEquipAffixStartNotify_proto_rawDesc = nil
	file_AvatarEquipAffixStartNotify_proto_goTypes = nil
	file_AvatarEquipAffixStartNotify_proto_depIdxs = nil
}
