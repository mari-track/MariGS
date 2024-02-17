package model

import (
	"time"

	"github.com/mari-track/MariGS/pkg/constant"
	"github.com/mari-track/MariGS/pkg/gdconf"
	"github.com/mari-track/MariGS/pkg/logger"
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
			AvatarType:    0,                                         // 类型
			AvatarId:      avatarId,                                  // id
			Guid:          avatarGuid,                                // 唯一id
			Level:         1,                                         // 等级
			LifeState:     constant.LIFE_STATE_ALIVE,                 // 状态
			CurHp:         conf.HpBase,                               // 当前血量
			CurElemEnergy: 0,                                         // 当前元素能量
			PromoteLevel:  0,                                         // 突破等级
			SkillDepotId:  0,                                         // 技能库id
			SkillMap:      make(map[uint32]*playerPb.AvatarSkillBin), // 技能
			// BuffMap:              make(map[uint32]*playerPb.AvatarBuffBin),       // buff
			// DepotMap:             make(map[uint32]*playerPb.AvatarSkillDepotBin), // ?
			WeaponGuid:           0,                         // 装备的武器唯一id
			SatiationVal:         0,                         // 饱食度
			SatiationPenaltyTime: 0,                         // 饱食度惩罚时间
			FlycloakId:           140001,                    // 风之翼id
			BornTime:             uint32(time.Now().Unix()), // 获取时间
			FightPropMap:         make(map[uint32]float32),
			Exp:                  0,
			FetterLevel:          0,
			FetterExp:            0,
			TalentIdList:         make([]uint32, 0),
		}
		// 获取技能

		avatarDb[avatarId] = avatarPb
		// 更新面板
		p.UpdateAvatarFightProp(avatarId)
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
func (p *Player) UpdateAvatarFightProp(avatarId uint32) {
	db := p.GetPbAvatarById(avatarId)
	if db == nil {
		return
	}
	conf := gdconf.GetAvatarDataById(avatarId)
	if conf == nil {
		logger.Error("conf error, avatarId: %v", avatarId)
		return
	}
	db.FightPropMap[constant.FIGHT_PROP_NONE] = 0.0
	// 白字攻防血
	db.FightPropMap[constant.FIGHT_PROP_BASE_ATTACK] = conf.GetBaseAttackByLevel(db.Level)
	db.FightPropMap[constant.FIGHT_PROP_BASE_DEFENSE] = conf.GetBaseDefenseByLevel(db.Level)
	db.FightPropMap[constant.FIGHT_PROP_BASE_HP] = conf.GetBaseHpByLevel(db.Level)
	// 白字+绿字攻防血
	db.FightPropMap[constant.FIGHT_PROP_CUR_ATTACK] = conf.GetBaseAttackByLevel(db.Level)
	db.FightPropMap[constant.FIGHT_PROP_CUR_DEFENSE] = conf.GetBaseDefenseByLevel(db.Level)
	db.FightPropMap[constant.FIGHT_PROP_MAX_HP] = conf.GetBaseHpByLevel(db.Level)
	// 双暴
	db.FightPropMap[constant.FIGHT_PROP_CRITICAL] = conf.Critical
	db.FightPropMap[constant.FIGHT_PROP_CRITICAL_HURT] = conf.CriticalHurt
	// 元素充能
	db.FightPropMap[constant.FIGHT_PROP_CHARGE_EFFICIENCY] = 1.0
}

func (p *Player) GetPbTeamList() map[uint32]*playerPb.AvatarTeamBin {
	db := p.GetPbPlayerAvatarCompBin()
	if db.TeamMap == nil {
		db.TeamMap = make(map[uint32]*playerPb.AvatarTeamBin)
		teamBin := &playerPb.AvatarTeamBin{
			AvatarGuidList:  make([]uint32, 0),
			TeamName:        "",
			LastCurAvatarId: 0,
		}
		db.TeamMap[1] = teamBin
		db.TeamMap[2] = teamBin
		db.TeamMap[3] = teamBin
		db.TeamMap[4] = teamBin
	}
	return db.TeamMap
}

func (p *Player) GetPbTeamById(teamId uint32) *playerPb.AvatarTeamBin {
	db := p.GetPbTeamList()
	if db[teamId] == nil {
		db[teamId] = &playerPb.AvatarTeamBin{
			AvatarGuidList:  make([]uint32, 0),
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
