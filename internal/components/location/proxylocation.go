package location

import "net/http"

type ProxyHandler struct{}

func (p *ProxyHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
	// TODO: implement proxy handler
}

type ProxyLocation struct {
	pattern string
	handler *ProxyHandler
}

func (pL *ProxyLocation) GetHandler() (string, http.Handler) {
	return pL.pattern, pL.handler
}

func (pL *ProxyLocation) SetPattern(pattern string) {
	pL.pattern = pattern
}

func (_ *ProxyLocation) _mustImplementLocation() {}
