package utils

import (
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

type Postgresql struct {
	Dsn string
}

func InitXorm(dbtype string, mc Postgresql) *xorm.Engine {

	engine, err := xorm.NewEngine(dbtype, mc.Dsn)
	logx.Infof("[XORM CONNECTING] Init Xorm DSN: %v", mc.Dsn)
	if err != nil {
		panic("[XORM ERROR] NewServiceContext 获取pgsql连接错误 " + err.Error())
	}
	err = engine.Ping()
	if err != nil {
		panic("[XORM ERROR] NewServiceContext ping pgsql 失败" + err.Error())
	}
	return engine
}
