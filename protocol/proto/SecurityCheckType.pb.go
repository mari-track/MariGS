// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: SecurityCheckType.proto

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

type SecurityCheckType int32

const (
	SecurityCheckType_SECURITY_CHECK_NONE  SecurityCheckType = 0
	SecurityCheckType_SECURITY_CHECK_LOGIN SecurityCheckType = 1
)

// Enum value maps for SecurityCheckType.
var (
	SecurityCheckType_name = map[int32]string{
		0: "SECURITY_CHECK_NONE",
		1: "SECURITY_CHECK_LOGIN",
	}
	SecurityCheckType_value = map[string]int32{
		"SECURITY_CHECK_NONE":  0,
		"SECURITY_CHECK_LOGIN": 1,
	}
)

func (x SecurityCheckType) Enum() *SecurityCheckType {
	p := new(SecurityCheckType)
	*p = x
	return p
}

func (x SecurityCheckType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SecurityCheckType) Descriptor() protoreflect.EnumDescriptor {
	return file_SecurityCheckType_proto_enumTypes[0].Descriptor()
}

func (SecurityCheckType) Type() protoreflect.EnumType {
	return &file_SecurityCheckType_proto_enumTypes[0]
}

func (x SecurityCheckType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SecurityCheckType.Descriptor instead.
func (SecurityCheckType) EnumDescriptor() ([]byte, []int) {
	return file_SecurityCheckType_proto_rawDescGZIP(), []int{0}
}

var File_SecurityCheckType_proto protoreflect.FileDescriptor

var file_SecurityCheckType_proto_rawDesc = []byte{
	0x0a, 0x17, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0x46, 0x0a, 0x11, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54,
	0x59, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x18,
	0x0a, 0x14, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b,
	0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SecurityCheckType_proto_rawDescOnce sync.Once
	file_SecurityCheckType_proto_rawDescData = file_SecurityCheckType_proto_rawDesc
)

func file_SecurityCheckType_proto_rawDescGZIP() []byte {
	file_SecurityCheckType_proto_rawDescOnce.Do(func() {
		file_SecurityCheckType_proto_rawDescData = protoimpl.X.CompressGZIP(file_SecurityCheckType_proto_rawDescData)
	})
	return file_SecurityCheckType_proto_rawDescData
}

var file_SecurityCheckType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SecurityCheckType_proto_goTypes = []interface{}{
	(SecurityCheckType)(0), // 0: proto.SecurityCheckType
}
var file_SecurityCheckType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SecurityCheckType_proto_init() }
func file_SecurityCheckType_proto_init() {
	if File_SecurityCheckType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_SecurityCheckType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SecurityCheckType_proto_goTypes,
		DependencyIndexes: file_SecurityCheckType_proto_depIdxs,
		EnumInfos:         file_SecurityCheckType_proto_enumTypes,
	}.Build()
	File_SecurityCheckType_proto = out.File
	file_SecurityCheckType_proto_rawDesc = nil
	file_SecurityCheckType_proto_goTypes = nil
	file_SecurityCheckType_proto_depIdxs = nil
}
