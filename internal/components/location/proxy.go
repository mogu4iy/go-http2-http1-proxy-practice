package location

import (
	"http2-http1.1-proxy/internal/config"
	"io"
	"log"
	"net/http"
)

func copyHeaders(dst http.Header, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func proxyHandler(rw http.ResponseWriter, req *http.Request) {
	//	TODO: Implement proxy handler based on location configuration
	//	client := &http.Client{}
	//
	//	resp, err := client.Do(req)
	//	if err != nil {
	//		http.Error(rw, "Server Error", http.StatusInternalServerError)
	//		log.Fatal("ServeHTTP:", err)
	//	}
	//			defer resp.Body.Close()
	//
	//	log.Println(req.RemoteAddr, " ", resp.Status)
	//	copyHeaders(rw.Header(), resp.Header)
	//
	//	rw.WriteHeader(resp.StatusCode)
	//	io.Copy(rw, resp.Body)
}

type ProxyLocation struct {
	headers config.ProxyHeaders
	UnimplementedLocation
}

func (l *ProxyLocation) init() {
	l.handler = proxyHandler
}
