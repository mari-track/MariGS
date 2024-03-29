// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: DropSubfieldType.proto

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

type DropSubfieldType int32

const (
	DropSubfieldType_DROP_SUBFIELD_NONE DropSubfieldType = 0
	DropSubfieldType_DROP_SUBFIELD_ONE  DropSubfieldType = 1
)

// Enum value maps for DropSubfieldType.
var (
	DropSubfieldType_name = map[int32]string{
		0: "DROP_SUBFIELD_NONE",
		1: "DROP_SUBFIELD_ONE",
	}
	DropSubfieldType_value = map[string]int32{
		"DROP_SUBFIELD_NONE": 0,
		"DROP_SUBFIELD_ONE":  1,
	}
)

func (x DropSubfieldType) Enum() *DropSubfieldType {
	p := new(DropSubfieldType)
	*p = x
	return p
}

func (x DropSubfieldType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DropSubfieldType) Descriptor() protoreflect.EnumDescriptor {
	return file_DropSubfieldType_proto_enumTypes[0].Descriptor()
}

func (DropSubfieldType) Type() protoreflect.EnumType {
	return &file_DropSubfieldType_proto_enumTypes[0]
}

func (x DropSubfieldType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DropSubfieldType.Descriptor instead.
func (DropSubfieldType) EnumDescriptor() ([]byte, []int) {
	return file_DropSubfieldType_proto_rawDescGZIP(), []int{0}
}

var File_DropSubfieldType_proto protoreflect.FileDescriptor

var file_DropSubfieldType_proto_rawDesc = []byte{
	0x0a, 0x16, 0x44, 0x72, 0x6f, 0x70, 0x53, 0x75, 0x62, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x41, 0x0a, 0x10, 0x44, 0x72, 0x6f, 0x70,
	0x53, 0x75, 0x62, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12,
	0x44, 0x52, 0x4f, 0x50, 0x5f, 0x53, 0x55, 0x42, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x52, 0x4f, 0x50, 0x5f, 0x53, 0x55, 0x42,
	0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DropSubfieldType_proto_rawDescOnce sync.Once
	file_DropSubfieldType_proto_rawDescData = file_DropSubfieldType_proto_rawDesc
)

func file_DropSubfieldType_proto_rawDescGZIP() []byte {
	file_DropSubfieldType_proto_rawDescOnce.Do(func() {
		file_DropSubfieldType_proto_rawDescData = protoimpl.X.CompressGZIP(file_DropSubfieldType_proto_rawDescData)
	})
	return file_DropSubfieldType_proto_rawDescData
}

var file_DropSubfieldType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DropSubfieldType_proto_goTypes = []interface{}{
	(DropSubfieldType)(0), // 0: DropSubfieldType
}
var file_DropSubfieldType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DropSubfieldType_proto_init() }
func file_DropSubfieldType_proto_init() {
	if File_DropSubfieldType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_DropSubfieldType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DropSubfieldType_proto_goTypes,
		DependencyIndexes: file_DropSubfieldType_proto_depIdxs,
		EnumInfos:         file_DropSubfieldType_proto_enumTypes,
	}.Build()
	File_DropSubfieldType_proto = out.File
	file_DropSubfieldType_proto_rawDesc = nil
	file_DropSubfieldType_proto_goTypes = nil
	file_DropSubfieldType_proto_depIdxs = nil
}
