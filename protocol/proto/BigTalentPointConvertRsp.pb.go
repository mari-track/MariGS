// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: BigTalentPointConvertRsp.proto

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

type BigTalentPointConvertRsp_CmdId int32

const (
	BigTalentPointConvertRsp_NONE             BigTalentPointConvertRsp_CmdId = 0
	BigTalentPointConvertRsp_CMD_ID           BigTalentPointConvertRsp_CmdId = 1006
	BigTalentPointConvertRsp_ENET_CHANNEL_ID  BigTalentPointConvertRsp_CmdId = 0
	BigTalentPointConvertRsp_ENET_IS_RELIABLE BigTalentPointConvertRsp_CmdId = 1
)

// Enum value maps for BigTalentPointConvertRsp_CmdId.
var (
	BigTalentPointConvertRsp_CmdId_name = map[int32]string{
		0:    "NONE",
		1006: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	BigTalentPointConvertRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           1006,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x BigTalentPointConvertRsp_CmdId) Enum() *BigTalentPointConvertRsp_CmdId {
	p := new(BigTalentPointConvertRsp_CmdId)
	*p = x
	return p
}

func (x BigTalentPointConvertRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BigTalentPointConvertRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_BigTalentPointConvertRsp_proto_enumTypes[0].Descriptor()
}

func (BigTalentPointConvertRsp_CmdId) Type() protoreflect.EnumType {
	return &file_BigTalentPointConvertRsp_proto_enumTypes[0]
}

func (x BigTalentPointConvertRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BigTalentPointConvertRsp_CmdId.Descriptor instead.
func (BigTalentPointConvertRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_BigTalentPointConvertRsp_proto_rawDescGZIP(), []int{0, 0}
}

type BigTalentPointConvertRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode    int32  `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	AvatarGuid uint64 `protobuf:"varint,2,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
}

func (x *BigTalentPointConvertRsp) Reset() {
	*x = BigTalentPointConvertRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BigTalentPointConvertRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigTalentPointConvertRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigTalentPointConvertRsp) ProtoMessage() {}

func (x *BigTalentPointConvertRsp) ProtoReflect() protoreflect.Message {
	mi := &file_BigTalentPointConvertRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigTalentPointConvertRsp.ProtoReflect.Descriptor instead.
func (*BigTalentPointConvertRsp) Descriptor() ([]byte, []int) {
	return file_BigTalentPointConvertRsp_proto_rawDescGZIP(), []int{0}
}

func (x *BigTalentPointConvertRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *BigTalentPointConvertRsp) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

var File_BigTalentPointConvertRsp_proto protoreflect.FileDescriptor

var file_BigTalentPointConvertRsp_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x42, 0x69, 0x67, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa4, 0x01, 0x0a, 0x18, 0x42, 0x69, 0x67, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x47, 0x75, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49,
	0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43,
	0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xee, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BigTalentPointConvertRsp_proto_rawDescOnce sync.Once
	file_BigTalentPointConvertRsp_proto_rawDescData = file_BigTalentPointConvertRsp_proto_rawDesc
)

func file_BigTalentPointConvertRsp_proto_rawDescGZIP() []byte {
	file_BigTalentPointConvertRsp_proto_rawDescOnce.Do(func() {
		file_BigTalentPointConvertRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_BigTalentPointConvertRsp_proto_rawDescData)
	})
	return file_BigTalentPointConvertRsp_proto_rawDescData
}

var file_BigTalentPointConvertRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_BigTalentPointConvertRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BigTalentPointConvertRsp_proto_goTypes = []interface{}{
	(BigTalentPointConvertRsp_CmdId)(0), // 0: BigTalentPointConvertRsp.CmdId
	(*BigTalentPointConvertRsp)(nil),    // 1: BigTalentPointConvertRsp
}
var file_BigTalentPointConvertRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BigTalentPointConvertRsp_proto_init() }
func file_BigTalentPointConvertRsp_proto_init() {
	if File_BigTalentPointConvertRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BigTalentPointConvertRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigTalentPointConvertRsp); i {
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
			RawDescriptor: file_BigTalentPointConvertRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BigTalentPointConvertRsp_proto_goTypes,
		DependencyIndexes: file_BigTalentPointConvertRsp_proto_depIdxs,
		EnumInfos:         file_BigTalentPointConvertRsp_proto_enumTypes,
		MessageInfos:      file_BigTalentPointConvertRsp_proto_msgTypes,
	}.Build()
	File_BigTalentPointConvertRsp_proto = out.File
	file_BigTalentPointConvertRsp_proto_rawDesc = nil
	file_BigTalentPointConvertRsp_proto_goTypes = nil
	file_BigTalentPointConvertRsp_proto_depIdxs = nil
}
