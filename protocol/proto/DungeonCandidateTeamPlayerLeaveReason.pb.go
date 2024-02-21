// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: DungeonCandidateTeamPlayerLeaveReason.proto

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

type DungeonCandidateTeamPlayerLeaveReason int32

const (
	DungeonCandidateTeamPlayerLeaveReason_DUNGEON_CANDIDATE_TPLR_NORMAL  DungeonCandidateTeamPlayerLeaveReason = 0
	DungeonCandidateTeamPlayerLeaveReason_DUNGEON_CANDIDATE_TPLR_DIE     DungeonCandidateTeamPlayerLeaveReason = 1
	DungeonCandidateTeamPlayerLeaveReason_DUNGEON_CANDIDATE_TPLR_BE_KICK DungeonCandidateTeamPlayerLeaveReason = 2
	DungeonCandidateTeamPlayerLeaveReason_DUNGEON_CANDIDATE_DISCONNECT   DungeonCandidateTeamPlayerLeaveReason = 3
)

// Enum value maps for DungeonCandidateTeamPlayerLeaveReason.
var (
	DungeonCandidateTeamPlayerLeaveReason_name = map[int32]string{
		0: "DUNGEON_CANDIDATE_TPLR_NORMAL",
		1: "DUNGEON_CANDIDATE_TPLR_DIE",
		2: "DUNGEON_CANDIDATE_TPLR_BE_KICK",
		3: "DUNGEON_CANDIDATE_DISCONNECT",
	}
	DungeonCandidateTeamPlayerLeaveReason_value = map[string]int32{
		"DUNGEON_CANDIDATE_TPLR_NORMAL":  0,
		"DUNGEON_CANDIDATE_TPLR_DIE":     1,
		"DUNGEON_CANDIDATE_TPLR_BE_KICK": 2,
		"DUNGEON_CANDIDATE_DISCONNECT":   3,
	}
)

func (x DungeonCandidateTeamPlayerLeaveReason) Enum() *DungeonCandidateTeamPlayerLeaveReason {
	p := new(DungeonCandidateTeamPlayerLeaveReason)
	*p = x
	return p
}

func (x DungeonCandidateTeamPlayerLeaveReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DungeonCandidateTeamPlayerLeaveReason) Descriptor() protoreflect.EnumDescriptor {
	return file_DungeonCandidateTeamPlayerLeaveReason_proto_enumTypes[0].Descriptor()
}

func (DungeonCandidateTeamPlayerLeaveReason) Type() protoreflect.EnumType {
	return &file_DungeonCandidateTeamPlayerLeaveReason_proto_enumTypes[0]
}

func (x DungeonCandidateTeamPlayerLeaveReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DungeonCandidateTeamPlayerLeaveReason.Descriptor instead.
func (DungeonCandidateTeamPlayerLeaveReason) EnumDescriptor() ([]byte, []int) {
	return file_DungeonCandidateTeamPlayerLeaveReason_proto_rawDescGZIP(), []int{0}
}

var File_DungeonCandidateTeamPlayerLeaveReason_proto protoreflect.FileDescriptor

var file_DungeonCandidateTeamPlayerLeaveReason_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76,
	0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xb0, 0x01,
	0x0a, 0x25, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76,
	0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x55, 0x4e, 0x47, 0x45,
	0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x44, 0x49, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x50, 0x4c,
	0x52, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x55,
	0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x44, 0x49, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x54, 0x50, 0x4c, 0x52, 0x5f, 0x44, 0x49, 0x45, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x44, 0x55,
	0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x44, 0x49, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x54, 0x50, 0x4c, 0x52, 0x5f, 0x42, 0x45, 0x5f, 0x4b, 0x49, 0x43, 0x4b, 0x10, 0x02, 0x12, 0x20,
	0x0a, 0x1c, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x44, 0x49, 0x44,
	0x41, 0x54, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x03,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DungeonCandidateTeamPlayerLeaveReason_proto_rawDescOnce sync.Once
	file_DungeonCandidateTeamPlayerLeaveReason_proto_rawDescData = file_DungeonCandidateTeamPlayerLeaveReason_proto_rawDesc
)

func file_DungeonCandidateTeamPlayerLeaveReason_proto_rawDescGZIP() []byte {
	file_DungeonCandidateTeamPlayerLeaveReason_proto_rawDescOnce.Do(func() {
		file_DungeonCandidateTeamPlayerLeaveReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_DungeonCandidateTeamPlayerLeaveReason_proto_rawDescData)
	})
	return file_DungeonCandidateTeamPlayerLeaveReason_proto_rawDescData
}

var file_DungeonCandidateTeamPlayerLeaveReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DungeonCandidateTeamPlayerLeaveReason_proto_goTypes = []interface{}{
	(DungeonCandidateTeamPlayerLeaveReason)(0), // 0: DungeonCandidateTeamPlayerLeaveReason
}
var file_DungeonCandidateTeamPlayerLeaveReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DungeonCandidateTeamPlayerLeaveReason_proto_init() }
func file_DungeonCandidateTeamPlayerLeaveReason_proto_init() {
	if File_DungeonCandidateTeamPlayerLeaveReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_DungeonCandidateTeamPlayerLeaveReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DungeonCandidateTeamPlayerLeaveReason_proto_goTypes,
		DependencyIndexes: file_DungeonCandidateTeamPlayerLeaveReason_proto_depIdxs,
		EnumInfos:         file_DungeonCandidateTeamPlayerLeaveReason_proto_enumTypes,
	}.Build()
	File_DungeonCandidateTeamPlayerLeaveReason_proto = out.File
	file_DungeonCandidateTeamPlayerLeaveReason_proto_rawDesc = nil
	file_DungeonCandidateTeamPlayerLeaveReason_proto_goTypes = nil
	file_DungeonCandidateTeamPlayerLeaveReason_proto_depIdxs = nil
}
