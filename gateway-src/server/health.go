package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

const (
	StatusUp   = 0
	StatusDown = -1
)

type HttpChecker struct {
	Servers          HttpServers
	RecoverThreshold int // 连续成功阀值
	FailThreshold    int // 连续失败阈值
	LoadBalance
}

func NewHttpChecker(servers HttpServers, recoverThreshold, failThreshold int) *HttpChecker {
	return &HttpChecker{
		Servers:          servers,
		FailThreshold:    failThreshold,
		RecoverThreshold: recoverThreshold,
	}
}

func (hc *HttpChecker) CheckServers(ctx context.Context) {
	var t = time.NewTicker(time.Second * 3)
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			hc.check()
			for _, server := range hc.Servers {
				log.Printf("[GATEWAY HEALTH CHECK] server: %v, status: %v, fail count: %v\n", server.Addr, server.Status)
			}
			log.Println("--------------------------------------------------------")
		}
	}
}

func (hc *HttpChecker) check() {
	var client = http.Client{}
	for _, server := range hc.Servers {
		res, err := client.Head(server.Addr)
		if res != nil {
			defer res.Body.Close()
		}
		if err != nil {
			hc.fail(server)
			continue
		}
		if res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusBadRequest {
			hc.success(server)
		} else {
			hc.fail(server)
		}
	}
}

func (hc *HttpChecker) fail(server *HttpServer) {
	if server.FailCount >= hc.FailThreshold {
		server.Status = StatusDown
	} else {
		server.FailCount++
	}
	server.RecoverCount = 0
}

func (hc *HttpChecker) success(server *HttpServer) {
	if server.FailCount > 0 {
		server.FailCount--
		server.RecoverCount++
		if server.RecoverCount == hc.RecoverThreshold {
			server.Status = StatusUp
			server.FailCount = 0
			server.RecoverCount = 0
		}
	} else {
		server.Status = StatusUp
	}
}

func (hc *HttpChecker) fail4SWRR(server *HttpServer) {
	if server.FailCount >= hc.FailThreshold {
		server.Status = StatusDown
	} else {
		server.FailCount++
		server.FailWeight++
		hc.LoadBalance.CalculateSum()
	}
	server.RecoverCount = 0
}
