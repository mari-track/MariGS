// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: LuaShellNotifyType.proto

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

type LuaShellNotifyType int32

const (
	LuaShellNotifyType_LUASHELL_NOTIFY LuaShellNotifyType = 0
	LuaShellNotifyType_LUASHELL_HIDE   LuaShellNotifyType = 1
)

// Enum value maps for LuaShellNotifyType.
var (
	LuaShellNotifyType_name = map[int32]string{
		0: "LUASHELL_NOTIFY",
		1: "LUASHELL_HIDE",
	}
	LuaShellNotifyType_value = map[string]int32{
		"LUASHELL_NOTIFY": 0,
		"LUASHELL_HIDE":   1,
	}
)

func (x LuaShellNotifyType) Enum() *LuaShellNotifyType {
	p := new(LuaShellNotifyType)
	*p = x
	return p
}

func (x LuaShellNotifyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LuaShellNotifyType) Descriptor() protoreflect.EnumDescriptor {
	return file_LuaShellNotifyType_proto_enumTypes[0].Descriptor()
}

func (LuaShellNotifyType) Type() protoreflect.EnumType {
	return &file_LuaShellNotifyType_proto_enumTypes[0]
}

func (x LuaShellNotifyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LuaShellNotifyType.Descriptor instead.
func (LuaShellNotifyType) EnumDescriptor() ([]byte, []int) {
	return file_LuaShellNotifyType_proto_rawDescGZIP(), []int{0}
}

var File_LuaShellNotifyType_proto protoreflect.FileDescriptor

var file_LuaShellNotifyType_proto_rawDesc = []byte{
	0x0a, 0x18, 0x4c, 0x75, 0x61, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0x3c, 0x0a, 0x12, 0x4c, 0x75, 0x61, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x55, 0x41, 0x53, 0x48,
	0x45, 0x4c, 0x4c, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d,
	0x4c, 0x55, 0x41, 0x53, 0x48, 0x45, 0x4c, 0x4c, 0x5f, 0x48, 0x49, 0x44, 0x45, 0x10, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_LuaShellNotifyType_proto_rawDescOnce sync.Once
	file_LuaShellNotifyType_proto_rawDescData = file_LuaShellNotifyType_proto_rawDesc
)

func file_LuaShellNotifyType_proto_rawDescGZIP() []byte {
	file_LuaShellNotifyType_proto_rawDescOnce.Do(func() {
		file_LuaShellNotifyType_proto_rawDescData = protoimpl.X.CompressGZIP(file_LuaShellNotifyType_proto_rawDescData)
	})
	return file_LuaShellNotifyType_proto_rawDescData
}

var file_LuaShellNotifyType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_LuaShellNotifyType_proto_goTypes = []interface{}{
	(LuaShellNotifyType)(0), // 0: proto.LuaShellNotifyType
}
var file_LuaShellNotifyType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_LuaShellNotifyType_proto_init() }
func file_LuaShellNotifyType_proto_init() {
	if File_LuaShellNotifyType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_LuaShellNotifyType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LuaShellNotifyType_proto_goTypes,
		DependencyIndexes: file_LuaShellNotifyType_proto_depIdxs,
		EnumInfos:         file_LuaShellNotifyType_proto_enumTypes,
	}.Build()
	File_LuaShellNotifyType_proto = out.File
	file_LuaShellNotifyType_proto_rawDesc = nil
	file_LuaShellNotifyType_proto_goTypes = nil
	file_LuaShellNotifyType_proto_depIdxs = nil
}
