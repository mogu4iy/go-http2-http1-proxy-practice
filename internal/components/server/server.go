package server

import (
	"fmt"
	"http2-http1.1-proxy/internal/components/location"
	"http2-http1.1-proxy/internal/config"
	"http2-http1.1-proxy/internal/provider"
	"net/http"
)

type Server interface {
	addLocation(l location.Location)
	serverHandlerDecorator(h http.HandlerFunc) http.HandlerFunc
	ListenAndServe() error
	_mustImplementServer()
}

type server struct {
	provider provider.HTTPProvider
}

func (s *server) serverHandlerDecorator(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}

func (s *server) addLocation(l location.Location) {
	pattern, handler := l.GetHandler()
	s.provider.SetHandler(pattern, s.serverHandlerDecorator(handler))
}

func (s *server) ListenAndServe() error {
	err := s.provider.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
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
		s.addLocation(l)
	}

	return s, nil
}
