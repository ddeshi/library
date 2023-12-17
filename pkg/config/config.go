package config

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppCfg = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerCfg = &Server{}

type Database struct {
	Type        string
	UserName    string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DBCfg = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisCfg = &Redis{}

var cfg *ini.File

// Init initialize the configuration instance
func Init() {
	var err error
	cfg, err = ini.Load("conf/cfg.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppCfg)
	mapTo("server", ServerCfg)
	mapTo("database", DBCfg)
	mapTo("redis", RedisCfg)

	//fmt.Println(AppCfg)

	AppCfg.ImageMaxSize = AppCfg.ImageMaxSize * 1024 * 1024
	ServerCfg.ReadTimeout = ServerCfg.ReadTimeout * time.Second
	ServerCfg.WriteTimeout = ServerCfg.WriteTimeout * time.Second
	RedisCfg.IdleTimeout = RedisCfg.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
