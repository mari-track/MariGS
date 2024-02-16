// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: BuyGoodsReq.proto

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

type BuyGoodsReq_CmdId int32

const (
	BuyGoodsReq_NONE             BuyGoodsReq_CmdId = 0
	BuyGoodsReq_CMD_ID           BuyGoodsReq_CmdId = 703
	BuyGoodsReq_ENET_CHANNEL_ID  BuyGoodsReq_CmdId = 0
	BuyGoodsReq_ENET_IS_RELIABLE BuyGoodsReq_CmdId = 1
	BuyGoodsReq_IS_ALLOW_CLIENT  BuyGoodsReq_CmdId = 1
)

// Enum value maps for BuyGoodsReq_CmdId.
var (
	BuyGoodsReq_CmdId_name = map[int32]string{
		0:   "NONE",
		703: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	BuyGoodsReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           703,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x BuyGoodsReq_CmdId) Enum() *BuyGoodsReq_CmdId {
	p := new(BuyGoodsReq_CmdId)
	*p = x
	return p
}

func (x BuyGoodsReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BuyGoodsReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_BuyGoodsReq_proto_enumTypes[0].Descriptor()
}

func (BuyGoodsReq_CmdId) Type() protoreflect.EnumType {
	return &file_BuyGoodsReq_proto_enumTypes[0]
}

func (x BuyGoodsReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BuyGoodsReq_CmdId.Descriptor instead.
func (BuyGoodsReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_BuyGoodsReq_proto_rawDescGZIP(), []int{0, 0}
}

type BuyGoodsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopType uint32     `protobuf:"varint,1,opt,name=shop_type,json=shopType,proto3" json:"shop_type,omitempty"`
	Goods    *ShopGoods `protobuf:"bytes,2,opt,name=goods,proto3" json:"goods,omitempty"`
	BuyCount uint32     `protobuf:"varint,3,opt,name=buy_count,json=buyCount,proto3" json:"buy_count,omitempty"`
}

func (x *BuyGoodsReq) Reset() {
	*x = BuyGoodsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BuyGoodsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyGoodsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyGoodsReq) ProtoMessage() {}

func (x *BuyGoodsReq) ProtoReflect() protoreflect.Message {
	mi := &file_BuyGoodsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyGoodsReq.ProtoReflect.Descriptor instead.
func (*BuyGoodsReq) Descriptor() ([]byte, []int) {
	return file_BuyGoodsReq_proto_rawDescGZIP(), []int{0}
}

func (x *BuyGoodsReq) GetShopType() uint32 {
	if x != nil {
		return x.ShopType
	}
	return 0
}

func (x *BuyGoodsReq) GetGoods() *ShopGoods {
	if x != nil {
		return x.Goods
	}
	return nil
}

func (x *BuyGoodsReq) GetBuyCount() uint32 {
	if x != nil {
		return x.BuyCount
	}
	return 0
}

var File_BuyGoodsReq_proto protoreflect.FileDescriptor

var file_BuyGoodsReq_proto_rawDesc = []byte{
	0x0a, 0x11, 0x42, 0x75, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x53, 0x68, 0x6f, 0x70,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x0b,
	0x42, 0x75, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x68, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x73, 0x68, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x68, 0x6f, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x75, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x62, 0x0a,
	0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xbf, 0x05, 0x12, 0x13, 0x0a,
	0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45,
	0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41,
	0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BuyGoodsReq_proto_rawDescOnce sync.Once
	file_BuyGoodsReq_proto_rawDescData = file_BuyGoodsReq_proto_rawDesc
)

func file_BuyGoodsReq_proto_rawDescGZIP() []byte {
	file_BuyGoodsReq_proto_rawDescOnce.Do(func() {
		file_BuyGoodsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_BuyGoodsReq_proto_rawDescData)
	})
	return file_BuyGoodsReq_proto_rawDescData
}

var file_BuyGoodsReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_BuyGoodsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BuyGoodsReq_proto_goTypes = []interface{}{
	(BuyGoodsReq_CmdId)(0), // 0: proto.BuyGoodsReq.CmdId
	(*BuyGoodsReq)(nil),    // 1: proto.BuyGoodsReq
	(*ShopGoods)(nil),      // 2: proto.ShopGoods
}
var file_BuyGoodsReq_proto_depIdxs = []int32{
	2, // 0: proto.BuyGoodsReq.goods:type_name -> proto.ShopGoods
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_BuyGoodsReq_proto_init() }
func file_BuyGoodsReq_proto_init() {
	if File_BuyGoodsReq_proto != nil {
		return
	}
	file_ShopGoods_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BuyGoodsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyGoodsReq); i {
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
			RawDescriptor: file_BuyGoodsReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BuyGoodsReq_proto_goTypes,
		DependencyIndexes: file_BuyGoodsReq_proto_depIdxs,
		EnumInfos:         file_BuyGoodsReq_proto_enumTypes,
		MessageInfos:      file_BuyGoodsReq_proto_msgTypes,
	}.Build()
	File_BuyGoodsReq_proto = out.File
	file_BuyGoodsReq_proto_rawDesc = nil
	file_BuyGoodsReq_proto_goTypes = nil
	file_BuyGoodsReq_proto_depIdxs = nil
}