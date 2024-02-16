// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: EntityAuthorityChangeNotify.proto

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

type EntityAuthorityChangeNotify_CmdId int32

const (
	EntityAuthorityChangeNotify_NONE             EntityAuthorityChangeNotify_CmdId = 0
	EntityAuthorityChangeNotify_CMD_ID           EntityAuthorityChangeNotify_CmdId = 326
	EntityAuthorityChangeNotify_ENET_CHANNEL_ID  EntityAuthorityChangeNotify_CmdId = 0
	EntityAuthorityChangeNotify_ENET_IS_RELIABLE EntityAuthorityChangeNotify_CmdId = 1
	EntityAuthorityChangeNotify_IS_ALLOW_CLIENT  EntityAuthorityChangeNotify_CmdId = 1
)

// Enum value maps for EntityAuthorityChangeNotify_CmdId.
var (
	EntityAuthorityChangeNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		326: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	EntityAuthorityChangeNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           326,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x EntityAuthorityChangeNotify_CmdId) Enum() *EntityAuthorityChangeNotify_CmdId {
	p := new(EntityAuthorityChangeNotify_CmdId)
	*p = x
	return p
}

func (x EntityAuthorityChangeNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntityAuthorityChangeNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_EntityAuthorityChangeNotify_proto_enumTypes[0].Descriptor()
}

func (EntityAuthorityChangeNotify_CmdId) Type() protoreflect.EnumType {
	return &file_EntityAuthorityChangeNotify_proto_enumTypes[0]
}

func (x EntityAuthorityChangeNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntityAuthorityChangeNotify_CmdId.Descriptor instead.
func (EntityAuthorityChangeNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_EntityAuthorityChangeNotify_proto_rawDescGZIP(), []int{0, 0}
}

type EntityAuthorityChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId            uint32                     `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	AuthorityPeerId     uint32                     `protobuf:"varint,2,opt,name=authority_peer_id,json=authorityPeerId,proto3" json:"authority_peer_id,omitempty"`
	AbilityInfo         *AbilitySyncStateInfo      `protobuf:"bytes,3,opt,name=ability_info,json=abilityInfo,proto3" json:"ability_info,omitempty"`
	RendererChangedInfo *EntityRendererChangedInfo `protobuf:"bytes,4,opt,name=renderer_changed_info,json=rendererChangedInfo,proto3" json:"renderer_changed_info,omitempty"`
}

func (x *EntityAuthorityChangeNotify) Reset() {
	*x = EntityAuthorityChangeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EntityAuthorityChangeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityAuthorityChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityAuthorityChangeNotify) ProtoMessage() {}

func (x *EntityAuthorityChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EntityAuthorityChangeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityAuthorityChangeNotify.ProtoReflect.Descriptor instead.
func (*EntityAuthorityChangeNotify) Descriptor() ([]byte, []int) {
	return file_EntityAuthorityChangeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EntityAuthorityChangeNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *EntityAuthorityChangeNotify) GetAuthorityPeerId() uint32 {
	if x != nil {
		return x.AuthorityPeerId
	}
	return 0
}

func (x *EntityAuthorityChangeNotify) GetAbilityInfo() *AbilitySyncStateInfo {
	if x != nil {
		return x.AbilityInfo
	}
	return nil
}

func (x *EntityAuthorityChangeNotify) GetRendererChangedInfo() *EntityRendererChangedInfo {
	if x != nil {
		return x.RendererChangedInfo
	}
	return nil
}

var File_EntityAuthorityChangeNotify_proto protoreflect.FileDescriptor

var file_EntityAuthorityChangeNotify_proto_rawDesc = []byte{
	0x0a, 0x21, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x41, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x02, 0x0a, 0x1b, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x3e, 0x0a, 0x0c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x54, 0x0a, 0x15, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x13, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44,
	0x5f, 0x49, 0x44, 0x10, 0xc6, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43,
	0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c,
	0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EntityAuthorityChangeNotify_proto_rawDescOnce sync.Once
	file_EntityAuthorityChangeNotify_proto_rawDescData = file_EntityAuthorityChangeNotify_proto_rawDesc
)

func file_EntityAuthorityChangeNotify_proto_rawDescGZIP() []byte {
	file_EntityAuthorityChangeNotify_proto_rawDescOnce.Do(func() {
		file_EntityAuthorityChangeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EntityAuthorityChangeNotify_proto_rawDescData)
	})
	return file_EntityAuthorityChangeNotify_proto_rawDescData
}

var file_EntityAuthorityChangeNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_EntityAuthorityChangeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EntityAuthorityChangeNotify_proto_goTypes = []interface{}{
	(EntityAuthorityChangeNotify_CmdId)(0), // 0: proto.EntityAuthorityChangeNotify.CmdId
	(*EntityAuthorityChangeNotify)(nil),    // 1: proto.EntityAuthorityChangeNotify
	(*AbilitySyncStateInfo)(nil),           // 2: proto.AbilitySyncStateInfo
	(*EntityRendererChangedInfo)(nil),      // 3: proto.EntityRendererChangedInfo
}
var file_EntityAuthorityChangeNotify_proto_depIdxs = []int32{
	2, // 0: proto.EntityAuthorityChangeNotify.ability_info:type_name -> proto.AbilitySyncStateInfo
	3, // 1: proto.EntityAuthorityChangeNotify.renderer_changed_info:type_name -> proto.EntityRendererChangedInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_EntityAuthorityChangeNotify_proto_init() }
func file_EntityAuthorityChangeNotify_proto_init() {
	if File_EntityAuthorityChangeNotify_proto != nil {
		return
	}
	file_EntityRendererChangedInfo_proto_init()
	file_AbilitySyncStateInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EntityAuthorityChangeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityAuthorityChangeNotify); i {
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
			RawDescriptor: file_EntityAuthorityChangeNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EntityAuthorityChangeNotify_proto_goTypes,
		DependencyIndexes: file_EntityAuthorityChangeNotify_proto_depIdxs,
		EnumInfos:         file_EntityAuthorityChangeNotify_proto_enumTypes,
		MessageInfos:      file_EntityAuthorityChangeNotify_proto_msgTypes,
	}.Build()
	File_EntityAuthorityChangeNotify_proto = out.File
	file_EntityAuthorityChangeNotify_proto_rawDesc = nil
	file_EntityAuthorityChangeNotify_proto_goTypes = nil
	file_EntityAuthorityChangeNotify_proto_depIdxs = nil
}
