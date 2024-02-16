// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: TakeFirstShareRewardRsp.proto

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

type TakeFirstShareRewardRsp_CmdId int32

const (
	TakeFirstShareRewardRsp_NONE             TakeFirstShareRewardRsp_CmdId = 0
	TakeFirstShareRewardRsp_CMD_ID           TakeFirstShareRewardRsp_CmdId = 4036
	TakeFirstShareRewardRsp_ENET_CHANNEL_ID  TakeFirstShareRewardRsp_CmdId = 0
	TakeFirstShareRewardRsp_ENET_IS_RELIABLE TakeFirstShareRewardRsp_CmdId = 1
	TakeFirstShareRewardRsp_IS_ALLOW_CLIENT  TakeFirstShareRewardRsp_CmdId = 1
)

// Enum value maps for TakeFirstShareRewardRsp_CmdId.
var (
	TakeFirstShareRewardRsp_CmdId_name = map[int32]string{
		0:    "NONE",
		4036: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	TakeFirstShareRewardRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           4036,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x TakeFirstShareRewardRsp_CmdId) Enum() *TakeFirstShareRewardRsp_CmdId {
	p := new(TakeFirstShareRewardRsp_CmdId)
	*p = x
	return p
}

func (x TakeFirstShareRewardRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TakeFirstShareRewardRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_TakeFirstShareRewardRsp_proto_enumTypes[0].Descriptor()
}

func (TakeFirstShareRewardRsp_CmdId) Type() protoreflect.EnumType {
	return &file_TakeFirstShareRewardRsp_proto_enumTypes[0]
}

func (x TakeFirstShareRewardRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TakeFirstShareRewardRsp_CmdId.Descriptor instead.
func (TakeFirstShareRewardRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_TakeFirstShareRewardRsp_proto_rawDescGZIP(), []int{0, 0}
}

type TakeFirstShareRewardRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32 `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *TakeFirstShareRewardRsp) Reset() {
	*x = TakeFirstShareRewardRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TakeFirstShareRewardRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeFirstShareRewardRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeFirstShareRewardRsp) ProtoMessage() {}

func (x *TakeFirstShareRewardRsp) ProtoReflect() protoreflect.Message {
	mi := &file_TakeFirstShareRewardRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeFirstShareRewardRsp.ProtoReflect.Descriptor instead.
func (*TakeFirstShareRewardRsp) Descriptor() ([]byte, []int) {
	return file_TakeFirstShareRewardRsp_proto_rawDescGZIP(), []int{0}
}

func (x *TakeFirstShareRewardRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_TakeFirstShareRewardRsp_proto protoreflect.FileDescriptor

var file_TakeFirstShareRewardRsp_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x54, 0x61, 0x6b, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x17, 0x54, 0x61, 0x6b, 0x65, 0x46,
	0x69, 0x72, 0x73, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x62, 0x0a, 0x05,
	0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xc4, 0x1f, 0x12, 0x13, 0x0a, 0x0f,
	0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c,
	0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c,
	0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TakeFirstShareRewardRsp_proto_rawDescOnce sync.Once
	file_TakeFirstShareRewardRsp_proto_rawDescData = file_TakeFirstShareRewardRsp_proto_rawDesc
)

func file_TakeFirstShareRewardRsp_proto_rawDescGZIP() []byte {
	file_TakeFirstShareRewardRsp_proto_rawDescOnce.Do(func() {
		file_TakeFirstShareRewardRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_TakeFirstShareRewardRsp_proto_rawDescData)
	})
	return file_TakeFirstShareRewardRsp_proto_rawDescData
}

var file_TakeFirstShareRewardRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_TakeFirstShareRewardRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TakeFirstShareRewardRsp_proto_goTypes = []interface{}{
	(TakeFirstShareRewardRsp_CmdId)(0), // 0: proto.TakeFirstShareRewardRsp.CmdId
	(*TakeFirstShareRewardRsp)(nil),    // 1: proto.TakeFirstShareRewardRsp
}
var file_TakeFirstShareRewardRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TakeFirstShareRewardRsp_proto_init() }
func file_TakeFirstShareRewardRsp_proto_init() {
	if File_TakeFirstShareRewardRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TakeFirstShareRewardRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeFirstShareRewardRsp); i {
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
			RawDescriptor: file_TakeFirstShareRewardRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TakeFirstShareRewardRsp_proto_goTypes,
		DependencyIndexes: file_TakeFirstShareRewardRsp_proto_depIdxs,
		EnumInfos:         file_TakeFirstShareRewardRsp_proto_enumTypes,
		MessageInfos:      file_TakeFirstShareRewardRsp_proto_msgTypes,
	}.Build()
	File_TakeFirstShareRewardRsp_proto = out.File
	file_TakeFirstShareRewardRsp_proto_rawDesc = nil
	file_TakeFirstShareRewardRsp_proto_goTypes = nil
	file_TakeFirstShareRewardRsp_proto_depIdxs = nil
}
