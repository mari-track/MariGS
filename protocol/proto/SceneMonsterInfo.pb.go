// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: SceneMonsterInfo.proto

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

type SceneMonsterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MonsterId       uint32             `protobuf:"varint,1,opt,name=monster_id,json=monsterId,proto3" json:"monster_id,omitempty"`
	GroupId         uint32             `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	ConfigId        uint32             `protobuf:"varint,3,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	WeaponList      []*SceneWeaponInfo `protobuf:"bytes,4,rep,name=weapon_list,json=weaponList,proto3" json:"weapon_list,omitempty"`
	AuthorityPeerId uint32             `protobuf:"varint,5,opt,name=authority_peer_id,json=authorityPeerId,proto3" json:"authority_peer_id,omitempty"`
	AffixList       []uint32           `protobuf:"varint,6,rep,packed,name=affix_list,json=affixList,proto3" json:"affix_list,omitempty"`
	IsElite         bool               `protobuf:"varint,7,opt,name=is_elite,json=isElite,proto3" json:"is_elite,omitempty"`
	OwnerEntityId   uint32             `protobuf:"varint,8,opt,name=owner_entity_id,json=ownerEntityId,proto3" json:"owner_entity_id,omitempty"`
	SummonedTag     uint32             `protobuf:"varint,9,opt,name=summoned_tag,json=summonedTag,proto3" json:"summoned_tag,omitempty"`
	SummonTagMap    map[uint32]uint32  `protobuf:"bytes,10,rep,name=summon_tag_map,json=summonTagMap,proto3" json:"summon_tag_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	PoseId          uint32             `protobuf:"varint,11,opt,name=pose_id,json=poseId,proto3" json:"pose_id,omitempty"`
	BornType        MonsterBornType    `protobuf:"varint,12,opt,name=born_type,json=bornType,proto3,enum=MonsterBornType" json:"born_type,omitempty"`
	BlockId         uint32             `protobuf:"varint,13,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	MarkFlag        uint32             `protobuf:"varint,14,opt,name=mark_flag,json=markFlag,proto3" json:"mark_flag,omitempty"`
	TitleId         uint32             `protobuf:"varint,15,opt,name=title_id,json=titleId,proto3" json:"title_id,omitempty"`
	SpecialNameId   uint32             `protobuf:"varint,16,opt,name=special_name_id,json=specialNameId,proto3" json:"special_name_id,omitempty"`
	AttackTargetId  uint32             `protobuf:"varint,17,opt,name=attack_target_id,json=attackTargetId,proto3" json:"attack_target_id,omitempty"`
}

func (x *SceneMonsterInfo) Reset() {
	*x = SceneMonsterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneMonsterInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneMonsterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneMonsterInfo) ProtoMessage() {}

func (x *SceneMonsterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SceneMonsterInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneMonsterInfo.ProtoReflect.Descriptor instead.
func (*SceneMonsterInfo) Descriptor() ([]byte, []int) {
	return file_SceneMonsterInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneMonsterInfo) GetMonsterId() uint32 {
	if x != nil {
		return x.MonsterId
	}
	return 0
}

func (x *SceneMonsterInfo) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *SceneMonsterInfo) GetConfigId() uint32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *SceneMonsterInfo) GetWeaponList() []*SceneWeaponInfo {
	if x != nil {
		return x.WeaponList
	}
	return nil
}

func (x *SceneMonsterInfo) GetAuthorityPeerId() uint32 {
	if x != nil {
		return x.AuthorityPeerId
	}
	return 0
}

func (x *SceneMonsterInfo) GetAffixList() []uint32 {
	if x != nil {
		return x.AffixList
	}
	return nil
}

func (x *SceneMonsterInfo) GetIsElite() bool {
	if x != nil {
		return x.IsElite
	}
	return false
}

func (x *SceneMonsterInfo) GetOwnerEntityId() uint32 {
	if x != nil {
		return x.OwnerEntityId
	}
	return 0
}

func (x *SceneMonsterInfo) GetSummonedTag() uint32 {
	if x != nil {
		return x.SummonedTag
	}
	return 0
}

func (x *SceneMonsterInfo) GetSummonTagMap() map[uint32]uint32 {
	if x != nil {
		return x.SummonTagMap
	}
	return nil
}

func (x *SceneMonsterInfo) GetPoseId() uint32 {
	if x != nil {
		return x.PoseId
	}
	return 0
}

