// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: DungeonPlayerDieNotify.proto

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

type DungeonPlayerDieNotify_CmdId int32

const (
	DungeonPlayerDieNotify_NONE             DungeonPlayerDieNotify_CmdId = 0
	DungeonPlayerDieNotify_CMD_ID           DungeonPlayerDieNotify_CmdId = 911
	DungeonPlayerDieNotify_ENET_CHANNEL_ID  DungeonPlayerDieNotify_CmdId = 0
	DungeonPlayerDieNotify_ENET_IS_RELIABLE DungeonPlayerDieNotify_CmdId = 1
)

// Enum value maps for DungeonPlayerDieNotify_CmdId.
var (
	DungeonPlayerDieNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		911: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	DungeonPlayerDieNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           911,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x DungeonPlayerDieNotify_CmdId) Enum() *DungeonPlayerDieNotify_CmdId {
	p := new(DungeonPlayerDieNotify_CmdId)
	*p = x
	return p
}

func (x DungeonPlayerDieNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DungeonPlayerDieNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_DungeonPlayerDieNotify_proto_enumTypes[0].Descriptor()
}

func (DungeonPlayerDieNotify_CmdId) Type() protoreflect.EnumType {
	return &file_DungeonPlayerDieNotify_proto_enumTypes[0]
}

func (x DungeonPlayerDieNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DungeonPlayerDieNotify_CmdId.Descriptor instead.
func (DungeonPlayerDieNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_DungeonPlayerDieNotify_proto_rawDescGZIP(), []int{0, 0}
}

type DungeonPlayerDieNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DungeonId        uint32        `protobuf:"varint,1,opt,name=dungeon_id,json=dungeonId,proto3" json:"dungeon_id,omitempty"`
	DieType          PlayerDieType `protobuf:"varint,2,opt,name=die_type,json=dieType,proto3,enum=PlayerDieType" json:"die_type,omitempty"`
	ReviveCount      uint32        `protobuf:"varint,3,opt,name=revive_count,json=reviveCount,proto3" json:"revive_count,omitempty"`
	WaitTime         uint32        `protobuf:"varint,4,opt,name=wait_time,json=waitTime,proto3" json:"wait_time,omitempty"`
	MurdererEntityId uint32        `protobuf:"varint,5,opt,name=murderer_entity_id,json=murdererEntityId,proto3" json:"murderer_entity_id,omitempty"`
	// Types that are assignable to Entity:
	//
	//	*DungeonPlayerDieNotify_MonsterId
	//	*DungeonPlayerDieNotify_GadgetId
	Entity isDungeonPlayerDieNotify_Entity `protobuf_oneof:"entity"`
}

func (x *DungeonPlayerDieNotify) Reset() {
	*x = DungeonPlayerDieNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DungeonPlayerDieNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DungeonPlayerDieNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DungeonPlayerDieNotify) ProtoMessage() {}

