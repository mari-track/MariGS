package Game

import (
	"github.com/mari-track/MariGS/protocol/cmd"
	"github.com/mari-track/MariGS/protocol/proto"
	pb "google.golang.org/protobuf/proto"
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
	basicDb := g.Player.GetPbPlayerBasicCompBin()
	sceneDb := g.Player.GetPbPlayerSceneCompBin()
	notify := &proto.PlayerEnterSceneNotify{
		SceneId: sceneDb.MyCurSceneId,
		Pos: &proto.Vector{
			X: sceneDb.MyCurPos.X,
			Y: sceneDb.MyCurPos.Y,
			Z: sceneDb.MyCurPos.Z,
		},
		PrevPos: &proto.Vector{
			X: sceneDb.MyCurRot.X,
			Y: sceneDb.MyCurRot.Y,
			Z: sceneDb.MyCurRot.Z,
		},
		SceneBeginTime:         uint64(GetServerTime()),
		Type:                   proto.EnterType_ENTER_SELF,
		TargetUid:              g.Uid,
		PrevSceneId:            0,
		DungeonId:              0,
		WorldLevel:             basicDb.WorldLevel,
		EnterSceneToken:        6070,
		IsFirstLoginEnterScene: true,
	}

	g.seed(cmd.PlayerEnterSceneNotify, notify)
}

// 场景情况通知
func (g *Game) EnterScenePeerNotify() {
	db := g.Player.GetPbPlayerSceneCompBin()
	notify := &proto.EnterScenePeerNotify{
		DestSceneId:     db.MyCurSceneId,
		PeerId:          1,
		HostPeerId:      1,
		EnterSceneToken: 6070,
	}
	g.seed(cmd.EnterScenePeerNotify, notify)
}

func (g *Game) WorldPlayerInfoNotify() {
	notify := &proto.WorldPlayerInfoNotify{
		PlayerInfoList: make([]*proto.OnlinePlayerInfo, 0),
		PlayerUidList:  []uint32{g.Uid},
	}
	// add 本玩家
	myDb := g.Player.GetPbPlayerBasicCompBin()
	notify.PlayerInfoList = append(notify.PlayerInfoList, &proto.OnlinePlayerInfo{
		Uid:                 g.Uid,
		Nickname:            myDb.Nickname,
		PlayerLevel:         myDb.Level,
		AvatarId:            myDb.AvatarId,
		MpSettingType:       proto.MpSettingType_MP_SETTING_ENTER_AFTER_APPLY,
		CurPlayerNumInWorld: 1,
		WorldLevel:          myDb.WorldLevel,
		OnlineId:            "",
		NameCardId:          myDb.NameCardId,
		BlacklistUidList:    nil,
		Signature:           myDb.Signature,
	})

	g.seed(cmd.WorldPlayerInfoNotify, notify)
}

func (g *Game) WorldDataNotify() {
	notify := &proto.WorldDataNotify{
		WorldPropMap: make(map[uint32]*proto.PropValue),
	}
	notify.WorldPropMap[1] = &proto.PropValue{
		Type:  1,
		Val:   0,
		Value: nil,
	}

	g.seed(cmd.WorldDataNotify, notify)
}

func (g *Game) SceneForceUnlockNotify() {
	notify := &proto.SceneForceUnlockNotify{}
	g.seed(cmd.SceneForceUnlockNotify, notify)
}

func (g *Game) HostPlayerNotify() {
	notify := &proto.HostPlayerNotify{
		HostUid:    g.Uid,
		HostPeerId: 1,
	}
	g.seed(cmd.HostPlayerNotify, notify)
}

func (g *Game) SceneTimeNotify() {
	db := g.Player.GetPbPlayerSceneCompBin()
	notify := &proto.SceneTimeNotify{
		SceneId:   db.MyCurSceneId,
		IsPaused:  false,
		SceneTime: uint64(GetServerTime()),
	}
	g.seed(cmd.SceneTimeNotify, notify)
}

func (g *Game) PlayerGameTimeNotify() {
	notify := &proto.PlayerGameTimeNotify{
		GameTime: uint32(GetServerTime()),
		Uid:      g.Uid,
	}
	g.seed(cmd.PlayerGameTimeNotify, notify)
}

