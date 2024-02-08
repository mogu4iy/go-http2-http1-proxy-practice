package upstream

import "http2-http1.1-proxy/internal/config"

// TODO: Implement upstream initialization and address resolving

type Upstream interface {
	RegisterServer(c config.UpstreamServer) error
	ResolveServer() error
	_mustImplementUpstream()
}

type Config struct {
	Name string `yaml:"name"`
}

type upstream struct {
	servers []any
}

func (u *upstream) RegisterServer(c config.UpstreamServer) error {
	return nil
}

func (u *upstream) ResolveServer() error {
	return nil
}

func (_ *upstream) _mustImplementUpstream() {}

func New(c config.Upstream) (Upstream, error) {
	u := &upstream{}
	for _, sC := range c.Addresses {
		err := u.RegisterServer(sC)
		if err != nil {
			return nil, err
		}
	}
	return u, nil
}
