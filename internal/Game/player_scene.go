package Game

import (
	"github.com/mari-track/MariGS/protocol/cmd"
	"github.com/mari-track/MariGS/protocol/proto"
)

/*************************************通知包******************************************/

func (g *Game) AllSeenMonsterNotify() {
	notify := &proto.AllSeenMonsterNotify{
		MonsterIdList: make([]uint32, 0),
	}

	g.seed(cmd.AllSeenMonsterNotify, notify)
}

func (g *Game) PlayerEnterSceneNotify() {
	// TODO
	db := g.Player.GetPbPlayerBasicCompBin()
	notify := &proto.PlayerEnterSceneNotify{
		SceneId: 3,
		Pos: &proto.Vector{
			X: 2747.562,
			Y: 194.633,
			Z: -1719.386,
		},
		PrevPos: &proto.Vector{
			X: 0,
			Y: 0,
			Z: 0,
		},
		SceneBeginTime:         uint64(GetServerTime()),
		Type:                   proto.EnterType_ENTER_SELF,
		TargetUid:              g.Uid,
		PrevSceneId:            0,
		DungeonId:              0,
		WorldLevel:             db.WorldLevel,
		EnterSceneToken:        6070,
		IsFirstLoginEnterScene: true,
	}

	g.seed(cmd.PlayerEnterSceneNotify, notify)
}
