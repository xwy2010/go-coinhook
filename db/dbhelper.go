package db

import (
	"fmt"
	"log"
	"sync"

	"github.com/xwy2010/go-coinhook/config"

	_ "github.com/go-sql-driver/mysql" //mysql驱动

	"xorm.io/xorm"
)

var (
	// MasterDB 主库引擎
	MasterDB *xorm.Engine
	lock     sync.Mutex
)

//TODO: 改成从配置文件中取
var mysqlConfig = map[string]string{
	"host":     config.Config().DB.Host,
	"port":     config.Config().DB.Port,
	"dbname":   config.Config().DB.DBname,
	"user":     config.Config().DB.User,
	"password": config.Config().DB.Password,
	"charset":  config.Config().DB.Charset,
}

func init() {

	// 启动时就打开数据库连接
	if err := initEngine(); err != nil {
		panic(err)
	}

	// 测试数据库连接是否 OK
	if err := MasterDB.Ping(); err != nil {
		panic(err)
	}

}

// initEngine 用于创建数据库实例
func initEngine() error {
	//TODO:
	if MasterDB != nil {
		return nil
	}
	lock.Lock()
	defer lock.Unlock()
	//再做一次判断，不然多节点时会多次创建
	if MasterDB != nil {
		return nil
	}
	c := mysqlConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
		c["user"], c["password"], c["host"], c["port"], c["dbname"])
	engine, err := xorm.NewEngine("mysql", driveSource)
	if err != nil {
		log.Fatal("dbConn error=", err)
		return err
	}
	MasterDB = engine
	//Debug模式，打印全部SQL语句。
	engine.ShowSQL(true)
	return nil
}
