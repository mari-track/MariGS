// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: PlayerApplyEnterMpAfterMatchAgreedNotify.proto

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

type PlayerApplyEnterMpAfterMatchAgreedNotify_CmdId int32

const (
	PlayerApplyEnterMpAfterMatchAgreedNotify_NONE             PlayerApplyEnterMpAfterMatchAgreedNotify_CmdId = 0
	PlayerApplyEnterMpAfterMatchAgreedNotify_CMD_ID           PlayerApplyEnterMpAfterMatchAgreedNotify_CmdId = 4162
	PlayerApplyEnterMpAfterMatchAgreedNotify_ENET_CHANNEL_ID  PlayerApplyEnterMpAfterMatchAgreedNotify_CmdId = 0
	PlayerApplyEnterMpAfterMatchAgreedNotify_ENET_IS_RELIABLE PlayerApplyEnterMpAfterMatchAgreedNotify_CmdId = 1
)

// Enum value maps for PlayerApplyEnterMpAfterMatchAgreedNotify_CmdId.
var (
	PlayerApplyEnterMpAfterMatchAgreedNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		4162: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	PlayerApplyEnterMpAfterMatchAgreedNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           4162,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x PlayerApplyEnterMpAfterMatchAgreedNotify_CmdId) Enum() *PlayerApplyEnterMpAfterMatchAgreedNotify_CmdId {
	p := new(PlayerApplyEnterMpAfterMatchAgreedNotify_CmdId)
	*p = x
	return p
}

func (x PlayerApplyEnterMpAfterMatchAgreedNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerApplyEnterMpAfterMatchAgreedNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_enumTypes[0].Descriptor()
}

func (PlayerApplyEnterMpAfterMatchAgreedNotify_CmdId) Type() protoreflect.EnumType {
	return &file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_enumTypes[0]
}

func (x PlayerApplyEnterMpAfterMatchAgreedNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerApplyEnterMpAfterMatchAgreedNotify_CmdId.Descriptor instead.
func (PlayerApplyEnterMpAfterMatchAgreedNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_rawDescGZIP(), []int{0, 0}
}

type PlayerApplyEnterMpAfterMatchAgreedNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchType     MatchType         `protobuf:"varint,1,opt,name=match_type,json=matchType,proto3,enum=MatchType" json:"match_type,omitempty"`
	SrcPlayerInfo *OnlinePlayerInfo `protobuf:"bytes,2,opt,name=src_player_info,json=srcPlayerInfo,proto3" json:"src_player_info,omitempty"`
	MatchserverId uint32            `protobuf:"varint,3,opt,name=matchserver_id,json=matchserverId,proto3" json:"matchserver_id,omitempty"`
}

func (x *PlayerApplyEnterMpAfterMatchAgreedNotify) Reset() {
	*x = PlayerApplyEnterMpAfterMatchAgreedNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerApplyEnterMpAfterMatchAgreedNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerApplyEnterMpAfterMatchAgreedNotify) ProtoMessage() {}

func (x *PlayerApplyEnterMpAfterMatchAgreedNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerApplyEnterMpAfterMatchAgreedNotify.ProtoReflect.Descriptor instead.
func (*PlayerApplyEnterMpAfterMatchAgreedNotify) Descriptor() ([]byte, []int) {
	return file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerApplyEnterMpAfterMatchAgreedNotify) GetMatchType() MatchType {
	if x != nil {
		return x.MatchType
	}
	return MatchType_MATCH_TYPE_NONE
}

func (x *PlayerApplyEnterMpAfterMatchAgreedNotify) GetSrcPlayerInfo() *OnlinePlayerInfo {
	if x != nil {
		return x.SrcPlayerInfo
	}
	return nil
}

func (x *PlayerApplyEnterMpAfterMatchAgreedNotify) GetMatchserverId() uint32 {
	if x != nil {
		return x.MatchserverId
	}
	return 0
}

var File_PlayerApplyEnterMpAfterMatchAgreedNotify_proto protoreflect.FileDescriptor

var file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x4d, 0x70, 0x41, 0x66, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x67,
	0x72, 0x65, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x02, 0x0a, 0x28, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x70,
	0x41, 0x66, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x67, 0x72, 0x65, 0x65, 0x64,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x29, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x39, 0x0a, 0x0f, 0x73, 0x72, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x73,
	0x72, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25, 0x0a, 0x0e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44,
	0x10, 0xc2, 0x20, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02,
	0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_rawDescOnce sync.Once
	file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_rawDescData = file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_rawDesc
)

func file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_rawDescGZIP() []byte {
	file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_rawDescOnce.Do(func() {
		file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_rawDescData)
	})
	return file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_rawDescData
}

var file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_goTypes = []interface{}{
	(PlayerApplyEnterMpAfterMatchAgreedNotify_CmdId)(0), // 0: PlayerApplyEnterMpAfterMatchAgreedNotify.CmdId
	(*PlayerApplyEnterMpAfterMatchAgreedNotify)(nil),    // 1: PlayerApplyEnterMpAfterMatchAgreedNotify
	(MatchType)(0),           // 2: MatchType
	(*OnlinePlayerInfo)(nil), // 3: OnlinePlayerInfo
}
var file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_depIdxs = []int32{
	2, // 0: PlayerApplyEnterMpAfterMatchAgreedNotify.match_type:type_name -> MatchType
	3, // 1: PlayerApplyEnterMpAfterMatchAgreedNotify.src_player_info:type_name -> OnlinePlayerInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_init() }
func file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_init() {
	if File_PlayerApplyEnterMpAfterMatchAgreedNotify_proto != nil {
		return
	}
	file_MatchType_proto_init()
	file_OnlinePlayerInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerApplyEnterMpAfterMatchAgreedNotify); i {
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
			RawDescriptor: file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_goTypes,
		DependencyIndexes: file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_depIdxs,
		EnumInfos:         file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_enumTypes,
		MessageInfos:      file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_msgTypes,
	}.Build()
	File_PlayerApplyEnterMpAfterMatchAgreedNotify_proto = out.File
	file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_rawDesc = nil
	file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_goTypes = nil
	file_PlayerApplyEnterMpAfterMatchAgreedNotify_proto_depIdxs = nil
}
