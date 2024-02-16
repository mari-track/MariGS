// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: DungeonCandidateTeamDismissNotify.proto

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

type DungeonCandidateTeamDismissNotify_CmdId int32

const (
	DungeonCandidateTeamDismissNotify_NONE             DungeonCandidateTeamDismissNotify_CmdId = 0
	DungeonCandidateTeamDismissNotify_CMD_ID           DungeonCandidateTeamDismissNotify_CmdId = 929
	DungeonCandidateTeamDismissNotify_ENET_CHANNEL_ID  DungeonCandidateTeamDismissNotify_CmdId = 0
	DungeonCandidateTeamDismissNotify_ENET_IS_RELIABLE DungeonCandidateTeamDismissNotify_CmdId = 1
)

// Enum value maps for DungeonCandidateTeamDismissNotify_CmdId.
var (
	DungeonCandidateTeamDismissNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		929: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	DungeonCandidateTeamDismissNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           929,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x DungeonCandidateTeamDismissNotify_CmdId) Enum() *DungeonCandidateTeamDismissNotify_CmdId {
	p := new(DungeonCandidateTeamDismissNotify_CmdId)
	*p = x
	return p
}

func (x DungeonCandidateTeamDismissNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DungeonCandidateTeamDismissNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_DungeonCandidateTeamDismissNotify_proto_enumTypes[0].Descriptor()
}

func (DungeonCandidateTeamDismissNotify_CmdId) Type() protoreflect.EnumType {
	return &file_DungeonCandidateTeamDismissNotify_proto_enumTypes[0]
}

func (x DungeonCandidateTeamDismissNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DungeonCandidateTeamDismissNotify_CmdId.Descriptor instead.
func (DungeonCandidateTeamDismissNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_DungeonCandidateTeamDismissNotify_proto_rawDescGZIP(), []int{0, 0}
}

type DungeonCandidateTeamDismissNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerUid uint32                            `protobuf:"varint,1,opt,name=player_uid,json=playerUid,proto3" json:"player_uid,omitempty"`
	Reason    DungeonCandidateTeamDismissReason `protobuf:"varint,2,opt,name=reason,proto3,enum=proto.DungeonCandidateTeamDismissReason" json:"reason,omitempty"`
}

func (x *DungeonCandidateTeamDismissNotify) Reset() {
	*x = DungeonCandidateTeamDismissNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DungeonCandidateTeamDismissNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DungeonCandidateTeamDismissNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DungeonCandidateTeamDismissNotify) ProtoMessage() {}

func (x *DungeonCandidateTeamDismissNotify) ProtoReflect() protoreflect.Message {
	mi := &file_DungeonCandidateTeamDismissNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DungeonCandidateTeamDismissNotify.ProtoReflect.Descriptor instead.
func (*DungeonCandidateTeamDismissNotify) Descriptor() ([]byte, []int) {
	return file_DungeonCandidateTeamDismissNotify_proto_rawDescGZIP(), []int{0}
}

func (x *DungeonCandidateTeamDismissNotify) GetPlayerUid() uint32 {
	if x != nil {
		return x.PlayerUid
	}
	return 0
}

func (x *DungeonCandidateTeamDismissNotify) GetReason() DungeonCandidateTeamDismissReason {
	if x != nil {
		return x.Reason
	}
	return DungeonCandidateTeamDismissReason_DUNGEON_CANDIDATE_TPDR_NORMAL
}

var File_DungeonCandidateTeamDismissNotify_proto protoreflect.FileDescriptor

var file_DungeonCandidateTeamDismissNotify_proto_rawDesc = []byte{
	0x0a, 0x27, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x44, 0x69, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x27, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x44, 0x69, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x21, 0x44, 0x75,
	0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x61, 0x6d, 0x44, 0x69, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x40,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x44, 0x69, 0x73, 0x6d, 0x69,
	0x73, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xa1, 0x07,
	0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c,
	0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53,
	0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_DungeonCandidateTeamDismissNotify_proto_rawDescOnce sync.Once
	file_DungeonCandidateTeamDismissNotify_proto_rawDescData = file_DungeonCandidateTeamDismissNotify_proto_rawDesc
)

func file_DungeonCandidateTeamDismissNotify_proto_rawDescGZIP() []byte {
	file_DungeonCandidateTeamDismissNotify_proto_rawDescOnce.Do(func() {
		file_DungeonCandidateTeamDismissNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_DungeonCandidateTeamDismissNotify_proto_rawDescData)
	})
	return file_DungeonCandidateTeamDismissNotify_proto_rawDescData
}

var file_DungeonCandidateTeamDismissNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DungeonCandidateTeamDismissNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DungeonCandidateTeamDismissNotify_proto_goTypes = []interface{}{
	(DungeonCandidateTeamDismissNotify_CmdId)(0), // 0: proto.DungeonCandidateTeamDismissNotify.CmdId
	(*DungeonCandidateTeamDismissNotify)(nil),    // 1: proto.DungeonCandidateTeamDismissNotify
	(DungeonCandidateTeamDismissReason)(0),       // 2: proto.DungeonCandidateTeamDismissReason
}
var file_DungeonCandidateTeamDismissNotify_proto_depIdxs = []int32{
	2, // 0: proto.DungeonCandidateTeamDismissNotify.reason:type_name -> proto.DungeonCandidateTeamDismissReason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_DungeonCandidateTeamDismissNotify_proto_init() }
func file_DungeonCandidateTeamDismissNotify_proto_init() {
	if File_DungeonCandidateTeamDismissNotify_proto != nil {
		return
	}
	file_DungeonCandidateTeamDismissReason_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DungeonCandidateTeamDismissNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DungeonCandidateTeamDismissNotify); i {
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
			RawDescriptor: file_DungeonCandidateTeamDismissNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DungeonCandidateTeamDismissNotify_proto_goTypes,
		DependencyIndexes: file_DungeonCandidateTeamDismissNotify_proto_depIdxs,
		EnumInfos:         file_DungeonCandidateTeamDismissNotify_proto_enumTypes,
		MessageInfos:      file_DungeonCandidateTeamDismissNotify_proto_msgTypes,
	}.Build()
	File_DungeonCandidateTeamDismissNotify_proto = out.File
	file_DungeonCandidateTeamDismissNotify_proto_rawDesc = nil
	file_DungeonCandidateTeamDismissNotify_proto_goTypes = nil
	file_DungeonCandidateTeamDismissNotify_proto_depIdxs = nil
}