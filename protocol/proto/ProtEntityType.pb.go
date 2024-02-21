// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: ProtEntityType.proto

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

type ProtEntityType int32

const (
	ProtEntityType_PROT_ENTITY_NONE           ProtEntityType = 0
	ProtEntityType_PROT_ENTITY_AVATAR         ProtEntityType = 1
	ProtEntityType_PROT_ENTITY_MONSTER        ProtEntityType = 2
	ProtEntityType_PROT_ENTITY_NPC            ProtEntityType = 3
	ProtEntityType_PROT_ENTITY_GADGET         ProtEntityType = 4
	ProtEntityType_PROT_ENTITY_REGION         ProtEntityType = 5
	ProtEntityType_PROT_ENTITY_WEAPON         ProtEntityType = 6
	ProtEntityType_PROT_ENTITY_WEATHER        ProtEntityType = 7
	ProtEntityType_PROT_ENTITY_SCENE          ProtEntityType = 8
	ProtEntityType_PROT_ENTITY_TEAM           ProtEntityType = 9
	ProtEntityType_PROT_ENTITY_MASSIVE_ENTITY ProtEntityType = 10
	ProtEntityType_PROT_ENTITY_MP_LEVEL       ProtEntityType = 11
	ProtEntityType_PROT_ENTITY_MAX            ProtEntityType = 12
)

// Enum value maps for ProtEntityType.
var (
	ProtEntityType_name = map[int32]string{
		0:  "PROT_ENTITY_NONE",
		1:  "PROT_ENTITY_AVATAR",
		2:  "PROT_ENTITY_MONSTER",
		3:  "PROT_ENTITY_NPC",
		4:  "PROT_ENTITY_GADGET",
		5:  "PROT_ENTITY_REGION",
		6:  "PROT_ENTITY_WEAPON",
		7:  "PROT_ENTITY_WEATHER",
		8:  "PROT_ENTITY_SCENE",
		9:  "PROT_ENTITY_TEAM",
		10: "PROT_ENTITY_MASSIVE_ENTITY",
		11: "PROT_ENTITY_MP_LEVEL",
		12: "PROT_ENTITY_MAX",
	}
	ProtEntityType_value = map[string]int32{
		"PROT_ENTITY_NONE":           0,
		"PROT_ENTITY_AVATAR":         1,
		"PROT_ENTITY_MONSTER":        2,
		"PROT_ENTITY_NPC":            3,
		"PROT_ENTITY_GADGET":         4,
		"PROT_ENTITY_REGION":         5,
		"PROT_ENTITY_WEAPON":         6,
		"PROT_ENTITY_WEATHER":        7,
		"PROT_ENTITY_SCENE":          8,
		"PROT_ENTITY_TEAM":           9,
		"PROT_ENTITY_MASSIVE_ENTITY": 10,
		"PROT_ENTITY_MP_LEVEL":       11,
		"PROT_ENTITY_MAX":            12,
	}
)

func (x ProtEntityType) Enum() *ProtEntityType {
	p := new(ProtEntityType)
	*p = x
	return p
}

func (x ProtEntityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtEntityType) Descriptor() protoreflect.EnumDescriptor {
	return file_ProtEntityType_proto_enumTypes[0].Descriptor()
}

func (ProtEntityType) Type() protoreflect.EnumType {
	return &file_ProtEntityType_proto_enumTypes[0]
}

func (x ProtEntityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProtEntityType.Descriptor instead.
func (ProtEntityType) EnumDescriptor() ([]byte, []int) {
	return file_ProtEntityType_proto_rawDescGZIP(), []int{0}
}

var File_ProtEntityType_proto protoreflect.FileDescriptor

var file_ProtEntityType_proto_rawDesc = []byte{
	0x0a, 0x14, 0x50, 0x72, 0x6f, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xc9, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x74, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x52, 0x4f,
	0x54, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x16, 0x0a, 0x12, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x41,
	0x56, 0x41, 0x54, 0x41, 0x52, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x4f, 0x54, 0x5f,
	0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x4d, 0x4f, 0x4e, 0x53, 0x54, 0x45, 0x52, 0x10, 0x02,
	0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f,
	0x4e, 0x50, 0x43, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x45, 0x4e,
	0x54, 0x49, 0x54, 0x59, 0x5f, 0x47, 0x41, 0x44, 0x47, 0x45, 0x54, 0x10, 0x04, 0x12, 0x16, 0x0a,
	0x12, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x52, 0x45, 0x47,
	0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x45, 0x4e,
	0x54, 0x49, 0x54, 0x59, 0x5f, 0x57, 0x45, 0x41, 0x50, 0x4f, 0x4e, 0x10, 0x06, 0x12, 0x17, 0x0a,
	0x13, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x57, 0x45, 0x41,
	0x54, 0x48, 0x45, 0x52, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x45,
	0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x43, 0x45, 0x4e, 0x45, 0x10, 0x08, 0x12, 0x14, 0x0a,
	0x10, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x45, 0x41,
	0x4d, 0x10, 0x09, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x54, 0x49,
	0x54, 0x59, 0x5f, 0x4d, 0x41, 0x53, 0x53, 0x49, 0x56, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54,
	0x59, 0x10, 0x0a, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x54, 0x49,
	0x54, 0x59, 0x5f, 0x4d, 0x50, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x10, 0x0b, 0x12, 0x13, 0x0a,
	0x0f, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x4d, 0x41, 0x58,
	0x10, 0x0c, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ProtEntityType_proto_rawDescOnce sync.Once
	file_ProtEntityType_proto_rawDescData = file_ProtEntityType_proto_rawDesc
)

func file_ProtEntityType_proto_rawDescGZIP() []byte {
	file_ProtEntityType_proto_rawDescOnce.Do(func() {
		file_ProtEntityType_proto_rawDescData = protoimpl.X.CompressGZIP(file_ProtEntityType_proto_rawDescData)
	})
	return file_ProtEntityType_proto_rawDescData
}

var file_ProtEntityType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ProtEntityType_proto_goTypes = []interface{}{
	(ProtEntityType)(0), // 0: ProtEntityType
}
var file_ProtEntityType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ProtEntityType_proto_init() }
func file_ProtEntityType_proto_init() {
	if File_ProtEntityType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ProtEntityType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ProtEntityType_proto_goTypes,
		DependencyIndexes: file_ProtEntityType_proto_depIdxs,
		EnumInfos:         file_ProtEntityType_proto_enumTypes,
	}.Build()
	File_ProtEntityType_proto = out.File
	file_ProtEntityType_proto_rawDesc = nil
	file_ProtEntityType_proto_goTypes = nil
	file_ProtEntityType_proto_depIdxs = nil
}
