package model

import (
	"github.com/mari-track/MariGS/pkg/alg"
	"github.com/mari-track/MariGS/pkg/logger"
	playerPb "github.com/mari-track/MariGS/protocol/server_only"
	pb "google.golang.org/protobuf/proto"
)

var SNOWFLAKE *alg.SnowflakeWorker

type Player struct {
	PlayerPb *playerPb.PlayerDataBin
}

func NewPlayer(bin []byte) *Player {
	var err error
	player := new(Player)
	player.PlayerPb = NewPlayerPb()
	if bin == nil {

	} else {
		// 解析数据
		err = pb.Unmarshal(bin, player.PlayerPb)
		if err != nil {
			logger.Error("player_pb unmarshal error")
			return nil
		}
	}

	return player
}

func NewPlayerPb() *playerPb.PlayerDataBin {
	PlayerPb := new(playerPb.PlayerDataBin)
	return PlayerPb
}

func (p *Player) GetPbPlayer() *playerPb.PlayerDataBin {
	if p.PlayerPb == nil {
		p.PlayerPb = NewPlayerPb()
	}
	return p.PlayerPb
}
