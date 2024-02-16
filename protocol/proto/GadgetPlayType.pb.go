// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GadgetPlayType.proto

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

type GadgetPlayType int32

const (
	GadgetPlayType_GADGET_PLAY_NONE     GadgetPlayType = 0
	GadgetPlayType_GADGET_PLAY_CRUSIBLE GadgetPlayType = 1
)

// Enum value maps for GadgetPlayType.
var (
	GadgetPlayType_name = map[int32]string{
		0: "GADGET_PLAY_NONE",
		1: "GADGET_PLAY_CRUSIBLE",
	}
	GadgetPlayType_value = map[string]int32{
		"GADGET_PLAY_NONE":     0,
		"GADGET_PLAY_CRUSIBLE": 1,
	}
)

func (x GadgetPlayType) Enum() *GadgetPlayType {
	p := new(GadgetPlayType)
	*p = x
	return p
}

func (x GadgetPlayType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GadgetPlayType) Descriptor() protoreflect.EnumDescriptor {
	return file_GadgetPlayType_proto_enumTypes[0].Descriptor()
}

func (GadgetPlayType) Type() protoreflect.EnumType {
	return &file_GadgetPlayType_proto_enumTypes[0]
}

func (x GadgetPlayType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GadgetPlayType.Descriptor instead.
func (GadgetPlayType) EnumDescriptor() ([]byte, []int) {
	return file_GadgetPlayType_proto_rawDescGZIP(), []int{0}
}

var File_GadgetPlayType_proto protoreflect.FileDescriptor

var file_GadgetPlayType_proto_rawDesc = []byte{
	0x0a, 0x14, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x40, 0x0a,
	0x0e, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x10, 0x47, 0x41, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x47, 0x41, 0x44, 0x47, 0x45, 0x54, 0x5f,
	0x50, 0x4c, 0x41, 0x59, 0x5f, 0x43, 0x52, 0x55, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_GadgetPlayType_proto_rawDescOnce sync.Once
	file_GadgetPlayType_proto_rawDescData = file_GadgetPlayType_proto_rawDesc
)

func file_GadgetPlayType_proto_rawDescGZIP() []byte {
	file_GadgetPlayType_proto_rawDescOnce.Do(func() {
		file_GadgetPlayType_proto_rawDescData = protoimpl.X.CompressGZIP(file_GadgetPlayType_proto_rawDescData)
	})
	return file_GadgetPlayType_proto_rawDescData
}

var file_GadgetPlayType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GadgetPlayType_proto_goTypes = []interface{}{
	(GadgetPlayType)(0), // 0: proto.GadgetPlayType
}
var file_GadgetPlayType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GadgetPlayType_proto_init() }
func file_GadgetPlayType_proto_init() {
	if File_GadgetPlayType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GadgetPlayType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GadgetPlayType_proto_goTypes,
		DependencyIndexes: file_GadgetPlayType_proto_depIdxs,
		EnumInfos:         file_GadgetPlayType_proto_enumTypes,
	}.Build()
	File_GadgetPlayType_proto = out.File
	file_GadgetPlayType_proto_rawDesc = nil
	file_GadgetPlayType_proto_goTypes = nil
	file_GadgetPlayType_proto_depIdxs = nil
}
