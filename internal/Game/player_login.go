package Game

import (
	"strconv"
	"sync"

	"github.com/mari-track/MariGS/internal/DataBase"
	"github.com/mari-track/MariGS/internal/Game/model"
	"github.com/mari-track/MariGS/pkg/logger"
	"github.com/mari-track/MariGS/pkg/random"
	"github.com/mari-track/MariGS/protocol/cmd"
	"github.com/mari-track/MariGS/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

var syncGD sync.Mutex

func GetPlayerTokenReq(g *Game, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetPlayerTokenReq)
	rsp := new(proto.GetPlayerTokenRsp)
	// 请求验证
	if req.AccountToken == "" || req.AccountUid == "" {
		return
	}
	accountUid, err := strconv.ParseUint(req.AccountUid, 10, 64)
	if err != nil {
		logger.Error("get token uid error")
		return
	}
	account := DataBase.DBASE.QueryUidPlayerUidByFieldPlayer(uint32(accountUid))
	if account.ComboToken != req.AccountToken {
		// rsp.Retcode = proto.
		return
	}
	// 封禁验证
	if account.IsBan {
		rsp.Uid = 0
		// rsp.Retcode = uint32(proto.Retcode_RET_IN_GM_BIND_ACCESS)
		rsp.Msg = "该账号正处于封禁状态，暂时无法登录，详情可联系客服。"
		// g.Send(cmd.PlayerGetTokenScRsp, rsp)
		logger.Info("登录账号:%v,已被封禁", accountUid)
		return
	}

	// 重复登录验证

	newuidPlayer := &DataBase.UidPlayer{
		AccountId:  account.AccountId,
		IsBan:      false,
		ComboToken: "",
	}

	syncGD.Lock()
	err = DataBase.DBASE.UpdateUidPlayer(account.AccountId, newuidPlayer)
	if err != nil {
		rsp.Uid = 0
		// rsp.Retcode = uint32(proto.Retcode_RET_ACCOUNT_PARA_ERROR)
		rsp.Msg = "账号刷新失败"
		// g.Send(cmd.PlayerGetTokenScRsp, rsp)
		logger.Error("登录账号:%v,账号刷新失败", accountUid)
		return
	}
	syncGD.Unlock()

	g.Uid = uint32(account.AccountUid)
	// 通过token验证
	g.IsToken = true

	// 初始化路由
	g.RouteManager = NewRouteManager(g)
	// 构造回复内容
	timeRand := random.GetTimeRand()
	serverSeedUint64 := timeRand.Uint64()
	g.Seed = serverSeedUint64

	rsp = &proto.GetPlayerTokenRsp{
		Retcode:            0,
		Msg:                "",
		Uid:                g.Uid,
		Token:              req.AccountToken,
		BlackUidEndTime:    0,
		AccountType:        req.AccountType,
		AccountUid:         req.AccountUid,
		IsProficientPlayer: false,
		SecretKey:          "",
		GmUid:              0,
		SecretKeySeed:      g.Seed,
		SecurityCmdBuffer:  nil,
		PlatformType:       req.PlatformType,
		ExtraBinData:       nil,
		IsGuest:            false,
		ChannelId:          req.ChannelId,
		SubChannelId:       req.SubChannelId,
		Tag:                0,
	}

	g.seed(cmd.GetPlayerTokenRsp, rsp)
}

func (g *Game) PlayerLoginReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.PlayerLoginReq)
	// 验证token
	logger.Info("[UID:%v]|登录客户端版本:%s", g.Uid, req.ClientVersion)
	dbPlayer := DataBase.DBASE.QueryAccountUidByFieldPlayer(g.Uid)
	// 去拉取数据
	g.Player = model.NewPlayer(dbPlayer.PlayerDataPb)

	if g.Player == nil {
		// 发送数据错误的信息
		return
	}

	if !g.Player.GetPbPlayerBasicCompBin().IsProficientPlayer {
		g.seed(cmd.DoSetPlayerBornDataNotify, nil)
	} else {
		// 发送登录通知包
		g.LoginNotify()
	}

	rsp := &proto.PlayerLoginRsp{
		AbilityHashCode:  0,
		GameBiz:          "hk4e_global",
		IsUseAbilityHash: true,
	}
	g.seed(cmd.PlayerLoginRsp, rsp)
}

func (g *Game) SetPlayerBornDataReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetPlayerBornDataReq)
	if req.NickName == "" || (req.AvatarId != 10000007 && req.AvatarId != 10000005) {
		return
	}
	// 更新昵称
	g.Player.UptoDateNickname(req.NickName)
	// 添加角色
	g.AddAvatar(req.AvatarId)
	// 将角色添加到队伍1
	avatarDb := g.Player.GetPbAvatarById(req.AvatarId)
	team1 := g.Player.GetPbTeamById(1)
	team1.LastCurAvatarGuid = avatarDb.Guid
	team1.AvatarGuidList = append(team1.AvatarGuidList, avatarDb.Guid)
	avatarBin := g.Player.GetPbPlayerAvatarCompBin()
	// 设置当前角色状态
	avatarBin.ChooseAvatarGuid = avatarDb.Guid
	avatarBin.CurTeamId = 1
	// 添加基础风之翼
	avatarBin.OwnedFlycloakList = append(avatarBin.OwnedFlycloakList, 140001)
	// 更新basic bin
	basicBin := g.Player.GetPbPlayerBasicCompBin()
	basicBin.AvatarId = req.AvatarId
	basicBin.NameCardId = 210001
	// 发送登录通知包
	g.LoginNotify()
	g.seed(cmd.SetPlayerBornDataRsp, nil)
}

// 登录通知包
func (g *Game) LoginNotify() {
	g.PlayerDataNotify()
	g.PlayerLevelRewardUpdateNotify()
	g.OpenStateUpdateNotify()
	g.AllSeenMonsterNotify()
	g.StoreWeightLimitNotify()
	g.PlayerStoreNotify()
	g.ItemCdGroupTimeNotify()
	g.AvatarDataNotify()
	/*
		AvatarExpeditionDataNotify
		AvatarSatiationDataNotify
	*/
	g.FinishedParentQuestNotify()
	g.QuestListNotify()
	g.PlayerEnterSceneNotify()
}
