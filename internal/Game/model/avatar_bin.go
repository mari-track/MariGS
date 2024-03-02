package model

import (
	"time"

	"github.com/mari-track/MariGS/pkg/constant"
	"github.com/mari-track/MariGS/pkg/endec"
	"github.com/mari-track/MariGS/pkg/gdconf"
	"github.com/mari-track/MariGS/pkg/logger"
	"github.com/mari-track/MariGS/protocol/proto"
	playerPb "github.com/mari-track/MariGS/protocol/server_only"
)

func (p *Player) GetPbPlayerAvatarCompBin() *playerPb.PlayerAvatarCompBin {
	if p.GetPbPlayer().AvatarBin == nil {
		p.GetPbPlayer().AvatarBin = &playerPb.PlayerAvatarCompBin{}
	}
	return p.GetPbPlayer().AvatarBin
}

func (p *Player) GetPbAvatarList() map[uint32]*playerPb.AvatarBin {
	avatarBin := p.GetPbPlayerAvatarCompBin()
	if avatarBin.AvatarList == nil {
		avatarBin.AvatarList = make(map[uint32]*playerPb.AvatarBin)
	}
	return avatarBin.AvatarList
}

func (p *Player) GetPbAvatarById(avatarId uint32) *playerPb.AvatarBin {
	db := p.GetPbAvatarList()
	return db[avatarId]
}

func (p *Player) AddAvatar(avatarId uint32) {
	avatarDb := p.GetPbAvatarList()
	if avatarDb[avatarId] != nil {
		logger.Error("重复添加角色")
		return
	} else {
		// 获取角色基础配置
		conf := gdconf.GetAvatarDataById(avatarId)
		if conf == nil {
			return
		}
		// 添加角色
		avatarGuid := uint64(SNOWFLAKE.GenId())
		avatarPb := &playerPb.AvatarBin{
			AvatarType:    0,                         // 类型
			AvatarId:      avatarId,                  // id
			Guid:          avatarGuid,                // 唯一id
			Level:         1,                         // 等级
			LifeState:     constant.LIFE_STATE_ALIVE, // 状态
			CurHp:         conf.HpBase,               // 当前血量
			CurElemEnergy: 0,                         // 当前元素能量
			PromoteLevel:  0,                         // 突破等级
			SkillDepotId:  0,                         // 技能库id
			SkillMap:      make(map[uint32]uint32),   // 技能
			// BuffMap:              make(map[uint32]*playerPb.AvatarBuffBin),       // buff
			// DepotMap:             make(map[uint32]*playerPb.AvatarSkillDepotBin), // ?
			WeaponGuid:           0,                         // 装备的武器唯一id
			SatiationVal:         0,                         // 饱食度
			SatiationPenaltyTime: 0,                         // 饱食度惩罚时间
			FlycloakId:           140001,                    // 风之翼id
			BornTime:             uint32(time.Now().Unix()), // 获取时间
			FightPropMap:         make(map[uint32]float32),
			Exp:                  0,
			FetterLevel:          1,
			FetterExp:            0,
			TalentIdList:         make([]uint32, 0),
		}
		// 获取技能
		switch avatarId {
		case 10000007:
			avatarPb.SkillDepotId = 701
		case 10000005:
			avatarPb.SkillDepotId = 501
		default:
			avatarPb.SkillDepotId = conf.SkillDepotId
		}
		avatarSkillDepotDataConfig := gdconf.GetAvatarSkillDepotDataById(avatarPb.SkillDepotId)
		if avatarSkillDepotDataConfig == nil {
			logger.Error("avatar skill depot data config is nil, skillDepotId: %v", avatarPb.SkillDepotId)
			return
		}
		// 元素爆发
		_, exist := avatarPb.SkillMap[uint32(avatarSkillDepotDataConfig.EnergySkill)]
		if !exist && avatarSkillDepotDataConfig.EnergySkill != 0 {
			avatarPb.SkillMap[uint32(avatarSkillDepotDataConfig.EnergySkill)] = 1
		}
		for _, skillId := range avatarSkillDepotDataConfig.Skills {
			if skillId == 0 {
				continue
			}
			// 小技能
			_, exist = avatarPb.SkillMap[uint32(skillId)]
			if !exist {
				avatarPb.SkillMap[uint32(skillId)] = 1
			}
		}
		// 更新面板
		avatarPb.FightPropMap = p.UpdateAvatarFightProp(avatarId, avatarPb.Level)
		avatarDb[avatarId] = avatarPb
	}
}

func (p *Player) UpdateAvatarWeapon(guid uint64, avatarId uint32) {
	db := p.GetPbAvatarById(avatarId)
	if db == nil {
		return
	}
	db.WeaponGuid = guid
}

