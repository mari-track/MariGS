// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: DungeonCandidateTeamKickRsp.proto

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

type DungeonCandidateTeamKickRsp_CmdId int32

const (
	DungeonCandidateTeamKickRsp_NONE             DungeonCandidateTeamKickRsp_CmdId = 0
	DungeonCandidateTeamKickRsp_CMD_ID           DungeonCandidateTeamKickRsp_CmdId = 935
	DungeonCandidateTeamKickRsp_ENET_CHANNEL_ID  DungeonCandidateTeamKickRsp_CmdId = 0
	DungeonCandidateTeamKickRsp_ENET_IS_RELIABLE DungeonCandidateTeamKickRsp_CmdId = 1
)

// Enum value maps for DungeonCandidateTeamKickRsp_CmdId.
var (
	DungeonCandidateTeamKickRsp_CmdId_name = map[int32]string{
		0:   "NONE",
		935: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	DungeonCandidateTeamKickRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           935,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x DungeonCandidateTeamKickRsp_CmdId) Enum() *DungeonCandidateTeamKickRsp_CmdId {
	p := new(DungeonCandidateTeamKickRsp_CmdId)
	*p = x
	return p
}

func (x DungeonCandidateTeamKickRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DungeonCandidateTeamKickRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_DungeonCandidateTeamKickRsp_proto_enumTypes[0].Descriptor()
}

func (DungeonCandidateTeamKickRsp_CmdId) Type() protoreflect.EnumType {
	return &file_DungeonCandidateTeamKickRsp_proto_enumTypes[0]
}

func (x DungeonCandidateTeamKickRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DungeonCandidateTeamKickRsp_CmdId.Descriptor instead.
func (DungeonCandidateTeamKickRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_DungeonCandidateTeamKickRsp_proto_rawDescGZIP(), []int{0, 0}
}

type DungeonCandidateTeamKickRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32 `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *DungeonCandidateTeamKickRsp) Reset() {
	*x = DungeonCandidateTeamKickRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DungeonCandidateTeamKickRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DungeonCandidateTeamKickRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DungeonCandidateTeamKickRsp) ProtoMessage() {}

func (x *DungeonCandidateTeamKickRsp) ProtoReflect() protoreflect.Message {
	mi := &file_DungeonCandidateTeamKickRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DungeonCandidateTeamKickRsp.ProtoReflect.Descriptor instead.
func (*DungeonCandidateTeamKickRsp) Descriptor() ([]byte, []int) {
	return file_DungeonCandidateTeamKickRsp_proto_rawDescGZIP(), []int{0}
}

func (x *DungeonCandidateTeamKickRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_DungeonCandidateTeamKickRsp_proto protoreflect.FileDescriptor

var file_DungeonCandidateTeamKickRsp_proto_rawDesc = []byte{
	0x0a, 0x21, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x4b, 0x69, 0x63, 0x6b, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x1b, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x4b, 0x69, 0x63, 0x6b,
	0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x4d, 0x0a,
	0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xa7, 0x07, 0x12, 0x13, 0x0a,
	0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45,
	0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DungeonCandidateTeamKickRsp_proto_rawDescOnce sync.Once
	file_DungeonCandidateTeamKickRsp_proto_rawDescData = file_DungeonCandidateTeamKickRsp_proto_rawDesc
)

func file_DungeonCandidateTeamKickRsp_proto_rawDescGZIP() []byte {
	file_DungeonCandidateTeamKickRsp_proto_rawDescOnce.Do(func() {
		file_DungeonCandidateTeamKickRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_DungeonCandidateTeamKickRsp_proto_rawDescData)
	})
	return file_DungeonCandidateTeamKickRsp_proto_rawDescData
}

var file_DungeonCandidateTeamKickRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DungeonCandidateTeamKickRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DungeonCandidateTeamKickRsp_proto_goTypes = []interface{}{
	(DungeonCandidateTeamKickRsp_CmdId)(0), // 0: DungeonCandidateTeamKickRsp.CmdId
	(*DungeonCandidateTeamKickRsp)(nil),    // 1: DungeonCandidateTeamKickRsp
}
var file_DungeonCandidateTeamKickRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DungeonCandidateTeamKickRsp_proto_init() }
func file_DungeonCandidateTeamKickRsp_proto_init() {
	if File_DungeonCandidateTeamKickRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DungeonCandidateTeamKickRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DungeonCandidateTeamKickRsp); i {
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
			RawDescriptor: file_DungeonCandidateTeamKickRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DungeonCandidateTeamKickRsp_proto_goTypes,
		DependencyIndexes: file_DungeonCandidateTeamKickRsp_proto_depIdxs,
		EnumInfos:         file_DungeonCandidateTeamKickRsp_proto_enumTypes,
		MessageInfos:      file_DungeonCandidateTeamKickRsp_proto_msgTypes,
	}.Build()
	File_DungeonCandidateTeamKickRsp_proto = out.File
	file_DungeonCandidateTeamKickRsp_proto_rawDesc = nil
	file_DungeonCandidateTeamKickRsp_proto_goTypes = nil
	file_DungeonCandidateTeamKickRsp_proto_depIdxs = nil
}
