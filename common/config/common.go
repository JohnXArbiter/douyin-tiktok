package config

import (
	"github.com/hashicorp/consul/api"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gopkg.in/yaml.v3"
	"os"
)

type (
	Consul struct {
		Addr string `yaml:"Addr"`
		Key  string `yaml:"Key"`
	}

	ApiConf struct {
		rest.RestConf
		//Mysql    utils.Mysql
		//Redis    utils.Redis
		//Mongo    utils.Mongo
		//RabbitMQ utils.RabbitMQConf
		//Idgen    struct {
		//	WorkerId uint16
		//}
	}

	RpcConf struct {
		zrpc.RpcServerConf
		//Consul   consul.Conf        `yaml:"Consul"`
		//Mysql    utils.Mysql        `yaml:"Mysql"`
		//Redis    utils.Redis        `yaml:"Redis"`
		//Mongo    utils.Mongo        `yaml:"Mongo"`
		//RabbitMQ utils.RabbitMQConf `yaml:"RabbitMQ"`
		//Idgen    struct {
		//	WorkerId uint16
		//}
	}

	ConsulConf struct {
		Consul Consul `yaml:"Consul"`
	}
)

func LoadConsulConf(filePath string) *Consul {
	var cc ConsulConf
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		panic("配置文件读取错误：" + err.Error())
	}
	if err = yaml.Unmarshal(yamlFile, &cc); err != nil {
		panic("配置文件解析错误：" + err.Error())
	}
	return &cc.Consul
}

func LoadApiConf(cc *Consul, ac interface{}) interface{} {
	var client, _ = api.NewClient(&api.Config{Address: cc.Addr})
	kv := client.KV()
	data, _, err := kv.Get(cc.Key, nil)
	logx.Must(err)
	_ = conf.LoadFromYamlBytes(data.Value, ac)
	return ac
}

func LoadRpcConf(cc *Consul, ttr interface{}) {
	var client, _ = api.NewClient(&api.Config{Address: cc.Addr})
	kv := client.KV()
	data, _, err := kv.Get(cc.Key, nil)
	logx.Must(err)
	_ = conf.LoadFromYamlBytes(data.Value, ttr)
}
