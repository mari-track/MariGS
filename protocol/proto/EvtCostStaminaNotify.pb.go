// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: EvtCostStaminaNotify.proto

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

type EvtCostStaminaNotify_CmdId int32

const (
	EvtCostStaminaNotify_NONE             EvtCostStaminaNotify_CmdId = 0
	EvtCostStaminaNotify_CMD_ID           EvtCostStaminaNotify_CmdId = 309
	EvtCostStaminaNotify_ENET_CHANNEL_ID  EvtCostStaminaNotify_CmdId = 0
	EvtCostStaminaNotify_ENET_IS_RELIABLE EvtCostStaminaNotify_CmdId = 1
	EvtCostStaminaNotify_IS_ALLOW_CLIENT  EvtCostStaminaNotify_CmdId = 1
)

// Enum value maps for EvtCostStaminaNotify_CmdId.
var (
	EvtCostStaminaNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		309: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	EvtCostStaminaNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           309,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x EvtCostStaminaNotify_CmdId) Enum() *EvtCostStaminaNotify_CmdId {
	p := new(EvtCostStaminaNotify_CmdId)
	*p = x
	return p
}

func (x EvtCostStaminaNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EvtCostStaminaNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_EvtCostStaminaNotify_proto_enumTypes[0].Descriptor()
}

func (EvtCostStaminaNotify_CmdId) Type() protoreflect.EnumType {
	return &file_EvtCostStaminaNotify_proto_enumTypes[0]
}

func (x EvtCostStaminaNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EvtCostStaminaNotify_CmdId.Descriptor instead.
func (EvtCostStaminaNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_EvtCostStaminaNotify_proto_rawDescGZIP(), []int{0, 0}
}

type EvtCostStaminaNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CostStamina float32 `protobuf:"fixed32,1,opt,name=cost_stamina,json=costStamina,proto3" json:"cost_stamina,omitempty"`
	SkillId     uint32  `protobuf:"varint,2,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
}

func (x *EvtCostStaminaNotify) Reset() {
	*x = EvtCostStaminaNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtCostStaminaNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtCostStaminaNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtCostStaminaNotify) ProtoMessage() {}

func (x *EvtCostStaminaNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EvtCostStaminaNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtCostStaminaNotify.ProtoReflect.Descriptor instead.
func (*EvtCostStaminaNotify) Descriptor() ([]byte, []int) {
	return file_EvtCostStaminaNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EvtCostStaminaNotify) GetCostStamina() float32 {
	if x != nil {
		return x.CostStamina
	}
	return 0
}

func (x *EvtCostStaminaNotify) GetSkillId() uint32 {
	if x != nil {
		return x.SkillId
	}
	return 0
}

var File_EvtCostStaminaNotify_proto protoreflect.FileDescriptor

var file_EvtCostStaminaNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x45, 0x76, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a,
	0x14, 0x45, 0x76, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x6d, 0x69, 0x6e, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x6f, 0x73,
	0x74, 0x53, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x49, 0x64, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44,
	0x10, 0xb5, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e,
	0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EvtCostStaminaNotify_proto_rawDescOnce sync.Once
	file_EvtCostStaminaNotify_proto_rawDescData = file_EvtCostStaminaNotify_proto_rawDesc
)

func file_EvtCostStaminaNotify_proto_rawDescGZIP() []byte {
	file_EvtCostStaminaNotify_proto_rawDescOnce.Do(func() {
		file_EvtCostStaminaNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtCostStaminaNotify_proto_rawDescData)
	})
	return file_EvtCostStaminaNotify_proto_rawDescData
}

var file_EvtCostStaminaNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_EvtCostStaminaNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvtCostStaminaNotify_proto_goTypes = []interface{}{
	(EvtCostStaminaNotify_CmdId)(0), // 0: EvtCostStaminaNotify.CmdId
	(*EvtCostStaminaNotify)(nil),    // 1: EvtCostStaminaNotify
}
var file_EvtCostStaminaNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EvtCostStaminaNotify_proto_init() }
func file_EvtCostStaminaNotify_proto_init() {
	if File_EvtCostStaminaNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EvtCostStaminaNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtCostStaminaNotify); i {
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
			RawDescriptor: file_EvtCostStaminaNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtCostStaminaNotify_proto_goTypes,
		DependencyIndexes: file_EvtCostStaminaNotify_proto_depIdxs,
		EnumInfos:         file_EvtCostStaminaNotify_proto_enumTypes,
		MessageInfos:      file_EvtCostStaminaNotify_proto_msgTypes,
	}.Build()
	File_EvtCostStaminaNotify_proto = out.File
	file_EvtCostStaminaNotify_proto_rawDesc = nil
	file_EvtCostStaminaNotify_proto_goTypes = nil
	file_EvtCostStaminaNotify_proto_depIdxs = nil
}
