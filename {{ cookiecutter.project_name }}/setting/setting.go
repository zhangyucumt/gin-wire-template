package setting

import (
	"github.com/go-ini/ini"
	"time"
)

type App struct {
	JwtSecret       string
	RuntimeRootPath string
	PageSize        int

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
	Version     string
	CookieName  string
	JwtTime     time.Duration
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpsModel   bool
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type string
	//User        string
	//Password    string
	//Host        string
	//Name        string
	//TablePrefix string
	Source string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Username    string
	Password    string
	MaxIdle     int
	MaxActive   int
	Db          int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File

func Load() error {
	var err error
	cfg, err = ini.Load("app.ini")
	if err != nil {
		return err
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)

	AppSetting.JwtTime = AppSetting.JwtTime * time.Second

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	//RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
	return nil
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		panic(err)
	}
}
