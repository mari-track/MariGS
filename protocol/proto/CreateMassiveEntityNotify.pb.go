// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: CreateMassiveEntityNotify.proto

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

type CreateMassiveEntityNotify_CmdId int32

const (
	CreateMassiveEntityNotify_NONE             CreateMassiveEntityNotify_CmdId = 0
	CreateMassiveEntityNotify_CMD_ID           CreateMassiveEntityNotify_CmdId = 345
	CreateMassiveEntityNotify_ENET_CHANNEL_ID  CreateMassiveEntityNotify_CmdId = 0
	CreateMassiveEntityNotify_ENET_IS_RELIABLE CreateMassiveEntityNotify_CmdId = 1
)

// Enum value maps for CreateMassiveEntityNotify_CmdId.
var (
	CreateMassiveEntityNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		345: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	CreateMassiveEntityNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           345,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x CreateMassiveEntityNotify_CmdId) Enum() *CreateMassiveEntityNotify_CmdId {
	p := new(CreateMassiveEntityNotify_CmdId)
	*p = x
	return p
}

func (x CreateMassiveEntityNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateMassiveEntityNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_CreateMassiveEntityNotify_proto_enumTypes[0].Descriptor()
}

func (CreateMassiveEntityNotify_CmdId) Type() protoreflect.EnumType {
	return &file_CreateMassiveEntityNotify_proto_enumTypes[0]
}

func (x CreateMassiveEntityNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateMassiveEntityNotify_CmdId.Descriptor instead.
func (CreateMassiveEntityNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_CreateMassiveEntityNotify_proto_rawDescGZIP(), []int{0, 0}
}

type CreateMassiveEntityNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MassiveEntityList []*ServerMassiveEntity `protobuf:"bytes,1,rep,name=massive_entity_list,json=massiveEntityList,proto3" json:"massive_entity_list,omitempty"`
}

func (x *CreateMassiveEntityNotify) Reset() {
	*x = CreateMassiveEntityNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CreateMassiveEntityNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMassiveEntityNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMassiveEntityNotify) ProtoMessage() {}

func (x *CreateMassiveEntityNotify) ProtoReflect() protoreflect.Message {
	mi := &file_CreateMassiveEntityNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMassiveEntityNotify.ProtoReflect.Descriptor instead.
func (*CreateMassiveEntityNotify) Descriptor() ([]byte, []int) {
	return file_CreateMassiveEntityNotify_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMassiveEntityNotify) GetMassiveEntityList() []*ServerMassiveEntity {
	if x != nil {
		return x.MassiveEntityList
	}
	return nil
}

var File_CreateMassiveEntityNotify_proto protoreflect.FileDescriptor

var file_CreateMassiveEntityNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a,
	0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x44, 0x0a, 0x13, 0x6d, 0x61,
	0x73, 0x73, 0x69, 0x76, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x11, 0x6d,
	0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xd9, 0x02,
	0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c,
	0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53,
	0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_CreateMassiveEntityNotify_proto_rawDescOnce sync.Once
	file_CreateMassiveEntityNotify_proto_rawDescData = file_CreateMassiveEntityNotify_proto_rawDesc
)

func file_CreateMassiveEntityNotify_proto_rawDescGZIP() []byte {
	file_CreateMassiveEntityNotify_proto_rawDescOnce.Do(func() {
		file_CreateMassiveEntityNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_CreateMassiveEntityNotify_proto_rawDescData)
	})
	return file_CreateMassiveEntityNotify_proto_rawDescData
}

var file_CreateMassiveEntityNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CreateMassiveEntityNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CreateMassiveEntityNotify_proto_goTypes = []interface{}{
	(CreateMassiveEntityNotify_CmdId)(0), // 0: CreateMassiveEntityNotify.CmdId
	(*CreateMassiveEntityNotify)(nil),    // 1: CreateMassiveEntityNotify
	(*ServerMassiveEntity)(nil),          // 2: ServerMassiveEntity
}
var file_CreateMassiveEntityNotify_proto_depIdxs = []int32{
	2, // 0: CreateMassiveEntityNotify.massive_entity_list:type_name -> ServerMassiveEntity
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CreateMassiveEntityNotify_proto_init() }
func file_CreateMassiveEntityNotify_proto_init() {
	if File_CreateMassiveEntityNotify_proto != nil {
		return
	}
	file_ServerMassiveEntity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CreateMassiveEntityNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMassiveEntityNotify); i {
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
			RawDescriptor: file_CreateMassiveEntityNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CreateMassiveEntityNotify_proto_goTypes,
		DependencyIndexes: file_CreateMassiveEntityNotify_proto_depIdxs,
		EnumInfos:         file_CreateMassiveEntityNotify_proto_enumTypes,
		MessageInfos:      file_CreateMassiveEntityNotify_proto_msgTypes,
	}.Build()
	File_CreateMassiveEntityNotify_proto = out.File
	file_CreateMassiveEntityNotify_proto_rawDesc = nil
	file_CreateMassiveEntityNotify_proto_goTypes = nil
	file_CreateMassiveEntityNotify_proto_depIdxs = nil
}
