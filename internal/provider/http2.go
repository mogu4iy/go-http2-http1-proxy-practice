package provider

import (
	"errors"
	"golang.org/x/net/http2"
	"net/http"
)

type HTTP2Provider struct {
	unimplementedHTTPProvider
}

func (h *HTTP2Provider) Init(addr string) error {
	if h.server != nil {
		return errors.New("server is already initialized")
	}
	h.mux = http.NewServeMux()
	h.server = &http.Server{
		Addr:    addr,
		Handler: h.mux,
	}
	http2Server := &http2.Server{}
	err := http2.ConfigureServer(h.server, http2Server)
	if err != nil {
		return err
	}
	return nil
}
