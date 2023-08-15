package registry

import (
	"gateway/routes"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type Conf struct {
	Key       string `yaml:"Key"`
	Host      string `yaml:"Host" json:",default=0.0.0.0"`
	Port      int    `yaml:"Port"`
	Token     string `yaml:"Token"`
	Frequency int64  `yaml:"Frequency"`
}

type (
	ServerInstance interface {
		GetHost() string                //
		GetPort() int                   //
		GetKey() string                 // 返回注册实例的服务名称
		IsSecure() bool                 // 返回注册实例是否使用 HTTPS
		GetMetadata() map[string]string // 返回关联注册实例的元数据键值对
	}

	DefaultServerInstance struct {
		Host     string
		Port     int
		Name     string
		Secure   bool
		Metadata map[string]string
	}
)

func NewDefaultServiceInstance(name string, host string, port int, secure bool, metadata map[string]string, instanceId string) (*DefaultServerInstance, error) {
	if len(host) == 0 {
		localIP, err := getLocalIP()
		if err != nil {
			return nil, err
		}
		host = localIP
	}

	if len(instanceId) == 0 {
		instanceId = strconv.FormatInt(time.Now().Unix(), 10) + "-" + strconv.Itoa(rand.Intn(9000)+1000)
	}
	return &DefaultServerInstance{Host: host, Port: port, Name: name, Secure: secure, Metadata: metadata}, nil
}

func (serviceInstance *DefaultServerInstance) GetHost() string {
	return serviceInstance.Host
}

func (serviceInstance *DefaultServerInstance) GetPort() int {
	return serviceInstance.Port
}

func (serviceInstance *DefaultServerInstance) GetKey() string {
	return serviceInstance.Name
}

func (serviceInstance *DefaultServerInstance) IsSecure() bool {
	return serviceInstance.Secure
}

func (serviceInstance *DefaultServerInstance) GetMetadata() map[string]string {
	return serviceInstance.Metadata
}

type ServerRegistry interface {
	Register(ServerInstance, string)
	Deregister()
	SetPredicates([]routes.Route)
	GetInstances()
}

func getLocalIP() (ipv4 string, err error) {
	addrs, err := net.InterfaceAddrs() // 获取所有网卡
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipNet, isIpNet := addr.(*net.IPNet)    // 这个网络地址是IP地址: ipv4, ipv6
		if isIpNet && !ipNet.IP.IsLoopback() { // 取第一个非lo的网卡IP
			if ipNet.IP.To4() != nil { // 跳过IPV6
				ipv4 = ipNet.IP.String() // 192.168.1.1
				return
			}
		}
	}
	return
}
