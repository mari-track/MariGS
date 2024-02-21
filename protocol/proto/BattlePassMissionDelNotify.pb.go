// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: BattlePassMissionDelNotify.proto

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

type BattlePassMissionDelNotify_CmdId int32

const (
	BattlePassMissionDelNotify_NONE             BattlePassMissionDelNotify_CmdId = 0
	BattlePassMissionDelNotify_CMD_ID           BattlePassMissionDelNotify_CmdId = 2603
	BattlePassMissionDelNotify_ENET_CHANNEL_ID  BattlePassMissionDelNotify_CmdId = 0
	BattlePassMissionDelNotify_ENET_IS_RELIABLE BattlePassMissionDelNotify_CmdId = 1
)

// Enum value maps for BattlePassMissionDelNotify_CmdId.
var (
	BattlePassMissionDelNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		2603: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	BattlePassMissionDelNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2603,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x BattlePassMissionDelNotify_CmdId) Enum() *BattlePassMissionDelNotify_CmdId {
	p := new(BattlePassMissionDelNotify_CmdId)
	*p = x
	return p
}

func (x BattlePassMissionDelNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BattlePassMissionDelNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_BattlePassMissionDelNotify_proto_enumTypes[0].Descriptor()
}

func (BattlePassMissionDelNotify_CmdId) Type() protoreflect.EnumType {
	return &file_BattlePassMissionDelNotify_proto_enumTypes[0]
}

func (x BattlePassMissionDelNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BattlePassMissionDelNotify_CmdId.Descriptor instead.
func (BattlePassMissionDelNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_BattlePassMissionDelNotify_proto_rawDescGZIP(), []int{0, 0}
}

type BattlePassMissionDelNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DelMissionIdList []uint32 `protobuf:"varint,1,rep,packed,name=del_mission_id_list,json=delMissionIdList,proto3" json:"del_mission_id_list,omitempty"`
}

func (x *BattlePassMissionDelNotify) Reset() {
	*x = BattlePassMissionDelNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BattlePassMissionDelNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattlePassMissionDelNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattlePassMissionDelNotify) ProtoMessage() {}

func (x *BattlePassMissionDelNotify) ProtoReflect() protoreflect.Message {
	mi := &file_BattlePassMissionDelNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattlePassMissionDelNotify.ProtoReflect.Descriptor instead.
func (*BattlePassMissionDelNotify) Descriptor() ([]byte, []int) {
	return file_BattlePassMissionDelNotify_proto_rawDescGZIP(), []int{0}
}

func (x *BattlePassMissionDelNotify) GetDelMissionIdList() []uint32 {
	if x != nil {
		return x.DelMissionIdList
	}
	return nil
}

var File_BattlePassMissionDelNotify_proto protoreflect.FileDescriptor

var file_BattlePassMissionDelNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x4d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x1a, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x2d, 0x0a, 0x13, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10,
	0x64, 0x65, 0x6c, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xab, 0x14,
	0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c,
	0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53,
	0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_BattlePassMissionDelNotify_proto_rawDescOnce sync.Once
	file_BattlePassMissionDelNotify_proto_rawDescData = file_BattlePassMissionDelNotify_proto_rawDesc
)

func file_BattlePassMissionDelNotify_proto_rawDescGZIP() []byte {
	file_BattlePassMissionDelNotify_proto_rawDescOnce.Do(func() {
		file_BattlePassMissionDelNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_BattlePassMissionDelNotify_proto_rawDescData)
	})
	return file_BattlePassMissionDelNotify_proto_rawDescData
}

var file_BattlePassMissionDelNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_BattlePassMissionDelNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BattlePassMissionDelNotify_proto_goTypes = []interface{}{
	(BattlePassMissionDelNotify_CmdId)(0), // 0: BattlePassMissionDelNotify.CmdId
	(*BattlePassMissionDelNotify)(nil),    // 1: BattlePassMissionDelNotify
}
var file_BattlePassMissionDelNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BattlePassMissionDelNotify_proto_init() }
func file_BattlePassMissionDelNotify_proto_init() {
	if File_BattlePassMissionDelNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BattlePassMissionDelNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattlePassMissionDelNotify); i {
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
			RawDescriptor: file_BattlePassMissionDelNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BattlePassMissionDelNotify_proto_goTypes,
		DependencyIndexes: file_BattlePassMissionDelNotify_proto_depIdxs,
		EnumInfos:         file_BattlePassMissionDelNotify_proto_enumTypes,
		MessageInfos:      file_BattlePassMissionDelNotify_proto_msgTypes,
	}.Build()
	File_BattlePassMissionDelNotify_proto = out.File
	file_BattlePassMissionDelNotify_proto_rawDesc = nil
	file_BattlePassMissionDelNotify_proto_goTypes = nil
	file_BattlePassMissionDelNotify_proto_depIdxs = nil
}
