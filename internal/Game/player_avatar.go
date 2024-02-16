package Game

import (
	"github.com/mari-track/MariGS/protocol/cmd"
	"github.com/mari-track/MariGS/protocol/proto"
)

/*************************************通知包******************************************/

func (g *Game) AvatarFightPropUpdateNotify(avatarId uint32) {
	db := g.Player.GetPbAvatarById(avatarId)
	notify := &proto.AvatarFightPropUpdateNotify{
		AvatarGuid:   db.Guid,
		FightPropMap: db.FightPropMap,
	}

	g.seed(cmd.AvatarFightPropUpdateNotify, notify)
}
