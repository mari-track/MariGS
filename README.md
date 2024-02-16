# MariGS

## 当前功能

* 登录

## 计划即将实现的功能

* 更换至MySQL数据库
* 祈愿
* 传送
* 战斗
* 好友
* 控制台指令
* ........

### 开始

- 安装[go1.20](https://go.dev/doc/install)或以上版本
- 安装[mysql](https://www.mysql.com/downloads/)
- 获取游戏1.0.0(很遗憾的是，已经无法在官方渠道找到此版本)
- 下载[Launcher](https://github.com/gucooing/Launcher/releases/latest)
- 以管理员身份打开 Launcher ,进入设置选择游戏路径后设置服务器地址

### 构建

##### Windows

```shell
git clone --recurse-submodules https://github.com/mari-track/MariGS.git
cd MariGS
go mod tidy # 安装依赖
go build cmd/server/main.go # 编译
```

##### Linux (GNU)

```shell
git clone --recurse-submodules https://github.com/mari-track/MariGS.git
cd MariGS
go mod tidy # 安装依赖
go build cmd/server/main.go # 编译
```

你可以在 cmd/server 目录中找到输出的可执行文件
