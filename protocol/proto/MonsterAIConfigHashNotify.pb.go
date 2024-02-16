// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: MonsterAIConfigHashNotify.proto

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

type MonsterAIConfigHashNotify_CmdId int32

const (
	MonsterAIConfigHashNotify_NONE             MonsterAIConfigHashNotify_CmdId = 0
	MonsterAIConfigHashNotify_CMD_ID           MonsterAIConfigHashNotify_CmdId = 3044
	MonsterAIConfigHashNotify_ENET_CHANNEL_ID  MonsterAIConfigHashNotify_CmdId = 0
	MonsterAIConfigHashNotify_ENET_IS_RELIABLE MonsterAIConfigHashNotify_CmdId = 1
	MonsterAIConfigHashNotify_IS_ALLOW_CLIENT  MonsterAIConfigHashNotify_CmdId = 1
)

// Enum value maps for MonsterAIConfigHashNotify_CmdId.
var (
	MonsterAIConfigHashNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		3044: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	MonsterAIConfigHashNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           3044,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x MonsterAIConfigHashNotify_CmdId) Enum() *MonsterAIConfigHashNotify_CmdId {
	p := new(MonsterAIConfigHashNotify_CmdId)
	*p = x
	return p
}

func (x MonsterAIConfigHashNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MonsterAIConfigHashNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_MonsterAIConfigHashNotify_proto_enumTypes[0].Descriptor()
}

func (MonsterAIConfigHashNotify_CmdId) Type() protoreflect.EnumType {
	return &file_MonsterAIConfigHashNotify_proto_enumTypes[0]
}

func (x MonsterAIConfigHashNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MonsterAIConfigHashNotify_CmdId.Descriptor instead.
func (MonsterAIConfigHashNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_MonsterAIConfigHashNotify_proto_rawDescGZIP(), []int{0, 0}
}

type MonsterAIConfigHashNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId  uint32 `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	HashValue int32  `protobuf:"varint,2,opt,name=hash_value,json=hashValue,proto3" json:"hash_value,omitempty"`
}

func (x *MonsterAIConfigHashNotify) Reset() {
	*x = MonsterAIConfigHashNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MonsterAIConfigHashNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonsterAIConfigHashNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonsterAIConfigHashNotify) ProtoMessage() {}

func (x *MonsterAIConfigHashNotify) ProtoReflect() protoreflect.Message {
	mi := &file_MonsterAIConfigHashNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonsterAIConfigHashNotify.ProtoReflect.Descriptor instead.
func (*MonsterAIConfigHashNotify) Descriptor() ([]byte, []int) {
	return file_MonsterAIConfigHashNotify_proto_rawDescGZIP(), []int{0}
}

func (x *MonsterAIConfigHashNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *MonsterAIConfigHashNotify) GetHashValue() int32 {
	if x != nil {
		return x.HashValue
	}
	return 0
}

var File_MonsterAIConfigHashNotify_proto protoreflect.FileDescriptor

var file_MonsterAIConfigHashNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x41, 0x49, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x48, 0x61, 0x73, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x19, 0x4d, 0x6f, 0x6e,
	0x73, 0x74, 0x65, 0x72, 0x41, 0x49, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x61, 0x73, 0x68,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x68, 0x61, 0x73, 0x68, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10,
	0xe4, 0x17, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e,
	0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f,
	0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a,
	0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54,
	0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MonsterAIConfigHashNotify_proto_rawDescOnce sync.Once
	file_MonsterAIConfigHashNotify_proto_rawDescData = file_MonsterAIConfigHashNotify_proto_rawDesc
)

func file_MonsterAIConfigHashNotify_proto_rawDescGZIP() []byte {
	file_MonsterAIConfigHashNotify_proto_rawDescOnce.Do(func() {
		file_MonsterAIConfigHashNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_MonsterAIConfigHashNotify_proto_rawDescData)
	})
	return file_MonsterAIConfigHashNotify_proto_rawDescData
}

var file_MonsterAIConfigHashNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MonsterAIConfigHashNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MonsterAIConfigHashNotify_proto_goTypes = []interface{}{
	(MonsterAIConfigHashNotify_CmdId)(0), // 0: proto.MonsterAIConfigHashNotify.CmdId
	(*MonsterAIConfigHashNotify)(nil),    // 1: proto.MonsterAIConfigHashNotify
}
var file_MonsterAIConfigHashNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MonsterAIConfigHashNotify_proto_init() }
func file_MonsterAIConfigHashNotify_proto_init() {
	if File_MonsterAIConfigHashNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MonsterAIConfigHashNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonsterAIConfigHashNotify); i {
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
			RawDescriptor: file_MonsterAIConfigHashNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MonsterAIConfigHashNotify_proto_goTypes,
		DependencyIndexes: file_MonsterAIConfigHashNotify_proto_depIdxs,
		EnumInfos:         file_MonsterAIConfigHashNotify_proto_enumTypes,
		MessageInfos:      file_MonsterAIConfigHashNotify_proto_msgTypes,
	}.Build()
	File_MonsterAIConfigHashNotify_proto = out.File
	file_MonsterAIConfigHashNotify_proto_rawDesc = nil
	file_MonsterAIConfigHashNotify_proto_goTypes = nil
	file_MonsterAIConfigHashNotify_proto_depIdxs = nil
}
