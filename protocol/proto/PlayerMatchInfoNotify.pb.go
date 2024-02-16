// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: PlayerMatchInfoNotify.proto

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

type PlayerMatchInfoNotify_CmdId int32

const (
	PlayerMatchInfoNotify_NONE             PlayerMatchInfoNotify_CmdId = 0
	PlayerMatchInfoNotify_CMD_ID           PlayerMatchInfoNotify_CmdId = 4153
	PlayerMatchInfoNotify_ENET_CHANNEL_ID  PlayerMatchInfoNotify_CmdId = 0
	PlayerMatchInfoNotify_ENET_IS_RELIABLE PlayerMatchInfoNotify_CmdId = 1
)

// Enum value maps for PlayerMatchInfoNotify_CmdId.
var (
	PlayerMatchInfoNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		4153: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	PlayerMatchInfoNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           4153,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x PlayerMatchInfoNotify_CmdId) Enum() *PlayerMatchInfoNotify_CmdId {
	p := new(PlayerMatchInfoNotify_CmdId)
	*p = x
	return p
}

func (x PlayerMatchInfoNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerMatchInfoNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerMatchInfoNotify_proto_enumTypes[0].Descriptor()
}

func (PlayerMatchInfoNotify_CmdId) Type() protoreflect.EnumType {
	return &file_PlayerMatchInfoNotify_proto_enumTypes[0]
}

func (x PlayerMatchInfoNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerMatchInfoNotify_CmdId.Descriptor instead.
func (PlayerMatchInfoNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PlayerMatchInfoNotify_proto_rawDescGZIP(), []int{0, 0}
}

type PlayerMatchInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostUid               uint32    `protobuf:"varint,1,opt,name=host_uid,json=hostUid,proto3" json:"host_uid,omitempty"`
	MatchType             MatchType `protobuf:"varint,2,opt,name=match_type,json=matchType,proto3,enum=proto.MatchType" json:"match_type,omitempty"`
	MatchBeginTime        uint32    `protobuf:"varint,3,opt,name=match_begin_time,json=matchBeginTime,proto3" json:"match_begin_time,omitempty"`
	EstimateMatchCostTime uint32    `protobuf:"varint,4,opt,name=estimate_match_cost_time,json=estimateMatchCostTime,proto3" json:"estimate_match_cost_time,omitempty"`
	DungeonId             uint32    `protobuf:"varint,11,opt,name=dungeon_id,json=dungeonId,proto3" json:"dungeon_id,omitempty"`
	MpPlayId              uint32    `protobuf:"varint,12,opt,name=mp_play_id,json=mpPlayId,proto3" json:"mp_play_id,omitempty"`
}

func (x *PlayerMatchInfoNotify) Reset() {
	*x = PlayerMatchInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerMatchInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerMatchInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerMatchInfoNotify) ProtoMessage() {}

func (x *PlayerMatchInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerMatchInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerMatchInfoNotify.ProtoReflect.Descriptor instead.
func (*PlayerMatchInfoNotify) Descriptor() ([]byte, []int) {
	return file_PlayerMatchInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerMatchInfoNotify) GetHostUid() uint32 {
	if x != nil {
		return x.HostUid
	}
	return 0
}

func (x *PlayerMatchInfoNotify) GetMatchType() MatchType {
	if x != nil {
		return x.MatchType
	}
	return MatchType_MATCH_TYPE_NONE
}

func (x *PlayerMatchInfoNotify) GetMatchBeginTime() uint32 {
	if x != nil {
		return x.MatchBeginTime
	}
	return 0
}

func (x *PlayerMatchInfoNotify) GetEstimateMatchCostTime() uint32 {
	if x != nil {
		return x.EstimateMatchCostTime
	}
	return 0
}

func (x *PlayerMatchInfoNotify) GetDungeonId() uint32 {
	if x != nil {
		return x.DungeonId
	}
	return 0
}

func (x *PlayerMatchInfoNotify) GetMpPlayId() uint32 {
	if x != nil {
		return x.MpPlayId
	}
	return 0
}

var File_PlayerMatchInfoNotify_proto protoreflect.FileDescriptor

var file_PlayerMatchInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66,
	0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x02, 0x0a, 0x15, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x19, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x55, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x0a, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x42, 0x65, 0x67, 0x69,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x18, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x0a, 0x6d, 0x70, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x6d, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x05, 0x43,
	0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xb9, 0x20, 0x12, 0x13, 0x0a, 0x0f, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerMatchInfoNotify_proto_rawDescOnce sync.Once
	file_PlayerMatchInfoNotify_proto_rawDescData = file_PlayerMatchInfoNotify_proto_rawDesc
)

func file_PlayerMatchInfoNotify_proto_rawDescGZIP() []byte {
	file_PlayerMatchInfoNotify_proto_rawDescOnce.Do(func() {
		file_PlayerMatchInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerMatchInfoNotify_proto_rawDescData)
	})
	return file_PlayerMatchInfoNotify_proto_rawDescData
}

var file_PlayerMatchInfoNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerMatchInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerMatchInfoNotify_proto_goTypes = []interface{}{
	(PlayerMatchInfoNotify_CmdId)(0), // 0: proto.PlayerMatchInfoNotify.CmdId
	(*PlayerMatchInfoNotify)(nil),    // 1: proto.PlayerMatchInfoNotify
	(MatchType)(0),                   // 2: proto.MatchType
}
var file_PlayerMatchInfoNotify_proto_depIdxs = []int32{
	2, // 0: proto.PlayerMatchInfoNotify.match_type:type_name -> proto.MatchType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerMatchInfoNotify_proto_init() }
func file_PlayerMatchInfoNotify_proto_init() {
	if File_PlayerMatchInfoNotify_proto != nil {
		return
	}
	file_MatchType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerMatchInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerMatchInfoNotify); i {
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
			RawDescriptor: file_PlayerMatchInfoNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerMatchInfoNotify_proto_goTypes,
		DependencyIndexes: file_PlayerMatchInfoNotify_proto_depIdxs,
		EnumInfos:         file_PlayerMatchInfoNotify_proto_enumTypes,
		MessageInfos:      file_PlayerMatchInfoNotify_proto_msgTypes,
	}.Build()
	File_PlayerMatchInfoNotify_proto = out.File
	file_PlayerMatchInfoNotify_proto_rawDesc = nil
	file_PlayerMatchInfoNotify_proto_goTypes = nil
	file_PlayerMatchInfoNotify_proto_depIdxs = nil
}
