// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: SeaLampFlyLampReq.proto

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

type SeaLampFlyLampReq_CmdId int32

const (
	SeaLampFlyLampReq_NONE             SeaLampFlyLampReq_CmdId = 0
	SeaLampFlyLampReq_CMD_ID           SeaLampFlyLampReq_CmdId = 2014
	SeaLampFlyLampReq_ENET_CHANNEL_ID  SeaLampFlyLampReq_CmdId = 0
	SeaLampFlyLampReq_ENET_IS_RELIABLE SeaLampFlyLampReq_CmdId = 1
	SeaLampFlyLampReq_IS_ALLOW_CLIENT  SeaLampFlyLampReq_CmdId = 1
)

// Enum value maps for SeaLampFlyLampReq_CmdId.
var (
	SeaLampFlyLampReq_CmdId_name = map[int32]string{
		0:    "NONE",
		2014: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	SeaLampFlyLampReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2014,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x SeaLampFlyLampReq_CmdId) Enum() *SeaLampFlyLampReq_CmdId {
	p := new(SeaLampFlyLampReq_CmdId)
	*p = x
	return p
}

func (x SeaLampFlyLampReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SeaLampFlyLampReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_SeaLampFlyLampReq_proto_enumTypes[0].Descriptor()
}

func (SeaLampFlyLampReq_CmdId) Type() protoreflect.EnumType {
	return &file_SeaLampFlyLampReq_proto_enumTypes[0]
}

func (x SeaLampFlyLampReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SeaLampFlyLampReq_CmdId.Descriptor instead.
func (SeaLampFlyLampReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_SeaLampFlyLampReq_proto_rawDescGZIP(), []int{0, 0}
}

type SeaLampFlyLampReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId uint32 `protobuf:"varint,4,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	ItemId     uint32 `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	WishText   string `protobuf:"bytes,2,opt,name=wish_text,json=wishText,proto3" json:"wish_text,omitempty"`
	ItemNum    uint32 `protobuf:"varint,3,opt,name=item_num,json=itemNum,proto3" json:"item_num,omitempty"`
}

func (x *SeaLampFlyLampReq) Reset() {
	*x = SeaLampFlyLampReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SeaLampFlyLampReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeaLampFlyLampReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeaLampFlyLampReq) ProtoMessage() {}

func (x *SeaLampFlyLampReq) ProtoReflect() protoreflect.Message {
	mi := &file_SeaLampFlyLampReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeaLampFlyLampReq.ProtoReflect.Descriptor instead.
func (*SeaLampFlyLampReq) Descriptor() ([]byte, []int) {
	return file_SeaLampFlyLampReq_proto_rawDescGZIP(), []int{0}
}

func (x *SeaLampFlyLampReq) GetActivityId() uint32 {
	if x != nil {
		return x.ActivityId
	}
	return 0
}

func (x *SeaLampFlyLampReq) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *SeaLampFlyLampReq) GetWishText() string {
	if x != nil {
		return x.WishText
	}
	return ""
}

func (x *SeaLampFlyLampReq) GetItemNum() uint32 {
	if x != nil {
		return x.ItemNum
	}
	return 0
}

var File_SeaLampFlyLampReq_proto protoreflect.FileDescriptor

var file_SeaLampFlyLampReq_proto_rawDesc = []byte{
	0x0a, 0x17, 0x53, 0x65, 0x61, 0x4c, 0x61, 0x6d, 0x70, 0x46, 0x6c, 0x79, 0x4c, 0x61, 0x6d, 0x70,
	0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x01, 0x0a, 0x11, 0x53, 0x65,
	0x61, 0x4c, 0x61, 0x6d, 0x70, 0x46, 0x6c, 0x79, 0x4c, 0x61, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x69, 0x73,
	0x68, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x69,
	0x73, 0x68, 0x54, 0x65, 0x78, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x75,
	0x6d, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xde,
	0x0f, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45,
	0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49,
	0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f,
	0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10,
	0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SeaLampFlyLampReq_proto_rawDescOnce sync.Once
	file_SeaLampFlyLampReq_proto_rawDescData = file_SeaLampFlyLampReq_proto_rawDesc
)

func file_SeaLampFlyLampReq_proto_rawDescGZIP() []byte {
	file_SeaLampFlyLampReq_proto_rawDescOnce.Do(func() {
		file_SeaLampFlyLampReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SeaLampFlyLampReq_proto_rawDescData)
	})
	return file_SeaLampFlyLampReq_proto_rawDescData
}

var file_SeaLampFlyLampReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SeaLampFlyLampReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SeaLampFlyLampReq_proto_goTypes = []interface{}{
	(SeaLampFlyLampReq_CmdId)(0), // 0: SeaLampFlyLampReq.CmdId
	(*SeaLampFlyLampReq)(nil),    // 1: SeaLampFlyLampReq
}
var file_SeaLampFlyLampReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SeaLampFlyLampReq_proto_init() }
func file_SeaLampFlyLampReq_proto_init() {
	if File_SeaLampFlyLampReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SeaLampFlyLampReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeaLampFlyLampReq); i {
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
			RawDescriptor: file_SeaLampFlyLampReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SeaLampFlyLampReq_proto_goTypes,
		DependencyIndexes: file_SeaLampFlyLampReq_proto_depIdxs,
		EnumInfos:         file_SeaLampFlyLampReq_proto_enumTypes,
		MessageInfos:      file_SeaLampFlyLampReq_proto_msgTypes,
	}.Build()
	File_SeaLampFlyLampReq_proto = out.File
	file_SeaLampFlyLampReq_proto_rawDesc = nil
	file_SeaLampFlyLampReq_proto_goTypes = nil
	file_SeaLampFlyLampReq_proto_depIdxs = nil
}
