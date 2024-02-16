// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GetBattlePassProductReq.proto

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

type GetBattlePassProductReq_CmdId int32

const (
	GetBattlePassProductReq_NONE             GetBattlePassProductReq_CmdId = 0
	GetBattlePassProductReq_CMD_ID           GetBattlePassProductReq_CmdId = 2609
	GetBattlePassProductReq_ENET_CHANNEL_ID  GetBattlePassProductReq_CmdId = 0
	GetBattlePassProductReq_ENET_IS_RELIABLE GetBattlePassProductReq_CmdId = 1
	GetBattlePassProductReq_IS_ALLOW_CLIENT  GetBattlePassProductReq_CmdId = 1
)

// Enum value maps for GetBattlePassProductReq_CmdId.
var (
	GetBattlePassProductReq_CmdId_name = map[int32]string{
		0:    "NONE",
		2609: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	GetBattlePassProductReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2609,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x GetBattlePassProductReq_CmdId) Enum() *GetBattlePassProductReq_CmdId {
	p := new(GetBattlePassProductReq_CmdId)
	*p = x
	return p
}

func (x GetBattlePassProductReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetBattlePassProductReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GetBattlePassProductReq_proto_enumTypes[0].Descriptor()
}

func (GetBattlePassProductReq_CmdId) Type() protoreflect.EnumType {
	return &file_GetBattlePassProductReq_proto_enumTypes[0]
}

func (x GetBattlePassProductReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetBattlePassProductReq_CmdId.Descriptor instead.
func (GetBattlePassProductReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GetBattlePassProductReq_proto_rawDescGZIP(), []int{0, 0}
}

type GetBattlePassProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BattlePassProductPlayType uint32 `protobuf:"varint,1,opt,name=battle_pass_product_play_type,json=battlePassProductPlayType,proto3" json:"battle_pass_product_play_type,omitempty"`
}

func (x *GetBattlePassProductReq) Reset() {
	*x = GetBattlePassProductReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetBattlePassProductReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBattlePassProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBattlePassProductReq) ProtoMessage() {}

func (x *GetBattlePassProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetBattlePassProductReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBattlePassProductReq.ProtoReflect.Descriptor instead.
func (*GetBattlePassProductReq) Descriptor() ([]byte, []int) {
	return file_GetBattlePassProductReq_proto_rawDescGZIP(), []int{0}
}

func (x *GetBattlePassProductReq) GetBattlePassProductPlayType() uint32 {
	if x != nil {
		return x.BattlePassProductPlayType
	}
	return 0
}

var File_GetBattlePassProductReq_proto protoreflect.FileDescriptor

var file_GetBattlePassProductReq_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x40, 0x0a, 0x1d, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x73,
	0x73, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x19, 0x62, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x6c, 0x61, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49,
	0x44, 0x10, 0xb1, 0x14, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45,
	0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12,
	0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45,
	0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetBattlePassProductReq_proto_rawDescOnce sync.Once
	file_GetBattlePassProductReq_proto_rawDescData = file_GetBattlePassProductReq_proto_rawDesc
)

func file_GetBattlePassProductReq_proto_rawDescGZIP() []byte {
	file_GetBattlePassProductReq_proto_rawDescOnce.Do(func() {
		file_GetBattlePassProductReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetBattlePassProductReq_proto_rawDescData)
	})
	return file_GetBattlePassProductReq_proto_rawDescData
}

var file_GetBattlePassProductReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GetBattlePassProductReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetBattlePassProductReq_proto_goTypes = []interface{}{
	(GetBattlePassProductReq_CmdId)(0), // 0: proto.GetBattlePassProductReq.CmdId
	(*GetBattlePassProductReq)(nil),    // 1: proto.GetBattlePassProductReq
}
var file_GetBattlePassProductReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetBattlePassProductReq_proto_init() }
func file_GetBattlePassProductReq_proto_init() {
	if File_GetBattlePassProductReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetBattlePassProductReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBattlePassProductReq); i {
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
			RawDescriptor: file_GetBattlePassProductReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetBattlePassProductReq_proto_goTypes,
		DependencyIndexes: file_GetBattlePassProductReq_proto_depIdxs,
		EnumInfos:         file_GetBattlePassProductReq_proto_enumTypes,
		MessageInfos:      file_GetBattlePassProductReq_proto_msgTypes,
	}.Build()
	File_GetBattlePassProductReq_proto = out.File
	file_GetBattlePassProductReq_proto_rawDesc = nil
	file_GetBattlePassProductReq_proto_goTypes = nil
	file_GetBattlePassProductReq_proto_depIdxs = nil
}
