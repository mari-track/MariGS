// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GetSceneNpcPositionRsp.proto

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

type GetSceneNpcPositionRsp_CmdId int32

const (
	GetSceneNpcPositionRsp_NONE             GetSceneNpcPositionRsp_CmdId = 0
	GetSceneNpcPositionRsp_CMD_ID           GetSceneNpcPositionRsp_CmdId = 505
	GetSceneNpcPositionRsp_ENET_CHANNEL_ID  GetSceneNpcPositionRsp_CmdId = 0
	GetSceneNpcPositionRsp_ENET_IS_RELIABLE GetSceneNpcPositionRsp_CmdId = 1
)

// Enum value maps for GetSceneNpcPositionRsp_CmdId.
var (
	GetSceneNpcPositionRsp_CmdId_name = map[int32]string{
		0:   "NONE",
		505: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	GetSceneNpcPositionRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           505,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x GetSceneNpcPositionRsp_CmdId) Enum() *GetSceneNpcPositionRsp_CmdId {
	p := new(GetSceneNpcPositionRsp_CmdId)
	*p = x
	return p
}

func (x GetSceneNpcPositionRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetSceneNpcPositionRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GetSceneNpcPositionRsp_proto_enumTypes[0].Descriptor()
}

func (GetSceneNpcPositionRsp_CmdId) Type() protoreflect.EnumType {
	return &file_GetSceneNpcPositionRsp_proto_enumTypes[0]
}

func (x GetSceneNpcPositionRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetSceneNpcPositionRsp_CmdId.Descriptor instead.
func (GetSceneNpcPositionRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GetSceneNpcPositionRsp_proto_rawDescGZIP(), []int{0, 0}
}

type GetSceneNpcPositionRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode     int32              `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	SceneId     uint32             `protobuf:"varint,2,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	NpcInfoList []*NpcPositionInfo `protobuf:"bytes,3,rep,name=npc_info_list,json=npcInfoList,proto3" json:"npc_info_list,omitempty"`
}

func (x *GetSceneNpcPositionRsp) Reset() {
	*x = GetSceneNpcPositionRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetSceneNpcPositionRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSceneNpcPositionRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSceneNpcPositionRsp) ProtoMessage() {}

func (x *GetSceneNpcPositionRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetSceneNpcPositionRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSceneNpcPositionRsp.ProtoReflect.Descriptor instead.
func (*GetSceneNpcPositionRsp) Descriptor() ([]byte, []int) {
	return file_GetSceneNpcPositionRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetSceneNpcPositionRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetSceneNpcPositionRsp) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *GetSceneNpcPositionRsp) GetNpcInfoList() []*NpcPositionInfo {
	if x != nil {
		return x.NpcInfoList
	}
	return nil
}

var File_GetSceneNpcPositionRsp_proto protoreflect.FileDescriptor

var file_GetSceneNpcPositionRsp_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4e, 0x70, 0x63, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x4e, 0x70, 0x63, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x01, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4e, 0x70, 0x63, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0d,
	0x6e, 0x70, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x70, 0x63, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x6e, 0x70, 0x63,
	0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49,
	0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43,
	0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xf9, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetSceneNpcPositionRsp_proto_rawDescOnce sync.Once
	file_GetSceneNpcPositionRsp_proto_rawDescData = file_GetSceneNpcPositionRsp_proto_rawDesc
)

func file_GetSceneNpcPositionRsp_proto_rawDescGZIP() []byte {
	file_GetSceneNpcPositionRsp_proto_rawDescOnce.Do(func() {
		file_GetSceneNpcPositionRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetSceneNpcPositionRsp_proto_rawDescData)
	})
	return file_GetSceneNpcPositionRsp_proto_rawDescData
}

var file_GetSceneNpcPositionRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GetSceneNpcPositionRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetSceneNpcPositionRsp_proto_goTypes = []interface{}{
	(GetSceneNpcPositionRsp_CmdId)(0), // 0: proto.GetSceneNpcPositionRsp.CmdId
	(*GetSceneNpcPositionRsp)(nil),    // 1: proto.GetSceneNpcPositionRsp
	(*NpcPositionInfo)(nil),           // 2: proto.NpcPositionInfo
}
var file_GetSceneNpcPositionRsp_proto_depIdxs = []int32{
	2, // 0: proto.GetSceneNpcPositionRsp.npc_info_list:type_name -> proto.NpcPositionInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetSceneNpcPositionRsp_proto_init() }
func file_GetSceneNpcPositionRsp_proto_init() {
	if File_GetSceneNpcPositionRsp_proto != nil {
		return
	}
	file_NpcPositionInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetSceneNpcPositionRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSceneNpcPositionRsp); i {
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
			RawDescriptor: file_GetSceneNpcPositionRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetSceneNpcPositionRsp_proto_goTypes,
		DependencyIndexes: file_GetSceneNpcPositionRsp_proto_depIdxs,
		EnumInfos:         file_GetSceneNpcPositionRsp_proto_enumTypes,
		MessageInfos:      file_GetSceneNpcPositionRsp_proto_msgTypes,
	}.Build()
	File_GetSceneNpcPositionRsp_proto = out.File
	file_GetSceneNpcPositionRsp_proto_rawDesc = nil
	file_GetSceneNpcPositionRsp_proto_goTypes = nil
	file_GetSceneNpcPositionRsp_proto_depIdxs = nil
}
