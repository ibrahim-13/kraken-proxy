package proxy

import (
	"fmt"
	"net/http"
)

type ProxyServer struct {
	http.Handler
	KrakenApiKey     string
	KrakenPrivateKey string
}

func StartProxyServer(host string, port string, apiKey string, privateKey string) {
	server := NewProxyServer(apiKey, privateKey)
	fmt.Println(http.ListenAndServe(host+":"+port, server))
}
