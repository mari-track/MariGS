// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GetRecentMpPlayerListRsp.proto

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

type GetRecentMpPlayerListRsp_CmdId int32

const (
	GetRecentMpPlayerListRsp_NONE             GetRecentMpPlayerListRsp_CmdId = 0
	GetRecentMpPlayerListRsp_CMD_ID           GetRecentMpPlayerListRsp_CmdId = 4033
	GetRecentMpPlayerListRsp_ENET_CHANNEL_ID  GetRecentMpPlayerListRsp_CmdId = 0
	GetRecentMpPlayerListRsp_ENET_IS_RELIABLE GetRecentMpPlayerListRsp_CmdId = 1
	GetRecentMpPlayerListRsp_IS_ALLOW_CLIENT  GetRecentMpPlayerListRsp_CmdId = 1
)

// Enum value maps for GetRecentMpPlayerListRsp_CmdId.
var (
	GetRecentMpPlayerListRsp_CmdId_name = map[int32]string{
		0:    "NONE",
		4033: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	GetRecentMpPlayerListRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           4033,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x GetRecentMpPlayerListRsp_CmdId) Enum() *GetRecentMpPlayerListRsp_CmdId {
	p := new(GetRecentMpPlayerListRsp_CmdId)
	*p = x
	return p
}

func (x GetRecentMpPlayerListRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetRecentMpPlayerListRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GetRecentMpPlayerListRsp_proto_enumTypes[0].Descriptor()
}

func (GetRecentMpPlayerListRsp_CmdId) Type() protoreflect.EnumType {
	return &file_GetRecentMpPlayerListRsp_proto_enumTypes[0]
}

func (x GetRecentMpPlayerListRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetRecentMpPlayerListRsp_CmdId.Descriptor instead.
func (GetRecentMpPlayerListRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GetRecentMpPlayerListRsp_proto_rawDescGZIP(), []int{0, 0}
}

type GetRecentMpPlayerListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode                 int32          `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	RecentMpPlayerBriefList []*FriendBrief `protobuf:"bytes,2,rep,name=recent_mp_player_brief_list,json=recentMpPlayerBriefList,proto3" json:"recent_mp_player_brief_list,omitempty"`
}

func (x *GetRecentMpPlayerListRsp) Reset() {
	*x = GetRecentMpPlayerListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetRecentMpPlayerListRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecentMpPlayerListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecentMpPlayerListRsp) ProtoMessage() {}

func (x *GetRecentMpPlayerListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetRecentMpPlayerListRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecentMpPlayerListRsp.ProtoReflect.Descriptor instead.
func (*GetRecentMpPlayerListRsp) Descriptor() ([]byte, []int) {
	return file_GetRecentMpPlayerListRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetRecentMpPlayerListRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetRecentMpPlayerListRsp) GetRecentMpPlayerBriefList() []*FriendBrief {
	if x != nil {
		return x.RecentMpPlayerBriefList
	}
	return nil
}

var File_GetRecentMpPlayerListRsp_proto protoreflect.FileDescriptor

var file_GetRecentMpPlayerListRsp_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x4d, 0x70, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x42,
	0x72, 0x69, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x01, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x4d, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x50, 0x0a, 0x1b, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x70, 0x5f, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x42, 0x72, 0x69, 0x65, 0x66, 0x52, 0x17, 0x72, 0x65, 0x63, 0x65,
	0x6e, 0x74, 0x4d, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x72, 0x69, 0x65, 0x66, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44,
	0x10, 0xc1, 0x1f, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e,
	0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetRecentMpPlayerListRsp_proto_rawDescOnce sync.Once
	file_GetRecentMpPlayerListRsp_proto_rawDescData = file_GetRecentMpPlayerListRsp_proto_rawDesc
)

func file_GetRecentMpPlayerListRsp_proto_rawDescGZIP() []byte {
	file_GetRecentMpPlayerListRsp_proto_rawDescOnce.Do(func() {
		file_GetRecentMpPlayerListRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetRecentMpPlayerListRsp_proto_rawDescData)
	})
	return file_GetRecentMpPlayerListRsp_proto_rawDescData
}

var file_GetRecentMpPlayerListRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GetRecentMpPlayerListRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetRecentMpPlayerListRsp_proto_goTypes = []interface{}{
	(GetRecentMpPlayerListRsp_CmdId)(0), // 0: proto.GetRecentMpPlayerListRsp.CmdId
	(*GetRecentMpPlayerListRsp)(nil),    // 1: proto.GetRecentMpPlayerListRsp
	(*FriendBrief)(nil),                 // 2: proto.FriendBrief
}
var file_GetRecentMpPlayerListRsp_proto_depIdxs = []int32{
	2, // 0: proto.GetRecentMpPlayerListRsp.recent_mp_player_brief_list:type_name -> proto.FriendBrief
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetRecentMpPlayerListRsp_proto_init() }
func file_GetRecentMpPlayerListRsp_proto_init() {
	if File_GetRecentMpPlayerListRsp_proto != nil {
		return
	}
	file_FriendBrief_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetRecentMpPlayerListRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecentMpPlayerListRsp); i {
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
			RawDescriptor: file_GetRecentMpPlayerListRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetRecentMpPlayerListRsp_proto_goTypes,
		DependencyIndexes: file_GetRecentMpPlayerListRsp_proto_depIdxs,
		EnumInfos:         file_GetRecentMpPlayerListRsp_proto_enumTypes,
		MessageInfos:      file_GetRecentMpPlayerListRsp_proto_msgTypes,
	}.Build()
	File_GetRecentMpPlayerListRsp_proto = out.File
	file_GetRecentMpPlayerListRsp_proto_rawDesc = nil
	file_GetRecentMpPlayerListRsp_proto_goTypes = nil
	file_GetRecentMpPlayerListRsp_proto_depIdxs = nil
}
