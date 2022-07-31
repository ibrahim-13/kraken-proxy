package proxy

import (
	"fmt"
	"net/http"
)

func StartProxyServer(host string, port string, apiKey string, privateKey string) {
	http.HandleFunc("/", createProxyHandler(apiKey, privateKey))
	fmt.Println(http.ListenAndServe(host+":"+port, nil))
}
