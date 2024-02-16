package model

import (
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
		RegisterCps:          "",
	}
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
