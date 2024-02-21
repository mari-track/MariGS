package Game

import (
	"time"

	"github.com/mari-track/MariGS/internal/Game/sceneEntity"
	"github.com/mari-track/MariGS/pkg/constant"
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
	// 设置场景token
	g.SceneToken = uint32(g.GetNextGameObjectGuid())
	notify := &proto.PlayerEnterSceneNotify{
		SceneId: sceneDb.MyCurSceneId,
		Pos: &proto.Vector{
			X: sceneDb.MyCurPos.X,
			Y: sceneDb.MyCurPos.Y,
			Z: sceneDb.MyCurPos.Z,
		},
		SceneBeginTime:         uint64(GetServerTime()),
		Type:                   proto.EnterType_ENTER_SELF,
		TargetUid:              g.Uid,
		PrevSceneId:            0,
		DungeonId:              0,
		WorldLevel:             basicDb.WorldLevel,
		EnterSceneToken:        g.SceneToken,
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
		EnterSceneToken: g.SceneToken,
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
		Value: &proto.PropValue_Ival{Ival: 0},
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

func (g *Game) SceneDataNotify() {
	notify := &proto.SceneDataNotify{
		LevelConfigNameList: []string{"Level_BigWorld"},
	}
	g.seed(cmd.SceneDataNotify, notify)
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
			TeamEntityId:    g.SceneEntity.GetNextWorldEntityId(constant.ENTITY_TYPE_TEAM),
			TeamAbilityInfo: &proto.AbilitySyncStateInfo{IsInited: false},
		},
		MpLevelEntityInfo: &proto.MPLevelEntityInfo{
			EntityId:        g.SceneEntity.GetNextWorldEntityId(constant.ENTITY_TYPE_MP_LEVEL),
			AuthorityPeerId: g.Uid,
			AbilityInfo: &proto.AbilitySyncStateInfo{
				IsInited: false,
			},
		},
		EnterSceneToken: g.SceneToken,
	}
	// 添加角色
	avatarEntity := g.SceneEntity.GetAvatarEntity()
	for _, avatarId := range curTeam.AvatarIdList {
		avatarEntityId := g.SceneEntity.GetNextWorldEntityId(constant.ENTITY_TYPE_AVATAR)
		weaponEntityId := g.SceneEntity.GetNextWorldEntityId(constant.ENTITY_TYPE_WEAPON)
		avatarDb := avatarList[avatarId]
		if avatarDb == nil {
			continue
		}
		avatarEnterSceneInfo := &proto.AvatarEnterSceneInfo{
			AvatarGuid:        avatarDb.Guid,
			AvatarEntityId:    avatarEntityId,
			AvatarAbilityInfo: &proto.AbilitySyncStateInfo{IsInited: false},
			BuffIdList:        nil,
			WeaponGuid:        avatarDb.WeaponGuid,
			WeaponEntityId:    weaponEntityId,
			WeaponAbilityInfo: &proto.AbilitySyncStateInfo{IsInited: false},
			ServerBuffList:    nil,
		}
		notify.AvatarEnterInfo = append(notify.AvatarEnterInfo, avatarEnterSceneInfo)
		if avatarId == curTeam.LastCurAvatarId {
			notify.CurAvatarEntityId = avatarEntityId
		}
		// 添加avatar实体
		avatarEntity[avatarEntityId] = &sceneEntity.AvatarEntity{
			AvatarId:       avatarId,
			WeaponEntityId: weaponEntityId,
		}
	}

	g.seed(cmd.PlayerEnterSceneInfoNotify, notify)
}

