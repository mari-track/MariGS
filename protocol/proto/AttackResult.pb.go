// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: AttackResult.proto

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

type AttackResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttackerId            uint32                 `protobuf:"varint,1,opt,name=attacker_id,json=attackerId,proto3" json:"attacker_id,omitempty"`
	DefenseId             uint32                 `protobuf:"varint,2,opt,name=defense_id,json=defenseId,proto3" json:"defense_id,omitempty"`
	AnimEventId           string                 `protobuf:"bytes,3,opt,name=anim_event_id,json=animEventId,proto3" json:"anim_event_id,omitempty"`
	AbilityIdentifier     *AbilityIdentifier     `protobuf:"bytes,4,opt,name=ability_identifier,json=abilityIdentifier,proto3" json:"ability_identifier,omitempty"`
	AttackTag             string                 `protobuf:"bytes,5,opt,name=attack_tag,json=attackTag,proto3" json:"attack_tag,omitempty"`
	Damage                float32                `protobuf:"fixed32,6,opt,name=damage,proto3" json:"damage,omitempty"`
	IsCrit                bool                   `protobuf:"varint,7,opt,name=is_crit,json=isCrit,proto3" json:"is_crit,omitempty"`
	HitCollision          *HitCollision          `protobuf:"bytes,8,opt,name=hit_collision,json=hitCollision,proto3" json:"hit_collision,omitempty"`
	HitPosType            uint32                 `protobuf:"varint,9,opt,name=hit_pos_type,json=hitPosType,proto3" json:"hit_pos_type,omitempty"`
	EndureBreak           uint32                 `protobuf:"varint,10,opt,name=endure_break,json=endureBreak,proto3" json:"endure_break,omitempty"`
	ResolvedDir           *Vector                `protobuf:"bytes,11,opt,name=resolved_dir,json=resolvedDir,proto3" json:"resolved_dir,omitempty"`
	HitRetreatAngleCompat int32                  `protobuf:"varint,12,opt,name=hit_retreat_angle_compat,json=hitRetreatAngleCompat,proto3" json:"hit_retreat_angle_compat,omitempty"`
	HitEffResult          *AttackHitEffectResult `protobuf:"bytes,13,opt,name=hit_eff_result,json=hitEffResult,proto3" json:"hit_eff_result,omitempty"`
	ElementType           uint32                 `protobuf:"varint,14,opt,name=element_type,json=elementType,proto3" json:"element_type,omitempty"`
	DamagePercentage      float32                `protobuf:"fixed32,15,opt,name=damage_percentage,json=damagePercentage,proto3" json:"damage_percentage,omitempty"`
	DamageExtra           float32                `protobuf:"fixed32,16,opt,name=damage_extra,json=damageExtra,proto3" json:"damage_extra,omitempty"`
	BonusCritical         float32                `protobuf:"fixed32,17,opt,name=bonus_critical,json=bonusCritical,proto3" json:"bonus_critical,omitempty"`
	BonusCriticalHurt     float32                `protobuf:"fixed32,18,opt,name=bonus_critical_hurt,json=bonusCriticalHurt,proto3" json:"bonus_critical_hurt,omitempty"`
	UseGadgetDamageAction bool                   `protobuf:"varint,19,opt,name=use_gadget_damage_action,json=useGadgetDamageAction,proto3" json:"use_gadget_damage_action,omitempty"`
	GadgetDamageActionIdx uint32                 `protobuf:"varint,20,opt,name=gadget_damage_action_idx,json=gadgetDamageActionIdx,proto3" json:"gadget_damage_action_idx,omitempty"`
	IsResistText          bool                   `protobuf:"varint,22,opt,name=is_resist_text,json=isResistText,proto3" json:"is_resist_text,omitempty"`
	CriticalRand          float32                `protobuf:"fixed32,23,opt,name=critical_rand,json=criticalRand,proto3" json:"critical_rand,omitempty"`
	ElementAmplifyRate    float32                `protobuf:"fixed32,24,opt,name=element_amplify_rate,json=elementAmplifyRate,proto3" json:"element_amplify_rate,omitempty"`
	AttenuationTag        string                 `protobuf:"bytes,25,opt,name=attenuation_tag,json=attenuationTag,proto3" json:"attenuation_tag,omitempty"`
	DamageShield          float32                `protobuf:"fixed32,26,opt,name=damage_shield,json=damageShield,proto3" json:"damage_shield,omitempty"`
	MuteElementHurt       bool                   `protobuf:"varint,27,opt,name=mute_element_hurt,json=muteElementHurt,proto3" json:"mute_element_hurt,omitempty"`
	TrueDamage            bool                   `protobuf:"varint,28,opt,name=true_damage,json=trueDamage,proto3" json:"true_damage,omitempty"`
	AttenuationGroup      string                 `protobuf:"bytes,29,opt,name=attenuation_group,json=attenuationGroup,proto3" json:"attenuation_group,omitempty"`
	AmplifyReactionType   uint32                 `protobuf:"varint,30,opt,name=amplify_reaction_type,json=amplifyReactionType,proto3" json:"amplify_reaction_type,omitempty"`
	AddhurtReactionType   uint32                 `protobuf:"varint,31,opt,name=addhurt_reaction_type,json=addhurtReactionType,proto3" json:"addhurt_reaction_type,omitempty"`
	BulletFlyTimeMs       uint32                 `protobuf:"varint,32,opt,name=bullet_fly_time_ms,json=bulletFlyTimeMs,proto3" json:"bullet_fly_time_ms,omitempty"`
	AttackCount           uint32                 `protobuf:"varint,33,opt,name=attack_count,json=attackCount,proto3" json:"attack_count,omitempty"`
	HashedAnimEventId     uint32                 `protobuf:"varint,34,opt,name=hashed_anim_event_id,json=hashedAnimEventId,proto3" json:"hashed_anim_event_id,omitempty"`
}

