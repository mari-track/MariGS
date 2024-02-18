// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: AvatarInfo.proto

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

type AvatarInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarId                uint32                      `protobuf:"varint,1,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	Guid                    uint64                      `protobuf:"varint,2,opt,name=guid,proto3" json:"guid,omitempty"`
	PropMap                 map[uint32]*PropValue       `protobuf:"bytes,3,rep,name=prop_map,json=propMap,proto3" json:"prop_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LifeState               uint32                      `protobuf:"varint,4,opt,name=life_state,json=lifeState,proto3" json:"life_state,omitempty"`
	EquipGuidList           []uint64                    `protobuf:"varint,5,rep,packed,name=equip_guid_list,json=equipGuidList,proto3" json:"equip_guid_list,omitempty"`
	TalentIdList            []uint32                    `protobuf:"varint,6,rep,packed,name=talent_id_list,json=talentIdList,proto3" json:"talent_id_list,omitempty"`
	FightPropMap            map[uint32]float32          `protobuf:"bytes,7,rep,name=fight_prop_map,json=fightPropMap,proto3" json:"fight_prop_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	TrialAvatarInfo         *TrialAvatarInfo            `protobuf:"bytes,9,opt,name=trial_avatar_info,json=trialAvatarInfo,proto3" json:"trial_avatar_info,omitempty"`
	SkillMap                map[uint32]*AvatarSkillInfo `protobuf:"bytes,10,rep,name=skill_map,json=skillMap,proto3" json:"skill_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SkillDepotId            uint32                      `protobuf:"varint,11,opt,name=skill_depot_id,json=skillDepotId,proto3" json:"skill_depot_id,omitempty"`
	FetterInfo              *AvatarFetterInfo           `protobuf:"bytes,12,opt,name=fetter_info,json=fetterInfo,proto3" json:"fetter_info,omitempty"`
	CoreProudSkillLevel     uint32                      `protobuf:"varint,13,opt,name=core_proud_skill_level,json=coreProudSkillLevel,proto3" json:"core_proud_skill_level,omitempty"`
	InherentProudSkillList  []uint32                    `protobuf:"varint,14,rep,packed,name=inherent_proud_skill_list,json=inherentProudSkillList,proto3" json:"inherent_proud_skill_list,omitempty"`
	SkillLevelMap           map[uint32]uint32           `protobuf:"bytes,15,rep,name=skill_level_map,json=skillLevelMap,proto3" json:"skill_level_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ExpeditionState         AvatarExpeditionState       `protobuf:"varint,16,opt,name=expedition_state,json=expeditionState,proto3,enum=proto.AvatarExpeditionState" json:"expedition_state,omitempty"`
	ProudSkillExtraLevelMap map[uint32]uint32           `protobuf:"bytes,17,rep,name=proud_skill_extra_level_map,json=proudSkillExtraLevelMap,proto3" json:"proud_skill_extra_level_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	IsFocus                 bool                        `protobuf:"varint,18,opt,name=is_focus,json=isFocus,proto3" json:"is_focus,omitempty"`
	AvatarType              AvatarType                  `protobuf:"varint,19,opt,name=avatar_type,json=avatarType,proto3,enum=proto.AvatarType" json:"avatar_type,omitempty"`
	TeamResonanceList       []uint32                    `protobuf:"varint,20,rep,packed,name=team_resonance_list,json=teamResonanceList,proto3" json:"team_resonance_list,omitempty"`
	WearingFlycloakId       uint32                      `protobuf:"varint,21,opt,name=wearing_flycloak_id,json=wearingFlycloakId,proto3" json:"wearing_flycloak_id,omitempty"`
	EquipAffixList          []*AvatarEquipAffixInfo     `protobuf:"bytes,22,rep,name=equip_affix_list,json=equipAffixList,proto3" json:"equip_affix_list,omitempty"`
}

func (x *AvatarInfo) Reset() {
	*x = AvatarInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarInfo) ProtoMessage() {}

func (x *AvatarInfo) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarInfo.ProtoReflect.Descriptor instead.
func (*AvatarInfo) Descriptor() ([]byte, []int) {
	return file_AvatarInfo_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarInfo) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *AvatarInfo) GetGuid() uint64 {
	if x != nil {
		return x.Guid
	}
	return 0
}

func (x *AvatarInfo) GetPropMap() map[uint32]*PropValue {
	if x != nil {
		return x.PropMap
	}
	return nil
}

func (x *AvatarInfo) GetLifeState() uint32 {
	if x != nil {
		return x.LifeState
	}
	return 0
}

func (x *AvatarInfo) GetEquipGuidList() []uint64 {
	if x != nil {
		return x.EquipGuidList
	}
	return nil
}

