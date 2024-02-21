// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: DungeonChallengeBeginNotify.proto

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

type DungeonChallengeBeginNotify_CmdId int32

const (
	DungeonChallengeBeginNotify_NONE             DungeonChallengeBeginNotify_CmdId = 0
	DungeonChallengeBeginNotify_CMD_ID           DungeonChallengeBeginNotify_CmdId = 918
	DungeonChallengeBeginNotify_ENET_CHANNEL_ID  DungeonChallengeBeginNotify_CmdId = 0
	DungeonChallengeBeginNotify_ENET_IS_RELIABLE DungeonChallengeBeginNotify_CmdId = 1
)

// Enum value maps for DungeonChallengeBeginNotify_CmdId.
var (
	DungeonChallengeBeginNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		918: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	DungeonChallengeBeginNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           918,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x DungeonChallengeBeginNotify_CmdId) Enum() *DungeonChallengeBeginNotify_CmdId {
	p := new(DungeonChallengeBeginNotify_CmdId)
	*p = x
	return p
}

func (x DungeonChallengeBeginNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DungeonChallengeBeginNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_DungeonChallengeBeginNotify_proto_enumTypes[0].Descriptor()
}

func (DungeonChallengeBeginNotify_CmdId) Type() protoreflect.EnumType {
	return &file_DungeonChallengeBeginNotify_proto_enumTypes[0]
}

func (x DungeonChallengeBeginNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DungeonChallengeBeginNotify_CmdId.Descriptor instead.
func (DungeonChallengeBeginNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_DungeonChallengeBeginNotify_proto_rawDescGZIP(), []int{0, 0}
}

type DungeonChallengeBeginNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeId    uint32   `protobuf:"varint,1,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	ChallengeIndex uint32   `protobuf:"varint,2,opt,name=challenge_index,json=challengeIndex,proto3" json:"challenge_index,omitempty"`
	ParamList      []uint32 `protobuf:"varint,3,rep,packed,name=param_list,json=paramList,proto3" json:"param_list,omitempty"`
	GroupId        uint32   `protobuf:"varint,4,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *DungeonChallengeBeginNotify) Reset() {
	*x = DungeonChallengeBeginNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DungeonChallengeBeginNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DungeonChallengeBeginNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DungeonChallengeBeginNotify) ProtoMessage() {}

func (x *DungeonChallengeBeginNotify) ProtoReflect() protoreflect.Message {
	mi := &file_DungeonChallengeBeginNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DungeonChallengeBeginNotify.ProtoReflect.Descriptor instead.
func (*DungeonChallengeBeginNotify) Descriptor() ([]byte, []int) {
	return file_DungeonChallengeBeginNotify_proto_rawDescGZIP(), []int{0}
}

func (x *DungeonChallengeBeginNotify) GetChallengeId() uint32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

func (x *DungeonChallengeBeginNotify) GetChallengeIndex() uint32 {
	if x != nil {
		return x.ChallengeIndex
	}
	return 0
}

func (x *DungeonChallengeBeginNotify) GetParamList() []uint32 {
	if x != nil {
		return x.ParamList
	}
	return nil
}

func (x *DungeonChallengeBeginNotify) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

var File_DungeonChallengeBeginNotify_proto protoreflect.FileDescriptor

var file_DungeonChallengeBeginNotify_proto_rawDesc = []byte{
	0x0a, 0x21, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x01, 0x0a, 0x1b, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64,
	0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06,
	0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x96, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45,
	0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DungeonChallengeBeginNotify_proto_rawDescOnce sync.Once
	file_DungeonChallengeBeginNotify_proto_rawDescData = file_DungeonChallengeBeginNotify_proto_rawDesc
)

func file_DungeonChallengeBeginNotify_proto_rawDescGZIP() []byte {
	file_DungeonChallengeBeginNotify_proto_rawDescOnce.Do(func() {
		file_DungeonChallengeBeginNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_DungeonChallengeBeginNotify_proto_rawDescData)
	})
	return file_DungeonChallengeBeginNotify_proto_rawDescData
}

var file_DungeonChallengeBeginNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DungeonChallengeBeginNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DungeonChallengeBeginNotify_proto_goTypes = []interface{}{
	(DungeonChallengeBeginNotify_CmdId)(0), // 0: DungeonChallengeBeginNotify.CmdId
	(*DungeonChallengeBeginNotify)(nil),    // 1: DungeonChallengeBeginNotify
}
var file_DungeonChallengeBeginNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DungeonChallengeBeginNotify_proto_init() }
func file_DungeonChallengeBeginNotify_proto_init() {
	if File_DungeonChallengeBeginNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DungeonChallengeBeginNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DungeonChallengeBeginNotify); i {
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
			RawDescriptor: file_DungeonChallengeBeginNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DungeonChallengeBeginNotify_proto_goTypes,
		DependencyIndexes: file_DungeonChallengeBeginNotify_proto_depIdxs,
		EnumInfos:         file_DungeonChallengeBeginNotify_proto_enumTypes,
		MessageInfos:      file_DungeonChallengeBeginNotify_proto_msgTypes,
	}.Build()
	File_DungeonChallengeBeginNotify_proto = out.File
	file_DungeonChallengeBeginNotify_proto_rawDesc = nil
	file_DungeonChallengeBeginNotify_proto_goTypes = nil
	file_DungeonChallengeBeginNotify_proto_depIdxs = nil
}
