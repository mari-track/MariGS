// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: MapMarkPointType.proto

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

type MapMarkPointType int32

const (
	MapMarkPointType_NPC        MapMarkPointType = 0
	MapMarkPointType_QUEST      MapMarkPointType = 1
	MapMarkPointType_SPECIAL    MapMarkPointType = 2
	MapMarkPointType_MINE       MapMarkPointType = 3
	MapMarkPointType_COLLECTION MapMarkPointType = 4
	MapMarkPointType_MONSTER    MapMarkPointType = 5
)

// Enum value maps for MapMarkPointType.
var (
	MapMarkPointType_name = map[int32]string{
		0: "NPC",
		1: "QUEST",
		2: "SPECIAL",
		3: "MINE",
		4: "COLLECTION",
		5: "MONSTER",
	}
	MapMarkPointType_value = map[string]int32{
		"NPC":        0,
		"QUEST":      1,
		"SPECIAL":    2,
		"MINE":       3,
		"COLLECTION": 4,
		"MONSTER":    5,
	}
)

func (x MapMarkPointType) Enum() *MapMarkPointType {
	p := new(MapMarkPointType)
	*p = x
	return p
}

func (x MapMarkPointType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MapMarkPointType) Descriptor() protoreflect.EnumDescriptor {
	return file_MapMarkPointType_proto_enumTypes[0].Descriptor()
}

func (MapMarkPointType) Type() protoreflect.EnumType {
	return &file_MapMarkPointType_proto_enumTypes[0]
}

func (x MapMarkPointType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MapMarkPointType.Descriptor instead.
func (MapMarkPointType) EnumDescriptor() ([]byte, []int) {
	return file_MapMarkPointType_proto_rawDescGZIP(), []int{0}
}

var File_MapMarkPointType_proto protoreflect.FileDescriptor

var file_MapMarkPointType_proto_rawDesc = []byte{
	0x0a, 0x16, 0x4d, 0x61, 0x70, 0x4d, 0x61, 0x72, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x5a, 0x0a, 0x10, 0x4d, 0x61, 0x70, 0x4d,
	0x61, 0x72, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03,
	0x4e, 0x50, 0x43, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x50, 0x45, 0x43, 0x49, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x08, 0x0a,
	0x04, 0x4d, 0x49, 0x4e, 0x45, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4c, 0x4c, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x4f, 0x4e, 0x53, 0x54,
	0x45, 0x52, 0x10, 0x05, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MapMarkPointType_proto_rawDescOnce sync.Once
	file_MapMarkPointType_proto_rawDescData = file_MapMarkPointType_proto_rawDesc
)

func file_MapMarkPointType_proto_rawDescGZIP() []byte {
	file_MapMarkPointType_proto_rawDescOnce.Do(func() {
		file_MapMarkPointType_proto_rawDescData = protoimpl.X.CompressGZIP(file_MapMarkPointType_proto_rawDescData)
	})
	return file_MapMarkPointType_proto_rawDescData
}

var file_MapMarkPointType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MapMarkPointType_proto_goTypes = []interface{}{
	(MapMarkPointType)(0), // 0: MapMarkPointType
}
var file_MapMarkPointType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MapMarkPointType_proto_init() }
func file_MapMarkPointType_proto_init() {
	if File_MapMarkPointType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MapMarkPointType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MapMarkPointType_proto_goTypes,
		DependencyIndexes: file_MapMarkPointType_proto_depIdxs,
		EnumInfos:         file_MapMarkPointType_proto_enumTypes,
	}.Build()
	File_MapMarkPointType_proto = out.File
	file_MapMarkPointType_proto_rawDesc = nil
	file_MapMarkPointType_proto_goTypes = nil
	file_MapMarkPointType_proto_depIdxs = nil
}
