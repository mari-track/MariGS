package Game

import (
	"github.com/mari-track/MariGS/pkg/logger"
	"github.com/mari-track/MariGS/protocol/cmd"
	pb "google.golang.org/protobuf/proto"
)

type HandlerFunc func(payloadMsg pb.Message)

type RouteManager struct {
	handlerFuncRouteMap map[uint16]HandlerFunc
}

func (r *RouteManager) initRoute(g *Game) {
	r.handlerFuncRouteMap = map[uint16]HandlerFunc{
		cmd.PingReq:                  g.PingReq,
		cmd.PlayerLoginReq:           g.PlayerLoginReq,           // 第二个登录包
		cmd.GetPlayerBlacklistReq:    g.GetPlayerBlacklistReq,    // 获取黑名单？
		cmd.SetPlayerBornDataReq:     g.SetPlayerBornDataReq,     // 选择主角
		cmd.GetPlayerSocialDetailReq: g.GetPlayerSocialDetailReq, // 获取账户信息
		cmd.SetOpenStateReq:          g.SetOpenStateReq,          // 设置开放状态
		cmd.EnterSceneReadyReq:       g.EnterSceneReadyReq,       // 获取场景信息
		cmd.PathfindingEnterSceneReq: g.PathfindingEnterSceneReq, // 进入场景
		cmd.GetScenePointReq:         g.GetScenePointReq,         // 获取场景id
		cmd.GetSceneAreaReq:          g.GetSceneAreaReq,          // 获取场景区域
		cmd.SceneInitFinishReq:       g.SceneInitFinishReq,       // 场景初始化请求
		cmd.EnterSceneDoneReq:        g.EnterSceneDoneReq,        // 场景进入完成
		cmd.PostEnterSceneReq:        g.PostEnterSceneReq,        // 进入场景后
	}
}

func NewRouteManager(g *Game) (r *RouteManager) {
	r = new(RouteManager)
	r.initRoute(g)
	return r
}

func (g *Game) RegisterMessage(cmdId uint16, payloadMsg pb.Message) {
	handlerFunc, ok := g.RouteManager.handlerFuncRouteMap[cmdId]
	if !ok {
		logger.Error("no route for msg, cmdId: %v", cmdId)
		return
	}
	handlerFunc(payloadMsg)
	return
}
