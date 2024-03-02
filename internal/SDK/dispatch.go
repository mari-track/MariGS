package SDK

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mari-track/MariGS/pkg/endec"
	"github.com/mari-track/MariGS/pkg/logger"
	"github.com/mari-track/MariGS/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

// ClientCustomConfig 客户端版本定义的配置 客户端版本号对应的配置 需要兼容老的json格式
type ClientCustomConfig struct {
	Visitor        bool              `json:"visitor"`        // 游客功能
	SdkEnv         string            `json:"sdkenv"`         // sdk环境类型
	DebugMenu      bool              `json:"debugmenu"`      // debug菜单
	DebugLogSwitch []int32           `json:"debuglogswitch"` // 打开的log类型
	DebugLog       bool              `json:"debuglog"`       // log总开关
	DeviceList     map[string]string `json:"devicelist"`
	LoadJsonData   bool              `json:"loadjsondata"`  // 用json读取InLevel数据
	ShowException  bool              `json:"showexception"` // 是否显示异常提示框 默认为true
	CheckDevice    bool              `json:"checkdevice"`
	LoadPatch      bool              `json:"loadPatch"`
	RegionConfig   string            `json:"regionConfig"`
	DownloadMode   int32             `json:"downloadMode"`
	CodeSwitch     []int32           `json:"codeSwitch"`
	CoverSwitch    []int32           `json:"coverSwitch"`
}

// RegionCustomConfig 区服相关的配置 避免在http中使用Json格式
type RegionCustomConfig struct {
	CloseAntiDebug   bool `json:"close_antidebug"`  // 默认打开反调开关 默认false
	ForceKill        bool `json:"force_kill"`       // 默认false
	AntiDebugPc      bool `json:"antidebug_pc"`     // pc默认不开启反调 默认false
	AntiDebugIos     bool `json:"antidubug_ios"`    // ios默认不开启反调 默认false
	AntiDebugAndroid bool `json:"antidubug_androd"` // android默认不开启反调 默认false
}

func (s *Server) GetRegionList(c *gin.Context) {
	rsp := &proto.QueryRegionListHttpRsp{
		Retcode:                     0,
		RegionList:                  make([]*proto.RegionSimpleInfo, 0),
		ClientSecretKey:             make([]byte, 0),
		ClientCustomConfigEncrypted: make([]byte, 0),
		EnableLoginPc:               false,
	}

	for _, dispatch := range s.Config.Dispatch {
		regionSimpleInfo := &proto.RegionSimpleInfo{
			Name:        dispatch.Name,
			Title:       dispatch.Title,
			Type:        dispatch.Type,
			DispatchUrl: dispatch.DispatchUrl,
		}
		rsp.RegionList = append(rsp.RegionList, regionSimpleInfo)
	}

	dispatchKey := s.Config.Ec2b.XorKey()
	rsp.ClientSecretKey = s.Config.Ec2b.Bytes()
	clientCustomConfigEncrypted, _ := json.Marshal(&ClientCustomConfig{
		SdkEnv:         "2",
		CheckDevice:    false,
		LoadPatch:      false,
		ShowException:  false,
		RegionConfig:   "pm|fk|add",
		DownloadMode:   0,
		DebugMenu:      true,
		DebugLogSwitch: []int32{0},
		DebugLog:       true,
		CodeSwitch:     []int32{3628},
		CoverSwitch:    []int32{40},
	})
	endec.Xor(clientCustomConfigEncrypted, dispatchKey)
	rsp.ClientCustomConfigEncrypted = clientCustomConfigEncrypted // 加密后的客户端版本定义的配置
	rsp.EnableLoginPc = true
	reqdata, err := pb.Marshal(rsp)
	if err != nil {
		logger.Error("pb marshal QueryRegionListHttpRsp error: %v", err)
		return
	}
	rspBase64 := base64.StdEncoding.EncodeToString(reqdata)

	c.String(200, rspBase64)
}

