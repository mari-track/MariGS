// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: WorldOwnerBlossomBriefInfoNotify.proto

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

type WorldOwnerBlossomBriefInfoNotify_CmdId int32

const (
	WorldOwnerBlossomBriefInfoNotify_NONE             WorldOwnerBlossomBriefInfoNotify_CmdId = 0
	WorldOwnerBlossomBriefInfoNotify_CMD_ID           WorldOwnerBlossomBriefInfoNotify_CmdId = 2704
	WorldOwnerBlossomBriefInfoNotify_ENET_CHANNEL_ID  WorldOwnerBlossomBriefInfoNotify_CmdId = 0
	WorldOwnerBlossomBriefInfoNotify_ENET_IS_RELIABLE WorldOwnerBlossomBriefInfoNotify_CmdId = 1
	WorldOwnerBlossomBriefInfoNotify_IS_ALLOW_CLIENT  WorldOwnerBlossomBriefInfoNotify_CmdId = 1
)

// Enum value maps for WorldOwnerBlossomBriefInfoNotify_CmdId.
var (
	WorldOwnerBlossomBriefInfoNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		2704: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	WorldOwnerBlossomBriefInfoNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2704,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x WorldOwnerBlossomBriefInfoNotify_CmdId) Enum() *WorldOwnerBlossomBriefInfoNotify_CmdId {
	p := new(WorldOwnerBlossomBriefInfoNotify_CmdId)
	*p = x
	return p
}

func (x WorldOwnerBlossomBriefInfoNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorldOwnerBlossomBriefInfoNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_WorldOwnerBlossomBriefInfoNotify_proto_enumTypes[0].Descriptor()
}

func (WorldOwnerBlossomBriefInfoNotify_CmdId) Type() protoreflect.EnumType {
	return &file_WorldOwnerBlossomBriefInfoNotify_proto_enumTypes[0]
}

func (x WorldOwnerBlossomBriefInfoNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorldOwnerBlossomBriefInfoNotify_CmdId.Descriptor instead.
func (WorldOwnerBlossomBriefInfoNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_WorldOwnerBlossomBriefInfoNotify_proto_rawDescGZIP(), []int{0, 0}
}

type WorldOwnerBlossomBriefInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BriefInfoList []*BlossomBriefInfo `protobuf:"bytes,1,rep,name=brief_info_list,json=briefInfoList,proto3" json:"brief_info_list,omitempty"`
}

func (x *WorldOwnerBlossomBriefInfoNotify) Reset() {
	*x = WorldOwnerBlossomBriefInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WorldOwnerBlossomBriefInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldOwnerBlossomBriefInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldOwnerBlossomBriefInfoNotify) ProtoMessage() {}

func (x *WorldOwnerBlossomBriefInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_WorldOwnerBlossomBriefInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldOwnerBlossomBriefInfoNotify.ProtoReflect.Descriptor instead.
func (*WorldOwnerBlossomBriefInfoNotify) Descriptor() ([]byte, []int) {
	return file_WorldOwnerBlossomBriefInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *WorldOwnerBlossomBriefInfoNotify) GetBriefInfoList() []*BlossomBriefInfo {
	if x != nil {
		return x.BriefInfoList
	}
	return nil
}

var File_WorldOwnerBlossomBriefInfoNotify_proto protoreflect.FileDescriptor

var file_WorldOwnerBlossomBriefInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x26, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x73,
	0x73, 0x6f, 0x6d, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x42, 0x6c, 0x6f, 0x73, 0x73, 0x6f,
	0x6d, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc1, 0x01, 0x0a, 0x20, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x42,
	0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x39, 0x0a, 0x0f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x42, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0d, 0x62, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x90, 0x15,
	0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c,
	0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53,
	0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49,
	0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01,
	0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WorldOwnerBlossomBriefInfoNotify_proto_rawDescOnce sync.Once
	file_WorldOwnerBlossomBriefInfoNotify_proto_rawDescData = file_WorldOwnerBlossomBriefInfoNotify_proto_rawDesc
)

func file_WorldOwnerBlossomBriefInfoNotify_proto_rawDescGZIP() []byte {
	file_WorldOwnerBlossomBriefInfoNotify_proto_rawDescOnce.Do(func() {
		file_WorldOwnerBlossomBriefInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_WorldOwnerBlossomBriefInfoNotify_proto_rawDescData)
	})
	return file_WorldOwnerBlossomBriefInfoNotify_proto_rawDescData
}

var file_WorldOwnerBlossomBriefInfoNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_WorldOwnerBlossomBriefInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WorldOwnerBlossomBriefInfoNotify_proto_goTypes = []interface{}{
	(WorldOwnerBlossomBriefInfoNotify_CmdId)(0), // 0: WorldOwnerBlossomBriefInfoNotify.CmdId
	(*WorldOwnerBlossomBriefInfoNotify)(nil),    // 1: WorldOwnerBlossomBriefInfoNotify
	(*BlossomBriefInfo)(nil),                    // 2: BlossomBriefInfo
}
var file_WorldOwnerBlossomBriefInfoNotify_proto_depIdxs = []int32{
	2, // 0: WorldOwnerBlossomBriefInfoNotify.brief_info_list:type_name -> BlossomBriefInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_WorldOwnerBlossomBriefInfoNotify_proto_init() }
func file_WorldOwnerBlossomBriefInfoNotify_proto_init() {
	if File_WorldOwnerBlossomBriefInfoNotify_proto != nil {
		return
	}
	file_BlossomBriefInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WorldOwnerBlossomBriefInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldOwnerBlossomBriefInfoNotify); i {
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
			RawDescriptor: file_WorldOwnerBlossomBriefInfoNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WorldOwnerBlossomBriefInfoNotify_proto_goTypes,
		DependencyIndexes: file_WorldOwnerBlossomBriefInfoNotify_proto_depIdxs,
		EnumInfos:         file_WorldOwnerBlossomBriefInfoNotify_proto_enumTypes,
		MessageInfos:      file_WorldOwnerBlossomBriefInfoNotify_proto_msgTypes,
	}.Build()
	File_WorldOwnerBlossomBriefInfoNotify_proto = out.File
	file_WorldOwnerBlossomBriefInfoNotify_proto_rawDesc = nil
	file_WorldOwnerBlossomBriefInfoNotify_proto_goTypes = nil
	file_WorldOwnerBlossomBriefInfoNotify_proto_depIdxs = nil
}
