package provider

import (
	"http2-http1.1-proxy/internal/constants"
	"net/http"
)

type HTTPProvider interface {
	Init(addr string) error
	SetHandler(pattern string, handler http.HandlerFunc)
	ListenAndServe() error
	_mustEmbedUnimplementedProvider()
}

type unimplementedHTTPProvider struct {
	server *http.Server
	mux    *http.ServeMux
}

func (h *unimplementedHTTPProvider) Init(addr string) error {
	return nil
}

func (h *unimplementedHTTPProvider) SetHandler(pattern string, handler http.HandlerFunc) {
	h.mux.HandleFunc(pattern, handler)
}

func (h *unimplementedHTTPProvider) ListenAndServe() error {
	err := h.server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (_ *unimplementedHTTPProvider) _mustEmbedUnimplementedProvider() {}

func New(version constants.HTTPVersion) HTTPProvider {
	switch version {
	case constants.HTTP11:
		return &HTTP11Provider{}
	case constants.HTTP2:
		return &HTTP2Provider{}
	}
	return nil
}
