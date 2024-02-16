// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: ExecuteGroupTriggerReq.proto

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

type ExecuteGroupTriggerReq_CmdId int32

const (
	ExecuteGroupTriggerReq_NONE             ExecuteGroupTriggerReq_CmdId = 0
	ExecuteGroupTriggerReq_CMD_ID           ExecuteGroupTriggerReq_CmdId = 253
	ExecuteGroupTriggerReq_ENET_CHANNEL_ID  ExecuteGroupTriggerReq_CmdId = 0
	ExecuteGroupTriggerReq_ENET_IS_RELIABLE ExecuteGroupTriggerReq_CmdId = 1
	ExecuteGroupTriggerReq_IS_ALLOW_CLIENT  ExecuteGroupTriggerReq_CmdId = 1
)

// Enum value maps for ExecuteGroupTriggerReq_CmdId.
var (
	ExecuteGroupTriggerReq_CmdId_name = map[int32]string{
		0:   "NONE",
		253: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	ExecuteGroupTriggerReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           253,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x ExecuteGroupTriggerReq_CmdId) Enum() *ExecuteGroupTriggerReq_CmdId {
	p := new(ExecuteGroupTriggerReq_CmdId)
	*p = x
	return p
}

func (x ExecuteGroupTriggerReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExecuteGroupTriggerReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_ExecuteGroupTriggerReq_proto_enumTypes[0].Descriptor()
}

func (ExecuteGroupTriggerReq_CmdId) Type() protoreflect.EnumType {
	return &file_ExecuteGroupTriggerReq_proto_enumTypes[0]
}

func (x ExecuteGroupTriggerReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExecuteGroupTriggerReq_CmdId.Descriptor instead.
func (ExecuteGroupTriggerReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_ExecuteGroupTriggerReq_proto_rawDescGZIP(), []int{0, 0}
}

type ExecuteGroupTriggerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceEntityId uint32 `protobuf:"varint,1,opt,name=source_entity_id,json=sourceEntityId,proto3" json:"source_entity_id,omitempty"`
	SourceName     string `protobuf:"bytes,2,opt,name=source_name,json=sourceName,proto3" json:"source_name,omitempty"`
	TargetEntityId uint32 `protobuf:"varint,3,opt,name=target_entity_id,json=targetEntityId,proto3" json:"target_entity_id,omitempty"`
	Param1         int32  `protobuf:"varint,4,opt,name=param1,proto3" json:"param1,omitempty"`
	Param2         int32  `protobuf:"varint,5,opt,name=param2,proto3" json:"param2,omitempty"`
	Param3         int32  `protobuf:"varint,6,opt,name=param3,proto3" json:"param3,omitempty"`
}

func (x *ExecuteGroupTriggerReq) Reset() {
	*x = ExecuteGroupTriggerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ExecuteGroupTriggerReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteGroupTriggerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteGroupTriggerReq) ProtoMessage() {}

func (x *ExecuteGroupTriggerReq) ProtoReflect() protoreflect.Message {
	mi := &file_ExecuteGroupTriggerReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteGroupTriggerReq.ProtoReflect.Descriptor instead.
func (*ExecuteGroupTriggerReq) Descriptor() ([]byte, []int) {
	return file_ExecuteGroupTriggerReq_proto_rawDescGZIP(), []int{0}
}

func (x *ExecuteGroupTriggerReq) GetSourceEntityId() uint32 {
	if x != nil {
		return x.SourceEntityId
	}
	return 0
}

func (x *ExecuteGroupTriggerReq) GetSourceName() string {
	if x != nil {
		return x.SourceName
	}
	return ""
}

func (x *ExecuteGroupTriggerReq) GetTargetEntityId() uint32 {
	if x != nil {
		return x.TargetEntityId
	}
	return 0
}

func (x *ExecuteGroupTriggerReq) GetParam1() int32 {
	if x != nil {
		return x.Param1
	}
	return 0
}

func (x *ExecuteGroupTriggerReq) GetParam2() int32 {
	if x != nil {
		return x.Param2
	}
	return 0
}

func (x *ExecuteGroupTriggerReq) GetParam3() int32 {
	if x != nil {
		return x.Param3
	}
	return 0
}

var File_ExecuteGroupTriggerReq_proto protoreflect.FileDescriptor

var file_ExecuteGroupTriggerReq_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x02, 0x0a, 0x16, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x28, 0x0a, 0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x31, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x31, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x33, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x33, 0x22, 0x62, 0x0a,
	0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xfd, 0x01, 0x12, 0x13, 0x0a,
	0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45,
	0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41,
	0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ExecuteGroupTriggerReq_proto_rawDescOnce sync.Once
	file_ExecuteGroupTriggerReq_proto_rawDescData = file_ExecuteGroupTriggerReq_proto_rawDesc
)

func file_ExecuteGroupTriggerReq_proto_rawDescGZIP() []byte {
	file_ExecuteGroupTriggerReq_proto_rawDescOnce.Do(func() {
		file_ExecuteGroupTriggerReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ExecuteGroupTriggerReq_proto_rawDescData)
	})
	return file_ExecuteGroupTriggerReq_proto_rawDescData
}

var file_ExecuteGroupTriggerReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ExecuteGroupTriggerReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ExecuteGroupTriggerReq_proto_goTypes = []interface{}{
	(ExecuteGroupTriggerReq_CmdId)(0), // 0: proto.ExecuteGroupTriggerReq.CmdId
	(*ExecuteGroupTriggerReq)(nil),    // 1: proto.ExecuteGroupTriggerReq
}
var file_ExecuteGroupTriggerReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ExecuteGroupTriggerReq_proto_init() }
func file_ExecuteGroupTriggerReq_proto_init() {
	if File_ExecuteGroupTriggerReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ExecuteGroupTriggerReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteGroupTriggerReq); i {
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
			RawDescriptor: file_ExecuteGroupTriggerReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ExecuteGroupTriggerReq_proto_goTypes,
		DependencyIndexes: file_ExecuteGroupTriggerReq_proto_depIdxs,
		EnumInfos:         file_ExecuteGroupTriggerReq_proto_enumTypes,
		MessageInfos:      file_ExecuteGroupTriggerReq_proto_msgTypes,
	}.Build()
	File_ExecuteGroupTriggerReq_proto = out.File
	file_ExecuteGroupTriggerReq_proto_rawDesc = nil
	file_ExecuteGroupTriggerReq_proto_goTypes = nil
	file_ExecuteGroupTriggerReq_proto_depIdxs = nil
}