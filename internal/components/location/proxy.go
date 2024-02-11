package location

import "net/http"

func proxyHandler(http.ResponseWriter, *http.Request) {
	// TODO: implement proxy handler
}

type ProxyLocation struct {
	UnimplementedLocation
}

func (l *ProxyLocation) init() {
	l.handler = proxyHandler
}
