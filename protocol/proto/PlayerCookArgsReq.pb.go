// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: PlayerCookArgsReq.proto

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

type PlayerCookArgsReq_CmdId int32

const (
	PlayerCookArgsReq_NONE             PlayerCookArgsReq_CmdId = 0
	PlayerCookArgsReq_CMD_ID           PlayerCookArgsReq_CmdId = 159
	PlayerCookArgsReq_ENET_CHANNEL_ID  PlayerCookArgsReq_CmdId = 0
	PlayerCookArgsReq_ENET_IS_RELIABLE PlayerCookArgsReq_CmdId = 1
	PlayerCookArgsReq_IS_ALLOW_CLIENT  PlayerCookArgsReq_CmdId = 1
)

// Enum value maps for PlayerCookArgsReq_CmdId.
var (
	PlayerCookArgsReq_CmdId_name = map[int32]string{
		0:   "NONE",
		159: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	PlayerCookArgsReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           159,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x PlayerCookArgsReq_CmdId) Enum() *PlayerCookArgsReq_CmdId {
	p := new(PlayerCookArgsReq_CmdId)
	*p = x
	return p
}

func (x PlayerCookArgsReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerCookArgsReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerCookArgsReq_proto_enumTypes[0].Descriptor()
}

func (PlayerCookArgsReq_CmdId) Type() protoreflect.EnumType {
	return &file_PlayerCookArgsReq_proto_enumTypes[0]
}

func (x PlayerCookArgsReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerCookArgsReq_CmdId.Descriptor instead.
func (PlayerCookArgsReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PlayerCookArgsReq_proto_rawDescGZIP(), []int{0, 0}
}

type PlayerCookArgsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipeId     uint32 `protobuf:"varint,1,opt,name=recipe_id,json=recipeId,proto3" json:"recipe_id,omitempty"`
	AssistAvatar uint32 `protobuf:"varint,2,opt,name=assist_avatar,json=assistAvatar,proto3" json:"assist_avatar,omitempty"`
}

func (x *PlayerCookArgsReq) Reset() {
	*x = PlayerCookArgsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerCookArgsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerCookArgsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerCookArgsReq) ProtoMessage() {}

func (x *PlayerCookArgsReq) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerCookArgsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerCookArgsReq.ProtoReflect.Descriptor instead.
func (*PlayerCookArgsReq) Descriptor() ([]byte, []int) {
	return file_PlayerCookArgsReq_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerCookArgsReq) GetRecipeId() uint32 {
	if x != nil {
		return x.RecipeId
	}
	return 0
}

func (x *PlayerCookArgsReq) GetAssistAvatar() uint32 {
	if x != nil {
		return x.AssistAvatar
	}
	return 0
}

var File_PlayerCookArgsReq_proto protoreflect.FileDescriptor

var file_PlayerCookArgsReq_proto_rawDesc = []byte{
	0x0a, 0x17, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6f, 0x6b, 0x41, 0x72, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb9, 0x01, 0x0a, 0x11, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6f, 0x6b, 0x41,
	0x72, 0x67, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x5f, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x61, 0x73, 0x73, 0x69,
	0x73, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49,
	0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43,
	0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x9f, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f,
	0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerCookArgsReq_proto_rawDescOnce sync.Once
	file_PlayerCookArgsReq_proto_rawDescData = file_PlayerCookArgsReq_proto_rawDesc
)

func file_PlayerCookArgsReq_proto_rawDescGZIP() []byte {
	file_PlayerCookArgsReq_proto_rawDescOnce.Do(func() {
		file_PlayerCookArgsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerCookArgsReq_proto_rawDescData)
	})
	return file_PlayerCookArgsReq_proto_rawDescData
}

var file_PlayerCookArgsReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerCookArgsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerCookArgsReq_proto_goTypes = []interface{}{
	(PlayerCookArgsReq_CmdId)(0), // 0: proto.PlayerCookArgsReq.CmdId
	(*PlayerCookArgsReq)(nil),    // 1: proto.PlayerCookArgsReq
}
var file_PlayerCookArgsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PlayerCookArgsReq_proto_init() }
func file_PlayerCookArgsReq_proto_init() {
	if File_PlayerCookArgsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PlayerCookArgsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerCookArgsReq); i {
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
			RawDescriptor: file_PlayerCookArgsReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerCookArgsReq_proto_goTypes,
		DependencyIndexes: file_PlayerCookArgsReq_proto_depIdxs,
		EnumInfos:         file_PlayerCookArgsReq_proto_enumTypes,
		MessageInfos:      file_PlayerCookArgsReq_proto_msgTypes,
	}.Build()
	File_PlayerCookArgsReq_proto = out.File
	file_PlayerCookArgsReq_proto_rawDesc = nil
	file_PlayerCookArgsReq_proto_goTypes = nil
	file_PlayerCookArgsReq_proto_depIdxs = nil
}