func (x *DungeonPlayerDieNotify) ProtoReflect() protoreflect.Message {
	mi := &file_DungeonPlayerDieNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DungeonPlayerDieNotify.ProtoReflect.Descriptor instead.
func (*DungeonPlayerDieNotify) Descriptor() ([]byte, []int) {
	return file_DungeonPlayerDieNotify_proto_rawDescGZIP(), []int{0}
}

func (x *DungeonPlayerDieNotify) GetDungeonId() uint32 {
	if x != nil {
		return x.DungeonId
	}
	return 0
}

func (x *DungeonPlayerDieNotify) GetDieType() PlayerDieType {
	if x != nil {
		return x.DieType
	}
	return PlayerDieType_PLAYER_DIE_NONE
}

func (x *DungeonPlayerDieNotify) GetReviveCount() uint32 {
	if x != nil {
		return x.ReviveCount
	}
	return 0
}

func (x *DungeonPlayerDieNotify) GetWaitTime() uint32 {
	if x != nil {
		return x.WaitTime
	}
	return 0
}

func (x *DungeonPlayerDieNotify) GetMurdererEntityId() uint32 {
	if x != nil {
		return x.MurdererEntityId
	}
	return 0
}

func (m *DungeonPlayerDieNotify) GetEntity() isDungeonPlayerDieNotify_Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (x *DungeonPlayerDieNotify) GetMonsterId() uint32 {
	if x, ok := x.GetEntity().(*DungeonPlayerDieNotify_MonsterId); ok {
		return x.MonsterId
	}
	return 0
}

func (x *DungeonPlayerDieNotify) GetGadgetId() uint32 {
	if x, ok := x.GetEntity().(*DungeonPlayerDieNotify_GadgetId); ok {
		return x.GadgetId
	}
	return 0
}

type isDungeonPlayerDieNotify_Entity interface {
	isDungeonPlayerDieNotify_Entity()
}

type DungeonPlayerDieNotify_MonsterId struct {
	MonsterId uint32 `protobuf:"varint,6,opt,name=monster_id,json=monsterId,proto3,oneof"`
}

type DungeonPlayerDieNotify_GadgetId struct {
	GadgetId uint32 `protobuf:"varint,7,opt,name=gadget_id,json=gadgetId,proto3,oneof"`
}

func (*DungeonPlayerDieNotify_MonsterId) isDungeonPlayerDieNotify_Entity() {}

func (*DungeonPlayerDieNotify_GadgetId) isDungeonPlayerDieNotify_Entity() {}

var File_DungeonPlayerDieNotify_proto protoreflect.FileDescriptor

var file_DungeonPlayerDieNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44,
	0x69, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x02, 0x0a, 0x16, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a,
	0x08, 0x64, 0x69, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0e, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x07, 0x64, 0x69, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x76, 0x69,
	0x76, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x72, 0x65, 0x76, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x77,
	0x61, 0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x77, 0x61, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x75, 0x72, 0x64,
	0x65, 0x72, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x6d, 0x75, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0a, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x09, 0x6d, 0x6f,
	0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x67, 0x61, 0x64, 0x67, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x08, 0x67, 0x61,
	0x64, 0x67, 0x65, 0x74, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44,
	0x5f, 0x49, 0x44, 0x10, 0x8f, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43,
	0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_DungeonPlayerDieNotify_proto_rawDescOnce sync.Once
	file_DungeonPlayerDieNotify_proto_rawDescData = file_DungeonPlayerDieNotify_proto_rawDesc
)

func file_DungeonPlayerDieNotify_proto_rawDescGZIP() []byte {
	file_DungeonPlayerDieNotify_proto_rawDescOnce.Do(func() {
		file_DungeonPlayerDieNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_DungeonPlayerDieNotify_proto_rawDescData)
	})
	return file_DungeonPlayerDieNotify_proto_rawDescData
}

var file_DungeonPlayerDieNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DungeonPlayerDieNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DungeonPlayerDieNotify_proto_goTypes = []interface{}{
	(DungeonPlayerDieNotify_CmdId)(0), // 0: DungeonPlayerDieNotify.CmdId
	(*DungeonPlayerDieNotify)(nil),    // 1: DungeonPlayerDieNotify
	(PlayerDieType)(0),                // 2: PlayerDieType
}
var file_DungeonPlayerDieNotify_proto_depIdxs = []int32{
	2, // 0: DungeonPlayerDieNotify.die_type:type_name -> PlayerDieType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_DungeonPlayerDieNotify_proto_init() }
func file_DungeonPlayerDieNotify_proto_init() {
	if File_DungeonPlayerDieNotify_proto != nil {
		return
	}
	file_PlayerDieType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DungeonPlayerDieNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DungeonPlayerDieNotify); i {
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
	file_DungeonPlayerDieNotify_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DungeonPlayerDieNotify_MonsterId)(nil),
		(*DungeonPlayerDieNotify_GadgetId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_DungeonPlayerDieNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DungeonPlayerDieNotify_proto_goTypes,
		DependencyIndexes: file_DungeonPlayerDieNotify_proto_depIdxs,
		EnumInfos:         file_DungeonPlayerDieNotify_proto_enumTypes,
		MessageInfos:      file_DungeonPlayerDieNotify_proto_msgTypes,
	}.Build()
	File_DungeonPlayerDieNotify_proto = out.File
	file_DungeonPlayerDieNotify_proto_rawDesc = nil
	file_DungeonPlayerDieNotify_proto_goTypes = nil
	file_DungeonPlayerDieNotify_proto_depIdxs = nil
}
