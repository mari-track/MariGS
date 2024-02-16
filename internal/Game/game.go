package Game

import (
	"github.com/mari-track/MariGS/internal/Game/model"
	"github.com/mari-track/MariGS/pkg/kcp"
	"github.com/mari-track/MariGS/pkg/logger"
	"github.com/mari-track/MariGS/protocol/cmd"
	"github.com/mari-track/MariGS/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

type Game struct {
	IsToken        bool // 是否通过token验证
	Uid            uint32
	Seed           uint64
	KcpConn        *kcp.UDPSession
	LastActiveTime int64 // 最近一次的活跃时间
	RouteManager   *RouteManager
	// 玩家数据
	Player *model.Player
	// 密钥
	XorKey []byte
}

func (g *Game) seed(cmdId uint16, playerMsg pb.Message) {
	name := cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId)
	logger.Debug("[UID:%v][S --> C][CMDID:%v][%s]MSG:%s", g.Uid, cmdId, name, playerMsg)
	SendHandle(g, cmdId, playerMsg)
}

func (g *Game) PingReq(playerMsg pb.Message) {
	req := playerMsg.(*proto.PingReq)
	rsp := &proto.PingRsp{
		Seq:        0,
		ClientTime: req.ClientTime,
	}
	g.seed(cmd.PingRsp, rsp)
}