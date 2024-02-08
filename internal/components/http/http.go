package http

import (
	"http2-http1.1-proxy/internal/components/server"
	"http2-http1.1-proxy/internal/components/upstream"
	"http2-http1.1-proxy/internal/config"
)

type HTTP interface {
	AddServer(s server.Server)
	AddUpstream(u upstream.Upstream)
	Init() error
	_mustImplementHTTP()
}

type http struct {
	Servers []server.Server
}

func (h *http) AddServer(s server.Server) {
	h.Servers = append(h.Servers, s)
}

func (h *http) AddUpstream(u upstream.Upstream) {}

func (h *http) Init() error {
	for _, s := range h.Servers {
		//	TOOD: Implement error handling
		go s.Init()
	}
	return nil
}

func (h *http) _mustImplementHTTP() {}

func New(c config.HTTP) (HTTP, error) {
	h := &http{}
	for _, sC := range c.Servers {
		s, err := server.New(sC)
		if err != nil {
			return nil, err
		}
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
