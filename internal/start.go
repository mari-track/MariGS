package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/mari-track/MariGS/internal/DataBase"
	"github.com/mari-track/MariGS/internal/SDK"
	"github.com/mari-track/MariGS/pkg/alg"
	"github.com/mari-track/MariGS/pkg/config"
)

// 初始化所有服务
func NewServer(cfg *config.Config) *SDK.Server {
	s := &SDK.Server{}
	s.Config = cfg
	s.Store = DataBase.NewStore(s.Config) // 初始化数据库连接
	gin.SetMode(gin.ReleaseMode)          // 初始化gin
	s.Router = gin.New()                  // gin.Default()
	s.Router.Use(gin.Recovery())
	cfg.Ec2b = alg.GetEc2b() // 读取ec2b密钥

	return s
}
