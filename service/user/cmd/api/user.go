package main

import (
	douyinConfig "douyin-tiktok/common/config"
	"flag"
	"fmt"

	"douyin-tiktok/service/user/cmd/api/internal/config"
	"douyin-tiktok/service/user/cmd/api/internal/handler"
	"douyin-tiktok/service/user/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "service/user/cmd/api/etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	cc := douyinConfig.LoadConsulConf("service/user/cmd/api/etc/user-api.yaml")
	douyinConfig.LoadApiConf(cc, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
