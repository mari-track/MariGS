package model

import (
	playerPb "github.com/mari-track/MariGS/protocol/server_only"
)

func NewPlayerItemCompBin() *playerPb.PlayerItemCompBin {
	bin := new(playerPb.PlayerItemCompBin)
	return bin
}

func (p *Player) GetPbPlayerItemCompBin() *playerPb.PlayerItemCompBin {
	db := p.GetPbPlayer()
	if db.ItemBin == nil {
		db.ItemBin = NewPlayerItemCompBin()
	}

	return db.ItemBin
}

func (p *Player) GetPbWeaponList() map[uint64]*playerPb.WeaponBin {
	itemBin := p.GetPbPlayerItemCompBin()
	if itemBin.WeaponList == nil {
		itemBin.WeaponList = make(map[uint64]*playerPb.WeaponBin)
	}
	return itemBin.WeaponList
}

func (p *Player) GetPbWeaponByGuid(guid uint64) *playerPb.WeaponBin {
	weaponList := p.GetPbWeaponList()
	return weaponList[guid]
}

func (p *Player) AddWeapon(guid uint64, weaponId uint32) {
	db := p.GetPbWeaponList()
	db[guid] = &playerPb.WeaponBin{
		Level:        1,
		Exp:          0,
		PromoteLevel: 0,
		AffixMap:     make(map[uint32]uint32),
		Guid:         guid,
		IsLock:       false,
		Refinement:   0,
		AvatarId:     0,
		WeaponId:     weaponId,
	}
}

func (p *Player) UpdateWeaponAvatar(guid uint64, avatarId uint32) {
	db := p.GetPbWeaponByGuid(guid)
	if db == nil {
		return
	}
	db.AvatarId = avatarId
}
