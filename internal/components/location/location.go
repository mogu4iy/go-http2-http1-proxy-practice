package location

import (
	"http2-http1.1-proxy/internal/config"
	"net/http"
)

type Location interface {
	GetHandler() (string, http.Handler)
	SetPattern(string)
	_mustImplementLocation()
}

func New(c config.Location) Location {
	//	TODO: Parse c and based on configuration return Location
	l := &ProxyLocation{}
	l.SetPattern(c.Path)
	return l
}
