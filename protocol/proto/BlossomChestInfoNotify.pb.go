// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: BlossomChestInfoNotify.proto

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

type BlossomChestInfoNotify_CmdId int32

const (
	BlossomChestInfoNotify_NONE             BlossomChestInfoNotify_CmdId = 0
	BlossomChestInfoNotify_CMD_ID           BlossomChestInfoNotify_CmdId = 808
	BlossomChestInfoNotify_ENET_CHANNEL_ID  BlossomChestInfoNotify_CmdId = 0
	BlossomChestInfoNotify_ENET_IS_RELIABLE BlossomChestInfoNotify_CmdId = 1
)

// Enum value maps for BlossomChestInfoNotify_CmdId.
var (
	BlossomChestInfoNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		808: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	BlossomChestInfoNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           808,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x BlossomChestInfoNotify_CmdId) Enum() *BlossomChestInfoNotify_CmdId {
	p := new(BlossomChestInfoNotify_CmdId)
	*p = x
	return p
}

func (x BlossomChestInfoNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlossomChestInfoNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_BlossomChestInfoNotify_proto_enumTypes[0].Descriptor()
}

func (BlossomChestInfoNotify_CmdId) Type() protoreflect.EnumType {
	return &file_BlossomChestInfoNotify_proto_enumTypes[0]
}

func (x BlossomChestInfoNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlossomChestInfoNotify_CmdId.Descriptor instead.
func (BlossomChestInfoNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_BlossomChestInfoNotify_proto_rawDescGZIP(), []int{0, 0}
}

type BlossomChestInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId         uint32            `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	BlossomChestInfo *BlossomChestInfo `protobuf:"bytes,2,opt,name=blossom_chest_info,json=blossomChestInfo,proto3" json:"blossom_chest_info,omitempty"`
}

func (x *BlossomChestInfoNotify) Reset() {
	*x = BlossomChestInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BlossomChestInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlossomChestInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlossomChestInfoNotify) ProtoMessage() {}

func (x *BlossomChestInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_BlossomChestInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlossomChestInfoNotify.ProtoReflect.Descriptor instead.
func (*BlossomChestInfoNotify) Descriptor() ([]byte, []int) {
	return file_BlossomChestInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *BlossomChestInfoNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *BlossomChestInfoNotify) GetBlossomChestInfo() *BlossomChestInfo {
	if x != nil {
		return x.BlossomChestInfo
	}
	return nil
}

var File_BlossomChestInfoNotify_proto protoreflect.FileDescriptor

var file_BlossomChestInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x42, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x43, 0x68, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x42, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x43, 0x68,
	0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x01,
	0x0a, 0x16, 0x42, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x43, 0x68, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x12, 0x62, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d,
	0x5f, 0x63, 0x68, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x73, 0x73, 0x6f,
	0x6d, 0x43, 0x68, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x62, 0x6c, 0x6f, 0x73,
	0x73, 0x6f, 0x6d, 0x43, 0x68, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4d, 0x0a, 0x05,
	0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xa8, 0x06, 0x12, 0x13, 0x0a, 0x0f,
	0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c,
	0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BlossomChestInfoNotify_proto_rawDescOnce sync.Once
	file_BlossomChestInfoNotify_proto_rawDescData = file_BlossomChestInfoNotify_proto_rawDesc
)

func file_BlossomChestInfoNotify_proto_rawDescGZIP() []byte {
	file_BlossomChestInfoNotify_proto_rawDescOnce.Do(func() {
		file_BlossomChestInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_BlossomChestInfoNotify_proto_rawDescData)
	})
	return file_BlossomChestInfoNotify_proto_rawDescData
}

var file_BlossomChestInfoNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_BlossomChestInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BlossomChestInfoNotify_proto_goTypes = []interface{}{
	(BlossomChestInfoNotify_CmdId)(0), // 0: proto.BlossomChestInfoNotify.CmdId
	(*BlossomChestInfoNotify)(nil),    // 1: proto.BlossomChestInfoNotify
	(*BlossomChestInfo)(nil),          // 2: proto.BlossomChestInfo
}
var file_BlossomChestInfoNotify_proto_depIdxs = []int32{
	2, // 0: proto.BlossomChestInfoNotify.blossom_chest_info:type_name -> proto.BlossomChestInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_BlossomChestInfoNotify_proto_init() }
func file_BlossomChestInfoNotify_proto_init() {
	if File_BlossomChestInfoNotify_proto != nil {
		return
	}
	file_BlossomChestInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BlossomChestInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlossomChestInfoNotify); i {
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
			RawDescriptor: file_BlossomChestInfoNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BlossomChestInfoNotify_proto_goTypes,
		DependencyIndexes: file_BlossomChestInfoNotify_proto_depIdxs,
		EnumInfos:         file_BlossomChestInfoNotify_proto_enumTypes,
		MessageInfos:      file_BlossomChestInfoNotify_proto_msgTypes,
	}.Build()
	File_BlossomChestInfoNotify_proto = out.File
	file_BlossomChestInfoNotify_proto_rawDesc = nil
	file_BlossomChestInfoNotify_proto_goTypes = nil
	file_BlossomChestInfoNotify_proto_depIdxs = nil
}