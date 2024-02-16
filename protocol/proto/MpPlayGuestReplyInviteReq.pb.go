// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: MpPlayGuestReplyInviteReq.proto

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

type MpPlayGuestReplyInviteReq_CmdId int32

const (
	MpPlayGuestReplyInviteReq_NONE             MpPlayGuestReplyInviteReq_CmdId = 0
	MpPlayGuestReplyInviteReq_CMD_ID           MpPlayGuestReplyInviteReq_CmdId = 1820
	MpPlayGuestReplyInviteReq_ENET_CHANNEL_ID  MpPlayGuestReplyInviteReq_CmdId = 0
	MpPlayGuestReplyInviteReq_ENET_IS_RELIABLE MpPlayGuestReplyInviteReq_CmdId = 1
	MpPlayGuestReplyInviteReq_IS_ALLOW_CLIENT  MpPlayGuestReplyInviteReq_CmdId = 1
)

// Enum value maps for MpPlayGuestReplyInviteReq_CmdId.
var (
	MpPlayGuestReplyInviteReq_CmdId_name = map[int32]string{
		0:    "NONE",
		1820: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	MpPlayGuestReplyInviteReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           1820,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x MpPlayGuestReplyInviteReq_CmdId) Enum() *MpPlayGuestReplyInviteReq_CmdId {
	p := new(MpPlayGuestReplyInviteReq_CmdId)
	*p = x
	return p
}

func (x MpPlayGuestReplyInviteReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MpPlayGuestReplyInviteReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_MpPlayGuestReplyInviteReq_proto_enumTypes[0].Descriptor()
}

func (MpPlayGuestReplyInviteReq_CmdId) Type() protoreflect.EnumType {
	return &file_MpPlayGuestReplyInviteReq_proto_enumTypes[0]
}

func (x MpPlayGuestReplyInviteReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MpPlayGuestReplyInviteReq_CmdId.Descriptor instead.
func (MpPlayGuestReplyInviteReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_MpPlayGuestReplyInviteReq_proto_rawDescGZIP(), []int{0, 0}
}

type MpPlayGuestReplyInviteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MpPlayId uint32 `protobuf:"varint,1,opt,name=mp_play_id,json=mpPlayId,proto3" json:"mp_play_id,omitempty"`
	IsAgree  bool   `protobuf:"varint,2,opt,name=is_agree,json=isAgree,proto3" json:"is_agree,omitempty"`
}

func (x *MpPlayGuestReplyInviteReq) Reset() {
	*x = MpPlayGuestReplyInviteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MpPlayGuestReplyInviteReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpPlayGuestReplyInviteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpPlayGuestReplyInviteReq) ProtoMessage() {}

func (x *MpPlayGuestReplyInviteReq) ProtoReflect() protoreflect.Message {
	mi := &file_MpPlayGuestReplyInviteReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpPlayGuestReplyInviteReq.ProtoReflect.Descriptor instead.
func (*MpPlayGuestReplyInviteReq) Descriptor() ([]byte, []int) {
	return file_MpPlayGuestReplyInviteReq_proto_rawDescGZIP(), []int{0}
}

func (x *MpPlayGuestReplyInviteReq) GetMpPlayId() uint32 {
	if x != nil {
		return x.MpPlayId
	}
	return 0
}

func (x *MpPlayGuestReplyInviteReq) GetIsAgree() bool {
	if x != nil {
		return x.IsAgree
	}
	return false
}

var File_MpPlayGuestReplyInviteReq_proto protoreflect.FileDescriptor

var file_MpPlayGuestReplyInviteReq_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x4d, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x47, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x19, 0x4d, 0x70, 0x50,
	0x6c, 0x61, 0x79, 0x47, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x0a, 0x6d, 0x70, 0x5f, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x70, 0x50, 0x6c,
	0x61, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x67, 0x72, 0x65, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x67, 0x72, 0x65, 0x65, 0x22,
	0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x9c, 0x0e, 0x12,
	0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f,
	0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53,
	0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a,
	0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MpPlayGuestReplyInviteReq_proto_rawDescOnce sync.Once
	file_MpPlayGuestReplyInviteReq_proto_rawDescData = file_MpPlayGuestReplyInviteReq_proto_rawDesc
)

func file_MpPlayGuestReplyInviteReq_proto_rawDescGZIP() []byte {
	file_MpPlayGuestReplyInviteReq_proto_rawDescOnce.Do(func() {
		file_MpPlayGuestReplyInviteReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_MpPlayGuestReplyInviteReq_proto_rawDescData)
	})
	return file_MpPlayGuestReplyInviteReq_proto_rawDescData
}

var file_MpPlayGuestReplyInviteReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MpPlayGuestReplyInviteReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MpPlayGuestReplyInviteReq_proto_goTypes = []interface{}{
	(MpPlayGuestReplyInviteReq_CmdId)(0), // 0: proto.MpPlayGuestReplyInviteReq.CmdId
	(*MpPlayGuestReplyInviteReq)(nil),    // 1: proto.MpPlayGuestReplyInviteReq
}
var file_MpPlayGuestReplyInviteReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MpPlayGuestReplyInviteReq_proto_init() }
func file_MpPlayGuestReplyInviteReq_proto_init() {
	if File_MpPlayGuestReplyInviteReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MpPlayGuestReplyInviteReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpPlayGuestReplyInviteReq); i {
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
			RawDescriptor: file_MpPlayGuestReplyInviteReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MpPlayGuestReplyInviteReq_proto_goTypes,
		DependencyIndexes: file_MpPlayGuestReplyInviteReq_proto_depIdxs,
		EnumInfos:         file_MpPlayGuestReplyInviteReq_proto_enumTypes,
		MessageInfos:      file_MpPlayGuestReplyInviteReq_proto_msgTypes,
	}.Build()
	File_MpPlayGuestReplyInviteReq_proto = out.File
	file_MpPlayGuestReplyInviteReq_proto_rawDesc = nil
	file_MpPlayGuestReplyInviteReq_proto_goTypes = nil
	file_MpPlayGuestReplyInviteReq_proto_depIdxs = nil
}
