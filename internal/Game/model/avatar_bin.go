package model

import (
	"github.com/mari-track/MariGS/pkg/constant"
	"github.com/mari-track/MariGS/pkg/gdconf"
	playerPb "github.com/mari-track/MariGS/protocol/server_only"
)

type Avatar struct {
	avatarList map[uint32]*playerPb.AvatarBin // id/bin
}

func NewPlayerAvatarCompBin() *playerPb.PlayerAvatarCompBin {
	bin := &playerPb.PlayerAvatarCompBin{
		AvatarList:                     make([]*playerPb.AvatarBin, 0),
		CurAvatarGuid:                  0,
		PbOnlyCurPos:                   new(playerPb.VectorBin),
		PbOnlyCurRot:                   new(playerPb.VectorBin),
		TeamMap:                        make(map[uint32]*playerPb.AvatarTeamBin),
		CurTeamId:                      0,
		LastChangeAvatarTime:           0,
		IsSpringAutoUse:                false,
		SpringAutoUsePercent:           0,
		LastBackgroudAvatarRecoverTime: 0,
		CurSpringVolume:                0,
		IsFlyable:                      false,
		IsTransferable:                 false,
		LastAvatarGuid:                 0,
		ChooseAvatarGuid:               0,
		DieType:                        0,
		TempAvatarGuidList:             make([]uint64, 0),
		AvatarTeamBuffMap:              make(map[uint32]*playerPb.AvatarBuffBin),
		LastServerBuffUid:              0,
		TotalExpeditionNum:             0,
		OwnedFlycloakList:              make([]uint32, 0),
	}

	return bin
}

func (p *Player) newOnlineAvatar() *Avatar {
	avatar := new(Avatar)
	avatar.avatarList = make(map[uint32]*playerPb.AvatarBin)
	for _, avatarList := range p.GetPbAvatarList() {
		avatar.avatarList[avatarList.AvatarId] = avatarList
	}
	return avatar
}

func (p *Player) GetOnlineAvatar() *Avatar {
	onlineData := p.GetOnlineData()
	if onlineData.avatar == nil {
		onlineData.avatar = new(Avatar)
	}
	return onlineData.avatar
}

func (p *Player) GetOnlineAvatarList() map[uint32]*playerPb.AvatarBin {
	avatar := p.GetOnlineAvatar()
	if avatar.avatarList == nil {
		avatar.avatarList = make(map[uint32]*playerPb.AvatarBin)
	}
	return avatar.avatarList
}

func (p *Player) GetPbPlayerAvatarCompBin() *playerPb.PlayerAvatarCompBin {
	if p.GetPbPlayer().AvatarBin == nil {
		p.GetPbPlayer().AvatarBin = NewPlayerAvatarCompBin()
	}
	return p.GetPbPlayer().AvatarBin
}

func (p *Player) GetPbAvatarList() []*playerPb.AvatarBin {
	avatarBin := p.GetPbPlayerAvatarCompBin()
	if avatarBin.AvatarList == nil {
		avatarBin.AvatarList = make([]*playerPb.AvatarBin, 0)
	}

	return avatarBin.AvatarList
}

func (p *Player) AddAvatar(avatarId uint32) {
	avatarDb := p.GetOnlineAvatarList()
	if avatarDb[avatarId] != nil {

	} else {
		// 获取角色基础配置
		conf := gdconf.GetAvatarDataById(avatarId)
		if conf == nil {
			return
		}
		// 添加角色
		avatarGuid := uint64(SNOWFLAKE.GenId())
		avatarPb := &playerPb.AvatarBin{
			Detail:               nil,
			AvatarType:           0, // 类型
			AvatarId:             avatarId,
			Guid:                 avatarGuid,                // 唯一id
			Level:                1,                         // 等级
			LifeState:            constant.LIFE_STATE_ALIVE, // 状态
			CurHp:                conf.HpBase,               // 当前血量
			CurElemEnergy:        0,                         // 当前元素能量
			PromoteLevel:         0,
			SkillDepotId:         0,                                         // 技能库id
			SkillMap:             make(map[uint32]*playerPb.AvatarSkillBin), // 技能
			BuffMap:              nil,                                       // buff
			DepotMap:             nil,
			EquipList:            make([]*playerPb.ItemBin, 0),             // 装备列表
			SatiationVal:         0,                                        // 饱食度
			SatiationPenaltyTime: 0,                                        // 饱食度惩罚时间
			FlycloakId:           0,                                        // 风之翼id
			AvatarEquipAffixList: make([]*playerPb.AvatarEquipAffixBin, 0), // 圣遗物主词条列表
			BornTime:             0,                                        // 获取时间
		}
		avatarDb[avatarId] = avatarPb
		// 添加角色对应武器
	}
}
