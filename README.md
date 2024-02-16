# MariGS

## 当前功能

* 登录

## 计划即将实现的功能

* 祈愿
* 传送
* 战斗
* 好友
* 活动
* 深渊
* 任务
* 大世界
* 联机
* ........

### 开始

- 安装[go1.20](https://go.dev/doc/install)或以上版本
- 安装[mysql](https://www.mysql.com/downloads/)
- 获取游戏1.0.0(很遗憾的是，已经无法在官方渠道找到此版本)

### 构建

##### Windows

```shell
git clone --recurse-submodules https://github.com/mari-track/MariGS.git
cd MariGS
go mod tidy # 安装依赖
go build main.go # 编译
```

##### Linux (GNU)

```shell
git clone --recurse-submodules https://github.com/mari-track/MariGS.git
cd MariGS
go mod tidy # 安装依赖
go build main.go # 编译
```

你可以在目录中找到输出的可执行文件
