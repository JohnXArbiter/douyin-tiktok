package main

import (
	"flag"
	"fmt"

	"douyin-tiktok/service/file/cmd/api/internal/config"
	"douyin-tiktok/service/file/cmd/api/internal/handler"
	"douyin-tiktok/service/file/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "service/file/cmd/api/etc/file-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
