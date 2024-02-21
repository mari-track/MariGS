// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: PlayerInvestigationNotify.proto

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

type PlayerInvestigationNotify_CmdId int32

const (
	PlayerInvestigationNotify_NONE             PlayerInvestigationNotify_CmdId = 0
	PlayerInvestigationNotify_CMD_ID           PlayerInvestigationNotify_CmdId = 1908
	PlayerInvestigationNotify_ENET_CHANNEL_ID  PlayerInvestigationNotify_CmdId = 0
	PlayerInvestigationNotify_ENET_IS_RELIABLE PlayerInvestigationNotify_CmdId = 1
)

// Enum value maps for PlayerInvestigationNotify_CmdId.
var (
	PlayerInvestigationNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		1908: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	PlayerInvestigationNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           1908,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x PlayerInvestigationNotify_CmdId) Enum() *PlayerInvestigationNotify_CmdId {
	p := new(PlayerInvestigationNotify_CmdId)
	*p = x
	return p
}

func (x PlayerInvestigationNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerInvestigationNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerInvestigationNotify_proto_enumTypes[0].Descriptor()
}

func (PlayerInvestigationNotify_CmdId) Type() protoreflect.EnumType {
	return &file_PlayerInvestigationNotify_proto_enumTypes[0]
}

func (x PlayerInvestigationNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerInvestigationNotify_CmdId.Descriptor instead.
func (PlayerInvestigationNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PlayerInvestigationNotify_proto_rawDescGZIP(), []int{0, 0}
}

type PlayerInvestigationNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvestigationList []*Investigation `protobuf:"bytes,1,rep,name=investigation_list,json=investigationList,proto3" json:"investigation_list,omitempty"`
}

func (x *PlayerInvestigationNotify) Reset() {
	*x = PlayerInvestigationNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerInvestigationNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInvestigationNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInvestigationNotify) ProtoMessage() {}

func (x *PlayerInvestigationNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerInvestigationNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInvestigationNotify.ProtoReflect.Descriptor instead.
func (*PlayerInvestigationNotify) Descriptor() ([]byte, []int) {
	return file_PlayerInvestigationNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerInvestigationNotify) GetInvestigationList() []*Investigation {
	if x != nil {
		return x.InvestigationList
	}
	return nil
}

var File_PlayerInvestigationNotify_proto protoreflect.FileDescriptor

var file_PlayerInvestigationNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x19, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x3d, 0x0a, 0x12, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x11, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44,
	0x10, 0xf4, 0x0e, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02,
	0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerInvestigationNotify_proto_rawDescOnce sync.Once
	file_PlayerInvestigationNotify_proto_rawDescData = file_PlayerInvestigationNotify_proto_rawDesc
)

func file_PlayerInvestigationNotify_proto_rawDescGZIP() []byte {
	file_PlayerInvestigationNotify_proto_rawDescOnce.Do(func() {
		file_PlayerInvestigationNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerInvestigationNotify_proto_rawDescData)
	})
	return file_PlayerInvestigationNotify_proto_rawDescData
}

var file_PlayerInvestigationNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerInvestigationNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerInvestigationNotify_proto_goTypes = []interface{}{
	(PlayerInvestigationNotify_CmdId)(0), // 0: PlayerInvestigationNotify.CmdId
	(*PlayerInvestigationNotify)(nil),    // 1: PlayerInvestigationNotify
	(*Investigation)(nil),                // 2: Investigation
}
var file_PlayerInvestigationNotify_proto_depIdxs = []int32{
	2, // 0: PlayerInvestigationNotify.investigation_list:type_name -> Investigation
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerInvestigationNotify_proto_init() }
func file_PlayerInvestigationNotify_proto_init() {
	if File_PlayerInvestigationNotify_proto != nil {
		return
	}
	file_Investigation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerInvestigationNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInvestigationNotify); i {
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
			RawDescriptor: file_PlayerInvestigationNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerInvestigationNotify_proto_goTypes,
		DependencyIndexes: file_PlayerInvestigationNotify_proto_depIdxs,
		EnumInfos:         file_PlayerInvestigationNotify_proto_enumTypes,
		MessageInfos:      file_PlayerInvestigationNotify_proto_msgTypes,
	}.Build()
	File_PlayerInvestigationNotify_proto = out.File
	file_PlayerInvestigationNotify_proto_rawDesc = nil
	file_PlayerInvestigationNotify_proto_goTypes = nil
	file_PlayerInvestigationNotify_proto_depIdxs = nil
}
