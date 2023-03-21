package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg           *ini.File
	RunMode       string
	HTTPPort      int
	ReadTimeout   time.Duration
	WriterTimeout time.Duration

	PageSize  int
	JwtSecret string
)

//初始化函数
func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini") //读取配置文件
	if err != nil {
		log.Fatalf("读取配置文件错误:%v", err)

	}
	LoadBase()
	LoadServer()
	LoadApp()

}
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

}
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriterTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECERT").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)

}
