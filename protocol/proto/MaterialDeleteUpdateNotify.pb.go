// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: MaterialDeleteUpdateNotify.proto

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

type MaterialDeleteUpdateNotify_CmdId int32

const (
	MaterialDeleteUpdateNotify_NONE             MaterialDeleteUpdateNotify_CmdId = 0
	MaterialDeleteUpdateNotify_CMD_ID           MaterialDeleteUpdateNotify_CmdId = 654
	MaterialDeleteUpdateNotify_ENET_CHANNEL_ID  MaterialDeleteUpdateNotify_CmdId = 0
	MaterialDeleteUpdateNotify_ENET_IS_RELIABLE MaterialDeleteUpdateNotify_CmdId = 1
)

// Enum value maps for MaterialDeleteUpdateNotify_CmdId.
var (
	MaterialDeleteUpdateNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		654: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	MaterialDeleteUpdateNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           654,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x MaterialDeleteUpdateNotify_CmdId) Enum() *MaterialDeleteUpdateNotify_CmdId {
	p := new(MaterialDeleteUpdateNotify_CmdId)
	*p = x
	return p
}

func (x MaterialDeleteUpdateNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MaterialDeleteUpdateNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_MaterialDeleteUpdateNotify_proto_enumTypes[0].Descriptor()
}

func (MaterialDeleteUpdateNotify_CmdId) Type() protoreflect.EnumType {
	return &file_MaterialDeleteUpdateNotify_proto_enumTypes[0]
}

func (x MaterialDeleteUpdateNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MaterialDeleteUpdateNotify_CmdId.Descriptor instead.
func (MaterialDeleteUpdateNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_MaterialDeleteUpdateNotify_proto_rawDescGZIP(), []int{0, 0}
}

type MaterialDeleteUpdateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MaterialDeleteUpdateNotify) Reset() {
	*x = MaterialDeleteUpdateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MaterialDeleteUpdateNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialDeleteUpdateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialDeleteUpdateNotify) ProtoMessage() {}

func (x *MaterialDeleteUpdateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_MaterialDeleteUpdateNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialDeleteUpdateNotify.ProtoReflect.Descriptor instead.
func (*MaterialDeleteUpdateNotify) Descriptor() ([]byte, []int) {
	return file_MaterialDeleteUpdateNotify_proto_rawDescGZIP(), []int{0}
}

var File_MaterialDeleteUpdateNotify_proto protoreflect.FileDescriptor

var file_MaterialDeleteUpdateNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x1a, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d,
	0x44, 0x5f, 0x49, 0x44, 0x10, 0x8e, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10,
	0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45,
	0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MaterialDeleteUpdateNotify_proto_rawDescOnce sync.Once
	file_MaterialDeleteUpdateNotify_proto_rawDescData = file_MaterialDeleteUpdateNotify_proto_rawDesc
)

func file_MaterialDeleteUpdateNotify_proto_rawDescGZIP() []byte {
	file_MaterialDeleteUpdateNotify_proto_rawDescOnce.Do(func() {
		file_MaterialDeleteUpdateNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_MaterialDeleteUpdateNotify_proto_rawDescData)
	})
	return file_MaterialDeleteUpdateNotify_proto_rawDescData
}

var file_MaterialDeleteUpdateNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MaterialDeleteUpdateNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MaterialDeleteUpdateNotify_proto_goTypes = []interface{}{
	(MaterialDeleteUpdateNotify_CmdId)(0), // 0: proto.MaterialDeleteUpdateNotify.CmdId
	(*MaterialDeleteUpdateNotify)(nil),    // 1: proto.MaterialDeleteUpdateNotify
}
var file_MaterialDeleteUpdateNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MaterialDeleteUpdateNotify_proto_init() }
func file_MaterialDeleteUpdateNotify_proto_init() {
	if File_MaterialDeleteUpdateNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MaterialDeleteUpdateNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialDeleteUpdateNotify); i {
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
			RawDescriptor: file_MaterialDeleteUpdateNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MaterialDeleteUpdateNotify_proto_goTypes,
		DependencyIndexes: file_MaterialDeleteUpdateNotify_proto_depIdxs,
		EnumInfos:         file_MaterialDeleteUpdateNotify_proto_enumTypes,
		MessageInfos:      file_MaterialDeleteUpdateNotify_proto_msgTypes,
	}.Build()
	File_MaterialDeleteUpdateNotify_proto = out.File
	file_MaterialDeleteUpdateNotify_proto_rawDesc = nil
	file_MaterialDeleteUpdateNotify_proto_goTypes = nil
	file_MaterialDeleteUpdateNotify_proto_depIdxs = nil
}
