package http

import (
	"http2-http1.1-proxy/internal/components/server"
	"http2-http1.1-proxy/internal/components/upstream"
	"http2-http1.1-proxy/internal/config"
)

type HTTP interface {
	Init() error
	addServer(s server.Server)
	addUpstream(u upstream.Upstream)
	_mustImplementHTTP()
}

type http struct {
	dns dns
	servers []server.Server
}

func (h *http) addServer(s server.Server) {
	h.servers = append(h.servers, s)
}

func (h *http) addUpstream(u upstream.Upstream) {}

func (h *http) Init() error {
	for _, s := range h.servers {
		//	TOOD: Implement error handling
		go s.ListenAndServe()
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
		h.addServer(s)
	}

	for _, uC := range c.Upstreams {
		u, err := upstream.New(uC)
		if err != nil {
			return nil, err
		}
		h.addUpstream(u)
	}

	return h, nil
}
