// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: ShopGoodsDisableType.proto

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

type ShopGoodsDisableType int32

const (
	ShopGoodsDisableType_SHOP_GOODS_DISABLE_NONE        ShopGoodsDisableType = 0
	ShopGoodsDisableType_SHOP_GOODS_DISABLE_TALENT_FULL ShopGoodsDisableType = 1
)

// Enum value maps for ShopGoodsDisableType.
var (
	ShopGoodsDisableType_name = map[int32]string{
		0: "SHOP_GOODS_DISABLE_NONE",
		1: "SHOP_GOODS_DISABLE_TALENT_FULL",
	}
	ShopGoodsDisableType_value = map[string]int32{
		"SHOP_GOODS_DISABLE_NONE":        0,
		"SHOP_GOODS_DISABLE_TALENT_FULL": 1,
	}
)

func (x ShopGoodsDisableType) Enum() *ShopGoodsDisableType {
	p := new(ShopGoodsDisableType)
	*p = x
	return p
}

func (x ShopGoodsDisableType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ShopGoodsDisableType) Descriptor() protoreflect.EnumDescriptor {
	return file_ShopGoodsDisableType_proto_enumTypes[0].Descriptor()
}

func (ShopGoodsDisableType) Type() protoreflect.EnumType {
	return &file_ShopGoodsDisableType_proto_enumTypes[0]
}

func (x ShopGoodsDisableType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ShopGoodsDisableType.Descriptor instead.
func (ShopGoodsDisableType) EnumDescriptor() ([]byte, []int) {
	return file_ShopGoodsDisableType_proto_rawDescGZIP(), []int{0}
}

var File_ShopGoodsDisableType_proto protoreflect.FileDescriptor

var file_ShopGoodsDisableType_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x53, 0x68, 0x6f, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x44, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x57, 0x0a, 0x14, 0x53, 0x68, 0x6f, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x53,
	0x48, 0x4f, 0x50, 0x5f, 0x47, 0x4f, 0x4f, 0x44, 0x53, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c,
	0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x48, 0x4f, 0x50,
	0x5f, 0x47, 0x4f, 0x4f, 0x44, 0x53, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x54,
	0x41, 0x4c, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ShopGoodsDisableType_proto_rawDescOnce sync.Once
	file_ShopGoodsDisableType_proto_rawDescData = file_ShopGoodsDisableType_proto_rawDesc
)

func file_ShopGoodsDisableType_proto_rawDescGZIP() []byte {
	file_ShopGoodsDisableType_proto_rawDescOnce.Do(func() {
		file_ShopGoodsDisableType_proto_rawDescData = protoimpl.X.CompressGZIP(file_ShopGoodsDisableType_proto_rawDescData)
	})
	return file_ShopGoodsDisableType_proto_rawDescData
}

var file_ShopGoodsDisableType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ShopGoodsDisableType_proto_goTypes = []interface{}{
	(ShopGoodsDisableType)(0), // 0: proto.ShopGoodsDisableType
}
var file_ShopGoodsDisableType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ShopGoodsDisableType_proto_init() }
func file_ShopGoodsDisableType_proto_init() {
	if File_ShopGoodsDisableType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ShopGoodsDisableType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ShopGoodsDisableType_proto_goTypes,
		DependencyIndexes: file_ShopGoodsDisableType_proto_depIdxs,
		EnumInfos:         file_ShopGoodsDisableType_proto_enumTypes,
	}.Build()
	File_ShopGoodsDisableType_proto = out.File
	file_ShopGoodsDisableType_proto_rawDesc = nil
	file_ShopGoodsDisableType_proto_goTypes = nil
	file_ShopGoodsDisableType_proto_depIdxs = nil
}
