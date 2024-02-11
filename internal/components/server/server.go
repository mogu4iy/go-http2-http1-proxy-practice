package server

import (
	"fmt"
	"http2-http1.1-proxy/internal/components/location"
	"http2-http1.1-proxy/internal/components/upstream"
	"http2-http1.1-proxy/internal/config"
	"http2-http1.1-proxy/internal/provider"
)

type Server interface {
	addLocation(l location.Location)
	ListenAndServe() error
	_mustImplementServer()
}

type server struct {
	provider provider.HTTPProvider
}

func (s *server) addLocation(l location.Location) {
	pattern, handler := l.GetHandler()
	s.provider.SetHandler(pattern, handler)
}

func (s *server) ListenAndServe() error {
	err := s.provider.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (_ *server) _mustImplementServer() {}

func New(c config.Server, upstreamClient upstream.Client) (Server, error) {
	s := &server{}

	s.provider = provider.New(c.Version)
	err := s.provider.Init(fmt.Sprintf(":%d", c.Listen))
	if err != nil {
		return nil, err
	}

	for _, lC := range c.Locations {
		l := location.New(lC, upstreamClient)
		s.addLocation(l)
	}

	return s, nil
}
