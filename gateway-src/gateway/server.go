package gateway

import (
	"bufio"
	"fmt"
	"gateway/config"
	"gateway/registry"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	Host     string
	Port     int
	ListenOn string
}

func MustNewServer(serverRegistry registry.ServerRegistry, conf *config.Conf) *Server {
	instance, _ := registry.NewDefaultServiceInstance(
		conf.RegistryConf.Key,
		conf.RegistryConf.Host,
		conf.RegistryConf.Port,
		false, nil, "",
	)

	listenOn := conf.GatewayConf.ListenOn
	if listenOn == "" {
		listenOn = conf.GatewayConf.Host + ":" + strconv.Itoa(conf.GatewayConf.Port)
	}

	serverRegistry.Register(instance, listenOn)
	serverRegistry.SetPredicates(conf.Routes)
	go serverRegistry.GetInstances()

	split := strings.Split(listenOn, ":")
	port, _ := strconv.Atoi(split[1])
	server := &Server{
		Host:     split[0],
		Port:     port,
		ListenOn: listenOn,
	}
	return server
}

func (s Server) Start() {
	ls, err := net.Listen("tcp", "0.0.0.0:10000")
	if err != nil {
		fmt.Printf("start tcp listener error: %v\n", err.Error())
		return
	}
	for {
		conn, err := ls.Accept()
		if err != nil {
			fmt.Printf("connect error: %v\n", err.Error())
		}
		go func(conn net.Conn) {
			_, err = bufio.NewWriter(conn).WriteString("ok")
			if err != nil {
				fmt.Printf("write conn error: %v\n", err)
			}
		}(conn)
	}
}

func (s Server) Stop() {

}
