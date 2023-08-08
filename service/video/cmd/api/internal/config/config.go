package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	UserRpc zrpc.RpcClientConf
	FileRpc zrpc.RpcClientConf

	Mysql struct {
		Dsn string
	}

	//Idgen struct {
	//	WorkerId uint16
	//}
}
