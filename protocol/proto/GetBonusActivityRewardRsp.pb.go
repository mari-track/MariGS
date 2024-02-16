// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GetBonusActivityRewardRsp.proto

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

type GetBonusActivityRewardRsp_CmdId int32

const (
	GetBonusActivityRewardRsp_NONE             GetBonusActivityRewardRsp_CmdId = 0
	GetBonusActivityRewardRsp_CMD_ID           GetBonusActivityRewardRsp_CmdId = 2516
	GetBonusActivityRewardRsp_ENET_CHANNEL_ID  GetBonusActivityRewardRsp_CmdId = 0
	GetBonusActivityRewardRsp_ENET_IS_RELIABLE GetBonusActivityRewardRsp_CmdId = 1
)

// Enum value maps for GetBonusActivityRewardRsp_CmdId.
var (
	GetBonusActivityRewardRsp_CmdId_name = map[int32]string{
		0:    "NONE",
		2516: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	GetBonusActivityRewardRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2516,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x GetBonusActivityRewardRsp_CmdId) Enum() *GetBonusActivityRewardRsp_CmdId {
	p := new(GetBonusActivityRewardRsp_CmdId)
	*p = x
	return p
}

func (x GetBonusActivityRewardRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetBonusActivityRewardRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GetBonusActivityRewardRsp_proto_enumTypes[0].Descriptor()
}

func (GetBonusActivityRewardRsp_CmdId) Type() protoreflect.EnumType {
	return &file_GetBonusActivityRewardRsp_proto_enumTypes[0]
}

func (x GetBonusActivityRewardRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetBonusActivityRewardRsp_CmdId.Descriptor instead.
func (GetBonusActivityRewardRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GetBonusActivityRewardRsp_proto_rawDescGZIP(), []int{0, 0}
}

type GetBonusActivityRewardRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode               int32              `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	BonusActivityInfoList *BonusActivityInfo `protobuf:"bytes,2,opt,name=bonus_activity_info_list,json=bonusActivityInfoList,proto3" json:"bonus_activity_info_list,omitempty"`
}

func (x *GetBonusActivityRewardRsp) Reset() {
	*x = GetBonusActivityRewardRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetBonusActivityRewardRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBonusActivityRewardRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBonusActivityRewardRsp) ProtoMessage() {}

func (x *GetBonusActivityRewardRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetBonusActivityRewardRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBonusActivityRewardRsp.ProtoReflect.Descriptor instead.
func (*GetBonusActivityRewardRsp) Descriptor() ([]byte, []int) {
	return file_GetBonusActivityRewardRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetBonusActivityRewardRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetBonusActivityRewardRsp) GetBonusActivityInfoList() *BonusActivityInfo {
	if x != nil {
		return x.BonusActivityInfoList
	}
	return nil
}

var File_GetBonusActivityRewardRsp_proto protoreflect.FileDescriptor

var file_GetBonusActivityRewardRsp_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x51, 0x0a, 0x18, 0x62, 0x6f, 0x6e,
	0x75, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x15, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x05,
	0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xd4, 0x13, 0x12, 0x13, 0x0a, 0x0f,
	0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c,
	0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetBonusActivityRewardRsp_proto_rawDescOnce sync.Once
	file_GetBonusActivityRewardRsp_proto_rawDescData = file_GetBonusActivityRewardRsp_proto_rawDesc
)

func file_GetBonusActivityRewardRsp_proto_rawDescGZIP() []byte {
	file_GetBonusActivityRewardRsp_proto_rawDescOnce.Do(func() {
		file_GetBonusActivityRewardRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetBonusActivityRewardRsp_proto_rawDescData)
	})
	return file_GetBonusActivityRewardRsp_proto_rawDescData
}

var file_GetBonusActivityRewardRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GetBonusActivityRewardRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetBonusActivityRewardRsp_proto_goTypes = []interface{}{
	(GetBonusActivityRewardRsp_CmdId)(0), // 0: proto.GetBonusActivityRewardRsp.CmdId
	(*GetBonusActivityRewardRsp)(nil),    // 1: proto.GetBonusActivityRewardRsp
	(*BonusActivityInfo)(nil),            // 2: proto.BonusActivityInfo
}
var file_GetBonusActivityRewardRsp_proto_depIdxs = []int32{
	2, // 0: proto.GetBonusActivityRewardRsp.bonus_activity_info_list:type_name -> proto.BonusActivityInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetBonusActivityRewardRsp_proto_init() }
func file_GetBonusActivityRewardRsp_proto_init() {
	if File_GetBonusActivityRewardRsp_proto != nil {
		return
	}
	file_BonusActivityInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetBonusActivityRewardRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBonusActivityRewardRsp); i {
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
			RawDescriptor: file_GetBonusActivityRewardRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetBonusActivityRewardRsp_proto_goTypes,
		DependencyIndexes: file_GetBonusActivityRewardRsp_proto_depIdxs,
		EnumInfos:         file_GetBonusActivityRewardRsp_proto_enumTypes,
		MessageInfos:      file_GetBonusActivityRewardRsp_proto_msgTypes,
	}.Build()
	File_GetBonusActivityRewardRsp_proto = out.File
	file_GetBonusActivityRewardRsp_proto_rawDesc = nil
	file_GetBonusActivityRewardRsp_proto_goTypes = nil
	file_GetBonusActivityRewardRsp_proto_depIdxs = nil
}