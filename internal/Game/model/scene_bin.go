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
			Y: 307,
			Z: 0,
		}
	}
	return db.SceneBin
}

func (p *Player) GetPbMyCurPos() *playerPb.VectorBin {
	bin := p.GetPbPlayerSceneCompBin()
	if bin.MyCurPos == nil {
		bin.MyCurPos = &playerPb.VectorBin{
			X: 2747.562,
			Y: 194.633,
			Z: -1719.386,
		}
	}
	return bin.MyCurPos
}

func (p *Player) GetPbMyCurRot() *playerPb.VectorBin {
	bin := p.GetPbPlayerSceneCompBin()
	if bin.MyCurRot == nil {
		bin.MyCurRot = &playerPb.VectorBin{
			X: 0,
			Y: 307,
			Z: 0,
		}
	}
	return bin.MyCurPos
}
