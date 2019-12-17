package app

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"os"
)

//初始化数据库
//一个Orm引擎称为Engine，一个Engine一般只对应一个数据库
var (
	Engine *xorm.Engine
	Log = logrus.New()
	RedisClient *redis.Client
	Cfg *ini.File
	APP_ROOT,_ = os.Getwd()
	CONFIG_PATH = APP_ROOT + "/config"

)


//初始化
func InitApp(){
	var err error

	Cfg, err = ini.Load(CONFIG_PATH + "/config.ini")
	if err != nil{
		fmt.Print("Fail to read file: %v", err)
		os.Exit(1)
	}

	initDb()

	initLog()

	initRedis()
}