func (s *Server) GetRegionCurr(c *gin.Context) {
	rsp := &proto.QueryCurrRegionHttpRsp{
		Retcode: 0,
		Msg:     "",
		RegionInfo: &proto.RegionInfo{
			GateserverIp:               s.Config.Game.Addr,
			GateserverPort:             s.Config.Game.Port,
			SecretKey:                  s.Config.Ec2b.Bytes(),
			PayCallbackUrl:             "",
			AreaType:                   "",
			ResourceUrl:                "",
			DataUrl:                    "",
			FeedbackUrl:                "",
			BulletinUrl:                "",
			ResourceUrlBak:             "",
			DataUrlBak:                 "",
			ClientDataVersion:          0,
			HandbookUrl:                "",
			ClientSilenceDataVersion:   0,
			ClientDataMd5:              "",
			ClientSilenceDataMd5:       "",
			ResVersionConfig:           nil,
			OfficialCommunityUrl:       "",
			ClientVersionSuffix:        "",
			ClientSilenceVersionSuffix: "",
			UseGateserverDomainName:    false,
			GateserverDomainName:       "",
			UserCenterUrl:              "",
			AccountBindUrl:             "",
			CdkeyUrl:                   "",
			PrivacyPolicyUrl:           "",
		},
		ClientSecretKey:                   s.Config.Ec2b.Bytes(),
		RegionCustomConfigEncrypted:       make([]byte, 0),
		ClientRegionCustomConfigEncrypted: make([]byte, 0),
		Detail:                            nil,
	}
	regionCustomConfig, _ := json.Marshal(&RegionCustomConfig{
		CloseAntiDebug:   true,
		ForceKill:        false,
		AntiDebugPc:      false,
		AntiDebugIos:     false,
		AntiDebugAndroid: false,
	})
	endec.Xor(regionCustomConfig, s.Config.Ec2b.XorKey())
	rsp.RegionCustomConfigEncrypted = regionCustomConfig
	clientCustomConfig, _ := json.Marshal(&ClientCustomConfig{
		SdkEnv:         "2",
		CheckDevice:    false,
		LoadPatch:      false,
		ShowException:  true,
		RegionConfig:   "pm|fk|add",
		DownloadMode:   0,
		DebugMenu:      true,
		DebugLogSwitch: []int32{0},
		DebugLog:       true,
		CodeSwitch:     []int32{3628},
		CoverSwitch:    []int32{40},
	})
	endec.Xor(clientCustomConfig, s.Config.Ec2b.XorKey())
	rsp.ClientRegionCustomConfigEncrypted = clientCustomConfig // 加密后的客户端区服定义的配置
	reqdata, err := pb.Marshal(rsp)
	if err != nil {
		logger.Error("pb marshal QueryCurrRegionHttpRsp error: %v", err)
		return
	}
	rspBase64 := base64.StdEncoding.EncodeToString(reqdata)
	c.String(200, rspBase64)
}

func (s *Server) GetRegionCurrDump(c *gin.Context) {
	urlPath := c.Request.URL.RawQuery

	rsps, err := http.Get("http://10.0.0.15:20001/query_cur_region?" + urlPath)
	if err != nil {
		logger.Error("Request failed:", err)
		return
	}
	defer rsps.Body.Close()

	data, err := io.ReadAll(rsps.Body)
	if err != nil {
		logger.Error("Read body failed:", err)
		return
	}

	datamsg, _ := base64.StdEncoding.DecodeString(string(data))

	dispatch := new(proto.QueryCurrRegionHttpRsp)

	err = pb.Unmarshal(datamsg, dispatch)
	if err != nil {
		logger.Error("", err)
	}

	/*
		if dispatch.ClientSecretKey != nil {
			dispatch.RegionInfo.GateserverIp = "127.0.0.1"
			dispatch.RegionInfo.GateserverPort = 20045
			ec2b, _ := random.LoadEc2bKey(dispatch.ClientSecretKey)
			logger.Info("%s", base64.StdEncoding.EncodeToString(dispatch.ClientSecretKey))
			logger.Info("%s", base64.StdEncoding.EncodeToString(ec2b.XorKey()))
		}
	*/
	rspbin, _ := pb.Marshal(dispatch)
	dispatchb64 := base64.StdEncoding.EncodeToString(rspbin)
	c.String(200, dispatchb64)
}
