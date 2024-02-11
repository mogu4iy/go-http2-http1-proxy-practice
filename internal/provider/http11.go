package provider

import (
	"errors"
	"net/http"
)

type HTTP11Provider struct {
	unimplementedHTTPProvider
}

func (h *HTTP11Provider) Init(addr string) error {
	if h.server != nil {
		return errors.New("server is already initialized")
	}
	h.mux = http.NewServeMux()
	h.server = &http.Server{
		Addr:    addr,
		Handler: h.mux,
	}
	return nil
}