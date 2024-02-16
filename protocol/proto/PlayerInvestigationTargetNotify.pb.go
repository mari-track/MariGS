// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: PlayerInvestigationTargetNotify.proto

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

type PlayerInvestigationTargetNotify_CmdId int32

const (
	PlayerInvestigationTargetNotify_NONE             PlayerInvestigationTargetNotify_CmdId = 0
	PlayerInvestigationTargetNotify_CMD_ID           PlayerInvestigationTargetNotify_CmdId = 1909
	PlayerInvestigationTargetNotify_ENET_CHANNEL_ID  PlayerInvestigationTargetNotify_CmdId = 0
	PlayerInvestigationTargetNotify_ENET_IS_RELIABLE PlayerInvestigationTargetNotify_CmdId = 1
)

// Enum value maps for PlayerInvestigationTargetNotify_CmdId.
var (
	PlayerInvestigationTargetNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		1909: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	PlayerInvestigationTargetNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           1909,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x PlayerInvestigationTargetNotify_CmdId) Enum() *PlayerInvestigationTargetNotify_CmdId {
	p := new(PlayerInvestigationTargetNotify_CmdId)
	*p = x
	return p
}

func (x PlayerInvestigationTargetNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerInvestigationTargetNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerInvestigationTargetNotify_proto_enumTypes[0].Descriptor()
}

func (PlayerInvestigationTargetNotify_CmdId) Type() protoreflect.EnumType {
	return &file_PlayerInvestigationTargetNotify_proto_enumTypes[0]
}

func (x PlayerInvestigationTargetNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerInvestigationTargetNotify_CmdId.Descriptor instead.
func (PlayerInvestigationTargetNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PlayerInvestigationTargetNotify_proto_rawDescGZIP(), []int{0, 0}
}

type PlayerInvestigationTargetNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvestigationTargetList []*InvestigationTarget `protobuf:"bytes,1,rep,name=investigation_target_list,json=investigationTargetList,proto3" json:"investigation_target_list,omitempty"`
}

func (x *PlayerInvestigationTargetNotify) Reset() {
	*x = PlayerInvestigationTargetNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerInvestigationTargetNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInvestigationTargetNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInvestigationTargetNotify) ProtoMessage() {}

func (x *PlayerInvestigationTargetNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerInvestigationTargetNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInvestigationTargetNotify.ProtoReflect.Descriptor instead.
func (*PlayerInvestigationTargetNotify) Descriptor() ([]byte, []int) {
	return file_PlayerInvestigationTargetNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerInvestigationTargetNotify) GetInvestigationTargetList() []*InvestigationTarget {
	if x != nil {
		return x.InvestigationTargetList
	}
	return nil
}

var File_PlayerInvestigationTargetNotify_proto protoreflect.FileDescriptor

var file_PlayerInvestigationTargetNotify_proto_rawDesc = []byte{
	0x0a, 0x25, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x1f, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x56, 0x0a,
	0x19, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x17, 0x69, 0x6e,
	0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f,
	0x49, 0x44, 0x10, 0xf5, 0x0e, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01,
	0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerInvestigationTargetNotify_proto_rawDescOnce sync.Once
	file_PlayerInvestigationTargetNotify_proto_rawDescData = file_PlayerInvestigationTargetNotify_proto_rawDesc
)

func file_PlayerInvestigationTargetNotify_proto_rawDescGZIP() []byte {
	file_PlayerInvestigationTargetNotify_proto_rawDescOnce.Do(func() {
		file_PlayerInvestigationTargetNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerInvestigationTargetNotify_proto_rawDescData)
	})
	return file_PlayerInvestigationTargetNotify_proto_rawDescData
}

var file_PlayerInvestigationTargetNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerInvestigationTargetNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerInvestigationTargetNotify_proto_goTypes = []interface{}{
	(PlayerInvestigationTargetNotify_CmdId)(0), // 0: proto.PlayerInvestigationTargetNotify.CmdId
	(*PlayerInvestigationTargetNotify)(nil),    // 1: proto.PlayerInvestigationTargetNotify
	(*InvestigationTarget)(nil),                // 2: proto.InvestigationTarget
}
var file_PlayerInvestigationTargetNotify_proto_depIdxs = []int32{
	2, // 0: proto.PlayerInvestigationTargetNotify.investigation_target_list:type_name -> proto.InvestigationTarget
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerInvestigationTargetNotify_proto_init() }
func file_PlayerInvestigationTargetNotify_proto_init() {
	if File_PlayerInvestigationTargetNotify_proto != nil {
		return
	}
	file_InvestigationTarget_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerInvestigationTargetNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInvestigationTargetNotify); i {
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
			RawDescriptor: file_PlayerInvestigationTargetNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerInvestigationTargetNotify_proto_goTypes,
		DependencyIndexes: file_PlayerInvestigationTargetNotify_proto_depIdxs,
		EnumInfos:         file_PlayerInvestigationTargetNotify_proto_enumTypes,
		MessageInfos:      file_PlayerInvestigationTargetNotify_proto_msgTypes,
	}.Build()
	File_PlayerInvestigationTargetNotify_proto = out.File
	file_PlayerInvestigationTargetNotify_proto_rawDesc = nil
	file_PlayerInvestigationTargetNotify_proto_goTypes = nil
	file_PlayerInvestigationTargetNotify_proto_depIdxs = nil
}
