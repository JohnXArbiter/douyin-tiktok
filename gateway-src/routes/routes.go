package routes

type (
	Route struct {
		Id     string   `yaml:"Id"` // 服务发现中心注册的 name/key
		Prefix []string `yaml:"Prefix"`
	}
)
