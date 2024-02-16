package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/mari-track/MariGS/internal"
	"github.com/mari-track/MariGS/internal/Game"
	"github.com/mari-track/MariGS/pkg"
	"github.com/mari-track/MariGS/pkg/config"
	"github.com/mari-track/MariGS/pkg/gdconf"
	"github.com/mari-track/MariGS/pkg/logger"
)

func main() {
	// 启动读取配置
	err := config.LoadConfig()
	if err != nil {
		if err == config.FileNotExist {
			p, _ := json.MarshalIndent(config.DefaultConfig, "", "  ")
			fmt.Printf("找不到配置文件\n已生成默认配置文件 config.json \n请修改后再启动\n按 'Enter' 键退出 ...\n")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			cf, _ := os.Create("./config.json")
			cf.Write(p)
			defer cf.Close()
			os.Exit(0)
		} else {
			panic(err)
		}
	}
	// 初始化日志
	logger.InitLogger("MariGS")
	logger.SetLogLevel(strings.ToUpper(config.GetConfig().LogLevel))
	logger.Info("MariGS")
	logger.Info("BuildVersion:%v", pkg.BuildVersion)
	logger.Info("ProtoVersion:%v", pkg.ProtoVersion)

	cfg := config.GetConfig()
	// 初始化
	newserver := internal.NewServer(cfg)
	if newserver == nil {
		logger.Error("服务器初始化失败")
		return
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// 启动SDK服务
	go func() {
		if err = newserver.Start(); err != nil {
			logger.Error("无法启动SDK服务器")
		}
	}()

	// 读取data
	gdconf.InitGameDataConfig()
	// 启动Game服务
	go func() {
		if err = Game.Run(); err != nil {
			logger.Error("无法启动Game服务器")
		}
	}()

	go func() {
		select {
		case <-done:
			logger.Info("SDK服务正在关闭")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := newserver.Shutdown(ctx); err != nil {
				logger.Error("无法正常关闭HTTP服务")
			}
			logger.Info("SDK服务已停止")
			logger.CloseLogger()
			os.Exit(0) // 将终止程序
		}
	}()
	select {}
}
