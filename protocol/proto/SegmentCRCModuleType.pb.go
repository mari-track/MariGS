// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: SegmentCRCModuleType.proto

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

type SegmentCRCModuleType int32

const (
	SegmentCRCModuleType_MODULE_NONE       SegmentCRCModuleType = 0
	SegmentCRCModuleType_PC_UNITYPLAYER    SegmentCRCModuleType = 1
	SegmentCRCModuleType_PC_USERASSEMBLY   SegmentCRCModuleType = 2
	SegmentCRCModuleType_PC_XLUA           SegmentCRCModuleType = 3
	SegmentCRCModuleType_ANDROID_LIBIL2CPP SegmentCRCModuleType = 11
	SegmentCRCModuleType_ANDROID_LIBUNITY  SegmentCRCModuleType = 12
	SegmentCRCModuleType_ANDROID_LIBXLUA   SegmentCRCModuleType = 13
	SegmentCRCModuleType_IOS_HK4E          SegmentCRCModuleType = 21
)

// Enum value maps for SegmentCRCModuleType.
var (
	SegmentCRCModuleType_name = map[int32]string{
		0:  "MODULE_NONE",
		1:  "PC_UNITYPLAYER",
		2:  "PC_USERASSEMBLY",
		3:  "PC_XLUA",
		11: "ANDROID_LIBIL2CPP",
		12: "ANDROID_LIBUNITY",
		13: "ANDROID_LIBXLUA",
		21: "IOS_HK4E",
	}
	SegmentCRCModuleType_value = map[string]int32{
		"MODULE_NONE":       0,
		"PC_UNITYPLAYER":    1,
		"PC_USERASSEMBLY":   2,
		"PC_XLUA":           3,
		"ANDROID_LIBIL2CPP": 11,
		"ANDROID_LIBUNITY":  12,
		"ANDROID_LIBXLUA":   13,
		"IOS_HK4E":          21,
	}
)

func (x SegmentCRCModuleType) Enum() *SegmentCRCModuleType {
	p := new(SegmentCRCModuleType)
	*p = x
	return p
}

func (x SegmentCRCModuleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SegmentCRCModuleType) Descriptor() protoreflect.EnumDescriptor {
	return file_SegmentCRCModuleType_proto_enumTypes[0].Descriptor()
}

func (SegmentCRCModuleType) Type() protoreflect.EnumType {
	return &file_SegmentCRCModuleType_proto_enumTypes[0]
}

func (x SegmentCRCModuleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SegmentCRCModuleType.Descriptor instead.
func (SegmentCRCModuleType) EnumDescriptor() ([]byte, []int) {
	return file_SegmentCRCModuleType_proto_rawDescGZIP(), []int{0}
}

var File_SegmentCRCModuleType_proto protoreflect.FileDescriptor

var file_SegmentCRCModuleType_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x52, 0x43, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xad, 0x01, 0x0a,
	0x14, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x52, 0x43, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x43, 0x5f, 0x55, 0x4e, 0x49,
	0x54, 0x59, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x43,
	0x5f, 0x55, 0x53, 0x45, 0x52, 0x41, 0x53, 0x53, 0x45, 0x4d, 0x42, 0x4c, 0x59, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x50, 0x43, 0x5f, 0x58, 0x4c, 0x55, 0x41, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11,
	0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x4c, 0x49, 0x42, 0x49, 0x4c, 0x32, 0x43, 0x50,
	0x50, 0x10, 0x0b, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x4c,
	0x49, 0x42, 0x55, 0x4e, 0x49, 0x54, 0x59, 0x10, 0x0c, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x4e, 0x44,
	0x52, 0x4f, 0x49, 0x44, 0x5f, 0x4c, 0x49, 0x42, 0x58, 0x4c, 0x55, 0x41, 0x10, 0x0d, 0x12, 0x0c,
	0x0a, 0x08, 0x49, 0x4f, 0x53, 0x5f, 0x48, 0x4b, 0x34, 0x45, 0x10, 0x15, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SegmentCRCModuleType_proto_rawDescOnce sync.Once
	file_SegmentCRCModuleType_proto_rawDescData = file_SegmentCRCModuleType_proto_rawDesc
)

func file_SegmentCRCModuleType_proto_rawDescGZIP() []byte {
	file_SegmentCRCModuleType_proto_rawDescOnce.Do(func() {
		file_SegmentCRCModuleType_proto_rawDescData = protoimpl.X.CompressGZIP(file_SegmentCRCModuleType_proto_rawDescData)
	})
	return file_SegmentCRCModuleType_proto_rawDescData
}

var file_SegmentCRCModuleType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SegmentCRCModuleType_proto_goTypes = []interface{}{
	(SegmentCRCModuleType)(0), // 0: SegmentCRCModuleType
}
var file_SegmentCRCModuleType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SegmentCRCModuleType_proto_init() }
func file_SegmentCRCModuleType_proto_init() {
	if File_SegmentCRCModuleType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_SegmentCRCModuleType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SegmentCRCModuleType_proto_goTypes,
		DependencyIndexes: file_SegmentCRCModuleType_proto_depIdxs,
		EnumInfos:         file_SegmentCRCModuleType_proto_enumTypes,
	}.Build()
	File_SegmentCRCModuleType_proto = out.File
	file_SegmentCRCModuleType_proto_rawDesc = nil
	file_SegmentCRCModuleType_proto_goTypes = nil
	file_SegmentCRCModuleType_proto_depIdxs = nil
}
