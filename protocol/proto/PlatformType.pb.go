// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: PlatformType.proto

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

type PlatformType int32

const (
	PlatformType_EDITOR        PlatformType = 0
	PlatformType_IOS           PlatformType = 1
	PlatformType_ANDROID       PlatformType = 2
	PlatformType_PC            PlatformType = 3
	PlatformType_PS4           PlatformType = 4
	PlatformType_SERVER        PlatformType = 5
	PlatformType_CLOUD_ANDROID PlatformType = 6
	PlatformType_CLOUD_IOS     PlatformType = 7
)

// Enum value maps for PlatformType.
var (
	PlatformType_name = map[int32]string{
		0: "EDITOR",
		1: "IOS",
		2: "ANDROID",
		3: "PC",
		4: "PS4",
		5: "SERVER",
		6: "CLOUD_ANDROID",
		7: "CLOUD_IOS",
	}
	PlatformType_value = map[string]int32{
		"EDITOR":        0,
		"IOS":           1,
		"ANDROID":       2,
		"PC":            3,
		"PS4":           4,
		"SERVER":        5,
		"CLOUD_ANDROID": 6,
		"CLOUD_IOS":     7,
	}
)

func (x PlatformType) Enum() *PlatformType {
	p := new(PlatformType)
	*p = x
	return p
}

func (x PlatformType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlatformType) Descriptor() protoreflect.EnumDescriptor {
	return file_PlatformType_proto_enumTypes[0].Descriptor()
}

func (PlatformType) Type() protoreflect.EnumType {
	return &file_PlatformType_proto_enumTypes[0]
}

func (x PlatformType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlatformType.Descriptor instead.
func (PlatformType) EnumDescriptor() ([]byte, []int) {
	return file_PlatformType_proto_rawDescGZIP(), []int{0}
}

var File_PlatformType_proto protoreflect.FileDescriptor

var file_PlatformType_proto_rawDesc = []byte{
	0x0a, 0x12, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x6f, 0x0a, 0x0c, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x45,
	0x44, 0x49, 0x54, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4f, 0x53, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x02, 0x12, 0x06, 0x0a,
	0x02, 0x50, 0x43, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x53, 0x34, 0x10, 0x04, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4c,
	0x4f, 0x55, 0x44, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x06, 0x12, 0x0d, 0x0a,
	0x09, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x49, 0x4f, 0x53, 0x10, 0x07, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlatformType_proto_rawDescOnce sync.Once
	file_PlatformType_proto_rawDescData = file_PlatformType_proto_rawDesc
)

func file_PlatformType_proto_rawDescGZIP() []byte {
	file_PlatformType_proto_rawDescOnce.Do(func() {
		file_PlatformType_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlatformType_proto_rawDescData)
	})
	return file_PlatformType_proto_rawDescData
}

var file_PlatformType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlatformType_proto_goTypes = []interface{}{
	(PlatformType)(0), // 0: proto.PlatformType
}
var file_PlatformType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PlatformType_proto_init() }
func file_PlatformType_proto_init() {
	if File_PlatformType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_PlatformType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlatformType_proto_goTypes,
		DependencyIndexes: file_PlatformType_proto_depIdxs,
		EnumInfos:         file_PlatformType_proto_enumTypes,
	}.Build()
	File_PlatformType_proto = out.File
	file_PlatformType_proto_rawDesc = nil
	file_PlatformType_proto_goTypes = nil
	file_PlatformType_proto_depIdxs = nil
}
