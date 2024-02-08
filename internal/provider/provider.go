package provider

import (
	"http2-http1.1-proxy/internal/constants"
	"net/http"
)

type HTTPProvider interface {
	Init(addr string) error
	SetHandler(pattern string, handler http.Handler)
	ListenAndServe() error
	_mustImplementProvider()
}

func New(version constants.HTTPVersion) HTTPProvider {
	switch version {
	case constants.HTTP11:
		return &HTTP11Provider{}
	case constants.HTTP2:
		return &HTTP2Provider{}
	}
	return nil
}
