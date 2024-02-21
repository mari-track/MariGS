// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: PlayerPropNotify.proto

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

type PlayerPropNotify_CmdId int32

const (
	PlayerPropNotify_NONE             PlayerPropNotify_CmdId = 0
	PlayerPropNotify_CMD_ID           PlayerPropNotify_CmdId = 112
	PlayerPropNotify_ENET_CHANNEL_ID  PlayerPropNotify_CmdId = 0
	PlayerPropNotify_ENET_IS_RELIABLE PlayerPropNotify_CmdId = 1
)

// Enum value maps for PlayerPropNotify_CmdId.
var (
	PlayerPropNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		112: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	PlayerPropNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           112,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x PlayerPropNotify_CmdId) Enum() *PlayerPropNotify_CmdId {
	p := new(PlayerPropNotify_CmdId)
	*p = x
	return p
}

func (x PlayerPropNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerPropNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerPropNotify_proto_enumTypes[0].Descriptor()
}

func (PlayerPropNotify_CmdId) Type() protoreflect.EnumType {
	return &file_PlayerPropNotify_proto_enumTypes[0]
}

func (x PlayerPropNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerPropNotify_CmdId.Descriptor instead.
func (PlayerPropNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PlayerPropNotify_proto_rawDescGZIP(), []int{0, 0}
}

type PlayerPropNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropMap map[uint32]*PropValue `protobuf:"bytes,1,rep,name=prop_map,json=propMap,proto3" json:"prop_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PlayerPropNotify) Reset() {
	*x = PlayerPropNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerPropNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerPropNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerPropNotify) ProtoMessage() {}

func (x *PlayerPropNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerPropNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerPropNotify.ProtoReflect.Descriptor instead.
func (*PlayerPropNotify) Descriptor() ([]byte, []int) {
	return file_PlayerPropNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerPropNotify) GetPropMap() map[uint32]*PropValue {
	if x != nil {
		return x.PropMap
	}
	return nil
}

var File_PlayerPropNotify_proto protoreflect.FileDescriptor

var file_PlayerPropNotify_proto_rawDesc = []byte{
	0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x01, 0x0a, 0x10, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x39,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x1a, 0x46, 0x0a, 0x0c, 0x50, 0x72, 0x6f,
	0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x50, 0x72, 0x6f,
	0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x4c, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x70,
	0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c,
	0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53,
	0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_PlayerPropNotify_proto_rawDescOnce sync.Once
	file_PlayerPropNotify_proto_rawDescData = file_PlayerPropNotify_proto_rawDesc
)

func file_PlayerPropNotify_proto_rawDescGZIP() []byte {
	file_PlayerPropNotify_proto_rawDescOnce.Do(func() {
		file_PlayerPropNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerPropNotify_proto_rawDescData)
	})
	return file_PlayerPropNotify_proto_rawDescData
}

var file_PlayerPropNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerPropNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_PlayerPropNotify_proto_goTypes = []interface{}{
	(PlayerPropNotify_CmdId)(0), // 0: PlayerPropNotify.CmdId
	(*PlayerPropNotify)(nil),    // 1: PlayerPropNotify
	nil,                         // 2: PlayerPropNotify.PropMapEntry
	(*PropValue)(nil),           // 3: PropValue
}
var file_PlayerPropNotify_proto_depIdxs = []int32{
	2, // 0: PlayerPropNotify.prop_map:type_name -> PlayerPropNotify.PropMapEntry
	3, // 1: PlayerPropNotify.PropMapEntry.value:type_name -> PropValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_PlayerPropNotify_proto_init() }
func file_PlayerPropNotify_proto_init() {
	if File_PlayerPropNotify_proto != nil {
		return
	}
	file_PropValue_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerPropNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerPropNotify); i {
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
			RawDescriptor: file_PlayerPropNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerPropNotify_proto_goTypes,
		DependencyIndexes: file_PlayerPropNotify_proto_depIdxs,
		EnumInfos:         file_PlayerPropNotify_proto_enumTypes,
		MessageInfos:      file_PlayerPropNotify_proto_msgTypes,
	}.Build()
	File_PlayerPropNotify_proto = out.File
	file_PlayerPropNotify_proto_rawDesc = nil
	file_PlayerPropNotify_proto_goTypes = nil
	file_PlayerPropNotify_proto_depIdxs = nil
}
