package location

import (
	"http2-http1.1-proxy/internal/components/upstream"
	"http2-http1.1-proxy/internal/config"
	"net/http"
)

type Location interface {
	init()
	setPattern(p string)
	setUpstreamClient(uC upstream.Client)
	GetHandler() (string, http.HandlerFunc)
	_mustEmbedUnimplementedLocation()
}

type UnimplementedLocation struct {
	upstreamClient upstream.Client
	pattern        string
	handler        http.HandlerFunc
}

func (l *UnimplementedLocation) init() {}

func (l *UnimplementedLocation) setPattern(pattern string) {
	l.pattern = pattern
}

func (l *UnimplementedLocation) setUpstreamClient(upstreamClient upstream.Client) {
	l.upstreamClient = upstreamClient
}

func (l *UnimplementedLocation) GetHandler() (string, http.HandlerFunc) {
	return l.pattern, l.handler
}

func (_ *UnimplementedLocation) _mustEmbedUnimplementedLocation() {}

func New(c config.Location, upstreamClient upstream.Client) Location {
	var l Location
	if c.ProxyPass != "" {
		l = &ProxyLocation{
			headers: c.ProxyHeaders,
		}
	}
	l.setUpstreamClient(upstreamClient)
	l.setPattern(c.Path)
	l.init()
	return l
}
