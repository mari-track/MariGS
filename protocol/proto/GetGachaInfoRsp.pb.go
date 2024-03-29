// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GetGachaInfoRsp.proto

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

type GetGachaInfoRsp_CmdId int32

const (
	GetGachaInfoRsp_NONE             GetGachaInfoRsp_CmdId = 0
	GetGachaInfoRsp_CMD_ID           GetGachaInfoRsp_CmdId = 1502
	GetGachaInfoRsp_ENET_CHANNEL_ID  GetGachaInfoRsp_CmdId = 0
	GetGachaInfoRsp_ENET_IS_RELIABLE GetGachaInfoRsp_CmdId = 1
)

// Enum value maps for GetGachaInfoRsp_CmdId.
var (
	GetGachaInfoRsp_CmdId_name = map[int32]string{
		0:    "NONE",
		1502: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	GetGachaInfoRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           1502,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x GetGachaInfoRsp_CmdId) Enum() *GetGachaInfoRsp_CmdId {
	p := new(GetGachaInfoRsp_CmdId)
	*p = x
	return p
}

func (x GetGachaInfoRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetGachaInfoRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GetGachaInfoRsp_proto_enumTypes[0].Descriptor()
}

func (GetGachaInfoRsp_CmdId) Type() protoreflect.EnumType {
	return &file_GetGachaInfoRsp_proto_enumTypes[0]
}

func (x GetGachaInfoRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetGachaInfoRsp_CmdId.Descriptor instead.
func (GetGachaInfoRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GetGachaInfoRsp_proto_rawDescGZIP(), []int{0, 0}
}

type GetGachaInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode       int32        `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	GachaInfoList []*GachaInfo `protobuf:"bytes,2,rep,name=gacha_info_list,json=gachaInfoList,proto3" json:"gacha_info_list,omitempty"`
	GachaRandom   uint32       `protobuf:"varint,3,opt,name=gacha_random,json=gachaRandom,proto3" json:"gacha_random,omitempty"`
}

func (x *GetGachaInfoRsp) Reset() {
	*x = GetGachaInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetGachaInfoRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGachaInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGachaInfoRsp) ProtoMessage() {}

func (x *GetGachaInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetGachaInfoRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGachaInfoRsp.ProtoReflect.Descriptor instead.
func (*GetGachaInfoRsp) Descriptor() ([]byte, []int) {
	return file_GetGachaInfoRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetGachaInfoRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetGachaInfoRsp) GetGachaInfoList() []*GachaInfo {
	if x != nil {
		return x.GachaInfoList
	}
	return nil
}

func (x *GetGachaInfoRsp) GetGachaRandom() uint32 {
	if x != nil {
		return x.GachaRandom
	}
	return 0
}

var File_GetGachaInfoRsp_proto protoreflect.FileDescriptor

var file_GetGachaInfoRsp_proto_rawDesc = []byte{
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x47, 0x61, 0x63, 0x68, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x47, 0x61, 0x63, 0x68, 0x61, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x47, 0x61, 0x63, 0x68, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x0f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x47, 0x61, 0x63, 0x68, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x67, 0x61, 0x63,
	0x68, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x61,
	0x63, 0x68, 0x61, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x67, 0x61, 0x63, 0x68, 0x61, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x22, 0x4d, 0x0a,
	0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xde, 0x0b, 0x12, 0x13, 0x0a,
	0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45,
	0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetGachaInfoRsp_proto_rawDescOnce sync.Once
	file_GetGachaInfoRsp_proto_rawDescData = file_GetGachaInfoRsp_proto_rawDesc
)

func file_GetGachaInfoRsp_proto_rawDescGZIP() []byte {
	file_GetGachaInfoRsp_proto_rawDescOnce.Do(func() {
		file_GetGachaInfoRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetGachaInfoRsp_proto_rawDescData)
	})
	return file_GetGachaInfoRsp_proto_rawDescData
}

var file_GetGachaInfoRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GetGachaInfoRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetGachaInfoRsp_proto_goTypes = []interface{}{
	(GetGachaInfoRsp_CmdId)(0), // 0: GetGachaInfoRsp.CmdId
	(*GetGachaInfoRsp)(nil),    // 1: GetGachaInfoRsp
	(*GachaInfo)(nil),          // 2: GachaInfo
}
var file_GetGachaInfoRsp_proto_depIdxs = []int32{
	2, // 0: GetGachaInfoRsp.gacha_info_list:type_name -> GachaInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetGachaInfoRsp_proto_init() }
func file_GetGachaInfoRsp_proto_init() {
	if File_GetGachaInfoRsp_proto != nil {
		return
	}
	file_GachaInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetGachaInfoRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGachaInfoRsp); i {
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
			RawDescriptor: file_GetGachaInfoRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetGachaInfoRsp_proto_goTypes,
		DependencyIndexes: file_GetGachaInfoRsp_proto_depIdxs,
		EnumInfos:         file_GetGachaInfoRsp_proto_enumTypes,
		MessageInfos:      file_GetGachaInfoRsp_proto_msgTypes,
	}.Build()
	File_GetGachaInfoRsp_proto = out.File
	file_GetGachaInfoRsp_proto_rawDesc = nil
	file_GetGachaInfoRsp_proto_goTypes = nil
	file_GetGachaInfoRsp_proto_depIdxs = nil
}
