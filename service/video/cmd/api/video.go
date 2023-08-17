package main

import (
	douyinConfig "douyin-tiktok/common/config"
	"douyin-tiktok/service/video/cmd/api/internal/logic/mq"
	"flag"
	"fmt"

	"douyin-tiktok/service/video/cmd/api/internal/config"
	"douyin-tiktok/service/video/cmd/api/internal/handler"
	"douyin-tiktok/service/video/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

var configFile = flag.String("f", "service/video/cmd/api/etc/video-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	cc := douyinConfig.LoadConsulConf("service/video/cmd/api/etc/video-api.yaml")
	douyinConfig.LoadApiConf(cc, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	mq.InitRabbitMQ(ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