func (x *SceneMonsterInfo) GetBornType() MonsterBornType {
	if x != nil {
		return x.BornType
	}
	return MonsterBornType_MONSTER_BORN_NONE
}

func (x *SceneMonsterInfo) GetBlockId() uint32 {
	if x != nil {
		return x.BlockId
	}
	return 0
}

func (x *SceneMonsterInfo) GetMarkFlag() uint32 {
	if x != nil {
		return x.MarkFlag
	}
	return 0
}

func (x *SceneMonsterInfo) GetTitleId() uint32 {
	if x != nil {
		return x.TitleId
	}
	return 0
}

func (x *SceneMonsterInfo) GetSpecialNameId() uint32 {
	if x != nil {
		return x.SpecialNameId
	}
	return 0
}

func (x *SceneMonsterInfo) GetAttackTargetId() uint32 {
	if x != nil {
		return x.AttackTargetId
	}
	return 0
}

var File_SceneMonsterInfo_proto protoreflect.FileDescriptor

var file_SceneMonsterInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x57,
	0x65, 0x61, 0x70, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x6f, 0x72, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x05, 0x0a, 0x10, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x49, 0x64, 0x12, 0x31, 0x0a, 0x0b, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x57,
	0x65, 0x61, 0x70, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x77, 0x65, 0x61, 0x70, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x50, 0x65, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x66, 0x66, 0x69, 0x78, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x66, 0x66, 0x69, 0x78, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x65, 0x6c, 0x69, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x45, 0x6c, 0x69, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x64, 0x5f,
	0x74, 0x61, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73, 0x75, 0x6d, 0x6d, 0x6f,
	0x6e, 0x65, 0x64, 0x54, 0x61, 0x67, 0x12, 0x49, 0x0a, 0x0e, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x74, 0x61, 0x67, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x54, 0x61, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0c, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x54, 0x61, 0x67, 0x4d, 0x61,
	0x70, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x65, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x09, 0x62, 0x6f,
	0x72, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x6f, 0x72, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x62, 0x6f, 0x72, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x66, 0x6c, 0x61,
	0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x46, 0x6c, 0x61,
	0x67, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x1a, 0x3f,
	0x0a, 0x11, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x54, 0x61, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_SceneMonsterInfo_proto_rawDescOnce sync.Once
	file_SceneMonsterInfo_proto_rawDescData = file_SceneMonsterInfo_proto_rawDesc
)

func file_SceneMonsterInfo_proto_rawDescGZIP() []byte {
	file_SceneMonsterInfo_proto_rawDescOnce.Do(func() {
		file_SceneMonsterInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneMonsterInfo_proto_rawDescData)
	})
	return file_SceneMonsterInfo_proto_rawDescData
}

var file_SceneMonsterInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_SceneMonsterInfo_proto_goTypes = []interface{}{
	(*SceneMonsterInfo)(nil), // 0: SceneMonsterInfo
	nil,                      // 1: SceneMonsterInfo.SummonTagMapEntry
	(*SceneWeaponInfo)(nil),  // 2: SceneWeaponInfo
	(MonsterBornType)(0),     // 3: MonsterBornType
}
var file_SceneMonsterInfo_proto_depIdxs = []int32{
	2, // 0: SceneMonsterInfo.weapon_list:type_name -> SceneWeaponInfo
	1, // 1: SceneMonsterInfo.summon_tag_map:type_name -> SceneMonsterInfo.SummonTagMapEntry
	3, // 2: SceneMonsterInfo.born_type:type_name -> MonsterBornType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_SceneMonsterInfo_proto_init() }
func file_SceneMonsterInfo_proto_init() {
	if File_SceneMonsterInfo_proto != nil {
		return
	}
	file_SceneWeaponInfo_proto_init()
	file_MonsterBornType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneMonsterInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneMonsterInfo); i {
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
			RawDescriptor: file_SceneMonsterInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneMonsterInfo_proto_goTypes,
		DependencyIndexes: file_SceneMonsterInfo_proto_depIdxs,
		MessageInfos:      file_SceneMonsterInfo_proto_msgTypes,
	}.Build()
	File_SceneMonsterInfo_proto = out.File
	file_SceneMonsterInfo_proto_rawDesc = nil
	file_SceneMonsterInfo_proto_goTypes = nil
	file_SceneMonsterInfo_proto_depIdxs = nil
}