func (x *AttackResult) Reset() {
	*x = AttackResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AttackResult_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttackResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttackResult) ProtoMessage() {}

func (x *AttackResult) ProtoReflect() protoreflect.Message {
	mi := &file_AttackResult_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttackResult.ProtoReflect.Descriptor instead.
func (*AttackResult) Descriptor() ([]byte, []int) {
	return file_AttackResult_proto_rawDescGZIP(), []int{0}
}

func (x *AttackResult) GetAttackerId() uint32 {
	if x != nil {
		return x.AttackerId
	}
	return 0
}

func (x *AttackResult) GetDefenseId() uint32 {
	if x != nil {
		return x.DefenseId
	}
	return 0
}

func (x *AttackResult) GetAnimEventId() string {
	if x != nil {
		return x.AnimEventId
	}
	return ""
}

func (x *AttackResult) GetAbilityIdentifier() *AbilityIdentifier {
	if x != nil {
		return x.AbilityIdentifier
	}
	return nil
}

func (x *AttackResult) GetAttackTag() string {
	if x != nil {
		return x.AttackTag
	}
	return ""
}

func (x *AttackResult) GetDamage() float32 {
	if x != nil {
		return x.Damage
	}
	return 0
}

func (x *AttackResult) GetIsCrit() bool {
	if x != nil {
		return x.IsCrit
	}
	return false
}

func (x *AttackResult) GetHitCollision() *HitCollision {
	if x != nil {
		return x.HitCollision
	}
	return nil
}

func (x *AttackResult) GetHitPosType() uint32 {
	if x != nil {
		return x.HitPosType
	}
	return 0
}

func (x *AttackResult) GetEndureBreak() uint32 {
	if x != nil {
		return x.EndureBreak
	}
	return 0
}

func (x *AttackResult) GetResolvedDir() *Vector {
	if x != nil {
		return x.ResolvedDir
	}
	return nil
}

func (x *AttackResult) GetHitRetreatAngleCompat() int32 {
	if x != nil {
		return x.HitRetreatAngleCompat
	}
	return 0
}

func (x *AttackResult) GetHitEffResult() *AttackHitEffectResult {
	if x != nil {
		return x.HitEffResult
	}
	return nil
}

func (x *AttackResult) GetElementType() uint32 {
	if x != nil {
		return x.ElementType
	}
	return 0
}

func (x *AttackResult) GetDamagePercentage() float32 {
	if x != nil {
		return x.DamagePercentage
	}
	return 0
}

func (x *AttackResult) GetDamageExtra() float32 {
	if x != nil {
		return x.DamageExtra
	}
	return 0
}

func (x *AttackResult) GetBonusCritical() float32 {
	if x != nil {
		return x.BonusCritical
	}
	return 0
}

func (x *AttackResult) GetBonusCriticalHurt() float32 {
	if x != nil {
		return x.BonusCriticalHurt
	}
	return 0
}

func (x *AttackResult) GetUseGadgetDamageAction() bool {
	if x != nil {
		return x.UseGadgetDamageAction
	}
	return false
}

func (x *AttackResult) GetGadgetDamageActionIdx() uint32 {
	if x != nil {
		return x.GadgetDamageActionIdx
	}
	return 0
}

func (x *AttackResult) GetIsResistText() bool {
	if x != nil {
		return x.IsResistText
	}
	return false
}

func (x *AttackResult) GetCriticalRand() float32 {
	if x != nil {
		return x.CriticalRand
	}
	return 0
}

func (x *AttackResult) GetElementAmplifyRate() float32 {
	if x != nil {
		return x.ElementAmplifyRate
	}
	return 0
}

func (x *AttackResult) GetAttenuationTag() string {
	if x != nil {
		return x.AttenuationTag
	}
	return ""
}

func (x *AttackResult) GetDamageShield() float32 {
	if x != nil {
		return x.DamageShield
	}
	return 0
}

func (x *AttackResult) GetMuteElementHurt() bool {
	if x != nil {
		return x.MuteElementHurt
	}
	return false
}

func (x *AttackResult) GetTrueDamage() bool {
	if x != nil {
		return x.TrueDamage
	}
	return false
}

func (x *AttackResult) GetAttenuationGroup() string {
	if x != nil {
		return x.AttenuationGroup
	}
	return ""
}

func (x *AttackResult) GetAmplifyReactionType() uint32 {
	if x != nil {
		return x.AmplifyReactionType
	}
	return 0
}

func (x *AttackResult) GetAddhurtReactionType() uint32 {
	if x != nil {
		return x.AddhurtReactionType
	}
	return 0
}

func (x *AttackResult) GetBulletFlyTimeMs() uint32 {
	if x != nil {
		return x.BulletFlyTimeMs
	}
	return 0
}

func (x *AttackResult) GetAttackCount() uint32 {
	if x != nil {
		return x.AttackCount
	}
	return 0
}

func (x *AttackResult) GetHashedAnimEventId() uint32 {
	if x != nil {
		return x.HashedAnimEventId
	}
	return 0
}

var File_AttackResult_proto protoreflect.FileDescriptor

var file_AttackResult_proto_rawDesc = []byte{
	0x0a, 0x12, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x6b, 0x48, 0x69, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x48, 0x69, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x41, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x0b, 0x0a, 0x0c, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x73, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x65, 0x66, 0x65, 0x6e,
	0x73, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x6e, 0x69, 0x6d, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x6e, 0x69,
	0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x12, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x11,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x61, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x61, 0x67,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x63,
	0x72, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x43, 0x72, 0x69,
	0x74, 0x12, 0x38, 0x0a, 0x0d, 0x68, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x48, 0x69, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x68,
	0x69, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x68,
	0x69, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x68, 0x69, 0x74, 0x50, 0x6f, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x65, 0x6e, 0x64, 0x75, 0x72, 0x65, 0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x65, 0x6e, 0x64, 0x75, 0x72, 0x65, 0x42, 0x72, 0x65, 0x61, 0x6b,
	0x12, 0x30, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x72,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x44,
	0x69, 0x72, 0x12, 0x37, 0x0a, 0x18, 0x68, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x65, 0x61,
	0x74, 0x5f, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x68, 0x69, 0x74, 0x52, 0x65, 0x74, 0x72, 0x65, 0x61, 0x74,
	0x41, 0x6e, 0x67, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x12, 0x42, 0x0a, 0x0e, 0x68,
	0x69, 0x74, 0x5f, 0x65, 0x66, 0x66, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x6b, 0x48, 0x69, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x0c, 0x68, 0x69, 0x74, 0x45, 0x66, 0x66, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x64,
	0x61, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x63, 0x72, 0x69, 0x74,
	0x69, 0x63, 0x61, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x62, 0x6f, 0x6e, 0x75,
	0x73, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x62, 0x6f, 0x6e,
	0x75, 0x73, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x68, 0x75, 0x72, 0x74,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x43, 0x72, 0x69,
	0x74, 0x69, 0x63, 0x61, 0x6c, 0x48, 0x75, 0x72, 0x74, 0x12, 0x37, 0x0a, 0x18, 0x75, 0x73, 0x65,
	0x5f, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x75, 0x73, 0x65,
	0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x18, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x44, 0x61, 0x6d, 0x61,
	0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x78, 0x12, 0x24, 0x0a, 0x0e, 0x69,
	0x73, 0x5f, 0x72, 0x65, 0x73, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x52, 0x65, 0x73, 0x69, 0x73, 0x74, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61,
	0x6e, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63,
	0x61, 0x6c, 0x52, 0x61, 0x6e, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x66, 0x79, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x18,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x70,
	0x6c, 0x69, 0x66, 0x79, 0x52, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x74, 0x74, 0x65,
	0x6e, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61,
	0x67, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x68, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65,
	0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x75, 0x74, 0x65, 0x5f, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x75, 0x72, 0x74, 0x18, 0x1b, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x6d, 0x75, 0x74, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x75,
	0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x75, 0x65, 0x5f, 0x64, 0x61, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x74, 0x72, 0x75, 0x65, 0x44, 0x61, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x61, 0x74, 0x74, 0x65, 0x6e, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x32, 0x0a, 0x15, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x66, 0x79, 0x5f, 0x72, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x13, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x66, 0x79, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x61, 0x64, 0x64, 0x68, 0x75, 0x72, 0x74, 0x5f,
	0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1f, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x13, 0x61, 0x64, 0x64, 0x68, 0x75, 0x72, 0x74, 0x52, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x12, 0x62, 0x75, 0x6c, 0x6c,
	0x65, 0x74, 0x5f, 0x66, 0x6c, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x20,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x46, 0x6c, 0x79, 0x54,
	0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x14, 0x68, 0x61, 0x73, 0x68,
	0x65, 0x64, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x22, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x68, 0x61, 0x73, 0x68, 0x65, 0x64, 0x41, 0x6e,
	0x69, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AttackResult_proto_rawDescOnce sync.Once
	file_AttackResult_proto_rawDescData = file_AttackResult_proto_rawDesc
)

func file_AttackResult_proto_rawDescGZIP() []byte {
	file_AttackResult_proto_rawDescOnce.Do(func() {
		file_AttackResult_proto_rawDescData = protoimpl.X.CompressGZIP(file_AttackResult_proto_rawDescData)
	})
	return file_AttackResult_proto_rawDescData
}

var file_AttackResult_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AttackResult_proto_goTypes = []interface{}{
	(*AttackResult)(nil),          // 0: proto.AttackResult
	(*AbilityIdentifier)(nil),     // 1: proto.AbilityIdentifier
	(*HitCollision)(nil),          // 2: proto.HitCollision
	(*Vector)(nil),                // 3: proto.Vector
	(*AttackHitEffectResult)(nil), // 4: proto.AttackHitEffectResult
}
var file_AttackResult_proto_depIdxs = []int32{
	1, // 0: proto.AttackResult.ability_identifier:type_name -> proto.AbilityIdentifier
	2, // 1: proto.AttackResult.hit_collision:type_name -> proto.HitCollision
	3, // 2: proto.AttackResult.resolved_dir:type_name -> proto.Vector
	4, // 3: proto.AttackResult.hit_eff_result:type_name -> proto.AttackHitEffectResult
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_AttackResult_proto_init() }
func file_AttackResult_proto_init() {
	if File_AttackResult_proto != nil {
		return
	}
	file_Vector_proto_init()
	file_AttackHitEffectResult_proto_init()
	file_HitCollision_proto_init()
	file_AbilityIdentifier_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AttackResult_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttackResult); i {
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
			RawDescriptor: file_AttackResult_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AttackResult_proto_goTypes,
		DependencyIndexes: file_AttackResult_proto_depIdxs,
		MessageInfos:      file_AttackResult_proto_msgTypes,
	}.Build()
	File_AttackResult_proto = out.File
	file_AttackResult_proto_rawDesc = nil
	file_AttackResult_proto_goTypes = nil
	file_AttackResult_proto_depIdxs = nil
}
