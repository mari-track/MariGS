package DataBase

import (
	"time"

	"github.com/mari-track/MariGS/pkg/config"
	"github.com/mari-track/MariGS/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DBASE *Store

func (s *Store) init() {
	var err error
	DBASE = s
	dsn := s.config.MysqlDsn
	DBASE.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	DBASE.config = s.config
	s.Db = DBASE.Db
	if err != nil {
		logger.Error("MySQL数据库连接失败,错误原因:%s", err)
		return
	}
	logger.Info("MySQL数据库连接成功")
	sqlDB, err := s.Db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 10 秒钟
	// 初始化表
	err = s.Db.AutoMigrate(&Account{}, &UidPlayer{}, &Player{})
	if err != nil {
		logger.Error("MySQL数据库初始化失败")
		return
	}
	logger.Info("MySQL数据库初始化成功")
}

// NewStore 创建一个新的 store。
func NewStore(config *config.Config) *Store {
	s := &Store{config: config}
	s.init()
	return s
}