// UpdateAvatarFightProp 更新角色面板
func (p *Player) UpdateAvatarFightProp(avatarId, level uint32) map[uint32]float32 {
	fightPropMap := make(map[uint32]float32, 0)
	conf := gdconf.GetAvatarDataById(avatarId)
	if conf == nil {
		logger.Error("conf error, avatarId: %v", avatarId)
		return nil
	}
	fightPropMap[constant.FIGHT_PROP_NONE] = 0.0
	// 白字攻防血
	fightPropMap[constant.FIGHT_PROP_BASE_ATTACK] = conf.GetBaseAttackByLevel(level)
	fightPropMap[constant.FIGHT_PROP_BASE_DEFENSE] = conf.GetBaseDefenseByLevel(level)
	fightPropMap[constant.FIGHT_PROP_BASE_HP] = conf.GetBaseHpByLevel(level)
	fightPropMap[constant.FIGHT_PROP_CUR_HP] = conf.GetBaseHpByLevel(level)
	// 白字+绿字攻防血
	fightPropMap[constant.FIGHT_PROP_CUR_ATTACK] = conf.GetBaseAttackByLevel(level)
	fightPropMap[constant.FIGHT_PROP_CUR_DEFENSE] = conf.GetBaseDefenseByLevel(level)
	fightPropMap[constant.FIGHT_PROP_MAX_HP] = conf.GetBaseHpByLevel(level)
	// 双暴
	fightPropMap[constant.FIGHT_PROP_CRITICAL] = conf.Critical
	fightPropMap[constant.FIGHT_PROP_CRITICAL_HURT] = conf.CriticalHurt
	// 元素充能
	fightPropMap[constant.FIGHT_PROP_CHARGE_EFFICIENCY] = 1.0

	return fightPropMap
}

func PacketAvatarAbilityControlBlock(avatarId uint32, skillDepotId uint32) *proto.AbilityControlBlock {
	acb := &proto.AbilityControlBlock{
		AbilityEmbryoList: make([]*proto.AbilityEmbryo, 0),
	}
	abilityId := 0
	// 默认ability
	for _, abilityHashCode := range constant.DEFAULT_ABILITY_HASH_CODE {
		abilityId++
		ae := &proto.AbilityEmbryo{
			AbilityId:               uint32(abilityId),
			AbilityNameHash:         uint32(abilityHashCode),
			AbilityOverrideNameHash: uint32(endec.Hk4eAbilityHashCode("Default")),
		}
		acb.AbilityEmbryoList = append(acb.AbilityEmbryoList, ae)
	}
	// 角色ability
	avatarDataConfig := gdconf.GetAvatarDataById(avatarId)
	if avatarDataConfig != nil {
		for _, abilityHashCode := range avatarDataConfig.AbilityHashCodeList {
			abilityId++
			ae := &proto.AbilityEmbryo{
				AbilityId:               uint32(abilityId),
				AbilityNameHash:         uint32(abilityHashCode),
				AbilityOverrideNameHash: uint32(endec.Hk4eAbilityHashCode("Default")),
			}
			acb.AbilityEmbryoList = append(acb.AbilityEmbryoList, ae)
		}
	}
	// 技能库ability
	skillDepot := gdconf.GetAvatarSkillDepotDataById(skillDepotId)
	if skillDepot != nil && len(skillDepot.AbilityHashCodeList) != 0 {
		for _, abilityHashCode := range skillDepot.AbilityHashCodeList {
			abilityId++
			ae := &proto.AbilityEmbryo{
				AbilityId:               uint32(abilityId),
				AbilityNameHash:         uint32(abilityHashCode),
				AbilityOverrideNameHash: uint32(endec.Hk4eAbilityHashCode("Default")),
			}
			acb.AbilityEmbryoList = append(acb.AbilityEmbryoList, ae)
		}
	}
	// TODO 队伍ability
	// TODO 装备ability
	return acb
}

func (p *Player) GetPbTeamList() map[uint32]*playerPb.AvatarTeamBin {
	db := p.GetPbPlayerAvatarCompBin()
	if db.TeamMap == nil {
		db.TeamMap = make(map[uint32]*playerPb.AvatarTeamBin)
	}
	return db.TeamMap
}

func (p *Player) GetPbTeamById(teamId uint32) *playerPb.AvatarTeamBin {
	db := p.GetPbTeamList()
	if db[teamId] == nil {
		db[teamId] = &playerPb.AvatarTeamBin{
			AvatarIdList:    make([]uint32, 0),
			TeamName:        "",
			LastCurAvatarId: 0,
		}
	}
	return db[teamId]
}

func (p *Player) GetPbCurTeam() *playerPb.AvatarTeamBin {
	db := p.GetPbPlayerAvatarCompBin()
	curTeam := p.GetPbTeamById(db.CurTeamId)

	return curTeam
}
