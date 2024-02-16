package Game

import (
	"github.com/mari-track/MariGS/protocol/cmd"
	"github.com/mari-track/MariGS/protocol/proto"
)

func (g *Game) QuestListNotify() {
	notify := &proto.QuestListNotify{
		QuestList: make([]*proto.Quest, 0),
	}
	g.seed(cmd.QuestListNotify, notify)
}
