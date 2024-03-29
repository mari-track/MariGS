// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: TakeAchievementRewardReq.proto

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

type TakeAchievementRewardReq_CmdId int32

const (
	TakeAchievementRewardReq_NONE             TakeAchievementRewardReq_CmdId = 0
	TakeAchievementRewardReq_CMD_ID           TakeAchievementRewardReq_CmdId = 2653
	TakeAchievementRewardReq_ENET_CHANNEL_ID  TakeAchievementRewardReq_CmdId = 0
	TakeAchievementRewardReq_ENET_IS_RELIABLE TakeAchievementRewardReq_CmdId = 1
	TakeAchievementRewardReq_IS_ALLOW_CLIENT  TakeAchievementRewardReq_CmdId = 1
)

// Enum value maps for TakeAchievementRewardReq_CmdId.
var (
	TakeAchievementRewardReq_CmdId_name = map[int32]string{
		0:    "NONE",
		2653: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	TakeAchievementRewardReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2653,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x TakeAchievementRewardReq_CmdId) Enum() *TakeAchievementRewardReq_CmdId {
	p := new(TakeAchievementRewardReq_CmdId)
	*p = x
	return p
}

func (x TakeAchievementRewardReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TakeAchievementRewardReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_TakeAchievementRewardReq_proto_enumTypes[0].Descriptor()
}

func (TakeAchievementRewardReq_CmdId) Type() protoreflect.EnumType {
	return &file_TakeAchievementRewardReq_proto_enumTypes[0]
}

func (x TakeAchievementRewardReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TakeAchievementRewardReq_CmdId.Descriptor instead.
func (TakeAchievementRewardReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_TakeAchievementRewardReq_proto_rawDescGZIP(), []int{0, 0}
}

type TakeAchievementRewardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdList []uint32 `protobuf:"varint,1,rep,packed,name=id_list,json=idList,proto3" json:"id_list,omitempty"`
}

func (x *TakeAchievementRewardReq) Reset() {
	*x = TakeAchievementRewardReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TakeAchievementRewardReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeAchievementRewardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeAchievementRewardReq) ProtoMessage() {}

func (x *TakeAchievementRewardReq) ProtoReflect() protoreflect.Message {
	mi := &file_TakeAchievementRewardReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeAchievementRewardReq.ProtoReflect.Descriptor instead.
func (*TakeAchievementRewardReq) Descriptor() ([]byte, []int) {
	return file_TakeAchievementRewardReq_proto_rawDescGZIP(), []int{0}
}

func (x *TakeAchievementRewardReq) GetIdList() []uint32 {
	if x != nil {
		return x.IdList
	}
	return nil
}

var File_TakeAchievementRewardReq_proto protoreflect.FileDescriptor

var file_TakeAchievementRewardReq_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x54, 0x61, 0x6b, 0x65, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x97, 0x01, 0x0a, 0x18, 0x54, 0x61, 0x6b, 0x65, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a,
	0x07, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06,
	0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44,
	0x5f, 0x49, 0x44, 0x10, 0xdd, 0x14, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43,
	0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c,
	0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TakeAchievementRewardReq_proto_rawDescOnce sync.Once
	file_TakeAchievementRewardReq_proto_rawDescData = file_TakeAchievementRewardReq_proto_rawDesc
)

func file_TakeAchievementRewardReq_proto_rawDescGZIP() []byte {
	file_TakeAchievementRewardReq_proto_rawDescOnce.Do(func() {
		file_TakeAchievementRewardReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_TakeAchievementRewardReq_proto_rawDescData)
	})
	return file_TakeAchievementRewardReq_proto_rawDescData
}

var file_TakeAchievementRewardReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_TakeAchievementRewardReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TakeAchievementRewardReq_proto_goTypes = []interface{}{
	(TakeAchievementRewardReq_CmdId)(0), // 0: TakeAchievementRewardReq.CmdId
	(*TakeAchievementRewardReq)(nil),    // 1: TakeAchievementRewardReq
}
var file_TakeAchievementRewardReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TakeAchievementRewardReq_proto_init() }
func file_TakeAchievementRewardReq_proto_init() {
	if File_TakeAchievementRewardReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TakeAchievementRewardReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeAchievementRewardReq); i {
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
			RawDescriptor: file_TakeAchievementRewardReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TakeAchievementRewardReq_proto_goTypes,
		DependencyIndexes: file_TakeAchievementRewardReq_proto_depIdxs,
		EnumInfos:         file_TakeAchievementRewardReq_proto_enumTypes,
		MessageInfos:      file_TakeAchievementRewardReq_proto_msgTypes,
	}.Build()
	File_TakeAchievementRewardReq_proto = out.File
	file_TakeAchievementRewardReq_proto_rawDesc = nil
	file_TakeAchievementRewardReq_proto_goTypes = nil
	file_TakeAchievementRewardReq_proto_depIdxs = nil
}
