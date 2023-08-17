package main

import (
	douyinConfig "douyin-tiktok/common/config"
	"flag"
	"fmt"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"

	"douyin-tiktok/service/file/cmd/rpc/internal/config"
	"douyin-tiktok/service/file/cmd/rpc/internal/server"
	"douyin-tiktok/service/file/cmd/rpc/internal/svc"
	"douyin-tiktok/service/file/cmd/rpc/types"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "service/file/cmd/rpc/etc/file-rpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	cc := douyinConfig.LoadConsulConf("service/file/cmd/rpc/etc/file-rpc.yaml")
	douyinConfig.LoadRpcConf(cc, &c)

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		__.RegisterFileServiceServer(grpcServer, server.NewFileServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	_ = consul.RegisterService(c.ListenOn, c.Consul)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
