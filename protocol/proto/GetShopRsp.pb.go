// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GetShopRsp.proto

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

type GetShopRsp_CmdId int32

const (
	GetShopRsp_NONE             GetShopRsp_CmdId = 0
	GetShopRsp_CMD_ID           GetShopRsp_CmdId = 702
	GetShopRsp_ENET_CHANNEL_ID  GetShopRsp_CmdId = 0
	GetShopRsp_ENET_IS_RELIABLE GetShopRsp_CmdId = 1
)

// Enum value maps for GetShopRsp_CmdId.
var (
	GetShopRsp_CmdId_name = map[int32]string{
		0:   "NONE",
		702: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	GetShopRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           702,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x GetShopRsp_CmdId) Enum() *GetShopRsp_CmdId {
	p := new(GetShopRsp_CmdId)
	*p = x
	return p
}

func (x GetShopRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetShopRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GetShopRsp_proto_enumTypes[0].Descriptor()
}

func (GetShopRsp_CmdId) Type() protoreflect.EnumType {
	return &file_GetShopRsp_proto_enumTypes[0]
}

func (x GetShopRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetShopRsp_CmdId.Descriptor instead.
func (GetShopRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GetShopRsp_proto_rawDescGZIP(), []int{0, 0}
}

type GetShopRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32 `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Shop    *Shop `protobuf:"bytes,2,opt,name=shop,proto3" json:"shop,omitempty"`
}

func (x *GetShopRsp) Reset() {
	*x = GetShopRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetShopRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShopRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShopRsp) ProtoMessage() {}

func (x *GetShopRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetShopRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShopRsp.ProtoReflect.Descriptor instead.
func (*GetShopRsp) Descriptor() ([]byte, []int) {
	return file_GetShopRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetShopRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetShopRsp) GetShop() *Shop {
	if x != nil {
		return x.Shop
	}
	return nil
}

var File_GetShopRsp_proto protoreflect.FileDescriptor

var file_GetShopRsp_proto_rawDesc = []byte{
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0a, 0x53, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90,
	0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x04, 0x73, 0x68,
	0x6f, 0x70, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10,
	0xbe, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e,
	0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f,
	0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetShopRsp_proto_rawDescOnce sync.Once
	file_GetShopRsp_proto_rawDescData = file_GetShopRsp_proto_rawDesc
)

func file_GetShopRsp_proto_rawDescGZIP() []byte {
	file_GetShopRsp_proto_rawDescOnce.Do(func() {
		file_GetShopRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetShopRsp_proto_rawDescData)
	})
	return file_GetShopRsp_proto_rawDescData
}

var file_GetShopRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GetShopRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetShopRsp_proto_goTypes = []interface{}{
	(GetShopRsp_CmdId)(0), // 0: GetShopRsp.CmdId
	(*GetShopRsp)(nil),    // 1: GetShopRsp
	(*Shop)(nil),          // 2: Shop
}
var file_GetShopRsp_proto_depIdxs = []int32{
	2, // 0: GetShopRsp.shop:type_name -> Shop
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetShopRsp_proto_init() }
func file_GetShopRsp_proto_init() {
	if File_GetShopRsp_proto != nil {
		return
	}
	file_Shop_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetShopRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShopRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GetShopRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetShopRsp_proto_goTypes,
		DependencyIndexes: file_GetShopRsp_proto_depIdxs,
		EnumInfos:         file_GetShopRsp_proto_enumTypes,
		MessageInfos:      file_GetShopRsp_proto_msgTypes,
	}.Build()
	File_GetShopRsp_proto = out.File
	file_GetShopRsp_proto_rawDesc = nil
	file_GetShopRsp_proto_goTypes = nil
	file_GetShopRsp_proto_depIdxs = nil
}
