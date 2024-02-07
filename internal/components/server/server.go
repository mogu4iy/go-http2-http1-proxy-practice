package server

import (
	"http2-http1.1-proxy/internal/components/location"
	"http2-http1.1-proxy/internal/config"
)

type Server interface {
	ListenAndServe() error
	_mustImplementServer()
}

type server struct {
	Locations []location.Location
}

func (s *server) ListenAndServe() error {
	return nil
}

func (_ *server) _mustImplementServer() {}

func New(c config.Server) Server {
	s := &server{}
	//			for _, lC := range sC.Locations {
	//				l := location.New(lC)
	//			}
	return s
}
