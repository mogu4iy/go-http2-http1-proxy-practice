package upstream

import (
	"http2-http1.1-proxy/internal/config"
	"sync"
)

type Upstream interface {
	getName() string
	registerServer(c config.UpstreamServer)
	NextServer() Server
}

type upstream struct {
	name         string
	servers      []Server
	backupServer Server
	totalWeight  int
	index        int
	sync.Mutex
}

func (u *upstream) getName() string {
	return u.name
}

func (u *upstream) registerServer(c config.UpstreamServer) {
	s := &server{
		addr:   c.Server,
		weight: c.Weight,
	}
	u.servers = append(u.servers, s)

	if c.Backup {
		u.backupServer = s
	}

	u.totalWeight += c.Weight
}

func (u *upstream) NextServer() Server {
	u.Lock()
	defer u.Unlock()

	u.index = (u.index + 1) % u.totalWeight

	weightSum := 0
	for _, s := range u.servers {
		weightSum += s.getWeight()
		if u.index < weightSum {
			return s
		}
	}

	return u.backupServer
}

func New(c config.Upstream) Upstream {
	u := &upstream{
		index: -1,
	}

	for _, sC := range c.Addresses {
		u.registerServer(sC)
	}
	
	return u
}
