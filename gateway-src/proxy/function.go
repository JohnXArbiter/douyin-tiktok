package proxy

import (
	"gateway/utils"
	"io"
	"log"
	"net/http"
)

func (p *ProxyHandler) DoRequest(w http.ResponseWriter, r *http.Request, url string) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
	}()

	serverReq, _ := http.NewRequest(r.Method, url, r.Body)
	utils.CopyHeader(&serverReq.Header, r.Header)
	serverReq.Header.Add("x-forwarded-for", r.RemoteAddr) // 将真实请求的 ip 传给服务器

	serverResp, _ := http.DefaultClient.Do(serverReq)

	w.WriteHeader(serverResp.StatusCode)
	header := w.Header()
	utils.CopyHeader(&header, serverResp.Header)

	defer serverResp.Body.Close()
	body, _ := io.ReadAll(serverResp.Body)
	w.Write(body)
	return

}
