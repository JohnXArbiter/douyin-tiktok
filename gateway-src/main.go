package main

import (
	"fmt"
	"gateway/config"
	"gateway/gateway"
	"gateway/registry"
)

func main() {

	//registryDiscoveryClient := registry.NewConsulRegistry(nil)
	//
	//serviceInstanceInfo, _ := registry.NewDefaultServiceInstance("go-user-gateway", "no", 51425,
	//	false, map[string]string{"user": "zyn"}, "")
	//
	//registryDiscoveryClient.Register(serviceInstanceInfo)
	//
	//r := gin.Default()
	//// 健康检测接口，其实只要是 200 就认为成功了
	//r.GET("/actuator/health", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//err := r.Run(":8500")
	//if err != nil {
	//	registryDiscoveryClient.Deregister()
	//}
	test()
}

func test() {
	var conf config.Conf
	config.MustLoad("config.yaml", &conf)
	fmt.Println(conf)
	consulRegistry := registry.NewConsulRegistry(&conf.RegistryConf)

	server := gateway.MustNewServer(consulRegistry, &conf)
	server.Start()
	defer server.Stop()
}
