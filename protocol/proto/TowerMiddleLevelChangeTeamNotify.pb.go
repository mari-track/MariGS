// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: TowerMiddleLevelChangeTeamNotify.proto

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

type TowerMiddleLevelChangeTeamNotify_CmdId int32

const (
	TowerMiddleLevelChangeTeamNotify_NONE             TowerMiddleLevelChangeTeamNotify_CmdId = 0
	TowerMiddleLevelChangeTeamNotify_CMD_ID           TowerMiddleLevelChangeTeamNotify_CmdId = 2432
	TowerMiddleLevelChangeTeamNotify_ENET_CHANNEL_ID  TowerMiddleLevelChangeTeamNotify_CmdId = 0
	TowerMiddleLevelChangeTeamNotify_ENET_IS_RELIABLE TowerMiddleLevelChangeTeamNotify_CmdId = 1
)

// Enum value maps for TowerMiddleLevelChangeTeamNotify_CmdId.
var (
	TowerMiddleLevelChangeTeamNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		2432: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	TowerMiddleLevelChangeTeamNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2432,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x TowerMiddleLevelChangeTeamNotify_CmdId) Enum() *TowerMiddleLevelChangeTeamNotify_CmdId {
	p := new(TowerMiddleLevelChangeTeamNotify_CmdId)
	*p = x
	return p
}

func (x TowerMiddleLevelChangeTeamNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TowerMiddleLevelChangeTeamNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_TowerMiddleLevelChangeTeamNotify_proto_enumTypes[0].Descriptor()
}

func (TowerMiddleLevelChangeTeamNotify_CmdId) Type() protoreflect.EnumType {
	return &file_TowerMiddleLevelChangeTeamNotify_proto_enumTypes[0]
}

func (x TowerMiddleLevelChangeTeamNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TowerMiddleLevelChangeTeamNotify_CmdId.Descriptor instead.
func (TowerMiddleLevelChangeTeamNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_TowerMiddleLevelChangeTeamNotify_proto_rawDescGZIP(), []int{0, 0}
}

type TowerMiddleLevelChangeTeamNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TowerMiddleLevelChangeTeamNotify) Reset() {
	*x = TowerMiddleLevelChangeTeamNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TowerMiddleLevelChangeTeamNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TowerMiddleLevelChangeTeamNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TowerMiddleLevelChangeTeamNotify) ProtoMessage() {}

func (x *TowerMiddleLevelChangeTeamNotify) ProtoReflect() protoreflect.Message {
	mi := &file_TowerMiddleLevelChangeTeamNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TowerMiddleLevelChangeTeamNotify.ProtoReflect.Descriptor instead.
func (*TowerMiddleLevelChangeTeamNotify) Descriptor() ([]byte, []int) {
	return file_TowerMiddleLevelChangeTeamNotify_proto_rawDescGZIP(), []int{0}
}

var File_TowerMiddleLevelChangeTeamNotify_proto protoreflect.FileDescriptor

var file_TowerMiddleLevelChangeTeamNotify_proto_rawDesc = []byte{
	0x0a, 0x26, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x71, 0x0a, 0x20, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44,
	0x10, 0x80, 0x13, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02,
	0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TowerMiddleLevelChangeTeamNotify_proto_rawDescOnce sync.Once
	file_TowerMiddleLevelChangeTeamNotify_proto_rawDescData = file_TowerMiddleLevelChangeTeamNotify_proto_rawDesc
)

func file_TowerMiddleLevelChangeTeamNotify_proto_rawDescGZIP() []byte {
	file_TowerMiddleLevelChangeTeamNotify_proto_rawDescOnce.Do(func() {
		file_TowerMiddleLevelChangeTeamNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_TowerMiddleLevelChangeTeamNotify_proto_rawDescData)
	})
	return file_TowerMiddleLevelChangeTeamNotify_proto_rawDescData
}

var file_TowerMiddleLevelChangeTeamNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_TowerMiddleLevelChangeTeamNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TowerMiddleLevelChangeTeamNotify_proto_goTypes = []interface{}{
	(TowerMiddleLevelChangeTeamNotify_CmdId)(0), // 0: proto.TowerMiddleLevelChangeTeamNotify.CmdId
	(*TowerMiddleLevelChangeTeamNotify)(nil),    // 1: proto.TowerMiddleLevelChangeTeamNotify
}
var file_TowerMiddleLevelChangeTeamNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TowerMiddleLevelChangeTeamNotify_proto_init() }
func file_TowerMiddleLevelChangeTeamNotify_proto_init() {
	if File_TowerMiddleLevelChangeTeamNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TowerMiddleLevelChangeTeamNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TowerMiddleLevelChangeTeamNotify); i {
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
			RawDescriptor: file_TowerMiddleLevelChangeTeamNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TowerMiddleLevelChangeTeamNotify_proto_goTypes,
		DependencyIndexes: file_TowerMiddleLevelChangeTeamNotify_proto_depIdxs,
		EnumInfos:         file_TowerMiddleLevelChangeTeamNotify_proto_enumTypes,
		MessageInfos:      file_TowerMiddleLevelChangeTeamNotify_proto_msgTypes,
	}.Build()
	File_TowerMiddleLevelChangeTeamNotify_proto = out.File
	file_TowerMiddleLevelChangeTeamNotify_proto_rawDesc = nil
	file_TowerMiddleLevelChangeTeamNotify_proto_goTypes = nil
	file_TowerMiddleLevelChangeTeamNotify_proto_depIdxs = nil
}