func (g *Game) PlayerEnterSceneInfoNotify() {
	avatarList := g.Player.GetPbAvatarList()
	curTeam := g.Player.GetPbCurTeam()
	notify := &proto.PlayerEnterSceneInfoNotify{
		CurAvatarEntityId: 0,
		AvatarEnterInfo:   make([]*proto.AvatarEnterSceneInfo, 0),
		TeamEnterInfo: &proto.TeamEnterSceneInfo{
			TeamEntityId:    uint32(g.GetNextGameObjectGuid()),
			TeamAbilityInfo: nil,
		},
		MpLevelEntityInfo: &proto.MPLevelEntityInfo{
			EntityId:        0,
			AuthorityPeerId: 0,
			AbilityInfo:     nil,
		},
		EnterSceneToken: 6070,
	}
	// 添加角色
	for _, avatarId := range curTeam.AvatarGuidList {
		entityId := uint32(g.GetNextGameObjectGuid())
		avatarDb := avatarList[avatarId]
		if avatarDb == nil {
			continue
		}
		avatarEnterSceneInfo := &proto.AvatarEnterSceneInfo{
			AvatarGuid:        avatarDb.Guid,
			AvatarEntityId:    entityId,
			AvatarAbilityInfo: nil,
			BuffIdList:        nil,
			WeaponGuid:        avatarDb.WeaponGuid,
			WeaponEntityId:    uint32(g.GetNextGameObjectGuid()),
			WeaponAbilityInfo: nil,
			ServerBuffList:    nil,
		}
		notify.AvatarEnterInfo = append(notify.AvatarEnterInfo, avatarEnterSceneInfo)
		if avatarId == curTeam.LastCurAvatarId {
			notify.CurAvatarEntityId = entityId
		}
	}

	g.seed(cmd.PlayerEnterSceneInfoNotify, notify)
}

/*************************************数据包处理******************************************/

func (g *Game) EnterSceneReadyReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.EnterSceneReadyReq)
	rsp := &proto.EnterSceneDoneRsp{
		Retcode:         0,
		EnterSceneToken: req.EnterSceneToken,
	}
	g.EnterScenePeerNotify() // 场景情况通知
	g.seed(cmd.EnterSceneReadyRsp, rsp)
}

func (g *Game) PathfindingEnterSceneReq(payloadMsg pb.Message) {
	// req := payloadMsg.(*proto.PathfindingEnterSceneReq)
	rsp := &proto.PathfindingEnterSceneRsp{}
	g.seed(cmd.PathfindingEnterSceneRsp, rsp)
}

func (g *Game) GetScenePointReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetScenePointReq)
	rsp := &proto.GetScenePointRsp{
		Retcode:                     0,
		SceneId:                     req.SceneId,
		UnlockedPointList:           nil,
		BelongUid:                   0,
		UnlockAreaList:              nil,
		LockedPointList:             nil,
		ToBeExploreDungeonEntryList: nil,
		NotExploredDungeonEntryList: nil,
		GroupUnlimitPointList:       nil,
		NotInteractDungeonEntryList: []uint32{7, 11, 2, 9, 1, 1001, 3, 10, 8, 6},
	}
	g.seed(cmd.GetScenePointRsp, rsp)
}

func (g *Game) GetSceneAreaReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetSceneAreaReq)
	rsp := &proto.GetSceneAreaRsp{
		Retcode:      0,
		SceneId:      req.SceneId,
		AreaIdList:   nil,
		CityInfoList: make([]*proto.CityInfo, 0),
	}
	rsp.CityInfoList = append(rsp.CityInfoList, &proto.CityInfo{
		CityId:     2,
		Level:      1,
		CrystalNum: 0,
	}, &proto.CityInfo{
		CityId:     1,
		Level:      1,
		CrystalNum: 0,
	},
	)

	g.seed(cmd.GetSceneAreaRsp, rsp)
}

func (g *Game) SceneInitFinishReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SceneInitFinishReq)
	g.WorldPlayerInfoNotify()
	g.ServerTimeNotify()
	// WorldOwnerDailyTaskNotify p1玩家每日任务通知
	// WorldOwnerBlossomBriefInfoNotify p1玩家世界
	g.WorldDataNotify() // 世界数据通知
	g.SceneForceUnlockNotify()
	g.HostPlayerNotify()
	g.SceneTimeNotify()
	g.PlayerGameTimeNotify()
	g.PlayerEnterSceneInfoNotify() // 实体通知
	g.ScenePlayerInfoNotify()
	/*
		SceneForceUnlockNotify
	*/
	rsp := &proto.SceneInitFinishRsp{
		Retcode:         0,
		EnterSceneToken: req.EnterSceneToken,
	}
	g.seed(cmd.SceneInitFinishRsp, rsp)
}

func (g *Game) EnterSceneDoneReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.EnterSceneDoneReq)
	rsp := &proto.EnterSceneDoneRsp{
		Retcode:         0,
		EnterSceneToken: req.EnterSceneToken,
	}
	g.seed(cmd.EnterSceneDoneRsp, rsp)
}

func (g *Game) PostEnterSceneReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.PostEnterSceneReq)
	rsp := &proto.PostEnterSceneRsp{
		Retcode:         0,
		EnterSceneToken: req.EnterSceneToken,
	}
	g.seed(cmd.PostEnterSceneRsp, rsp)
}
