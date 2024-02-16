package model

import (
	"github.com/mari-track/MariGS/pkg/alg"
	"github.com/mari-track/MariGS/pkg/logger"
	playerPb "github.com/mari-track/MariGS/protocol/server_only"
	pb "google.golang.org/protobuf/proto"
)

var SNOWFLAKE *alg.SnowflakeWorker

type Player struct {
	OnlineData *OnlineData             // 在线数据
	PlayerPb   *playerPb.PlayerDataBin // 离线数据只读,禁止修改
}

type OnlineData struct {
	avatar *Avatar // 角色配置
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
		// 进入玩家数据预处理
		player.newOnlineData()
		if player.OnlineData == nil {
			logger.Error("onlinePlayerData unmarshal error")
			return nil
		}
	}

	return player
}

// 生成在线数据
func (p *Player) newOnlineData() {
	player := new(OnlineData)
	if p.PlayerPb == nil {
		return
	}
	player.avatar = p.newOnlineAvatar()

	p.OnlineData = player
}

func (p *Player) GetOnlineData() *OnlineData {
	return p.OnlineData
}

func NewPlayerPb() *playerPb.PlayerDataBin {
	PlayerPb := &playerPb.PlayerDataBin{
		BasicBin:  new(playerPb.PlayerBasicCompBin),
		AvatarBin: new(playerPb.PlayerAvatarCompBin),
	}
	return PlayerPb
}

func (p *Player) GetPbPlayer() *playerPb.PlayerDataBin {
	if p.PlayerPb == nil {
		p.PlayerPb = NewPlayerPb()
	}
	return p.PlayerPb
}
