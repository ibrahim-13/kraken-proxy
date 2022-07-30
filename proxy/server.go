package proxy

import (
	"fmt"
	"net/http"
)

func StartProxyServer(host string, port string, privateKey string) {
	http.HandleFunc("/", createProxyHandler(privateKey))
	fmt.Println(http.ListenAndServe(host+":"+port, nil))
}
