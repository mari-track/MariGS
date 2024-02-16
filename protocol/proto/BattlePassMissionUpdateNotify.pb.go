// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: BattlePassMissionUpdateNotify.proto

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

type BattlePassMissionUpdateNotify_CmdId int32

const (
	BattlePassMissionUpdateNotify_NONE             BattlePassMissionUpdateNotify_CmdId = 0
	BattlePassMissionUpdateNotify_CMD_ID           BattlePassMissionUpdateNotify_CmdId = 2602
	BattlePassMissionUpdateNotify_ENET_CHANNEL_ID  BattlePassMissionUpdateNotify_CmdId = 0
	BattlePassMissionUpdateNotify_ENET_IS_RELIABLE BattlePassMissionUpdateNotify_CmdId = 1
)

// Enum value maps for BattlePassMissionUpdateNotify_CmdId.
var (
	BattlePassMissionUpdateNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		2602: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	BattlePassMissionUpdateNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2602,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x BattlePassMissionUpdateNotify_CmdId) Enum() *BattlePassMissionUpdateNotify_CmdId {
	p := new(BattlePassMissionUpdateNotify_CmdId)
	*p = x
	return p
}

func (x BattlePassMissionUpdateNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BattlePassMissionUpdateNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_BattlePassMissionUpdateNotify_proto_enumTypes[0].Descriptor()
}

func (BattlePassMissionUpdateNotify_CmdId) Type() protoreflect.EnumType {
	return &file_BattlePassMissionUpdateNotify_proto_enumTypes[0]
}

func (x BattlePassMissionUpdateNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BattlePassMissionUpdateNotify_CmdId.Descriptor instead.
func (BattlePassMissionUpdateNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_BattlePassMissionUpdateNotify_proto_rawDescGZIP(), []int{0, 0}
}

type BattlePassMissionUpdateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MissionList []*BattlePassMission `protobuf:"bytes,1,rep,name=mission_list,json=missionList,proto3" json:"mission_list,omitempty"`
}

func (x *BattlePassMissionUpdateNotify) Reset() {
	*x = BattlePassMissionUpdateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BattlePassMissionUpdateNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattlePassMissionUpdateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattlePassMissionUpdateNotify) ProtoMessage() {}

func (x *BattlePassMissionUpdateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_BattlePassMissionUpdateNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattlePassMissionUpdateNotify.ProtoReflect.Descriptor instead.
func (*BattlePassMissionUpdateNotify) Descriptor() ([]byte, []int) {
	return file_BattlePassMissionUpdateNotify_proto_rawDescGZIP(), []int{0}
}

func (x *BattlePassMissionUpdateNotify) GetMissionList() []*BattlePassMission {
	if x != nil {
		return x.MissionList
	}
	return nil
}

var File_BattlePassMissionUpdateNotify_proto protoreflect.FileDescriptor

var file_BattlePassMissionUpdateNotify_proto_rawDesc = []byte{
	0x0a, 0x23, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x4d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x1d, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x3b, 0x0a, 0x0c, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49,
	0x44, 0x10, 0xaa, 0x14, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45,
	0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a,
	0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BattlePassMissionUpdateNotify_proto_rawDescOnce sync.Once
	file_BattlePassMissionUpdateNotify_proto_rawDescData = file_BattlePassMissionUpdateNotify_proto_rawDesc
)

func file_BattlePassMissionUpdateNotify_proto_rawDescGZIP() []byte {
	file_BattlePassMissionUpdateNotify_proto_rawDescOnce.Do(func() {
		file_BattlePassMissionUpdateNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_BattlePassMissionUpdateNotify_proto_rawDescData)
	})
	return file_BattlePassMissionUpdateNotify_proto_rawDescData
}

var file_BattlePassMissionUpdateNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_BattlePassMissionUpdateNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BattlePassMissionUpdateNotify_proto_goTypes = []interface{}{
	(BattlePassMissionUpdateNotify_CmdId)(0), // 0: proto.BattlePassMissionUpdateNotify.CmdId
	(*BattlePassMissionUpdateNotify)(nil),    // 1: proto.BattlePassMissionUpdateNotify
	(*BattlePassMission)(nil),                // 2: proto.BattlePassMission
}
var file_BattlePassMissionUpdateNotify_proto_depIdxs = []int32{
	2, // 0: proto.BattlePassMissionUpdateNotify.mission_list:type_name -> proto.BattlePassMission
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_BattlePassMissionUpdateNotify_proto_init() }
func file_BattlePassMissionUpdateNotify_proto_init() {
	if File_BattlePassMissionUpdateNotify_proto != nil {
		return
	}
	file_BattlePassMission_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BattlePassMissionUpdateNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattlePassMissionUpdateNotify); i {
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
			RawDescriptor: file_BattlePassMissionUpdateNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BattlePassMissionUpdateNotify_proto_goTypes,
		DependencyIndexes: file_BattlePassMissionUpdateNotify_proto_depIdxs,
		EnumInfos:         file_BattlePassMissionUpdateNotify_proto_enumTypes,
		MessageInfos:      file_BattlePassMissionUpdateNotify_proto_msgTypes,
	}.Build()
	File_BattlePassMissionUpdateNotify_proto = out.File
	file_BattlePassMissionUpdateNotify_proto_rawDesc = nil
	file_BattlePassMissionUpdateNotify_proto_goTypes = nil
	file_BattlePassMissionUpdateNotify_proto_depIdxs = nil
}