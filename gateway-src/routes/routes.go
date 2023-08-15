package routes

import "gateway/server"

type (
	Routes     map[string]*server.LoadBalance // 对应的负载均衡类
	Predicates map[string]*server.LoadBalance // url pattern prefix 匹配

	Route struct {
		Key        string   `yaml:"Key"` // 服务发现中心注册的 name/key
		Uri        string   `yaml:"Uri"`
		Predicates []string `yaml:"Predicates"`
	}
)
