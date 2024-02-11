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
}

type http struct {
	servers        []server.Server
	upstreamClient upstream.Client
}

func (h *http) Init() error {
	for _, s := range h.servers {
		//	TOOD: Implement error handling
		go s.ListenAndServe()
	}
	return nil
}

func (h *http) addServer(s server.Server) {
	h.servers = append(h.servers, s)
}

func (h *http) addUpstream(u upstream.Upstream) {
	h.upstreamClient.AddUpstream(u)
}

func New(c config.HTTP) (HTTP, error) {
	upstreamClient := upstream.NewClient()

	h := &http{
		upstreamClient: upstreamClient,
	}
	for _, sC := range c.Servers {
		s, err := server.New(sC, upstreamClient)
		if err != nil {
			return nil, err
		}
		h.addServer(s)
	}

	for _, uC := range c.Upstreams {
		u := upstream.New(uC)
		h.addUpstream(u)
	}

	return h, nil
}
