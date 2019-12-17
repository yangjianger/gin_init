package app

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"os"
)

//初始化数据库引擎
func initDb(){
	var err error

	mysqlConfig := Cfg.Section("mysql")

	dbDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		mysqlConfig.Key("username").String(),
		mysqlConfig.Key("password").String(),
		mysqlConfig.Key("host").String(),
		mysqlConfig.Key("port").String(),
		mysqlConfig.Key("dbname").String(),
		mysqlConfig.Key("charset").String(),
	)

	Engine, err = xorm.NewEngine("mysql", dbDsn)
	if err != nil{
		panic(fmt.Errorf("xorm.NewEngine err = %v", err))
	}

	if ok, _ := mysqlConfig.Key("db_has_write_log").Bool(); ok{
		db_sql_file := mysqlConfig.Key("db_sql_file").String()
		f, err := os.OpenFile(db_sql_file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil{
			println(err.Error())
			return
		}
		//如果希望将信息不仅打印到控制台，而是保存为文件，那么可以通过类似如下的代码实现，NewSimpleLogger(w io.Writer)接收一个io.Writer接口来将数据写入到对应的设施中。
		Engine.SetLogger(xorm.NewSimpleLogger(f))
	}

	//则会在控制台打印出生成的SQL语句
	Engine.ShowSQL(true)

	//则会在控制台打印调试及以上的信息
	Engine.Logger().SetLevel(core.LOG_DEBUG)

}
