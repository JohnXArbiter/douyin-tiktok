package server

import (
	"sort"
	"sync"
)

const (
	SmoothWeightedRoundRobinStrategy = iota
	RoundRobinStrategy
)

type LoadBalance interface {
	SelectServer() *HttpServer
	SetServers(servers HttpServers)
	Lock() *sync.RWMutex
}

func NewLoadBalanceStrategy(key string, strategy int) LoadBalance {

	switch strategy {
	case SmoothWeightedRoundRobinStrategy:
		return NewSmoothWeightedRoundRobin(key)
	case RoundRobinStrategy:
		return nil
	}
	return nil
}

//func init() {
//	strategy := NewLoadBalance()
//
//	for _, server := range strategy.Servers {
//		weightSum += server.Weight
//		weightSumList = append(weightSumList, weightSum)
//	}

// 加权随机 deprecated
//for idx, gateway := range strategy.Servers {
//	if gateway.Weight > 0 {
//		for i := 0; i < gateway.Weight; i++ {
//			ServerIndices = append(ServerIndices, idx)
//		}
//	}
//}
//fmt.Println(ServerIndices)
//}

// LoadBalance 负载均衡类
//type LoadBalance struct {
//	DiscoveryKey string
//	Servers   HttpServers // 服务器实例
//	ServerNum int
//	CurIndex  int // 轮询下标
//}
//
//func NewLoadBalance() *LoadBalance {
//	return &LoadBalance{Servers: make([]*HttpServer, 0)}
//}
//
//func (l *LoadBalance) AddServer(server *HttpServer) {
//	l.Servers = append(l.Servers, server)
//}
//
//// SelectByRand 随机
//func (l *LoadBalance) SelectByRand() *HttpServer {
//	var index = rand.Intn(len(l.Servers))
//	return l.Servers[index]
//}
//
//// SelectByIPHash ip 哈希
//func (l *LoadBalance) SelectByIPHash(ip string) *HttpServer {
//	var index = int(crc32.ChecksumIEEE([]byte(ip))) % len(l.Servers)
//	return l.Servers[index]
//}
//
//// SelectByWeightRand 加权随机 deprecated
////func (l *LoadBalance) SelectByWeightRand() *HttpServer {
////	var index = rand.Intn(len(ServerIndices))
////	return l.Servers[ServerIndices[index]]
////}
//
//// SelectByWeightRand 加权随机
//func (l *LoadBalance) SelectByWeightRand() *HttpServer {
//	var index = rand.Intn(weightSum)
//	for idx, v := range weightSumList {
//		if index < v {
//			return l.Servers[idx]
//		}
//	}
//	return l.Servers[0]
//}
//
//// SelectByRoundRobin 轮询
//func (l *LoadBalance) SelectByRoundRobin() *HttpServer {
//	var server = l.Servers[l.CurIndex]
//	l.CurIndex = (l.CurIndex + 1) % l.ServerNum
//	return server
//}
//
//func (l *LoadBalance) SelectByWeightRoundRobin() *HttpServer {
//	var server = l.Servers[0]
//	sum := 0
//	for i := 0; i < ServerNum; i++ {
//		sum += l.Servers[i].Weight
//		if l.CurIndex < sum {
//			server = l.Servers[i]
//			if l.CurIndex == sum-1 && l.CurIndex != l.ServerNum-1 {
//				l.CurIndex++ // 当 curIndex 还没走到最后一个服务器也不是当前权值-1的位置直接加1
//			} else {
//				l.CurIndex = (l.CurIndex + 1) % sum //
//			}
//			break
//		}
//	}
//	return server
//}
//
//// SelectBySmoothWeightedRoundRobin 平加权滑轮询
//func (l *LoadBalance) SelectBySmoothWeightedRoundRobin() *HttpServer {
//	for _, server := range l.Servers {
//		server.CWeight += server.Weight
//	}
//	sort.Sort(l.Servers)
//	target := l.Servers[0]
//	target.CWeight = target.CWeight - l.WeightSum
//	return target
//}

// SmoothWeightedRoundRobin 平加权滑轮询
type SmoothWeightedRoundRobin struct {
	DiscoveryKey string
	Servers      HttpServers   // 服务器实例
	ServerNum    int           // 服务实例数量
	CurIndex     int           // 轮询下标
	WeightSum    int           // 权重和
	lock         *sync.RWMutex // 防止并发问题
}

func NewSmoothWeightedRoundRobin(key string) LoadBalance {
	return &SmoothWeightedRoundRobin{
		DiscoveryKey: key,
		lock:         new(sync.RWMutex),
	}
}

func (l *SmoothWeightedRoundRobin) SelectServer() *HttpServer {
	for _, server := range l.Servers {
		server.CWeight += server.Weight - server.FailWeight
	}
	sort.Sort(l.Servers)
	target := l.Servers[0]
	target.CWeight = target.CWeight - l.WeightSum
	return target
}

func (l *SmoothWeightedRoundRobin) SetServers(servers HttpServers) {
	l.Servers = servers
	l.ServerNum = len(servers)
	for _, server := range l.Servers {
		l.WeightSum += server.Weight
	}
}

func (l *SmoothWeightedRoundRobin) Lock() *sync.RWMutex {
	return l.lock
}