func (x *AvatarInfo) GetTalentIdList() []uint32 {
	if x != nil {
		return x.TalentIdList
	}
	return nil
}

func (x *AvatarInfo) GetFightPropMap() map[uint32]float32 {
	if x != nil {
		return x.FightPropMap
	}
	return nil
}

func (x *AvatarInfo) GetTrialAvatarInfo() *TrialAvatarInfo {
	if x != nil {
		return x.TrialAvatarInfo
	}
	return nil
}

func (x *AvatarInfo) GetSkillMap() map[uint32]*AvatarSkillInfo {
	if x != nil {
		return x.SkillMap
	}
	return nil
}

func (x *AvatarInfo) GetSkillDepotId() uint32 {
	if x != nil {
		return x.SkillDepotId
	}
	return 0
}

func (x *AvatarInfo) GetFetterInfo() *AvatarFetterInfo {
	if x != nil {
		return x.FetterInfo
	}
	return nil
}

func (x *AvatarInfo) GetCoreProudSkillLevel() uint32 {
	if x != nil {
		return x.CoreProudSkillLevel
	}
	return 0
}

func (x *AvatarInfo) GetInherentProudSkillList() []uint32 {
	if x != nil {
		return x.InherentProudSkillList
	}
	return nil
}

func (x *AvatarInfo) GetSkillLevelMap() map[uint32]uint32 {
	if x != nil {
		return x.SkillLevelMap
	}
	return nil
}

func (x *AvatarInfo) GetExpeditionState() AvatarExpeditionState {
	if x != nil {
		return x.ExpeditionState
	}
	return AvatarExpeditionState_AVATAR_EXPEDITION_NONE
}

func (x *AvatarInfo) GetProudSkillExtraLevelMap() map[uint32]uint32 {
	if x != nil {
		return x.ProudSkillExtraLevelMap
	}
	return nil
}

func (x *AvatarInfo) GetIsFocus() bool {
	if x != nil {
		return x.IsFocus
	}
	return false
}

func (x *AvatarInfo) GetAvatarType() AvatarType {
	if x != nil {
		return x.AvatarType
	}
	return AvatarType_AVATAR_TYPE_NONE
}

func (x *AvatarInfo) GetTeamResonanceList() []uint32 {
	if x != nil {
		return x.TeamResonanceList
	}
	return nil
}

func (x *AvatarInfo) GetWearingFlycloakId() uint32 {
	if x != nil {
		return x.WearingFlycloakId
	}
	return 0
}

func (x *AvatarInfo) GetEquipAffixList() []*AvatarEquipAffixInfo {
	if x != nil {
		return x.EquipAffixList
	}
	return nil
}

var File_AvatarInfo_proto protoreflect.FileDescriptor

var file_AvatarInfo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x54, 0x72, 0x69, 0x61, 0x6c,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x46, 0x65, 0x74, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x71, 0x75, 0x69, 0x70, 0x41, 0x66, 0x66, 0x69, 0x78, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x0b, 0x0a, 0x0a,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x70, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x50, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x66, 0x65, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x5f, 0x67,
	0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0d,
	0x65, 0x71, 0x75, 0x69, 0x70, 0x47, 0x75, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x0e, 0x66, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x70, 0x72, 0x6f,
	0x70, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x46,
	0x69, 0x67, 0x68, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0c, 0x66, 0x69, 0x67, 0x68, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x12, 0x42,
	0x0a, 0x11, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0f, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x09, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x4d, 0x61, 0x70,
	0x12, 0x24, 0x0a, 0x0e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x44,
	0x65, 0x70, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0b, 0x66, 0x65, 0x74, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x46, 0x65, 0x74, 0x74, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x66, 0x65, 0x74, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x33, 0x0a, 0x16, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x75, 0x64, 0x5f, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x13, 0x63, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x39, 0x0a, 0x19, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x16, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x4c, 0x0a, 0x0f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f,
	0x6d, 0x61, 0x70, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0d, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x12, 0x47,
	0x0a, 0x10, 0x65, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x6c, 0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x75, 0x64,
	0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x50, 0x72, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x17, 0x70, 0x72,
	0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x4d, 0x61, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x63, 0x75,
	0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x46, 0x6f, 0x63, 0x75, 0x73,
	0x12, 0x32, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x73,
	0x6f, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x14, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x11, 0x74, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x77, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x66, 0x6c, 0x79, 0x63, 0x6c, 0x6f, 0x61, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x11, 0x77, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x6c, 0x79, 0x63, 0x6c, 0x6f,
	0x61, 0x6b, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x10, 0x65, 0x71, 0x75, 0x69, 0x70, 0x5f, 0x61, 0x66,
	0x66, 0x69, 0x78, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x16, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x71, 0x75,
	0x69, 0x70, 0x41, 0x66, 0x66, 0x69, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x65, 0x71, 0x75,
	0x69, 0x70, 0x41, 0x66, 0x66, 0x69, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x4c, 0x0a, 0x0c, 0x50,
	0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3f, 0x0a, 0x11, 0x46, 0x69, 0x67,
	0x68, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x53, 0x0a, 0x0d, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x40, 0x0a, 0x12, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x4a, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_AvatarInfo_proto_rawDescOnce sync.Once
	file_AvatarInfo_proto_rawDescData = file_AvatarInfo_proto_rawDesc
)

