// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: CanUseSkillNotify.proto

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

type CanUseSkillNotify_CmdId int32

const (
	CanUseSkillNotify_NONE             CanUseSkillNotify_CmdId = 0
	CanUseSkillNotify_CMD_ID           CanUseSkillNotify_CmdId = 1016
	CanUseSkillNotify_ENET_CHANNEL_ID  CanUseSkillNotify_CmdId = 0
	CanUseSkillNotify_ENET_IS_RELIABLE CanUseSkillNotify_CmdId = 1
)

// Enum value maps for CanUseSkillNotify_CmdId.
var (
	CanUseSkillNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		1016: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	CanUseSkillNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           1016,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x CanUseSkillNotify_CmdId) Enum() *CanUseSkillNotify_CmdId {
	p := new(CanUseSkillNotify_CmdId)
	*p = x
	return p
}

func (x CanUseSkillNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CanUseSkillNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_CanUseSkillNotify_proto_enumTypes[0].Descriptor()
}

func (CanUseSkillNotify_CmdId) Type() protoreflect.EnumType {
	return &file_CanUseSkillNotify_proto_enumTypes[0]
}

func (x CanUseSkillNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CanUseSkillNotify_CmdId.Descriptor instead.
func (CanUseSkillNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_CanUseSkillNotify_proto_rawDescGZIP(), []int{0, 0}
}

type CanUseSkillNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsCanUseSkill bool `protobuf:"varint,1,opt,name=is_can_use_skill,json=isCanUseSkill,proto3" json:"is_can_use_skill,omitempty"`
}

func (x *CanUseSkillNotify) Reset() {
	*x = CanUseSkillNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CanUseSkillNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanUseSkillNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanUseSkillNotify) ProtoMessage() {}

func (x *CanUseSkillNotify) ProtoReflect() protoreflect.Message {
	mi := &file_CanUseSkillNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanUseSkillNotify.ProtoReflect.Descriptor instead.
func (*CanUseSkillNotify) Descriptor() ([]byte, []int) {
	return file_CanUseSkillNotify_proto_rawDescGZIP(), []int{0}
}

func (x *CanUseSkillNotify) GetIsCanUseSkill() bool {
	if x != nil {
		return x.IsCanUseSkill
	}
	return false
}

var File_CanUseSkillNotify_proto protoreflect.FileDescriptor

var file_CanUseSkillNotify_proto_rawDesc = []byte{
	0x0a, 0x17, 0x43, 0x61, 0x6e, 0x55, 0x73, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8b, 0x01, 0x0a, 0x11, 0x43, 0x61, 0x6e, 0x55, 0x73, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x27, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x63, 0x61, 0x6e,
	0x5f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x69, 0x73, 0x43, 0x61, 0x6e, 0x55, 0x73, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x22,
	0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xf8, 0x07, 0x12,
	0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f,
	0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_CanUseSkillNotify_proto_rawDescOnce sync.Once
	file_CanUseSkillNotify_proto_rawDescData = file_CanUseSkillNotify_proto_rawDesc
)

func file_CanUseSkillNotify_proto_rawDescGZIP() []byte {
	file_CanUseSkillNotify_proto_rawDescOnce.Do(func() {
		file_CanUseSkillNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_CanUseSkillNotify_proto_rawDescData)
	})
	return file_CanUseSkillNotify_proto_rawDescData
}

var file_CanUseSkillNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CanUseSkillNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CanUseSkillNotify_proto_goTypes = []interface{}{
	(CanUseSkillNotify_CmdId)(0), // 0: proto.CanUseSkillNotify.CmdId
	(*CanUseSkillNotify)(nil),    // 1: proto.CanUseSkillNotify
}
var file_CanUseSkillNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CanUseSkillNotify_proto_init() }
func file_CanUseSkillNotify_proto_init() {
	if File_CanUseSkillNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CanUseSkillNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CanUseSkillNotify); i {
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
			RawDescriptor: file_CanUseSkillNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CanUseSkillNotify_proto_goTypes,
		DependencyIndexes: file_CanUseSkillNotify_proto_depIdxs,
		EnumInfos:         file_CanUseSkillNotify_proto_enumTypes,
		MessageInfos:      file_CanUseSkillNotify_proto_msgTypes,
	}.Build()
	File_CanUseSkillNotify_proto = out.File
	file_CanUseSkillNotify_proto_rawDesc = nil
	file_CanUseSkillNotify_proto_goTypes = nil
	file_CanUseSkillNotify_proto_depIdxs = nil
}
