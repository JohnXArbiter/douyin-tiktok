package server

type HttpServer struct {
	Addr         string
	Weight       int
	CWeight      int
	Status       int
	FailWeight   int
	FailCount    int
	RecoverCount int
}

func NewHttpServer(addr string, weight int) *HttpServer {
	return &HttpServer{
		Addr:    addr,
		Weight:  weight,
		Status:  StatusUp,
		CWeight: weight,
	}
}

type HttpServers []*HttpServer

func (s HttpServers) Len() int {
	return len(s)
}

func (s HttpServers) Less(i, j int) bool {
	return s[i].CWeight < s[j].CWeight
}

func (s HttpServers) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
