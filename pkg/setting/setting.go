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
	Cfg, err := ini.Load("conf/app.ini") //读取配置文件
	if err != nil {
		log.Fatalf("读取配置文件错误:%v", err)

	}

}
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

}
