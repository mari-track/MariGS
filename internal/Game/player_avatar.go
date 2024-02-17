package Game

import (
	"github.com/mari-track/MariGS/pkg/constant"
	"github.com/mari-track/MariGS/pkg/gdconf"
	"github.com/mari-track/MariGS/protocol/cmd"
	"github.com/mari-track/MariGS/protocol/proto"
)

/*************************************通知包******************************************/

func (g *Game) AvatarFightPropUpdateNotify(avatarId uint32) {
	db := g.Player.GetPbAvatarById(avatarId)
	notify := &proto.AvatarFightPropUpdateNotify{
		AvatarGuid:   db.Guid,
		FightPropMap: db.FightPropMap,
	}

	g.seed(cmd.AvatarFightPropUpdateNotify, notify)
}

func (g *Game) AvatarDataNotify() {
	db := g.Player.GetPbPlayerAvatarCompBin()
	notify := &proto.AvatarDataNotify{
		AvatarList:         make([]*proto.AvatarInfo, 0),
		AvatarTeamMap:      make(map[uint32]*proto.AvatarTeam),
		CurAvatarTeamId:    db.CurTeamId,
		ChooseAvatarGuid:   db.ChooseAvatarGuid,
		TempAvatarGuidList: make([]uint64, 0),
		OwnedFlycloakList:  db.OwnedFlycloakList,
	}

	// 获取角色
	avatarList := g.Player.GetPbAvatarList()
	for _, avatar := range avatarList {
		avatarPb := g.PacketAvatarInfo(avatar.AvatarId)
		notify.AvatarList = append(notify.AvatarList, avatarPb)
	}

	// 获取队伍
	teamList := g.Player.GetPbTeamList()
	for id, team := range teamList {
		avatarTeam := &proto.AvatarTeam{
			AvatarGuidList: make([]uint64, 0),
			TeamName:       team.TeamName,
		}
		for _, avatarId := range team.AvatarGuidList {
			avatar := g.Player.GetPbAvatarById(avatarId)
			avatarTeam.AvatarGuidList = append(avatarTeam.AvatarGuidList, avatar.Guid)
		}
		notify.AvatarTeamMap[id] = avatarTeam
	}

	g.seed(cmd.AvatarDataNotify, notify)
}

/*************************************封装******************************************/

func (g *Game) PacketAvatarInfo(avatarId uint32) *proto.AvatarInfo {
	avatar := g.Player.GetPbAvatarById(avatarId)
	pbAvatar := &proto.AvatarInfo{
		IsFocus:  false,
		AvatarId: avatar.AvatarId,
		Guid:     avatar.Guid,
		PropMap: map[uint32]*proto.PropValue{
			uint32(constant.PLAYER_PROP_LEVEL): {
				Type:  uint32(constant.PLAYER_PROP_LEVEL),
				Val:   int64(avatar.Level),
				Value: &proto.PropValue_Ival{Ival: int64(avatar.Level)},
			},
			uint32(constant.PLAYER_PROP_EXP): {
				Type:  uint32(constant.PLAYER_PROP_EXP),
				Val:   int64(avatar.Exp),
				Value: &proto.PropValue_Ival{Ival: int64(avatar.Exp)},
			},
			uint32(constant.PLAYER_PROP_BREAK_LEVEL): {
				Type:  uint32(constant.PLAYER_PROP_BREAK_LEVEL),
				Val:   int64(avatar.PromoteLevel),
				Value: &proto.PropValue_Ival{Ival: int64(avatar.PromoteLevel)},
			},
			uint32(constant.PLAYER_PROP_SATIATION_VAL): {
				Type:  uint32(constant.PLAYER_PROP_SATIATION_VAL),
				Val:   int64(avatar.SatiationVal),
				Value: &proto.PropValue_Ival{Ival: int64(avatar.SatiationVal)},
			},
			uint32(constant.PLAYER_PROP_SATIATION_PENALTY_TIME): {
				Type:  uint32(constant.PLAYER_PROP_SATIATION_PENALTY_TIME),
				Val:   int64(avatar.SatiationPenaltyTime),
				Value: &proto.PropValue_Ival{Ival: int64(avatar.SatiationPenaltyTime)},
			},
		},
		LifeState:     avatar.LifeState,
		EquipGuidList: make([]uint64, 0),
		FightPropMap:  avatar.FightPropMap,
		SkillDepotId:  avatar.SkillDepotId,
		FetterInfo: &proto.AvatarFetterInfo{
			ExpLevel:                avatar.FetterLevel,
			ExpNumber:               avatar.FetterExp,
			FetterList:              nil,
			RewardedFetterLevelList: []uint32{10},
		},
		// SkillLevelMap:          avatar.SkillLevelMap,
		TalentIdList: avatar.TalentIdList,
		// InherentProudSkillList: gdconf.GetAvatarInherentProudSkillList(avatar.SkillDepotId, avatar.PromoteLevel),
		AvatarType:        1,
		WearingFlycloakId: avatar.FlycloakId,
	}
	// 解锁全部资料
	for _, v := range gdconf.GetFetterIdListByAvatarId(int32(avatar.AvatarId)) {
		pbAvatar.FetterInfo.FetterList = append(pbAvatar.FetterInfo.FetterList, &proto.FetterData{
			FetterId:    uint32(v),
			FetterState: constant.FETTER_STATE_FINISH,
		})
	}
	/*
		// 突破等级奖励
		for promoteLevel, isTaken := range avatar.PromoteRewardMap {
			if !isTaken {
				pbAvatar.PendingPromoteRewardList = append(pbAvatar.PendingPromoteRewardList, promoteLevel)
			}
		}
	*/
	return pbAvatar
}
