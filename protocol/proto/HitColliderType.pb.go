// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: HitColliderType.proto

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

type HitColliderType int32

const (
	HitColliderType_HIT_COLLIDER_INVALID     HitColliderType = 0
	HitColliderType_HIT_COLLIDER_HIT_BOX     HitColliderType = 1
	HitColliderType_HIT_COLLIDER_WET_HIT_BOX HitColliderType = 2
	HitColliderType_HIT_COLLIDER_HEAD_BOX    HitColliderType = 3
)

// Enum value maps for HitColliderType.
var (
	HitColliderType_name = map[int32]string{
		0: "HIT_COLLIDER_INVALID",
		1: "HIT_COLLIDER_HIT_BOX",
		2: "HIT_COLLIDER_WET_HIT_BOX",
		3: "HIT_COLLIDER_HEAD_BOX",
	}
	HitColliderType_value = map[string]int32{
		"HIT_COLLIDER_INVALID":     0,
		"HIT_COLLIDER_HIT_BOX":     1,
		"HIT_COLLIDER_WET_HIT_BOX": 2,
		"HIT_COLLIDER_HEAD_BOX":    3,
	}
)

func (x HitColliderType) Enum() *HitColliderType {
	p := new(HitColliderType)
	*p = x
	return p
}

func (x HitColliderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HitColliderType) Descriptor() protoreflect.EnumDescriptor {
	return file_HitColliderType_proto_enumTypes[0].Descriptor()
}

func (HitColliderType) Type() protoreflect.EnumType {
	return &file_HitColliderType_proto_enumTypes[0]
}

func (x HitColliderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HitColliderType.Descriptor instead.
func (HitColliderType) EnumDescriptor() ([]byte, []int) {
	return file_HitColliderType_proto_rawDescGZIP(), []int{0}
}

var File_HitColliderType_proto protoreflect.FileDescriptor

var file_HitColliderType_proto_rawDesc = []byte{
	0x0a, 0x15, 0x48, 0x69, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x7e,
	0x0a, 0x0f, 0x48, 0x69, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x14, 0x48, 0x49, 0x54, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x49, 0x44, 0x45,
	0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x48,
	0x49, 0x54, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x48, 0x49, 0x54, 0x5f,
	0x42, 0x4f, 0x58, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x48, 0x49, 0x54, 0x5f, 0x43, 0x4f, 0x4c,
	0x4c, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x57, 0x45, 0x54, 0x5f, 0x48, 0x49, 0x54, 0x5f, 0x42, 0x4f,
	0x58, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x48, 0x49, 0x54, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x49,
	0x44, 0x45, 0x52, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x5f, 0x42, 0x4f, 0x58, 0x10, 0x03, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_HitColliderType_proto_rawDescOnce sync.Once
	file_HitColliderType_proto_rawDescData = file_HitColliderType_proto_rawDesc
)

func file_HitColliderType_proto_rawDescGZIP() []byte {
	file_HitColliderType_proto_rawDescOnce.Do(func() {
		file_HitColliderType_proto_rawDescData = protoimpl.X.CompressGZIP(file_HitColliderType_proto_rawDescData)
	})
	return file_HitColliderType_proto_rawDescData
}

var file_HitColliderType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_HitColliderType_proto_goTypes = []interface{}{
	(HitColliderType)(0), // 0: proto.HitColliderType
}
var file_HitColliderType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HitColliderType_proto_init() }
func file_HitColliderType_proto_init() {
	if File_HitColliderType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_HitColliderType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HitColliderType_proto_goTypes,
		DependencyIndexes: file_HitColliderType_proto_depIdxs,
		EnumInfos:         file_HitColliderType_proto_enumTypes,
	}.Build()
	File_HitColliderType_proto = out.File
	file_HitColliderType_proto_rawDesc = nil
	file_HitColliderType_proto_goTypes = nil
	file_HitColliderType_proto_depIdxs = nil
}
