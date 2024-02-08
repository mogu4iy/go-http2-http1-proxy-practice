package provider

import (
	"errors"
	"net/http"
)

type HTTP11Provider struct {
	server *http.Server
	mux    *http.ServeMux
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

func (h *HTTP11Provider) SetHandler(pattern string, handler http.Handler) {
	h.mux.Handle(pattern, handler)
}

func (h *HTTP11Provider) ListenAndServe() error {
	err := h.server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (h *HTTP11Provider) _mustImplementProvider() {}
