package Game

import (
	"reflect"

	"github.com/mari-track/MariGS/pkg/logger"
	"github.com/mari-track/MariGS/protocol/cmd"
	"github.com/mari-track/MariGS/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

/*************************************通知包******************************************/

func (g *Game) ServerTimeNotify() {
	notify := &proto.ServerTimeNotify{
		ServerTime: uint64(GetServerTime()),
	}

	g.seed(cmd.ServerTimeNotify, notify)
}
func (g *Game) PlayerDataNotify() {
	db := g.Player.GetPbPlayerBasicCompBin()
	notify := &proto.PlayerDataNotify{
		NickName:          db.Nickname,
		ServerTime:        uint64(GetServerTime()),
		IsFirstLoginToday: true,
		RegionId:          1,
		PropMap:           make(map[uint32]*proto.PropValue),
	}

	for k, v := range db.PropMap {
		notify.PropMap[k] = g.PacketPropValue(k, v)
	}
	g.seed(cmd.PlayerDataNotify, notify)
}

func (g *Game) PlayerLevelRewardUpdateNotify() {
	notify := &proto.PlayerLevelRewardUpdateNotify{
		LevelList: make([]uint32, 0),
	}
	g.seed(cmd.PlayerLevelRewardUpdateNotify, notify)
}

func (g *Game) OpenStateUpdateNotify() {
	db := g.Player.GetPbPlayerBasicCompBin()
	notify := &proto.OpenStateUpdateNotify{
		OpenStateMap: db.OpenStateMap,
	}
	g.seed(cmd.OpenStateUpdateNotify, notify)
}

func (g *Game) OpenStateChangeNotify(stateMap map[uint32]uint32) {
	notify := &proto.OpenStateChangeNotify{OpenStateMap: make(map[uint32]uint32)}
	for key, value := range stateMap {
		notify.OpenStateMap[key] = value
	}
	g.seed(cmd.OpenStateChangeNotify, notify)
}

func (g *Game) FinishedParentQuestNotify() {
	notify := &proto.FinishedParentQuestNotify{
		ParentQuestList: make([]*proto.ParentQuest, 0),
	}
	g.seed(cmd.FinishedParentQuestNotify, notify)
}

func (g *Game) ScenePlayerInfoNotify() {
	notify := &proto.ScenePlayerInfoNotify{
		PlayerInfoList: make([]*proto.ScenePlayerInfo, 0),
	}
	// 添加当前角色
	db := g.Player.GetPbPlayerBasicCompBin()
	notify.PlayerInfoList = append(notify.PlayerInfoList, &proto.ScenePlayerInfo{
		Uid:         g.Uid,
		PeerId:      1,
		Name:        db.Nickname,
		IsConnected: false,
		SceneId:     3,
		OnlinePlayerInfo: &proto.OnlinePlayerInfo{
			Uid:                 g.Uid,
			Nickname:            db.Nickname,
			PlayerLevel:         db.Level,
			AvatarId:            db.AvatarId,
			MpSettingType:       proto.MpSettingType_MP_SETTING_ENTER_AFTER_APPLY,
			CurPlayerNumInWorld: 1,
			WorldLevel:          db.WorldLevel,
			OnlineId:            "",
			NameCardId:          db.NameCardId,
			BlacklistUidList:    nil,
			Signature:           db.Signature,
		},
	})

	g.seed(cmd.ScenePlayerInfoNotify, notify)
}

func (g *Game) SceneTeamUpdateNotify() {
	notify := &proto.SceneTeamUpdateNotify{
		SceneTeamAvatarList: make([]*proto.SceneTeamAvatar, 0),
		IsInMp:              false,
	}
	// 添加当前角色
	db := g.Player.GetPbPlayerSceneCompBin()
	avatarEntity := g.SceneEntity.GetAvatarEntity()
	for entityId, entity := range avatarEntity {
		avatarDb := g.Player.GetPbAvatarById(entity.AvatarId)
		sceneTeamAvatar := &proto.SceneTeamAvatar{
			PlayerUid:         g.Uid,
			AvatarGuid:        avatarDb.Guid,
			SceneId:           db.MyCurSceneId,
			EntityId:          entityId,
			AvatarInfo:        nil,
			SceneAvatarInfo:   nil,
			AvatarAbilityInfo: &proto.AbilitySyncStateInfo{IsInited: false},
			ServerBuffList:    nil,
			WeaponGuid:        avatarDb.WeaponGuid,
			WeaponEntityId:    entity.WeaponEntityId,
			WeaponAbilityInfo: &proto.AbilitySyncStateInfo{IsInited: false},
			AbilityControlBlock: &proto.AbilityControlBlock{AbilityEmbryoList: []*proto.AbilityEmbryo{
				{AbilityId: 1, AbilityNameHash: 4291357363, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 2, AbilityNameHash: 518324758, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 3, AbilityNameHash: 3276790745, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 4, AbilityNameHash: 3429175060, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 5, AbilityNameHash: 3429175061, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 6, AbilityNameHash: 4253958193, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 7, AbilityNameHash: 209033715, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 8, AbilityNameHash: 900298203, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 9, AbilityNameHash: 1049300416, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 10, AbilityNameHash: 872734369, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 11, AbilityNameHash: 3283199419, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 12, AbilityNameHash: 3157276645, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 13, AbilityNameHash: 664564586, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 14, AbilityNameHash: 4172444990, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 15, AbilityNameHash: 1578170661, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 16, AbilityNameHash: 918348879, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 17, AbilityNameHash: 1410219662, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 18, AbilityNameHash: 1474894886, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 19, AbilityNameHash: 3832178184, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 20, AbilityNameHash: 2306062007, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 21, AbilityNameHash: 3105629177, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 22, AbilityNameHash: 3771526669, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 23, AbilityNameHash: 3032325400, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 24, AbilityNameHash: 100636247, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 25, AbilityNameHash: 1564404322, AbilityOverrideNameHash: 1178079449},
				{AbilityId: 26, AbilityNameHash: 3374327026, AbilityOverrideNameHash: 1178079449},
			}},
			IsReconnect: false,
		}
		notify.SceneTeamAvatarList = append(notify.SceneTeamAvatarList, sceneTeamAvatar)
	}
}

/*************************************封装******************************************/

func (g *Game) PacketPropValue(key uint32, value any) *proto.PropValue {
	propValue := new(proto.PropValue)
	propValue.Type = key
	switch value.(type) {
	case int:
		v := value.(int)
		propValue.Val = int64(v)
		propValue.Value = &proto.PropValue_Ival{Ival: int64(v)}
	case int8:
		v := value.(int8)
		propValue.Val = int64(v)
		propValue.Value = &proto.PropValue_Ival{Ival: int64(v)}
	case int16:
		v := value.(int16)
		propValue.Val = int64(v)
		propValue.Value = &proto.PropValue_Ival{Ival: int64(v)}
	case int32:
		v := value.(int32)
		propValue.Val = int64(v)
		propValue.Value = &proto.PropValue_Ival{Ival: int64(v)}
	case int64:
		v := value.(int64)
		propValue.Val = v
		propValue.Value = &proto.PropValue_Ival{Ival: v}
	case float32:
		v := value.(float32)
		propValue.Value = &proto.PropValue_Fval{Fval: v}
	case float64:
		v := value.(float64)
		propValue.Value = &proto.PropValue_Fval{Fval: float32(v)}
	case uint8:
		v := value.(uint8)
		propValue.Val = int64(v)
		propValue.Value = &proto.PropValue_Ival{Ival: int64(v)}
	case uint16:
		v := value.(uint16)
		propValue.Val = int64(v)
		propValue.Value = &proto.PropValue_Ival{Ival: int64(v)}
	case uint32:
		v := value.(uint32)
		propValue.Val = int64(v)
		propValue.Value = &proto.PropValue_Ival{Ival: int64(v)}
	case uint64:
		v := value.(uint64)
		propValue.Val = int64(v)
		propValue.Value = &proto.PropValue_Ival{Ival: int64(v)}
	default:
		logger.Error("unknown value type: %v, value: %v", reflect.TypeOf(value), value)
		return nil
	}
	return propValue
}

/*************************************数据包处理******************************************/

func (g *Game) GetPlayerBlacklistReq(payloadMsg pb.Message) {
	g.seed(cmd.GetPlayerBlacklistRsp, &proto.GetPlayerBlacklistRsp{
		Retcode:   0,
		Blacklist: make([]*proto.FriendBrief, 0),
	})
}

func (g *Game) GetPlayerSocialDetailReq(payloadMsg pb.Message) {
	db := g.Player.GetPbPlayerBasicCompBin()
	req := payloadMsg.(*proto.GetPlayerSocialDetailReq)
	if req.Uid != g.Uid {
		return
	}
	rsp := &proto.GetPlayerSocialDetailRsp{
		Retcode: 0,
		DetailData: &proto.SocialDetail{
			Uid:       g.Uid,
			Nickname:  db.Nickname,
			Level:     db.Level,
			AvatarId:  db.AvatarId,
			Signature: db.Signature,
			Birthday: &proto.Birthday{
				Month: 0,
				Day:   0,
			},
			WorldLevel:        db.WorldLevel,
			ReservedList:      nil,
			OnlineState:       proto.FriendOnlineState_FRIEND_ONLINE,
			Param:             0,
			IsFriend:          true,
			IsMpModeAvailable: true,
			OnlineId:          "",
			NameCardId:        db.NameCardId,
			IsInBlacklist:     false,
		},
	}
	g.seed(cmd.GetPlayerSocialDetailRsp, rsp)
}

func (g *Game) SetOpenStateReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetOpenStateReq)
	rsp := &proto.SetOpenStateRsp{
		Retcode: 0,
		Key:     req.Key,
		Value:   req.Value,
	}
	g.OpenStateChangeNotify(map[uint32]uint32{req.Key: req.Value})
	g.seed(cmd.SetOpenStateRsp, rsp)
}
