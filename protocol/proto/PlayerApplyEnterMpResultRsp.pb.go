// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: PlayerApplyEnterMpResultRsp.proto

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

type PlayerApplyEnterMpResultRsp_CmdId int32

const (
	PlayerApplyEnterMpResultRsp_NONE             PlayerApplyEnterMpResultRsp_CmdId = 0
	PlayerApplyEnterMpResultRsp_CMD_ID           PlayerApplyEnterMpResultRsp_CmdId = 1806
	PlayerApplyEnterMpResultRsp_ENET_CHANNEL_ID  PlayerApplyEnterMpResultRsp_CmdId = 0
	PlayerApplyEnterMpResultRsp_ENET_IS_RELIABLE PlayerApplyEnterMpResultRsp_CmdId = 1
)

// Enum value maps for PlayerApplyEnterMpResultRsp_CmdId.
var (
	PlayerApplyEnterMpResultRsp_CmdId_name = map[int32]string{
		0:    "NONE",
		1806: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	PlayerApplyEnterMpResultRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           1806,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x PlayerApplyEnterMpResultRsp_CmdId) Enum() *PlayerApplyEnterMpResultRsp_CmdId {
	p := new(PlayerApplyEnterMpResultRsp_CmdId)
	*p = x
	return p
}

func (x PlayerApplyEnterMpResultRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerApplyEnterMpResultRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerApplyEnterMpResultRsp_proto_enumTypes[0].Descriptor()
}

func (PlayerApplyEnterMpResultRsp_CmdId) Type() protoreflect.EnumType {
	return &file_PlayerApplyEnterMpResultRsp_proto_enumTypes[0]
}

func (x PlayerApplyEnterMpResultRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerApplyEnterMpResultRsp_CmdId.Descriptor instead.
func (PlayerApplyEnterMpResultRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PlayerApplyEnterMpResultRsp_proto_rawDescGZIP(), []int{0, 0}
}

type PlayerApplyEnterMpResultRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode  int32  `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ApplyUid uint32 `protobuf:"varint,2,opt,name=apply_uid,json=applyUid,proto3" json:"apply_uid,omitempty"`
	IsAgreed bool   `protobuf:"varint,3,opt,name=is_agreed,json=isAgreed,proto3" json:"is_agreed,omitempty"`
}

func (x *PlayerApplyEnterMpResultRsp) Reset() {
	*x = PlayerApplyEnterMpResultRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerApplyEnterMpResultRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerApplyEnterMpResultRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerApplyEnterMpResultRsp) ProtoMessage() {}

func (x *PlayerApplyEnterMpResultRsp) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerApplyEnterMpResultRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerApplyEnterMpResultRsp.ProtoReflect.Descriptor instead.
func (*PlayerApplyEnterMpResultRsp) Descriptor() ([]byte, []int) {
	return file_PlayerApplyEnterMpResultRsp_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerApplyEnterMpResultRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *PlayerApplyEnterMpResultRsp) GetApplyUid() uint32 {
	if x != nil {
		return x.ApplyUid
	}
	return 0
}

func (x *PlayerApplyEnterMpResultRsp) GetIsAgreed() bool {
	if x != nil {
		return x.IsAgreed
	}
	return false
}

var File_PlayerApplyEnterMpResultRsp_proto protoreflect.FileDescriptor

var file_PlayerApplyEnterMpResultRsp_proto_rawDesc = []byte{
	0x0a, 0x21, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x4d, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x1b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x55, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x61, 0x67, 0x72, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x41, 0x67, 0x72, 0x65, 0x65, 0x64, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d,
	0x44, 0x5f, 0x49, 0x44, 0x10, 0x8e, 0x0e, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10,
	0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45,
	0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerApplyEnterMpResultRsp_proto_rawDescOnce sync.Once
	file_PlayerApplyEnterMpResultRsp_proto_rawDescData = file_PlayerApplyEnterMpResultRsp_proto_rawDesc
)

func file_PlayerApplyEnterMpResultRsp_proto_rawDescGZIP() []byte {
	file_PlayerApplyEnterMpResultRsp_proto_rawDescOnce.Do(func() {
		file_PlayerApplyEnterMpResultRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerApplyEnterMpResultRsp_proto_rawDescData)
	})
	return file_PlayerApplyEnterMpResultRsp_proto_rawDescData
}

var file_PlayerApplyEnterMpResultRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerApplyEnterMpResultRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerApplyEnterMpResultRsp_proto_goTypes = []interface{}{
	(PlayerApplyEnterMpResultRsp_CmdId)(0), // 0: PlayerApplyEnterMpResultRsp.CmdId
	(*PlayerApplyEnterMpResultRsp)(nil),    // 1: PlayerApplyEnterMpResultRsp
}
var file_PlayerApplyEnterMpResultRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PlayerApplyEnterMpResultRsp_proto_init() }
func file_PlayerApplyEnterMpResultRsp_proto_init() {
	if File_PlayerApplyEnterMpResultRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PlayerApplyEnterMpResultRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerApplyEnterMpResultRsp); i {
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
			RawDescriptor: file_PlayerApplyEnterMpResultRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerApplyEnterMpResultRsp_proto_goTypes,
		DependencyIndexes: file_PlayerApplyEnterMpResultRsp_proto_depIdxs,
		EnumInfos:         file_PlayerApplyEnterMpResultRsp_proto_enumTypes,
		MessageInfos:      file_PlayerApplyEnterMpResultRsp_proto_msgTypes,
	}.Build()
	File_PlayerApplyEnterMpResultRsp_proto = out.File
	file_PlayerApplyEnterMpResultRsp_proto_rawDesc = nil
	file_PlayerApplyEnterMpResultRsp_proto_goTypes = nil
	file_PlayerApplyEnterMpResultRsp_proto_depIdxs = nil
}
