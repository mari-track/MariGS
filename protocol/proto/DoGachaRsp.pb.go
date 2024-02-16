// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: DoGachaRsp.proto

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

type DoGachaRsp_CmdId int32

const (
	DoGachaRsp_NONE             DoGachaRsp_CmdId = 0
	DoGachaRsp_CMD_ID           DoGachaRsp_CmdId = 1504
	DoGachaRsp_ENET_CHANNEL_ID  DoGachaRsp_CmdId = 0
	DoGachaRsp_ENET_IS_RELIABLE DoGachaRsp_CmdId = 1
)

// Enum value maps for DoGachaRsp_CmdId.
var (
	DoGachaRsp_CmdId_name = map[int32]string{
		0:    "NONE",
		1504: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	DoGachaRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           1504,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x DoGachaRsp_CmdId) Enum() *DoGachaRsp_CmdId {
	p := new(DoGachaRsp_CmdId)
	*p = x
	return p
}

func (x DoGachaRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DoGachaRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_DoGachaRsp_proto_enumTypes[0].Descriptor()
}

func (DoGachaRsp_CmdId) Type() protoreflect.EnumType {
	return &file_DoGachaRsp_proto_enumTypes[0]
}

func (x DoGachaRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DoGachaRsp_CmdId.Descriptor instead.
func (DoGachaRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_DoGachaRsp_proto_rawDescGZIP(), []int{0, 0}
}

type DoGachaRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode         int32        `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	GachaType       uint32       `protobuf:"varint,2,opt,name=gacha_type,json=gachaType,proto3" json:"gacha_type,omitempty"`
	GachaTimes      uint32       `protobuf:"varint,3,opt,name=gacha_times,json=gachaTimes,proto3" json:"gacha_times,omitempty"`
	GachaScheduleId uint32       `protobuf:"varint,4,opt,name=gacha_schedule_id,json=gachaScheduleId,proto3" json:"gacha_schedule_id,omitempty"`
	GachaItemList   []*GachaItem `protobuf:"bytes,5,rep,name=gacha_item_list,json=gachaItemList,proto3" json:"gacha_item_list,omitempty"`
	NewGachaRandom  uint32       `protobuf:"varint,6,opt,name=new_gacha_random,json=newGachaRandom,proto3" json:"new_gacha_random,omitempty"`
	CostItemId      uint32       `protobuf:"varint,7,opt,name=cost_item_id,json=costItemId,proto3" json:"cost_item_id,omitempty"`
	CostItemNum     uint32       `protobuf:"varint,8,opt,name=cost_item_num,json=costItemNum,proto3" json:"cost_item_num,omitempty"`
	TenCostItemId   uint32       `protobuf:"varint,9,opt,name=ten_cost_item_id,json=tenCostItemId,proto3" json:"ten_cost_item_id,omitempty"`
	TenCostItemNum  uint32       `protobuf:"varint,10,opt,name=ten_cost_item_num,json=tenCostItemNum,proto3" json:"ten_cost_item_num,omitempty"`
	LeftGachaTimes  uint32       `protobuf:"varint,11,opt,name=left_gacha_times,json=leftGachaTimes,proto3" json:"left_gacha_times,omitempty"`
	GachaTimesLimit uint32       `protobuf:"varint,12,opt,name=gacha_times_limit,json=gachaTimesLimit,proto3" json:"gacha_times_limit,omitempty"`
}

func (x *DoGachaRsp) Reset() {
	*x = DoGachaRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DoGachaRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoGachaRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoGachaRsp) ProtoMessage() {}

func (x *DoGachaRsp) ProtoReflect() protoreflect.Message {
	mi := &file_DoGachaRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoGachaRsp.ProtoReflect.Descriptor instead.
func (*DoGachaRsp) Descriptor() ([]byte, []int) {
	return file_DoGachaRsp_proto_rawDescGZIP(), []int{0}
}

func (x *DoGachaRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *DoGachaRsp) GetGachaType() uint32 {
	if x != nil {
		return x.GachaType
	}
	return 0
}

func (x *DoGachaRsp) GetGachaTimes() uint32 {
	if x != nil {
		return x.GachaTimes
	}
	return 0
}

func (x *DoGachaRsp) GetGachaScheduleId() uint32 {
	if x != nil {
		return x.GachaScheduleId
	}
	return 0
}

func (x *DoGachaRsp) GetGachaItemList() []*GachaItem {
	if x != nil {
		return x.GachaItemList
	}
	return nil
}

func (x *DoGachaRsp) GetNewGachaRandom() uint32 {
	if x != nil {
		return x.NewGachaRandom
	}
	return 0
}

func (x *DoGachaRsp) GetCostItemId() uint32 {
	if x != nil {
		return x.CostItemId
	}
	return 0
}

func (x *DoGachaRsp) GetCostItemNum() uint32 {
	if x != nil {
		return x.CostItemNum
	}
	return 0
}

func (x *DoGachaRsp) GetTenCostItemId() uint32 {
	if x != nil {
		return x.TenCostItemId
	}
	return 0
}

func (x *DoGachaRsp) GetTenCostItemNum() uint32 {
	if x != nil {
		return x.TenCostItemNum
	}
	return 0
}

func (x *DoGachaRsp) GetLeftGachaTimes() uint32 {
	if x != nil {
		return x.LeftGachaTimes
	}
	return 0
}

func (x *DoGachaRsp) GetGachaTimesLimit() uint32 {
	if x != nil {
		return x.GachaTimesLimit
	}
	return 0
}

var File_DoGachaRsp_proto protoreflect.FileDescriptor

var file_DoGachaRsp_proto_rawDesc = []byte{
	0x0a, 0x10, 0x44, 0x6f, 0x47, 0x61, 0x63, 0x68, 0x61, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x47, 0x61, 0x63, 0x68, 0x61,
	0x49, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x04, 0x0a, 0x0a, 0x44,
	0x6f, 0x47, 0x61, 0x63, 0x68, 0x61, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x67, 0x61, 0x63, 0x68, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x67, 0x61, 0x63, 0x68, 0x61, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f,
	0x67, 0x61, 0x63, 0x68, 0x61, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x38, 0x0a, 0x0f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x61, 0x63, 0x68, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0d, 0x67, 0x61, 0x63, 0x68,
	0x61, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x65, 0x77,
	0x5f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x47, 0x61, 0x63, 0x68, 0x61, 0x52, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6f,
	0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4e, 0x75, 0x6d, 0x12, 0x27, 0x0a, 0x10, 0x74, 0x65, 0x6e,
	0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0d, 0x74, 0x65, 0x6e, 0x43, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x49, 0x64, 0x12, 0x29, 0x0a, 0x11, 0x74, 0x65, 0x6e, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x74,
	0x65, 0x6e, 0x43, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4e, 0x75, 0x6d, 0x12, 0x28, 0x0a,
	0x10, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6c, 0x65, 0x66, 0x74, 0x47, 0x61, 0x63,
	0x68, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x61, 0x63, 0x68, 0x61,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44,
	0x10, 0xe0, 0x0b, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02,
	0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DoGachaRsp_proto_rawDescOnce sync.Once
	file_DoGachaRsp_proto_rawDescData = file_DoGachaRsp_proto_rawDesc
)

func file_DoGachaRsp_proto_rawDescGZIP() []byte {
	file_DoGachaRsp_proto_rawDescOnce.Do(func() {
		file_DoGachaRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_DoGachaRsp_proto_rawDescData)
	})
	return file_DoGachaRsp_proto_rawDescData
}

var file_DoGachaRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DoGachaRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DoGachaRsp_proto_goTypes = []interface{}{
	(DoGachaRsp_CmdId)(0), // 0: proto.DoGachaRsp.CmdId
	(*DoGachaRsp)(nil),    // 1: proto.DoGachaRsp
	(*GachaItem)(nil),     // 2: proto.GachaItem
}
var file_DoGachaRsp_proto_depIdxs = []int32{
	2, // 0: proto.DoGachaRsp.gacha_item_list:type_name -> proto.GachaItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_DoGachaRsp_proto_init() }
func file_DoGachaRsp_proto_init() {
	if File_DoGachaRsp_proto != nil {
		return
	}
	file_GachaItem_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DoGachaRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoGachaRsp); i {
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
			RawDescriptor: file_DoGachaRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DoGachaRsp_proto_goTypes,
		DependencyIndexes: file_DoGachaRsp_proto_depIdxs,
		EnumInfos:         file_DoGachaRsp_proto_enumTypes,
		MessageInfos:      file_DoGachaRsp_proto_msgTypes,
	}.Build()
	File_DoGachaRsp_proto = out.File
	file_DoGachaRsp_proto_rawDesc = nil
	file_DoGachaRsp_proto_goTypes = nil
	file_DoGachaRsp_proto_depIdxs = nil
}
