// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: LanguageType.proto

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

type LanguageType int32

const (
	LanguageType_LANGUAGE_NONE LanguageType = 0
	LanguageType_LANGUAGE_EN   LanguageType = 1
	LanguageType_LANGUAGE_SC   LanguageType = 2
	LanguageType_LANGUAGE_TC   LanguageType = 3
	LanguageType_LANGUAGE_FR   LanguageType = 4
	LanguageType_LANGUAGE_DE   LanguageType = 5
	LanguageType_LANGUAGE_ES   LanguageType = 6
	LanguageType_LANGUAGE_PT   LanguageType = 7
	LanguageType_LANGUAGE_RU   LanguageType = 8
	LanguageType_LANGUAGE_JP   LanguageType = 9
	LanguageType_LANGUAGE_KR   LanguageType = 10
	LanguageType_LANGUAGE_TH   LanguageType = 11
	LanguageType_LANGUAGE_VN   LanguageType = 12
	LanguageType_LANGUAGE_ID   LanguageType = 13
)

// Enum value maps for LanguageType.
var (
	LanguageType_name = map[int32]string{
		0:  "LANGUAGE_NONE",
		1:  "LANGUAGE_EN",
		2:  "LANGUAGE_SC",
		3:  "LANGUAGE_TC",
		4:  "LANGUAGE_FR",
		5:  "LANGUAGE_DE",
		6:  "LANGUAGE_ES",
		7:  "LANGUAGE_PT",
		8:  "LANGUAGE_RU",
		9:  "LANGUAGE_JP",
		10: "LANGUAGE_KR",
		11: "LANGUAGE_TH",
		12: "LANGUAGE_VN",
		13: "LANGUAGE_ID",
	}
	LanguageType_value = map[string]int32{
		"LANGUAGE_NONE": 0,
		"LANGUAGE_EN":   1,
		"LANGUAGE_SC":   2,
		"LANGUAGE_TC":   3,
		"LANGUAGE_FR":   4,
		"LANGUAGE_DE":   5,
		"LANGUAGE_ES":   6,
		"LANGUAGE_PT":   7,
		"LANGUAGE_RU":   8,
		"LANGUAGE_JP":   9,
		"LANGUAGE_KR":   10,
		"LANGUAGE_TH":   11,
		"LANGUAGE_VN":   12,
		"LANGUAGE_ID":   13,
	}
)

func (x LanguageType) Enum() *LanguageType {
	p := new(LanguageType)
	*p = x
	return p
}

func (x LanguageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LanguageType) Descriptor() protoreflect.EnumDescriptor {
	return file_LanguageType_proto_enumTypes[0].Descriptor()
}

func (LanguageType) Type() protoreflect.EnumType {
	return &file_LanguageType_proto_enumTypes[0]
}

func (x LanguageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LanguageType.Descriptor instead.
func (LanguageType) EnumDescriptor() ([]byte, []int) {
	return file_LanguageType_proto_rawDescGZIP(), []int{0}
}

var File_LanguageType_proto protoreflect.FileDescriptor

var file_LanguageType_proto_rawDesc = []byte{
	0x0a, 0x12, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xfe, 0x01, 0x0a, 0x0c,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d,
	0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x45, 0x4e, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x43, 0x10,
	0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x43,
	0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x46,
	0x52, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f,
	0x44, 0x45, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45,
	0x5f, 0x45, 0x53, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47,
	0x45, 0x5f, 0x50, 0x54, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41,
	0x47, 0x45, 0x5f, 0x52, 0x55, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x55,
	0x41, 0x47, 0x45, 0x5f, 0x4a, 0x50, 0x10, 0x09, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47,
	0x55, 0x41, 0x47, 0x45, 0x5f, 0x4b, 0x52, 0x10, 0x0a, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e,
	0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x48, 0x10, 0x0b, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41,
	0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x56, 0x4e, 0x10, 0x0c, 0x12, 0x0f, 0x0a, 0x0b, 0x4c,
	0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x0d, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LanguageType_proto_rawDescOnce sync.Once
	file_LanguageType_proto_rawDescData = file_LanguageType_proto_rawDesc
)

func file_LanguageType_proto_rawDescGZIP() []byte {
	file_LanguageType_proto_rawDescOnce.Do(func() {
		file_LanguageType_proto_rawDescData = protoimpl.X.CompressGZIP(file_LanguageType_proto_rawDescData)
	})
	return file_LanguageType_proto_rawDescData
}

var file_LanguageType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_LanguageType_proto_goTypes = []interface{}{
	(LanguageType)(0), // 0: proto.LanguageType
}
var file_LanguageType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_LanguageType_proto_init() }
func file_LanguageType_proto_init() {
	if File_LanguageType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_LanguageType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LanguageType_proto_goTypes,
		DependencyIndexes: file_LanguageType_proto_depIdxs,
		EnumInfos:         file_LanguageType_proto_enumTypes,
	}.Build()
	File_LanguageType_proto = out.File
	file_LanguageType_proto_rawDesc = nil
	file_LanguageType_proto_goTypes = nil
	file_LanguageType_proto_depIdxs = nil
}