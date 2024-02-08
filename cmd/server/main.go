package main

import (
	"http2-http1.1-proxy/internal/components/http"
	"http2-http1.1-proxy/internal/config"
	"log"
)

func main() {
	c, err := config.New("../configs")
	if err != nil {
		log.Fatalf("error parsing config: %v", err)
	}
	var HTTPs []http.HTTP
	for _, hC := range c.HTTPs {
		h, err := http.New(hC)
		if err != nil {
			log.Fatalf("error creating http: %v", err)
		}
		HTTPs = append(HTTPs, h)
	}
	for _, h := range HTTPs {
		//	TOOD: Implement error handling
		go h.Init()
	}
}