func file_AvatarInfo_proto_rawDescGZIP() []byte {
	file_AvatarInfo_proto_rawDescOnce.Do(func() {
		file_AvatarInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarInfo_proto_rawDescData)
	})
	return file_AvatarInfo_proto_rawDescData
}

var file_AvatarInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_AvatarInfo_proto_goTypes = []interface{}{
	(*AvatarInfo)(nil),           // 0: proto.AvatarInfo
	nil,                          // 1: proto.AvatarInfo.PropMapEntry
	nil,                          // 2: proto.AvatarInfo.FightPropMapEntry
	nil,                          // 3: proto.AvatarInfo.SkillMapEntry
	nil,                          // 4: proto.AvatarInfo.SkillLevelMapEntry
	nil,                          // 5: proto.AvatarInfo.ProudSkillExtraLevelMapEntry
	(*TrialAvatarInfo)(nil),      // 6: proto.TrialAvatarInfo
	(*AvatarFetterInfo)(nil),     // 7: proto.AvatarFetterInfo
	(AvatarExpeditionState)(0),   // 8: proto.AvatarExpeditionState
	(AvatarType)(0),              // 9: proto.AvatarType
	(*AvatarEquipAffixInfo)(nil), // 10: proto.AvatarEquipAffixInfo
	(*PropValue)(nil),            // 11: proto.PropValue
	(*AvatarSkillInfo)(nil),      // 12: proto.AvatarSkillInfo
}
var file_AvatarInfo_proto_depIdxs = []int32{
	1,  // 0: proto.AvatarInfo.prop_map:type_name -> proto.AvatarInfo.PropMapEntry
	2,  // 1: proto.AvatarInfo.fight_prop_map:type_name -> proto.AvatarInfo.FightPropMapEntry
	6,  // 2: proto.AvatarInfo.trial_avatar_info:type_name -> proto.TrialAvatarInfo
	3,  // 3: proto.AvatarInfo.skill_map:type_name -> proto.AvatarInfo.SkillMapEntry
	7,  // 4: proto.AvatarInfo.fetter_info:type_name -> proto.AvatarFetterInfo
	4,  // 5: proto.AvatarInfo.skill_level_map:type_name -> proto.AvatarInfo.SkillLevelMapEntry
	8,  // 6: proto.AvatarInfo.expedition_state:type_name -> proto.AvatarExpeditionState
	5,  // 7: proto.AvatarInfo.proud_skill_extra_level_map:type_name -> proto.AvatarInfo.ProudSkillExtraLevelMapEntry
	9,  // 8: proto.AvatarInfo.avatar_type:type_name -> proto.AvatarType
	10, // 9: proto.AvatarInfo.equip_affix_list:type_name -> proto.AvatarEquipAffixInfo
	11, // 10: proto.AvatarInfo.PropMapEntry.value:type_name -> proto.PropValue
	12, // 11: proto.AvatarInfo.SkillMapEntry.value:type_name -> proto.AvatarSkillInfo
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_AvatarInfo_proto_init() }
func file_AvatarInfo_proto_init() {
	if File_AvatarInfo_proto != nil {
		return
	}
	file_TrialAvatarInfo_proto_init()
	file_PropValue_proto_init()
	file_AvatarExpeditionState_proto_init()
	file_AvatarFetterInfo_proto_init()
	file_AvatarSkillInfo_proto_init()
	file_AvatarEquipAffixInfo_proto_init()
	file_AvatarType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AvatarInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarInfo); i {
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
			RawDescriptor: file_AvatarInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarInfo_proto_goTypes,
		DependencyIndexes: file_AvatarInfo_proto_depIdxs,
		MessageInfos:      file_AvatarInfo_proto_msgTypes,
	}.Build()
	File_AvatarInfo_proto = out.File
	file_AvatarInfo_proto_rawDesc = nil
	file_AvatarInfo_proto_goTypes = nil
	file_AvatarInfo_proto_depIdxs = nil
}
