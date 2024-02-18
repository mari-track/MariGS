package Game

import (
	"time"

	"github.com/mari-track/MariGS/internal/Game/model"
	"github.com/mari-track/MariGS/internal/Game/sceneEntity"
	"github.com/mari-track/MariGS/pkg/kcp"
	"github.com/mari-track/MariGS/pkg/logger"
	"github.com/mari-track/MariGS/protocol/cmd"
	"github.com/mari-track/MariGS/protocol/proto"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

type Game struct {
	IsToken      bool // 是否通过token验证
	Uid          uint32
	Seed         uint64
	KcpConn      *kcp.UDPSession
	RouteManager *RouteManager
	// 玩家数据
	Player *model.Player
	// 临时数据
	SceneEntity           *sceneEntity.SceneEntity // 场景实体数据
	XorKey                []byte                   // 密钥
	LastActiveTime        int64                    // 最近一次的活跃时间
	GameObjectGuidCounter uint64                   // 玩家在线唯一id
	SceneToken            uint32                   // 场景token
}

func (g *Game) GetNextGameObjectGuid() uint64 {
	g.GameObjectGuidCounter++
	return uint64(g.Uid)<<32 + g.GameObjectGuidCounter
}

func (g *Game) seed(cmdId uint16, playerMsg pb.Message) {
	name := cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId)
	logger.Debug("[UID:%v][S --> C][CMDID:%v][%s]MSG:\n%s", g.Uid, cmdId, name, protojson.Format(playerMsg))
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

func GetServerTime() int64 {
	serverTime := time.Now().UnixMilli()
	return serverTime
}
