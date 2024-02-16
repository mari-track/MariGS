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
		cmd.PingReq:              g.PingReq,
		cmd.PlayerLoginReq:       g.PlayerLoginReq,       // 第二个登录包
		cmd.SetPlayerBornDataReq: g.SetPlayerBornDataReq, // 选择主角
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

func (g *Game) GMRegisterMessage(cmdId uint16, payloadMsg pb.Message) {
	switch cmdId {

	}
}