func (g *Game) SceneEntityAppearNotify() {
	notify := &proto.SceneEntityAppearNotify{
		EntityList: make([]*proto.SceneEntityInfo, 0),
		AppearType: proto.VisionType_VISION_MEET,
		Param:      0,
	}
	// add avatar
	avatarEntity := g.SceneEntity.GetAvatarEntity()
	for entityId, entity := range avatarEntity {
		avatarPb := g.Player.GetPbAvatarById(entity.AvatarId)
		weaponPb := g.Player.GetPbWeaponByGuid(avatarPb.WeaponGuid)
		sceneEntityInfo := &proto.SceneEntityInfo{
			AbilityInfo: &proto.AbilitySyncStateInfo{IsInited: false},
			AiInfo: &proto.SceneEntityAiInfo{
				IsAiOpen: true,
				BornPos: &proto.Vector{
					X: 0,
					Y: 0,
					Z: 0,
				},
				SkillCdMap:  nil,
				ServantInfo: nil,
				AiThreatMap: nil,
			},
			AnimatorParaList: make([]*proto.AnimatorParameterValueInfoPair, 0),
			Entity: &proto.SceneEntityInfo_Avatar{Avatar: &proto.SceneAvatarInfo{
				AvatarId:      entity.AvatarId,
				BornTime:      uint32(time.Now().Unix()),
				EquipIdList:   []uint32{11101},
				Guid:          avatarPb.Guid,
				PeerId:        1,
				SkillDepotId:  avatarPb.SkillDepotId,
				SkillLevelMap: avatarPb.SkillMap,
				Uid:           g.Uid,
				Weapon: &proto.SceneWeaponInfo{
					EntityId:     entity.WeaponEntityId,
					GadgetId:     50011101,
					ItemId:       weaponPb.WeaponId,
					Guid:         weaponPb.Guid,
					Level:        weaponPb.Level,
					PromoteLevel: weaponPb.PromoteLevel,
					AbilityInfo:  &proto.AbilitySyncStateInfo{IsInited: false},
					AffixMap:     nil,
				},
				WearingFlycloakId: avatarPb.FlycloakId,
			}},
			EntityClientData: &proto.EntityClientData{
				WindChangeSceneTime:   0,
				WindmillSyncAngle:     0,
				WindChangeTargetLevel: 0,
			},
			EntityId:      entityId,
			EntityType:    proto.ProtEntityType_PROT_ENTITY_AVATAR,
			FightPropList: make([]*proto.FightPropPair, 0),
			LifeState:     1,
			MotionInfo: &proto.MotionInfo{
				Pos: &proto.Vector{
					X: 2739.658,
					Y: 194.6,
					Z: -1711.08,
				},
				Rot: &proto.Vector{
					X: 0,
					Y: 316.9,
					Z: 0,
				},
				Speed: &proto.Vector{
					X: 0,
					Y: 0,
					Z: 0,
				},
				State:  0,
				Params: nil,
			},
			PropList: []*proto.PropPair{{
				PropValue: &proto.PropValue{
					Type:  4001,
					Val:   1,
					Value: &proto.PropValue_Ival{Ival: 1},
				},
				Type: 4001,
			}},
			RendererChangedInfo: &proto.EntityRendererChangedInfo{
				ChangedRenderers: nil,
				VisibilityCount:  0,
			},
		}

		for id, value := range avatarPb.FightPropMap {
			fightPropList := &proto.FightPropPair{
				PropType:  id,
				PropValue: value,
			}
			sceneEntityInfo.FightPropList = append(sceneEntityInfo.FightPropList, fightPropList)
		}
		notify.EntityList = append(notify.EntityList, sceneEntityInfo)
	}

	g.seed(cmd.SceneEntityAppearNotify, notify)
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
	g.HostPlayerNotify()
	g.SceneDataNotify()
	g.SceneTimeNotify()
	g.seed(cmd.SceneAreaWeatherNotify, &proto.SceneAreaWeatherNotify{
		WeatherAreaId: 1,
		ClimateType:   1,
	})
	g.PlayerGameTimeNotify()
	g.PlayerEnterSceneInfoNotify() // 实体通知
	g.ScenePlayerInfoNotify()
	g.SceneTeamUpdateNotify()
	g.seed(cmd.SyncTeamEntityNotify, &proto.SyncTeamEntityNotify{
		SceneId: 3,
	})
	g.SceneForceUnlockNotify()
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
	g.SceneEntityAppearNotify()
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
