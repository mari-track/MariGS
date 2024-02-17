package model

import (
	playerPb "github.com/mari-track/MariGS/protocol/server_only"
)

func (p *Player) GetPbPlayerSceneCompBin() *playerPb.PlayerSceneCompBin {
	db := p.GetPbPlayer()
	if db.SceneBin == nil {
		db.SceneBin = new(playerPb.PlayerSceneCompBin)
		db.SceneBin.MyCurSceneId = 3
		db.SceneBin.MyCurPos = &playerPb.VectorBin{
			X: 2747.562,
			Y: 194.633,
			Z: -1719.386,
		}
		db.SceneBin.MyCurRot = &playerPb.VectorBin{
			X: 0,
			Y: 0,
			Z: 0,
		}
	}
	return db.SceneBin
}
