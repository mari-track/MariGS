// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: PlayerChatReq.proto

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

type PlayerChatReq_CmdId int32

const (
	PlayerChatReq_NONE             PlayerChatReq_CmdId = 0
	PlayerChatReq_CMD_ID           PlayerChatReq_CmdId = 3018
	PlayerChatReq_ENET_CHANNEL_ID  PlayerChatReq_CmdId = 0
	PlayerChatReq_ENET_IS_RELIABLE PlayerChatReq_CmdId = 1
	PlayerChatReq_IS_ALLOW_CLIENT  PlayerChatReq_CmdId = 1
)

// Enum value maps for PlayerChatReq_CmdId.
var (
	PlayerChatReq_CmdId_name = map[int32]string{
		0:    "NONE",
		3018: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	PlayerChatReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           3018,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x PlayerChatReq_CmdId) Enum() *PlayerChatReq_CmdId {
	p := new(PlayerChatReq_CmdId)
	*p = x
	return p
}

func (x PlayerChatReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerChatReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerChatReq_proto_enumTypes[0].Descriptor()
}

func (PlayerChatReq_CmdId) Type() protoreflect.EnumType {
	return &file_PlayerChatReq_proto_enumTypes[0]
}

func (x PlayerChatReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerChatReq_CmdId.Descriptor instead.
func (PlayerChatReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PlayerChatReq_proto_rawDescGZIP(), []int{0, 0}
}

type PlayerChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId uint32    `protobuf:"varint,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	ChatInfo  *ChatInfo `protobuf:"bytes,2,opt,name=chat_info,json=chatInfo,proto3" json:"chat_info,omitempty"`
}

func (x *PlayerChatReq) Reset() {
	*x = PlayerChatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerChatReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerChatReq) ProtoMessage() {}

func (x *PlayerChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerChatReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerChatReq.ProtoReflect.Descriptor instead.
func (*PlayerChatReq) Descriptor() ([]byte, []int) {
	return file_PlayerChatReq_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerChatReq) GetChannelId() uint32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *PlayerChatReq) GetChatInfo() *ChatInfo {
	if x != nil {
		return x.ChatInfo
	}
	return nil
}

var File_PlayerChatReq_proto protoreflect.FileDescriptor

var file_PlayerChatReq_proto_rawDesc = []byte{
	0x0a, 0x13, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x43, 0x68,
	0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a,
	0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x62, 0x0a, 0x05, 0x43,
	0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xca, 0x17, 0x12, 0x13, 0x0a, 0x0f, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c,
	0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_PlayerChatReq_proto_rawDescOnce sync.Once
	file_PlayerChatReq_proto_rawDescData = file_PlayerChatReq_proto_rawDesc
)

func file_PlayerChatReq_proto_rawDescGZIP() []byte {
	file_PlayerChatReq_proto_rawDescOnce.Do(func() {
		file_PlayerChatReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerChatReq_proto_rawDescData)
	})
	return file_PlayerChatReq_proto_rawDescData
}

var file_PlayerChatReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerChatReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerChatReq_proto_goTypes = []interface{}{
	(PlayerChatReq_CmdId)(0), // 0: proto.PlayerChatReq.CmdId
	(*PlayerChatReq)(nil),    // 1: proto.PlayerChatReq
	(*ChatInfo)(nil),         // 2: proto.ChatInfo
}
var file_PlayerChatReq_proto_depIdxs = []int32{
	2, // 0: proto.PlayerChatReq.chat_info:type_name -> proto.ChatInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerChatReq_proto_init() }
func file_PlayerChatReq_proto_init() {
	if File_PlayerChatReq_proto != nil {
		return
	}
	file_ChatInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerChatReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerChatReq); i {
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
			RawDescriptor: file_PlayerChatReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerChatReq_proto_goTypes,
		DependencyIndexes: file_PlayerChatReq_proto_depIdxs,
		EnumInfos:         file_PlayerChatReq_proto_enumTypes,
		MessageInfos:      file_PlayerChatReq_proto_msgTypes,
	}.Build()
	File_PlayerChatReq_proto = out.File
	file_PlayerChatReq_proto_rawDesc = nil
	file_PlayerChatReq_proto_goTypes = nil
	file_PlayerChatReq_proto_depIdxs = nil
}
