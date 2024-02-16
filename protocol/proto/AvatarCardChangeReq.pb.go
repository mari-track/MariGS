// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: AvatarCardChangeReq.proto

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

type AvatarCardChangeReq_CmdId int32

const (
	AvatarCardChangeReq_NONE             AvatarCardChangeReq_CmdId = 0
	AvatarCardChangeReq_CMD_ID           AvatarCardChangeReq_CmdId = 627
	AvatarCardChangeReq_ENET_CHANNEL_ID  AvatarCardChangeReq_CmdId = 0
	AvatarCardChangeReq_ENET_IS_RELIABLE AvatarCardChangeReq_CmdId = 1
	AvatarCardChangeReq_IS_ALLOW_CLIENT  AvatarCardChangeReq_CmdId = 1
)

// Enum value maps for AvatarCardChangeReq_CmdId.
var (
	AvatarCardChangeReq_CmdId_name = map[int32]string{
		0:   "NONE",
		627: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	AvatarCardChangeReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           627,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x AvatarCardChangeReq_CmdId) Enum() *AvatarCardChangeReq_CmdId {
	p := new(AvatarCardChangeReq_CmdId)
	*p = x
	return p
}

func (x AvatarCardChangeReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AvatarCardChangeReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_AvatarCardChangeReq_proto_enumTypes[0].Descriptor()
}

func (AvatarCardChangeReq_CmdId) Type() protoreflect.EnumType {
	return &file_AvatarCardChangeReq_proto_enumTypes[0]
}

func (x AvatarCardChangeReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AvatarCardChangeReq_CmdId.Descriptor instead.
func (AvatarCardChangeReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_AvatarCardChangeReq_proto_rawDescGZIP(), []int{0, 0}
}

type AvatarCardChangeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarGuid uint64 `protobuf:"varint,1,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
	ItemId     uint32 `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Count      uint32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *AvatarCardChangeReq) Reset() {
	*x = AvatarCardChangeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarCardChangeReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarCardChangeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarCardChangeReq) ProtoMessage() {}

func (x *AvatarCardChangeReq) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarCardChangeReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarCardChangeReq.ProtoReflect.Descriptor instead.
func (*AvatarCardChangeReq) Descriptor() ([]byte, []int) {
	return file_AvatarCardChangeReq_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarCardChangeReq) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

func (x *AvatarCardChangeReq) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *AvatarCardChangeReq) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_AvatarCardChangeReq_proto protoreflect.FileDescriptor

var file_AvatarCardChangeReq_proto_rawDesc = []byte{
	0x0a, 0x19, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x43, 0x61, 0x72, 0x64, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x13, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x43, 0x61, 0x72,
	0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x47, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x69, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d,
	0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xf3, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41,
	0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f,
	0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_AvatarCardChangeReq_proto_rawDescOnce sync.Once
	file_AvatarCardChangeReq_proto_rawDescData = file_AvatarCardChangeReq_proto_rawDesc
)

func file_AvatarCardChangeReq_proto_rawDescGZIP() []byte {
	file_AvatarCardChangeReq_proto_rawDescOnce.Do(func() {
		file_AvatarCardChangeReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarCardChangeReq_proto_rawDescData)
	})
	return file_AvatarCardChangeReq_proto_rawDescData
}

var file_AvatarCardChangeReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_AvatarCardChangeReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarCardChangeReq_proto_goTypes = []interface{}{
	(AvatarCardChangeReq_CmdId)(0), // 0: proto.AvatarCardChangeReq.CmdId
	(*AvatarCardChangeReq)(nil),    // 1: proto.AvatarCardChangeReq
}
var file_AvatarCardChangeReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AvatarCardChangeReq_proto_init() }
func file_AvatarCardChangeReq_proto_init() {
	if File_AvatarCardChangeReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AvatarCardChangeReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarCardChangeReq); i {
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
			RawDescriptor: file_AvatarCardChangeReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarCardChangeReq_proto_goTypes,
		DependencyIndexes: file_AvatarCardChangeReq_proto_depIdxs,
		EnumInfos:         file_AvatarCardChangeReq_proto_enumTypes,
		MessageInfos:      file_AvatarCardChangeReq_proto_msgTypes,
	}.Build()
	File_AvatarCardChangeReq_proto = out.File
	file_AvatarCardChangeReq_proto_rawDesc = nil
	file_AvatarCardChangeReq_proto_goTypes = nil
	file_AvatarCardChangeReq_proto_depIdxs = nil
}