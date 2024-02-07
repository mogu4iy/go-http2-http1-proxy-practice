package http

import (
	"http2-http1.1-proxy/internal/components/server"
	"http2-http1.1-proxy/internal/components/upstream"
	"http2-http1.1-proxy/internal/config"
)

type HTTP interface {
	AddServer(s server.Server)
	AddUpstream(u upstream.Upstream)
	Init()
	_mustImplementHTTP()
}

type http struct {
	Upstreams []upstream.Upstream
	Servers   []server.Server
}

func (h *http) AddServer(s server.Server) {
	h.Servers = append(h.Servers, s)
}

func (h *http) AddUpstream(u upstream.Upstream) {
	h.Upstreams = append(h.Upstreams, u)
}

func (h *http) Init() {}

func (h *http) _mustImplementHTTP() {}

func New(c config.HTTP) (HTTP, error) {
	h := &http{}
	for _, sC := range c.Servers {
		s := server.New(sC)
		h.AddServer(s)
	}
	for _, uC := range c.Upstreams {
		u, err := upstream.New(uC)
		if err != nil {
			return nil, err
		}
		h.AddUpstream(u)
	}
	return h, nil
}
