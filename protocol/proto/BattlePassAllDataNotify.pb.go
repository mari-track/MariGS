// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: BattlePassAllDataNotify.proto

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

type BattlePassAllDataNotify_CmdId int32

const (
	BattlePassAllDataNotify_NONE             BattlePassAllDataNotify_CmdId = 0
	BattlePassAllDataNotify_CMD_ID           BattlePassAllDataNotify_CmdId = 2601
	BattlePassAllDataNotify_ENET_CHANNEL_ID  BattlePassAllDataNotify_CmdId = 0
	BattlePassAllDataNotify_ENET_IS_RELIABLE BattlePassAllDataNotify_CmdId = 1
)

// Enum value maps for BattlePassAllDataNotify_CmdId.
var (
	BattlePassAllDataNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		2601: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	BattlePassAllDataNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2601,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x BattlePassAllDataNotify_CmdId) Enum() *BattlePassAllDataNotify_CmdId {
	p := new(BattlePassAllDataNotify_CmdId)
	*p = x
	return p
}

func (x BattlePassAllDataNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BattlePassAllDataNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_BattlePassAllDataNotify_proto_enumTypes[0].Descriptor()
}

func (BattlePassAllDataNotify_CmdId) Type() protoreflect.EnumType {
	return &file_BattlePassAllDataNotify_proto_enumTypes[0]
}

func (x BattlePassAllDataNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BattlePassAllDataNotify_CmdId.Descriptor instead.
func (BattlePassAllDataNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_BattlePassAllDataNotify_proto_rawDescGZIP(), []int{0, 0}
}

type BattlePassAllDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HaveCurSchedule bool                 `protobuf:"varint,1,opt,name=have_cur_schedule,json=haveCurSchedule,proto3" json:"have_cur_schedule,omitempty"`
	CurSchedule     *BattlePassSchedule  `protobuf:"bytes,2,opt,name=cur_schedule,json=curSchedule,proto3" json:"cur_schedule,omitempty"`
	MissionList     []*BattlePassMission `protobuf:"bytes,3,rep,name=mission_list,json=missionList,proto3" json:"mission_list,omitempty"`
}

func (x *BattlePassAllDataNotify) Reset() {
	*x = BattlePassAllDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BattlePassAllDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattlePassAllDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattlePassAllDataNotify) ProtoMessage() {}

func (x *BattlePassAllDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_BattlePassAllDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattlePassAllDataNotify.ProtoReflect.Descriptor instead.
func (*BattlePassAllDataNotify) Descriptor() ([]byte, []int) {
	return file_BattlePassAllDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *BattlePassAllDataNotify) GetHaveCurSchedule() bool {
	if x != nil {
		return x.HaveCurSchedule
	}
	return false
}

func (x *BattlePassAllDataNotify) GetCurSchedule() *BattlePassSchedule {
	if x != nil {
		return x.CurSchedule
	}
	return nil
}

func (x *BattlePassAllDataNotify) GetMissionList() []*BattlePassMission {
	if x != nil {
		return x.MissionList
	}
	return nil
}

var File_BattlePassAllDataNotify_proto protoreflect.FileDescriptor

var file_BattlePassAllDataNotify_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x41, 0x6c, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x83, 0x02, 0x0a, 0x17, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2a,
	0x0a, 0x11, 0x68, 0x61, 0x76, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x68, 0x61, 0x76, 0x65, 0x43,
	0x75, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x63, 0x75,
	0x72, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64,
	0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06,
	0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xa9, 0x14, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45,
	0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BattlePassAllDataNotify_proto_rawDescOnce sync.Once
	file_BattlePassAllDataNotify_proto_rawDescData = file_BattlePassAllDataNotify_proto_rawDesc
)

func file_BattlePassAllDataNotify_proto_rawDescGZIP() []byte {
	file_BattlePassAllDataNotify_proto_rawDescOnce.Do(func() {
		file_BattlePassAllDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_BattlePassAllDataNotify_proto_rawDescData)
	})
	return file_BattlePassAllDataNotify_proto_rawDescData
}

var file_BattlePassAllDataNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_BattlePassAllDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BattlePassAllDataNotify_proto_goTypes = []interface{}{
	(BattlePassAllDataNotify_CmdId)(0), // 0: BattlePassAllDataNotify.CmdId
	(*BattlePassAllDataNotify)(nil),    // 1: BattlePassAllDataNotify
	(*BattlePassSchedule)(nil),         // 2: BattlePassSchedule
	(*BattlePassMission)(nil),          // 3: BattlePassMission
}
var file_BattlePassAllDataNotify_proto_depIdxs = []int32{
	2, // 0: BattlePassAllDataNotify.cur_schedule:type_name -> BattlePassSchedule
	3, // 1: BattlePassAllDataNotify.mission_list:type_name -> BattlePassMission
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_BattlePassAllDataNotify_proto_init() }
func file_BattlePassAllDataNotify_proto_init() {
	if File_BattlePassAllDataNotify_proto != nil {
		return
	}
	file_BattlePassSchedule_proto_init()
	file_BattlePassMission_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BattlePassAllDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattlePassAllDataNotify); i {
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
			RawDescriptor: file_BattlePassAllDataNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BattlePassAllDataNotify_proto_goTypes,
		DependencyIndexes: file_BattlePassAllDataNotify_proto_depIdxs,
		EnumInfos:         file_BattlePassAllDataNotify_proto_enumTypes,
		MessageInfos:      file_BattlePassAllDataNotify_proto_msgTypes,
	}.Build()
	File_BattlePassAllDataNotify_proto = out.File
	file_BattlePassAllDataNotify_proto_rawDesc = nil
	file_BattlePassAllDataNotify_proto_goTypes = nil
	file_BattlePassAllDataNotify_proto_depIdxs = nil
}
