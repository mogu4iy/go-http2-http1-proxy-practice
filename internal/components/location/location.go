package location

import (
	"http2-http1.1-proxy/internal/config"
	"net/http"
)

type Location interface {
	init()
	setPattern(pattern string)
	GetHandler() (string, http.HandlerFunc)
	_mustEmbedUnimplementedLocation()
}

type UnimplementedLocation struct {
	pattern string
	handler http.HandlerFunc
}

func (l *UnimplementedLocation) init() {}

func (l *UnimplementedLocation) setPattern(pattern string) {
	l.pattern = pattern
}

func (l *UnimplementedLocation) GetHandler() (string, http.HandlerFunc) {
	return l.pattern, l.handler
}

func (_ *UnimplementedLocation) _mustEmbedUnimplementedLocation() {}

func New(c config.Location) Location {
	var l Location
	if c.ProxyPass != "" {
		l = &ProxyLocation{}
	}
	l.setPattern(c.Path)
	l.init()
	return l
}
