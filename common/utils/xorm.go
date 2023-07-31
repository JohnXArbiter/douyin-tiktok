package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

type Mysql struct {
	Dsn string
}

func InitXorm(dbtype string, mc Mysql) *xorm.Engine {
	engine, err := xorm.NewEngine(dbtype, mc.Dsn)
	logx.Infof("[XORM CONNECTING] Init Xorm DSN: %v", mc.Dsn)
	if err != nil {
		panic("[XORM ERROR] NewServiceContext 获取mysql连接错误 " + err.Error())
	}
	err = engine.Ping()
	if err != nil {
		panic("[XORM ERROR] NewServiceContext ping mysql 失败" + err.Error())
	}
	return engine
}
