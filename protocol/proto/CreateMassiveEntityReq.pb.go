// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: CreateMassiveEntityReq.proto

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

type CreateMassiveEntityReq_CmdId int32

const (
	CreateMassiveEntityReq_NONE             CreateMassiveEntityReq_CmdId = 0
	CreateMassiveEntityReq_CMD_ID           CreateMassiveEntityReq_CmdId = 343
	CreateMassiveEntityReq_ENET_CHANNEL_ID  CreateMassiveEntityReq_CmdId = 0
	CreateMassiveEntityReq_ENET_IS_RELIABLE CreateMassiveEntityReq_CmdId = 1
	CreateMassiveEntityReq_IS_ALLOW_CLIENT  CreateMassiveEntityReq_CmdId = 1
)

// Enum value maps for CreateMassiveEntityReq_CmdId.
var (
	CreateMassiveEntityReq_CmdId_name = map[int32]string{
		0:   "NONE",
		343: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	CreateMassiveEntityReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           343,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x CreateMassiveEntityReq_CmdId) Enum() *CreateMassiveEntityReq_CmdId {
	p := new(CreateMassiveEntityReq_CmdId)
	*p = x
	return p
}

func (x CreateMassiveEntityReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateMassiveEntityReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_CreateMassiveEntityReq_proto_enumTypes[0].Descriptor()
}

func (CreateMassiveEntityReq_CmdId) Type() protoreflect.EnumType {
	return &file_CreateMassiveEntityReq_proto_enumTypes[0]
}

func (x CreateMassiveEntityReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateMassiveEntityReq_CmdId.Descriptor instead.
func (CreateMassiveEntityReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_CreateMassiveEntityReq_proto_rawDescGZIP(), []int{0, 0}
}

type CreateMassiveEntityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MassiveEntityList []*ClientMassiveEntity `protobuf:"bytes,1,rep,name=massive_entity_list,json=massiveEntityList,proto3" json:"massive_entity_list,omitempty"`
}

func (x *CreateMassiveEntityReq) Reset() {
	*x = CreateMassiveEntityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CreateMassiveEntityReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMassiveEntityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMassiveEntityReq) ProtoMessage() {}

func (x *CreateMassiveEntityReq) ProtoReflect() protoreflect.Message {
	mi := &file_CreateMassiveEntityReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMassiveEntityReq.ProtoReflect.Descriptor instead.
func (*CreateMassiveEntityReq) Descriptor() ([]byte, []int) {
	return file_CreateMassiveEntityReq_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMassiveEntityReq) GetMassiveEntityList() []*ClientMassiveEntity {
	if x != nil {
		return x.MassiveEntityList
	}
	return nil
}

var File_CreateMassiveEntityReq_proto protoreflect.FileDescriptor

var file_CreateMassiveEntityReq_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x12, 0x44, 0x0a, 0x13, 0x6d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x5f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x73, 0x73, 0x69, 0x76,
	0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x11, 0x6d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d,
	0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xd7, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41,
	0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f,
	0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_CreateMassiveEntityReq_proto_rawDescOnce sync.Once
	file_CreateMassiveEntityReq_proto_rawDescData = file_CreateMassiveEntityReq_proto_rawDesc
)

func file_CreateMassiveEntityReq_proto_rawDescGZIP() []byte {
	file_CreateMassiveEntityReq_proto_rawDescOnce.Do(func() {
		file_CreateMassiveEntityReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_CreateMassiveEntityReq_proto_rawDescData)
	})
	return file_CreateMassiveEntityReq_proto_rawDescData
}

var file_CreateMassiveEntityReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CreateMassiveEntityReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CreateMassiveEntityReq_proto_goTypes = []interface{}{
	(CreateMassiveEntityReq_CmdId)(0), // 0: CreateMassiveEntityReq.CmdId
	(*CreateMassiveEntityReq)(nil),    // 1: CreateMassiveEntityReq
	(*ClientMassiveEntity)(nil),       // 2: ClientMassiveEntity
}
var file_CreateMassiveEntityReq_proto_depIdxs = []int32{
	2, // 0: CreateMassiveEntityReq.massive_entity_list:type_name -> ClientMassiveEntity
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CreateMassiveEntityReq_proto_init() }
func file_CreateMassiveEntityReq_proto_init() {
	if File_CreateMassiveEntityReq_proto != nil {
		return
	}
	file_ClientMassiveEntity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CreateMassiveEntityReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMassiveEntityReq); i {
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
			RawDescriptor: file_CreateMassiveEntityReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CreateMassiveEntityReq_proto_goTypes,
		DependencyIndexes: file_CreateMassiveEntityReq_proto_depIdxs,
		EnumInfos:         file_CreateMassiveEntityReq_proto_enumTypes,
		MessageInfos:      file_CreateMassiveEntityReq_proto_msgTypes,
	}.Build()
	File_CreateMassiveEntityReq_proto = out.File
	file_CreateMassiveEntityReq_proto_rawDesc = nil
	file_CreateMassiveEntityReq_proto_goTypes = nil
	file_CreateMassiveEntityReq_proto_depIdxs = nil
}
