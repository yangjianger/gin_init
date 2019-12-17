package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

//初始化 日志
func initLog(){
	//记录json 格式数据
	Log.Formatter = &logrus.JSONFormatter{}

	//指定日志输入
	logConfig := Cfg.Section("log")
	// 典型读取操作，默认分区可以使用空字符串表示
	log_file := logConfig.Key("log_file").String()

	file, err := os.OpenFile(log_file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil{
		fmt.Println("os.OpenFile err = ", err)
		return
	}
	Log.Out = file

	//告诉gin框架，把他的日志记录到文件中
	gin.SetMode(gin.ReleaseMode) //设置ReleaseMode模式
	//可以往多个地方写
	//gin.DefaultWriter = io.MultiWriter(log.Out, os.Stdout)
	gin.DefaultWriter = Log.Out

	//设置日志级别
	Log.Level = logrus.DebugLevel
}