// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GadgetStateNotify.proto

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

type GadgetStateNotify_CmdId int32

const (
	GadgetStateNotify_NONE             GadgetStateNotify_CmdId = 0
	GadgetStateNotify_CMD_ID           GadgetStateNotify_CmdId = 803
	GadgetStateNotify_ENET_CHANNEL_ID  GadgetStateNotify_CmdId = 0
	GadgetStateNotify_ENET_IS_RELIABLE GadgetStateNotify_CmdId = 1
)

// Enum value maps for GadgetStateNotify_CmdId.
var (
	GadgetStateNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		803: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	GadgetStateNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           803,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x GadgetStateNotify_CmdId) Enum() *GadgetStateNotify_CmdId {
	p := new(GadgetStateNotify_CmdId)
	*p = x
	return p
}

func (x GadgetStateNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GadgetStateNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GadgetStateNotify_proto_enumTypes[0].Descriptor()
}

func (GadgetStateNotify_CmdId) Type() protoreflect.EnumType {
	return &file_GadgetStateNotify_proto_enumTypes[0]
}

func (x GadgetStateNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GadgetStateNotify_CmdId.Descriptor instead.
func (GadgetStateNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GadgetStateNotify_proto_rawDescGZIP(), []int{0, 0}
}

type GadgetStateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GadgetEntityId   uint32 `protobuf:"varint,1,opt,name=gadget_entity_id,json=gadgetEntityId,proto3" json:"gadget_entity_id,omitempty"`
	GadgetState      uint32 `protobuf:"varint,2,opt,name=gadget_state,json=gadgetState,proto3" json:"gadget_state,omitempty"`
	IsEnableInteract bool   `protobuf:"varint,3,opt,name=is_enable_interact,json=isEnableInteract,proto3" json:"is_enable_interact,omitempty"`
}

func (x *GadgetStateNotify) Reset() {
	*x = GadgetStateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GadgetStateNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GadgetStateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GadgetStateNotify) ProtoMessage() {}

func (x *GadgetStateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_GadgetStateNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GadgetStateNotify.ProtoReflect.Descriptor instead.
func (*GadgetStateNotify) Descriptor() ([]byte, []int) {
	return file_GadgetStateNotify_proto_rawDescGZIP(), []int{0}
}

func (x *GadgetStateNotify) GetGadgetEntityId() uint32 {
	if x != nil {
		return x.GadgetEntityId
	}
	return 0
}

func (x *GadgetStateNotify) GetGadgetState() uint32 {
	if x != nil {
		return x.GadgetState
	}
	return 0
}

func (x *GadgetStateNotify) GetIsEnableInteract() bool {
	if x != nil {
		return x.IsEnableInteract
	}
	return false
}

var File_GadgetStateNotify_proto protoreflect.FileDescriptor

var file_GadgetStateNotify_proto_rawDesc = []byte{
	0x0a, 0x17, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdd, 0x01, 0x0a, 0x11, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74,
	0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0e, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x10, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xa3,
	0x06, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45,
	0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49,
	0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GadgetStateNotify_proto_rawDescOnce sync.Once
	file_GadgetStateNotify_proto_rawDescData = file_GadgetStateNotify_proto_rawDesc
)

func file_GadgetStateNotify_proto_rawDescGZIP() []byte {
	file_GadgetStateNotify_proto_rawDescOnce.Do(func() {
		file_GadgetStateNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_GadgetStateNotify_proto_rawDescData)
	})
	return file_GadgetStateNotify_proto_rawDescData
}

var file_GadgetStateNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GadgetStateNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GadgetStateNotify_proto_goTypes = []interface{}{
	(GadgetStateNotify_CmdId)(0), // 0: proto.GadgetStateNotify.CmdId
	(*GadgetStateNotify)(nil),    // 1: proto.GadgetStateNotify
}
var file_GadgetStateNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GadgetStateNotify_proto_init() }
func file_GadgetStateNotify_proto_init() {
	if File_GadgetStateNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GadgetStateNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GadgetStateNotify); i {
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
			RawDescriptor: file_GadgetStateNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GadgetStateNotify_proto_goTypes,
		DependencyIndexes: file_GadgetStateNotify_proto_depIdxs,
		EnumInfos:         file_GadgetStateNotify_proto_enumTypes,
		MessageInfos:      file_GadgetStateNotify_proto_msgTypes,
	}.Build()
	File_GadgetStateNotify_proto = out.File
	file_GadgetStateNotify_proto_rawDesc = nil
	file_GadgetStateNotify_proto_goTypes = nil
	file_GadgetStateNotify_proto_depIdxs = nil
}
