package server

// import (
//
//	"log"
//	"net/http"
//	"time"
//
// )
const (
	StatusUp   = 0
	StatusDown = -1
)

//
//type HttpChecker struct {
//	Servers          HttpServers
//	FailThreshold    int // 连续失败阈值
//	RecoverThreshold int // 连续成功阀值
//}
//
//func NewHttpChecker(servers HttpServers, failThreshold, recoverThreshold int) *HttpChecker {
//	return &HttpChecker{
//		Servers:          servers,
//		FailThreshold:    failThreshold,
//		RecoverThreshold: recoverThreshold,
//	}
//}
//
//func checkServers(servers HttpServers, failThreshold, recoverThreshold int) {
//	var t = time.NewTicker(time.Second * 3)
//	check := NewHttpChecker(servers, failThreshold, recoverThreshold)
//	for {
//		select {
//		case <-t.C:
//			check.check()
//			for _, server := range servers {
//				log.Printf("[GATEWAY HEALTH CHECK] server: %v, status: %v, fail count: %v\n", server.Host, server.Status, server.FailCount)
//			}
//			log.Println("--------------------------------------------------------")
//		}
//	}
//}
//
//func (hc *HttpChecker) check() {
//	var client = http.Client{}
//	for _, server := range hc.Servers {
//		res, err := client.Head(server.Host)
//		if res != nil {
//			defer res.Body.Close()
//		}
//		if err != nil {
//			hc.fail(server)
//			continue
//		}
//		if res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusBadRequest {
//			hc.success(server)
//		} else {
//			hc.fail(server)
//		}
//	}
//}
//
//func (hc *HttpChecker) fail(server *HttpServer) {
//	if server.FailCount >= hc.FailThreshold {
//		server.Status = StatusDown
//	} else {
//		server.FailCount++
//	}
//	server.RecoverCount = 0
//}
//
//func (hc *HttpChecker) success(server *HttpServer) {
//	if server.FailCount > 0 {
//		server.FailCount--
//		server.RecoverCount++
//		if server.RecoverCount == hc.RecoverThreshold {
//			server.Status = StatusUp
//			server.FailCount = 0
//			server.RecoverCount = 0
//		}
//	} else {
//		server.Status = StatusUp
//	}
//}
//
//func (hc *HttpChecker) fail4SWRR(server *HttpServer) {
//	hc.Servers
//	if server.FailCount >= hc.FailThreshold {
//		server.Status = StatusDown
//	} else {
//		server.FailCount++
//	}
//	server.RecoverCount = 0
//}
