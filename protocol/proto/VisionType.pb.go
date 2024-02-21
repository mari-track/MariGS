// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: VisionType.proto

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

type VisionType int32

const (
	VisionType_VISION_NONE            VisionType = 0
	VisionType_VISION_MEET            VisionType = 1
	VisionType_VISION_REBORN          VisionType = 2
	VisionType_VISION_REPLACE         VisionType = 3
	VisionType_VISION_WAYPOINT_REBORN VisionType = 4
	VisionType_VISION_MISS            VisionType = 5
	VisionType_VISION_DIE             VisionType = 6
	VisionType_VISION_GATHER_ESCAPE   VisionType = 7
	VisionType_VISION_REFRESH         VisionType = 8
	VisionType_VISION_TRANSPORT       VisionType = 9
	VisionType_VISION_REPLACE_DIE     VisionType = 10
)

// Enum value maps for VisionType.
var (
	VisionType_name = map[int32]string{
		0:  "VISION_NONE",
		1:  "VISION_MEET",
		2:  "VISION_REBORN",
		3:  "VISION_REPLACE",
		4:  "VISION_WAYPOINT_REBORN",
		5:  "VISION_MISS",
		6:  "VISION_DIE",
		7:  "VISION_GATHER_ESCAPE",
		8:  "VISION_REFRESH",
		9:  "VISION_TRANSPORT",
		10: "VISION_REPLACE_DIE",
	}
	VisionType_value = map[string]int32{
		"VISION_NONE":            0,
		"VISION_MEET":            1,
		"VISION_REBORN":          2,
		"VISION_REPLACE":         3,
		"VISION_WAYPOINT_REBORN": 4,
		"VISION_MISS":            5,
		"VISION_DIE":             6,
		"VISION_GATHER_ESCAPE":   7,
		"VISION_REFRESH":         8,
		"VISION_TRANSPORT":       9,
		"VISION_REPLACE_DIE":     10,
	}
)

func (x VisionType) Enum() *VisionType {
	p := new(VisionType)
	*p = x
	return p
}

func (x VisionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VisionType) Descriptor() protoreflect.EnumDescriptor {
	return file_VisionType_proto_enumTypes[0].Descriptor()
}

func (VisionType) Type() protoreflect.EnumType {
	return &file_VisionType_proto_enumTypes[0]
}

func (x VisionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VisionType.Descriptor instead.
func (VisionType) EnumDescriptor() ([]byte, []int) {
	return file_VisionType_proto_rawDescGZIP(), []int{0}
}

var File_VisionType_proto protoreflect.FileDescriptor

var file_VisionType_proto_rawDesc = []byte{
	0x0a, 0x10, 0x56, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0xee, 0x01, 0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x45,
	0x54, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45,
	0x42, 0x4f, 0x52, 0x4e, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x56, 0x49,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x57, 0x41, 0x59, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x52, 0x45,
	0x42, 0x4f, 0x52, 0x4e, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x4d, 0x49, 0x53, 0x53, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x49, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x44, 0x49, 0x45, 0x10, 0x06, 0x12, 0x18, 0x0a, 0x14, 0x56, 0x49, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x47, 0x41, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x45, 0x53, 0x43, 0x41, 0x50, 0x45, 0x10,
	0x07, 0x12, 0x12, 0x0a, 0x0e, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x46, 0x52,
	0x45, 0x53, 0x48, 0x10, 0x08, 0x12, 0x14, 0x0a, 0x10, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x09, 0x12, 0x16, 0x0a, 0x12, 0x56,
	0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x5f, 0x44, 0x49,
	0x45, 0x10, 0x0a, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_VisionType_proto_rawDescOnce sync.Once
	file_VisionType_proto_rawDescData = file_VisionType_proto_rawDesc
)

func file_VisionType_proto_rawDescGZIP() []byte {
	file_VisionType_proto_rawDescOnce.Do(func() {
		file_VisionType_proto_rawDescData = protoimpl.X.CompressGZIP(file_VisionType_proto_rawDescData)
	})
	return file_VisionType_proto_rawDescData
}

var file_VisionType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_VisionType_proto_goTypes = []interface{}{
	(VisionType)(0), // 0: VisionType
}
var file_VisionType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_VisionType_proto_init() }
func file_VisionType_proto_init() {
	if File_VisionType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_VisionType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_VisionType_proto_goTypes,
		DependencyIndexes: file_VisionType_proto_depIdxs,
		EnumInfos:         file_VisionType_proto_enumTypes,
	}.Build()
	File_VisionType_proto = out.File
	file_VisionType_proto_rawDesc = nil
	file_VisionType_proto_goTypes = nil
	file_VisionType_proto_depIdxs = nil
}
