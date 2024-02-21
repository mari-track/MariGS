// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: EntityAiSyncNotify.proto

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

type EntityAiSyncNotify_CmdId int32

const (
	EntityAiSyncNotify_NONE             EntityAiSyncNotify_CmdId = 0
	EntityAiSyncNotify_CMD_ID           EntityAiSyncNotify_CmdId = 354
	EntityAiSyncNotify_ENET_CHANNEL_ID  EntityAiSyncNotify_CmdId = 0
	EntityAiSyncNotify_ENET_IS_RELIABLE EntityAiSyncNotify_CmdId = 1
	EntityAiSyncNotify_IS_ALLOW_CLIENT  EntityAiSyncNotify_CmdId = 1
)

// Enum value maps for EntityAiSyncNotify_CmdId.
var (
	EntityAiSyncNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		354: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	EntityAiSyncNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           354,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x EntityAiSyncNotify_CmdId) Enum() *EntityAiSyncNotify_CmdId {
	p := new(EntityAiSyncNotify_CmdId)
	*p = x
	return p
}

func (x EntityAiSyncNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntityAiSyncNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_EntityAiSyncNotify_proto_enumTypes[0].Descriptor()
}

func (EntityAiSyncNotify_CmdId) Type() protoreflect.EnumType {
	return &file_EntityAiSyncNotify_proto_enumTypes[0]
}

func (x EntityAiSyncNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntityAiSyncNotify_CmdId.Descriptor instead.
func (EntityAiSyncNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_EntityAiSyncNotify_proto_rawDescGZIP(), []int{0, 0}
}

type EntityAiSyncNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InfoList                      []*AiSyncInfo `protobuf:"bytes,1,rep,name=info_list,json=infoList,proto3" json:"info_list,omitempty"`
	LocalAvatarAlertedMonsterList []uint32      `protobuf:"varint,2,rep,packed,name=local_avatar_alerted_monster_list,json=localAvatarAlertedMonsterList,proto3" json:"local_avatar_alerted_monster_list,omitempty"`
}

func (x *EntityAiSyncNotify) Reset() {
	*x = EntityAiSyncNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EntityAiSyncNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityAiSyncNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityAiSyncNotify) ProtoMessage() {}

func (x *EntityAiSyncNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EntityAiSyncNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityAiSyncNotify.ProtoReflect.Descriptor instead.
func (*EntityAiSyncNotify) Descriptor() ([]byte, []int) {
	return file_EntityAiSyncNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EntityAiSyncNotify) GetInfoList() []*AiSyncInfo {
	if x != nil {
		return x.InfoList
	}
	return nil
}

func (x *EntityAiSyncNotify) GetLocalAvatarAlertedMonsterList() []uint32 {
	if x != nil {
		return x.LocalAvatarAlertedMonsterList
	}
	return nil
}

var File_EntityAiSyncNotify_proto protoreflect.FileDescriptor

var file_EntityAiSyncNotify_proto_rawDesc = []byte{
	0x0a, 0x18, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x69, 0x53, 0x79, 0x6e, 0x63, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x41, 0x69, 0x53, 0x79,
	0x6e, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a,
	0x12, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x69, 0x53, 0x79, 0x6e, 0x63, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x28, 0x0a, 0x09, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x69, 0x53, 0x79, 0x6e, 0x63, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x48, 0x0a,
	0x21, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x1d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x6e, 0x73,
	0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d,
	0x44, 0x5f, 0x49, 0x44, 0x10, 0xe2, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10,
	0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45,
	0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43,
	0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EntityAiSyncNotify_proto_rawDescOnce sync.Once
	file_EntityAiSyncNotify_proto_rawDescData = file_EntityAiSyncNotify_proto_rawDesc
)

func file_EntityAiSyncNotify_proto_rawDescGZIP() []byte {
	file_EntityAiSyncNotify_proto_rawDescOnce.Do(func() {
		file_EntityAiSyncNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EntityAiSyncNotify_proto_rawDescData)
	})
	return file_EntityAiSyncNotify_proto_rawDescData
}

var file_EntityAiSyncNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_EntityAiSyncNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EntityAiSyncNotify_proto_goTypes = []interface{}{
	(EntityAiSyncNotify_CmdId)(0), // 0: EntityAiSyncNotify.CmdId
	(*EntityAiSyncNotify)(nil),    // 1: EntityAiSyncNotify
	(*AiSyncInfo)(nil),            // 2: AiSyncInfo
}
var file_EntityAiSyncNotify_proto_depIdxs = []int32{
	2, // 0: EntityAiSyncNotify.info_list:type_name -> AiSyncInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_EntityAiSyncNotify_proto_init() }
func file_EntityAiSyncNotify_proto_init() {
	if File_EntityAiSyncNotify_proto != nil {
		return
	}
	file_AiSyncInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EntityAiSyncNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityAiSyncNotify); i {
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
			RawDescriptor: file_EntityAiSyncNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EntityAiSyncNotify_proto_goTypes,
		DependencyIndexes: file_EntityAiSyncNotify_proto_depIdxs,
		EnumInfos:         file_EntityAiSyncNotify_proto_enumTypes,
		MessageInfos:      file_EntityAiSyncNotify_proto_msgTypes,
	}.Build()
	File_EntityAiSyncNotify_proto = out.File
	file_EntityAiSyncNotify_proto_rawDesc = nil
	file_EntityAiSyncNotify_proto_goTypes = nil
	file_EntityAiSyncNotify_proto_depIdxs = nil
}
