package upstream

import (
	"net"
)

type Client interface {
	AddUpstream(upstream Upstream)
	ResolveAddr(addr string) string
}

type client struct {
	upstreams map[string]Upstream
}

func (c *client) AddUpstream(upstream Upstream) {
	c.upstreams[upstream.getName()] = upstream
}

func (c *client) ResolveAddr(addr string) string {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return ""
	}

	u, ok := c.upstreams[host]
	if !ok {
		return addr
	}

	return u.NextServer().GetAddress()
}

func NewClient() Client {
	return &client{}
}
