package gdconf

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/mari-track/MariGS/pkg/config"
	"github.com/mari-track/MariGS/pkg/logger"

	"github.com/jszwec/csvutil"
)

var CONF *GameDataConfig = nil

type GameDataConfig struct {
	// 配置表路径前缀
	jsonPrefix string
	luaPrefix  string
	txtPrefix  string
	// 配置表数据
	AvatarDataMap         map[uint32]*AvatarData   // 角色
	FetterDataMap         map[int32]*FetterData    // 角色资料解锁
	FetterDataAvatarIdMap map[int32][]int32        // 角色资料解锁角色id索引
	OpenStateDataMap      map[int32]*OpenStateData // 开放状态
}

func InitGameDataConfig() {
	logger.Info("读取资源文件")
	CONF = new(GameDataConfig)
	startTime := time.Now().Unix()
	CONF.loadAll()
	runtime.GC()
	endTime := time.Now().Unix()
	logger.Info("load all game data config finish, cost: %v(s)", endTime-startTime)
}

func (g *GameDataConfig) loadAll() {
	pathPrefix := config.GetConfig().GameDataConfigPath

	dirInfo, err := os.Stat(pathPrefix)
	if err != nil || !dirInfo.IsDir() {
		info := fmt.Sprintf("open game data config dir error: %v", err)
		panic(info)
	}

	g.jsonPrefix = pathPrefix + "/json"
	dirInfo, err = os.Stat(g.jsonPrefix)
	if err != nil || !dirInfo.IsDir() {
		info := fmt.Sprintf("open game data config json dir error: %v", err)
		panic(info)
	}
	g.jsonPrefix += "/"

	g.luaPrefix = pathPrefix + "/lua"
	dirInfo, err = os.Stat(g.luaPrefix)
	if err != nil || !dirInfo.IsDir() {
		info := fmt.Sprintf("open game data config lua dir error: %v", err)
		panic(info)
	}
	g.luaPrefix += "/"

	g.txtPrefix = pathPrefix + "/txt"
	dirInfo, err = os.Stat(g.txtPrefix)
	if err != nil || !dirInfo.IsDir() {
		info := fmt.Sprintf("open game data config txt dir error: %v", err)
		panic(info)
	}
	g.txtPrefix += "/"

	g.load()
}

func (g *GameDataConfig) load() {
	g.loadAvatarData()    // 角色
	g.loadFetterData()    // 角色资料解锁
	g.loadOpenStateData() // 开放状态
}

func splitStringArray(str string) []string {
	if str == "" {
		return make([]string, 0)
	} else if strings.Contains(str, ";") {
		return strings.Split(str, ";")
	} else if strings.Contains(str, ",") {
		return strings.Split(str, ",")
	} else {
		return []string{str}
	}
}

type IntArray []int32

func (a *IntArray) UnmarshalCSV(data []byte) error {
	str := string(data)
	str = strings.ReplaceAll(str, " ", "")
	intStrList := splitStringArray(str)
	for _, intStr := range intStrList {
		v, err := strconv.ParseInt(intStr, 10, 32)
		if err != nil {
			panic(err)
		}
		*a = append(*a, int32(v))
	}
	return nil
}

func readTable[T any](tablePath string, table *[]*T) {
	fileData, err := os.ReadFile(tablePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	reader := csv.NewReader(bytes.NewBuffer(fileData))
	reader.Comma = '\t'
	reader.LazyQuotes = true
	dec, err := csvutil.NewDecoder(reader)
	if err != nil {
		info := fmt.Sprintf("create decoder error: %v", err)
		panic(info)
	}
	for {
		t := new(T)
		err := dec.Decode(t)
		if err == io.EOF {
			break
		}
		if err != nil {
			info := fmt.Sprintf("decode file error: %v", err)
			panic(info)
		}
		*table = append(*table, t)
	}
}
