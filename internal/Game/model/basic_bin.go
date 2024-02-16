package model

import (
	"github.com/mari-track/MariGS/pkg/constant"
	playerPb "github.com/mari-track/MariGS/protocol/server_only"
)

func NewPlayerBasicCompBin() *playerPb.PlayerBasicCompBin {
	bin := &playerPb.PlayerBasicCompBin{
		Level:                1,
		Exp:                  0,
		Nickname:             "",
		TotalGameTime:        0,
		TransNoCount:         0,
		IsWeatherLocked:      false,
		IsGameTimeLocked:     false,
		PersistStaminaLimit:  0,
		CurPersistStamina:    0,
		CurTemporaryStamina:  0,
		OpenStateMap:         make(map[uint32]uint32),
		PlayerTimeMs:         0,
		LastLoginTime:        0,
		RegisterTime:         0,
		TotalLoginDays:       0,
		UpdateLoginDaysTime:  0,
		RewardTakenLevelList: make([]uint32, 0),
		LanguageType:         0,
		ClientAppVersion:     "",
		ClientDeviceInfo:     "",
		ClientSystemVersion:  "",
		IsGuest:              false,
		OnlineTime:           0,
		IsProficientPlayer:   false,
		SetLanguageTag:       0,
		HeadImageAvatarId:    0,
		GuidSeqId:            0,
		IsRebateMailSent:     false,
		IsRebateMailReceived: false,
		PropMap:              make(map[uint32]uint32),
	}

	bin.PropMap[constant.PLAYER_PROP_PLAYER_WORLD_LEVEL] = 0
	bin.PropMap[constant.PLAYER_PROP_CUR_PERSIST_STAMINA] = 10000
	bin.PropMap[constant.PLAYER_PROP_CUR_TEMPORARY_STAMINA] = 0
	bin.PropMap[constant.PLAYER_PROP_PLAYER_LEVEL] = 1
	bin.PropMap[constant.PLAYER_PROP_PLAYER_EXP] = 0
	bin.PropMap[constant.PLAYER_PROP_PLAYER_HCOIN] = 0
	bin.PropMap[constant.PLAYER_PROP_PLAYER_SCOIN] = 0
	bin.PropMap[constant.PLAYER_PROP_PLAYER_MP_SETTING_TYPE] = 2
	bin.PropMap[constant.PLAYER_PROP_PLAYER_RESIN] = 160
	bin.PropMap[constant.PLAYER_PROP_PLAYER_WAIT_SUB_HCOIN] = 0
	bin.PropMap[constant.PLAYER_PROP_PLAYER_WAIT_SUB_SCOIN] = 0
	bin.PropMap[constant.PLAYER_PROP_PLAYER_MCOIN] = 0
	bin.PropMap[constant.PLAYER_PROP_PLAYER_WAIT_SUB_MCOIN] = 0
	bin.PropMap[constant.PLAYER_PROP_PLAYER_LEGENDARY_KEY] = 0
	bin.PropMap[constant.PLAYER_PROP_CUR_CLIMATE_METER] = 0
	bin.PropMap[constant.PLAYER_PROP_CUR_CLIMATE_TYPE] = 0
	bin.PropMap[constant.PLAYER_PROP_CUR_CLIMATE_AREA_ID] = 0
	bin.PropMap[constant.PLAYER_PROP_CUR_CLIMATE_AREA_CLIMATE_TYPE] = 0
	bin.PropMap[constant.PLAYER_PROP_PLAYER_WORLD_LEVEL_LIMIT] = 0
	bin.PropMap[constant.PLAYER_PROP_PLAYER_WORLD_LEVEL_ADJUST_CD] = 0
	bin.PropMap[constant.PLAYER_PROP_PLAYER_LEGENDARY_DAILY_TASK_NUM] = 0
	bin.PropMap[constant.PLAYER_PROP_PLAYER_HOME_COIN] = 0
	bin.PropMap[constant.PLAYER_PROP_PLAYER_WAIT_SUB_HOME_COIN] = 0
	bin.PropMap[constant.PLAYER_PROP_IS_AUTO_UNLOCK_SPECIFIC_EQUIP] = 0
	bin.PropMap[constant.PLAYER_PROP_IS_SPRING_AUTO_USE] = 1
	bin.PropMap[constant.PLAYER_PROP_SPRING_AUTO_USE_PERCENT] = 50
	bin.PropMap[constant.PLAYER_PROP_IS_FLYABLE] = 0
	bin.PropMap[constant.PLAYER_PROP_IS_WEATHER_LOCKED] = 1
	bin.PropMap[constant.PLAYER_PROP_IS_GAME_TIME_LOCKED] = 1
	bin.PropMap[constant.PLAYER_PROP_IS_TRANSFERABLE] = 1
	bin.PropMap[constant.PLAYER_PROP_MAX_STAMINA] = 10000

	return bin
}

func (p *Player) GetPbPlayerBasicCompBin() *playerPb.PlayerBasicCompBin {
	if p.GetPbPlayer().BasicBin == nil {
		p.GetPbPlayer().BasicBin = NewPlayerBasicCompBin()
	}
	return p.GetPbPlayer().BasicBin
}

func (p *Player) UptoDateNickname(nickname string) {
	playerBasicCompBin := p.GetPbPlayerBasicCompBin()
	playerBasicCompBin.Nickname = nickname
}
