// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: DungeonCandidateTeamPlayerLeaveNotify.proto

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

type DungeonCandidateTeamPlayerLeaveNotify_CmdId int32

const (
	DungeonCandidateTeamPlayerLeaveNotify_NONE             DungeonCandidateTeamPlayerLeaveNotify_CmdId = 0
	DungeonCandidateTeamPlayerLeaveNotify_CMD_ID           DungeonCandidateTeamPlayerLeaveNotify_CmdId = 928
	DungeonCandidateTeamPlayerLeaveNotify_ENET_CHANNEL_ID  DungeonCandidateTeamPlayerLeaveNotify_CmdId = 0
	DungeonCandidateTeamPlayerLeaveNotify_ENET_IS_RELIABLE DungeonCandidateTeamPlayerLeaveNotify_CmdId = 1
)

// Enum value maps for DungeonCandidateTeamPlayerLeaveNotify_CmdId.
var (
	DungeonCandidateTeamPlayerLeaveNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		928: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	DungeonCandidateTeamPlayerLeaveNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           928,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x DungeonCandidateTeamPlayerLeaveNotify_CmdId) Enum() *DungeonCandidateTeamPlayerLeaveNotify_CmdId {
	p := new(DungeonCandidateTeamPlayerLeaveNotify_CmdId)
	*p = x
	return p
}

func (x DungeonCandidateTeamPlayerLeaveNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DungeonCandidateTeamPlayerLeaveNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_DungeonCandidateTeamPlayerLeaveNotify_proto_enumTypes[0].Descriptor()
}

func (DungeonCandidateTeamPlayerLeaveNotify_CmdId) Type() protoreflect.EnumType {
	return &file_DungeonCandidateTeamPlayerLeaveNotify_proto_enumTypes[0]
}

func (x DungeonCandidateTeamPlayerLeaveNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DungeonCandidateTeamPlayerLeaveNotify_CmdId.Descriptor instead.
func (DungeonCandidateTeamPlayerLeaveNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDescGZIP(), []int{0, 0}
}

type DungeonCandidateTeamPlayerLeaveNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerUid uint32                                `protobuf:"varint,1,opt,name=player_uid,json=playerUid,proto3" json:"player_uid,omitempty"`
	Reason    DungeonCandidateTeamPlayerLeaveReason `protobuf:"varint,2,opt,name=reason,proto3,enum=DungeonCandidateTeamPlayerLeaveReason" json:"reason,omitempty"`
}

func (x *DungeonCandidateTeamPlayerLeaveNotify) Reset() {
	*x = DungeonCandidateTeamPlayerLeaveNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DungeonCandidateTeamPlayerLeaveNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DungeonCandidateTeamPlayerLeaveNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DungeonCandidateTeamPlayerLeaveNotify) ProtoMessage() {}

func (x *DungeonCandidateTeamPlayerLeaveNotify) ProtoReflect() protoreflect.Message {
	mi := &file_DungeonCandidateTeamPlayerLeaveNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DungeonCandidateTeamPlayerLeaveNotify.ProtoReflect.Descriptor instead.
func (*DungeonCandidateTeamPlayerLeaveNotify) Descriptor() ([]byte, []int) {
	return file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDescGZIP(), []int{0}
}

func (x *DungeonCandidateTeamPlayerLeaveNotify) GetPlayerUid() uint32 {
	if x != nil {
		return x.PlayerUid
	}
	return 0
}

func (x *DungeonCandidateTeamPlayerLeaveNotify) GetReason() DungeonCandidateTeamPlayerLeaveReason {
	if x != nil {
		return x.Reason
	}
	return DungeonCandidateTeamPlayerLeaveReason_DUNGEON_CANDIDATE_TPLR_NORMAL
}

var File_DungeonCandidateTeamPlayerLeaveNotify_proto protoreflect.FileDescriptor

var file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76,
	0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x44,
	0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x61, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x25, 0x44,
	0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x61, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x55, 0x69, 0x64, 0x12, 0x3e, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44,
	0x10, 0xa0, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02,
	0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDescOnce sync.Once
	file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDescData = file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDesc
)

func file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDescGZIP() []byte {
	file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDescOnce.Do(func() {
		file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDescData)
	})
	return file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDescData
}

var file_DungeonCandidateTeamPlayerLeaveNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DungeonCandidateTeamPlayerLeaveNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DungeonCandidateTeamPlayerLeaveNotify_proto_goTypes = []interface{}{
	(DungeonCandidateTeamPlayerLeaveNotify_CmdId)(0), // 0: DungeonCandidateTeamPlayerLeaveNotify.CmdId
	(*DungeonCandidateTeamPlayerLeaveNotify)(nil),    // 1: DungeonCandidateTeamPlayerLeaveNotify
	(DungeonCandidateTeamPlayerLeaveReason)(0),       // 2: DungeonCandidateTeamPlayerLeaveReason
}
var file_DungeonCandidateTeamPlayerLeaveNotify_proto_depIdxs = []int32{
	2, // 0: DungeonCandidateTeamPlayerLeaveNotify.reason:type_name -> DungeonCandidateTeamPlayerLeaveReason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_DungeonCandidateTeamPlayerLeaveNotify_proto_init() }
func file_DungeonCandidateTeamPlayerLeaveNotify_proto_init() {
	if File_DungeonCandidateTeamPlayerLeaveNotify_proto != nil {
		return
	}
	file_DungeonCandidateTeamPlayerLeaveReason_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DungeonCandidateTeamPlayerLeaveNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DungeonCandidateTeamPlayerLeaveNotify); i {
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
			RawDescriptor: file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DungeonCandidateTeamPlayerLeaveNotify_proto_goTypes,
		DependencyIndexes: file_DungeonCandidateTeamPlayerLeaveNotify_proto_depIdxs,
		EnumInfos:         file_DungeonCandidateTeamPlayerLeaveNotify_proto_enumTypes,
		MessageInfos:      file_DungeonCandidateTeamPlayerLeaveNotify_proto_msgTypes,
	}.Build()
	File_DungeonCandidateTeamPlayerLeaveNotify_proto = out.File
	file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDesc = nil
	file_DungeonCandidateTeamPlayerLeaveNotify_proto_goTypes = nil
	file_DungeonCandidateTeamPlayerLeaveNotify_proto_depIdxs = nil
}
