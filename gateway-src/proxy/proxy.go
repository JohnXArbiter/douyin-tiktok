package proxy

import (
	"gateway/config"
	"gateway/server"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
)

type ProxyHandler struct{}

func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
	}()

	_ = server.NewLoadBalance()

	for k, v := range config.ProxyConfig {
		if matched, _ := regexp.MatchString(k, r.URL.Path); matched {
			// go 自带的反向代理
			target, _ := url.Parse(v)
			proxy := httputil.NewSingleHostReverseProxy(target)
			proxy.ServeHTTP(w, r)
			//p.DoRequest(w, r, v)
			return
		}
	}
}

//服务器取出 ip 方式
//func (web1handler) GetIP(request *http.Request) string {
//	ips := request.Header.Get("x-forwarded-for")
//	if ips != "" {
//		ips_list := strings.Split(ips, ",")
//		if len(ips_list) > 0 && ips_list[0] != "" {
//			return ips_list[0]
//		}
//	}
//	return request.RemoteAddr
//}
