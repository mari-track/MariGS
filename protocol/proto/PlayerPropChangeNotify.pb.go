// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: PlayerPropChangeNotify.proto

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

type PlayerPropChangeNotify_CmdId int32

const (
	PlayerPropChangeNotify_NONE             PlayerPropChangeNotify_CmdId = 0
	PlayerPropChangeNotify_CMD_ID           PlayerPropChangeNotify_CmdId = 119
	PlayerPropChangeNotify_ENET_CHANNEL_ID  PlayerPropChangeNotify_CmdId = 0
	PlayerPropChangeNotify_ENET_IS_RELIABLE PlayerPropChangeNotify_CmdId = 1
)

// Enum value maps for PlayerPropChangeNotify_CmdId.
var (
	PlayerPropChangeNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		119: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	PlayerPropChangeNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           119,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x PlayerPropChangeNotify_CmdId) Enum() *PlayerPropChangeNotify_CmdId {
	p := new(PlayerPropChangeNotify_CmdId)
	*p = x
	return p
}

func (x PlayerPropChangeNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerPropChangeNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerPropChangeNotify_proto_enumTypes[0].Descriptor()
}

func (PlayerPropChangeNotify_CmdId) Type() protoreflect.EnumType {
	return &file_PlayerPropChangeNotify_proto_enumTypes[0]
}

func (x PlayerPropChangeNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerPropChangeNotify_CmdId.Descriptor instead.
func (PlayerPropChangeNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PlayerPropChangeNotify_proto_rawDescGZIP(), []int{0, 0}
}

type PlayerPropChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropType  uint32 `protobuf:"varint,1,opt,name=prop_type,json=propType,proto3" json:"prop_type,omitempty"`
	PropDelta uint32 `protobuf:"varint,2,opt,name=prop_delta,json=propDelta,proto3" json:"prop_delta,omitempty"`
}

func (x *PlayerPropChangeNotify) Reset() {
	*x = PlayerPropChangeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerPropChangeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerPropChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerPropChangeNotify) ProtoMessage() {}

func (x *PlayerPropChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerPropChangeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerPropChangeNotify.ProtoReflect.Descriptor instead.
func (*PlayerPropChangeNotify) Descriptor() ([]byte, []int) {
	return file_PlayerPropChangeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerPropChangeNotify) GetPropType() uint32 {
	if x != nil {
		return x.PropType
	}
	return 0
}

func (x *PlayerPropChangeNotify) GetPropDelta() uint32 {
	if x != nil {
		return x.PropDelta
	}
	return 0
}

var File_PlayerPropChangeNotify_proto protoreflect.FileDescriptor

var file_PlayerPropChangeNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x22, 0x4c, 0x0a, 0x05,
	0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x77, 0x12, 0x13, 0x0a, 0x0f, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerPropChangeNotify_proto_rawDescOnce sync.Once
	file_PlayerPropChangeNotify_proto_rawDescData = file_PlayerPropChangeNotify_proto_rawDesc
)

func file_PlayerPropChangeNotify_proto_rawDescGZIP() []byte {
	file_PlayerPropChangeNotify_proto_rawDescOnce.Do(func() {
		file_PlayerPropChangeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerPropChangeNotify_proto_rawDescData)
	})
	return file_PlayerPropChangeNotify_proto_rawDescData
}

var file_PlayerPropChangeNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerPropChangeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerPropChangeNotify_proto_goTypes = []interface{}{
	(PlayerPropChangeNotify_CmdId)(0), // 0: proto.PlayerPropChangeNotify.CmdId
	(*PlayerPropChangeNotify)(nil),    // 1: proto.PlayerPropChangeNotify
}
var file_PlayerPropChangeNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PlayerPropChangeNotify_proto_init() }
func file_PlayerPropChangeNotify_proto_init() {
	if File_PlayerPropChangeNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PlayerPropChangeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerPropChangeNotify); i {
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
			RawDescriptor: file_PlayerPropChangeNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerPropChangeNotify_proto_goTypes,
		DependencyIndexes: file_PlayerPropChangeNotify_proto_depIdxs,
		EnumInfos:         file_PlayerPropChangeNotify_proto_enumTypes,
		MessageInfos:      file_PlayerPropChangeNotify_proto_msgTypes,
	}.Build()
	File_PlayerPropChangeNotify_proto = out.File
	file_PlayerPropChangeNotify_proto_rawDesc = nil
	file_PlayerPropChangeNotify_proto_goTypes = nil
	file_PlayerPropChangeNotify_proto_depIdxs = nil
}
