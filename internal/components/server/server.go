package server

import (
	"fmt"
	"http2-http1.1-proxy/internal/components/location"
	"http2-http1.1-proxy/internal/config"
	"http2-http1.1-proxy/internal/provider"
)

type Server interface {
	Init() error
	AddLocation(l location.Location)
	_mustImplementServer()
}

type server struct {
	provider provider.HTTPProvider
}

func (s *server) Init() error {
	err := s.provider.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *server) AddLocation(l location.Location) {
	s.provider.SetHandler(l.GetHandler())
}

func (_ *server) _mustImplementServer() {}

func New(c config.Server) (Server, error) {
	s := &server{}

	s.provider = provider.New(c.Version)
	err := s.provider.Init(fmt.Sprintf(":%d", c.Listen))
	if err != nil {
		return nil, err
	}

	for _, lC := range c.Locations {
		l := location.New(lC)
		s.AddLocation(l)
	}
	return s, nil
}
