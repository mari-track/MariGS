// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: TeamEntityInfo.proto

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

type TeamEntityInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamEntityId    uint32                `protobuf:"varint,1,opt,name=team_entity_id,json=teamEntityId,proto3" json:"team_entity_id,omitempty"`
	AuthorityPeerId uint32                `protobuf:"varint,2,opt,name=authority_peer_id,json=authorityPeerId,proto3" json:"authority_peer_id,omitempty"`
	TeamAbilityInfo *AbilitySyncStateInfo `protobuf:"bytes,3,opt,name=team_ability_info,json=teamAbilityInfo,proto3" json:"team_ability_info,omitempty"`
}

func (x *TeamEntityInfo) Reset() {
	*x = TeamEntityInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TeamEntityInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamEntityInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamEntityInfo) ProtoMessage() {}

func (x *TeamEntityInfo) ProtoReflect() protoreflect.Message {
	mi := &file_TeamEntityInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamEntityInfo.ProtoReflect.Descriptor instead.
func (*TeamEntityInfo) Descriptor() ([]byte, []int) {
	return file_TeamEntityInfo_proto_rawDescGZIP(), []int{0}
}

func (x *TeamEntityInfo) GetTeamEntityId() uint32 {
	if x != nil {
		return x.TeamEntityId
	}
	return 0
}

func (x *TeamEntityInfo) GetAuthorityPeerId() uint32 {
	if x != nil {
		return x.AuthorityPeerId
	}
	return 0
}

func (x *TeamEntityInfo) GetTeamAbilityInfo() *AbilitySyncStateInfo {
	if x != nil {
		return x.TeamAbilityInfo
	}
	return nil
}

var File_TeamEntityInfo_proto protoreflect.FileDescriptor

var file_TeamEntityInfo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x54, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53,
	0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x0e, 0x54, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x74,
	0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x11, 0x74, 0x65, 0x61, 0x6d, 0x5f,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x79, 0x6e, 0x63,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x74, 0x65, 0x61, 0x6d, 0x41,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TeamEntityInfo_proto_rawDescOnce sync.Once
	file_TeamEntityInfo_proto_rawDescData = file_TeamEntityInfo_proto_rawDesc
)

func file_TeamEntityInfo_proto_rawDescGZIP() []byte {
	file_TeamEntityInfo_proto_rawDescOnce.Do(func() {
		file_TeamEntityInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_TeamEntityInfo_proto_rawDescData)
	})
	return file_TeamEntityInfo_proto_rawDescData
}

var file_TeamEntityInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TeamEntityInfo_proto_goTypes = []interface{}{
	(*TeamEntityInfo)(nil),       // 0: TeamEntityInfo
	(*AbilitySyncStateInfo)(nil), // 1: AbilitySyncStateInfo
}
var file_TeamEntityInfo_proto_depIdxs = []int32{
	1, // 0: TeamEntityInfo.team_ability_info:type_name -> AbilitySyncStateInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_TeamEntityInfo_proto_init() }
func file_TeamEntityInfo_proto_init() {
	if File_TeamEntityInfo_proto != nil {
		return
	}
	file_AbilitySyncStateInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TeamEntityInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamEntityInfo); i {
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
			RawDescriptor: file_TeamEntityInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TeamEntityInfo_proto_goTypes,
		DependencyIndexes: file_TeamEntityInfo_proto_depIdxs,
		MessageInfos:      file_TeamEntityInfo_proto_msgTypes,
	}.Build()
	File_TeamEntityInfo_proto = out.File
	file_TeamEntityInfo_proto_rawDesc = nil
	file_TeamEntityInfo_proto_goTypes = nil
	file_TeamEntityInfo_proto_depIdxs = nil
}
