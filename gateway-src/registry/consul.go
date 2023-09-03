package registry

import (
	"context"
	"gateway/routes"
	"gateway/server"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/consul/api"
)

type ConsulRegistry struct {
	ListenOn            string
	client              *api.Client
	localServerInstance ServerInstance
	FetchInterval       int64
	Routes
	Predicates
}

func (c *ConsulRegistry) Register(serverInstance ServerInstance, listenOn string) {
	c.ListenOn = listenOn
	c.localServerInstance = serverInstance
	schema, tags := "http", make([]string, 0)

	id := serverInstance.GetKey() + "-" + serverInstance.GetHost() + "-" + strconv.Itoa(serverInstance.GetPort())
	registration := &api.AgentServiceRegistration{
		ID:      id,
		Name:    serverInstance.GetKey(),
		Address: serverInstance.GetHost(),
		Port:    serverInstance.GetPort(),
	}

	if serverInstance.IsSecure() {
		tags = append(tags, "secure=true")
	} else {
		tags = append(tags, "secure=false")
	}
	if serverInstance.GetMetadata() != nil {
		var tags []string
		for key, value := range serverInstance.GetMetadata() {
			tags = append(tags, key+"="+value)
		}
		registration.Tags = tags
	}
	registration.Tags = tags

	// consul健康检查回调函数
	registration.Check = &api.AgentServiceCheck{
		TCP:                            c.ListenOn, // 使用 TCP
		Timeout:                        "10s",
		Interval:                       "30s",
		DeregisterCriticalServiceAfter: "20s", // 故障检查失败20s后 consul自动将注册服务删除
		//HTTP:
	}

	if serverInstance.IsSecure() {
		schema = "https"
	}
	_ = schema + "://" + registration.Address + ":" + strconv.Itoa(registration.Port) + "/actuator/health"
	//check.HTTP = _ // 指定健康监测为 HTTP
	// 注册服务到consul
	if err := c.client.Agent().ServiceRegister(registration); err != nil {
		log.Fatalf("[FATAL REGISTRY] consul 网关注册失败 %v", err)
	}
}

func (c *ConsulRegistry) Deregister() {
	if c.localServerInstance == nil {
		return
	}
	_ = c.client.Agent().ServiceDeregister(c.localServerInstance.GetKey())
	c.localServerInstance = nil
}

func NewConsulRegistry(conf *Conf) *ConsulRegistry {
	if len(conf.Host) < 3 {
		log.Fatalf("[FATAL REGISTRY] consul 网关注册失败 check host")
	}

	if conf.Port <= 0 || conf.Port > 65535 {
		log.Fatalf("[FATAL REGISTRY] consul 网关注册失败 check port, port should between 1 and 65535")
	}

	apiConfig := api.DefaultConfig()
	apiConfig.Address = conf.Host + ":" + strconv.Itoa(conf.Port)
	apiConfig.Token = conf.Token
	client, err := api.NewClient(apiConfig)
	if err != nil {
		log.Fatalf("[FATAL REGISTRY] 网关注册失败 %v", err)
	}

	return &ConsulRegistry{client: client, FetchInterval: conf.Frequency}
}

func (c *ConsulRegistry) SetPredicates(rs []routes.Route, strategy int) {
	c.Routes = make(Routes)
	for _, route := range rs {
		lb := server.NewLoadBalanceStrategy(route.Id, strategy, route.RecoverThreshold, route.FailThreshold)
		c.Predicates = make(Predicates)
		c.Routes[route.Id] = lb
		for _, path := range route.Prefix {
			c.Predicates[path] = lb
		}
	}
}

func (c *ConsulRegistry) FetchInstances() {
	interval := time.Duration(c.FetchInterval) * time.Second
	ticker := time.NewTicker(interval)
	ctx, cancelFunc := context.WithTimeout(context.Background(), interval)
	defer cancelFunc()
	for {
		select {
		case <-ticker.C:
			for location, lb := range c.Routes {
				location1 := location
				lb1 := lb
				go func() {
					c.discovery(location1)
					lb1.GetChecker().CheckServers(ctx)
				}()

			}
		}
	}
}

func (c *ConsulRegistry) discovery(serviceName string) {
	var httpServers server.HttpServers
	services, _, err := c.client.Health().Service(serviceName, "", false, nil) // c.client.Catalog().Service() 这个是获取所有
	if err != nil {
		log.Printf("[ERROR DISCOVERY] 获取 %v 服务失败 %v\n", serviceName, err)
		return
	}

	for _, service := range services {
		weight := 10
		if service.Service.Meta != nil {
			if w := service.Service.Meta["weight"]; w != "" {
				weight, _ = strconv.Atoi(w)
			}
		}
		addr := service.Service.Address + ":" + strconv.Itoa(service.Service.Port)
		httpServer := server.NewHttpServer(addr, weight)
		httpServers = append(httpServers, httpServer)
	}

	lb := c.Routes[serviceName]
	lb.Lock().Lock()
	lb.SetServers(httpServers)
	lb.Lock().Unlock()
}
