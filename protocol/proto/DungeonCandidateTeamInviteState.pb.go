// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: DungeonCandidateTeamInviteState.proto

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

type DungeonCandidateTeamInviteState int32

const (
	DungeonCandidateTeamInviteState_DUNGEON_CANDIDATE_TEAM_INVITE_STATE_NONE   DungeonCandidateTeamInviteState = 0
	DungeonCandidateTeamInviteState_DUNGEON_CANDIDATE_TEAM_INVITE_STATE_SEND   DungeonCandidateTeamInviteState = 1
	DungeonCandidateTeamInviteState_DUNGEON_CANDIDATE_TEAM_INVITE_STATE_ACCEPT DungeonCandidateTeamInviteState = 2
	DungeonCandidateTeamInviteState_DUNGEON_CANDIDATE_TEAM_INVITE_STATE_REFUSE DungeonCandidateTeamInviteState = 3
)

// Enum value maps for DungeonCandidateTeamInviteState.
var (
	DungeonCandidateTeamInviteState_name = map[int32]string{
		0: "DUNGEON_CANDIDATE_TEAM_INVITE_STATE_NONE",
		1: "DUNGEON_CANDIDATE_TEAM_INVITE_STATE_SEND",
		2: "DUNGEON_CANDIDATE_TEAM_INVITE_STATE_ACCEPT",
		3: "DUNGEON_CANDIDATE_TEAM_INVITE_STATE_REFUSE",
	}
	DungeonCandidateTeamInviteState_value = map[string]int32{
		"DUNGEON_CANDIDATE_TEAM_INVITE_STATE_NONE":   0,
		"DUNGEON_CANDIDATE_TEAM_INVITE_STATE_SEND":   1,
		"DUNGEON_CANDIDATE_TEAM_INVITE_STATE_ACCEPT": 2,
		"DUNGEON_CANDIDATE_TEAM_INVITE_STATE_REFUSE": 3,
	}
)

func (x DungeonCandidateTeamInviteState) Enum() *DungeonCandidateTeamInviteState {
	p := new(DungeonCandidateTeamInviteState)
	*p = x
	return p
}

func (x DungeonCandidateTeamInviteState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DungeonCandidateTeamInviteState) Descriptor() protoreflect.EnumDescriptor {
	return file_DungeonCandidateTeamInviteState_proto_enumTypes[0].Descriptor()
}

func (DungeonCandidateTeamInviteState) Type() protoreflect.EnumType {
	return &file_DungeonCandidateTeamInviteState_proto_enumTypes[0]
}

func (x DungeonCandidateTeamInviteState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DungeonCandidateTeamInviteState.Descriptor instead.
func (DungeonCandidateTeamInviteState) EnumDescriptor() ([]byte, []int) {
	return file_DungeonCandidateTeamInviteState_proto_rawDescGZIP(), []int{0}
}

var File_DungeonCandidateTeamInviteState_proto protoreflect.FileDescriptor

var file_DungeonCandidateTeamInviteState_proto_rawDesc = []byte{
	0x0a, 0x25, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xdd,
	0x01, 0x0a, 0x1f, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x2c, 0x0a, 0x28, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x43, 0x41,
	0x4e, 0x44, 0x49, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x49, 0x4e, 0x56,
	0x49, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x2c, 0x0a, 0x28, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x44,
	0x49, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x49, 0x4e, 0x56, 0x49, 0x54,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x45, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x2e,
	0x0a, 0x2a, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x44, 0x49, 0x44,
	0x41, 0x54, 0x45, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x10, 0x02, 0x12, 0x2e,
	0x0a, 0x2a, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x44, 0x49, 0x44,
	0x41, 0x54, 0x45, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x46, 0x55, 0x53, 0x45, 0x10, 0x03, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_DungeonCandidateTeamInviteState_proto_rawDescOnce sync.Once
	file_DungeonCandidateTeamInviteState_proto_rawDescData = file_DungeonCandidateTeamInviteState_proto_rawDesc
)

func file_DungeonCandidateTeamInviteState_proto_rawDescGZIP() []byte {
	file_DungeonCandidateTeamInviteState_proto_rawDescOnce.Do(func() {
		file_DungeonCandidateTeamInviteState_proto_rawDescData = protoimpl.X.CompressGZIP(file_DungeonCandidateTeamInviteState_proto_rawDescData)
	})
	return file_DungeonCandidateTeamInviteState_proto_rawDescData
}

var file_DungeonCandidateTeamInviteState_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DungeonCandidateTeamInviteState_proto_goTypes = []interface{}{
	(DungeonCandidateTeamInviteState)(0), // 0: proto.DungeonCandidateTeamInviteState
}
var file_DungeonCandidateTeamInviteState_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DungeonCandidateTeamInviteState_proto_init() }
func file_DungeonCandidateTeamInviteState_proto_init() {
	if File_DungeonCandidateTeamInviteState_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_DungeonCandidateTeamInviteState_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DungeonCandidateTeamInviteState_proto_goTypes,
		DependencyIndexes: file_DungeonCandidateTeamInviteState_proto_depIdxs,
		EnumInfos:         file_DungeonCandidateTeamInviteState_proto_enumTypes,
	}.Build()
	File_DungeonCandidateTeamInviteState_proto = out.File
	file_DungeonCandidateTeamInviteState_proto_rawDesc = nil
	file_DungeonCandidateTeamInviteState_proto_goTypes = nil
	file_DungeonCandidateTeamInviteState_proto_depIdxs = nil
}
