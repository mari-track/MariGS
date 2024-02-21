package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/mari-track/MariGS/pkg/random"
)

type Config struct {
	LogLevel           string `json:"logLevel"`
	MysqlDsn           string `json:"MysqlDsn"`
	GameDataConfigPath string `json:"GameDataConfigPath"`
	SignRsaKey         []byte
	EncRsaKeyMap       map[string][]byte
	PwdRsaKey          []byte
	Dispatch           []Dispatch
	Account            *Account
	Http               *Http
	Game               *Game
	Ec2b               *random.Ec2b `json:"Ec2B"`
}
type Dispatch struct {
	Name        string
	Title       string
	Type        string
	DispatchUrl string
}
type Account struct {
	AutoCreate bool  `json:"autoCreate"`
	MaxPlayer  int64 `json:"maxPlayer"`
}
type Http struct {
	Addr string `json:"addr"`
	Port int64  `json:"port"`
}
type Game struct {
	Addr             string `json:"addr"`
	Port             uint32 `json:"port"`
	IsInitialization bool   `json:"isInitialization"`
}

var CONF *Config = nil

func GetConfig() *Config {
	return CONF
}

var FileNotExist = errors.New("config file not found")

func LoadConfig() error {
	filePath := "./config.json"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}
	f, err := os.Open(filePath)
	if err != nil {
		return FileNotExist
	}
	defer func() {
		_ = f.Close()
	}()
	c := new(Config)
	d := json.NewDecoder(f)
	if err := d.Decode(c); err != nil {
		return err
	}
	CONF = c
	return nil
}

var DefaultConfig = &Config{
	LogLevel:           "Info",
	MysqlDsn:           "root:password@tcp(127.0.0.1:3306)/MariGS?charset=utf8mb4&parseTime=True&loc=Local",
	GameDataConfigPath: "data",
	Dispatch: []Dispatch{
		{
			Name:        "os_usa",
			Title:       "MariGS",
			Type:        "os_usa",
			DispatchUrl: "http://127.0.0.1:8080/query_cur_region",
		},
	},
	Account: &Account{
		AutoCreate: true,
		MaxPlayer:  -1,
	},
	Http: &Http{
		Addr: "0.0.0.0",
		Port: 8080,
	},
	Game: &Game{
		Addr:             "127.0.0.1",
		Port:             22102,
		IsInitialization: true,
	},
}
