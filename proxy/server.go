package proxy

import (
	"fmt"
	"kraken-proxy/util"
	"net/http"
)

func StartProxyServer(conf *util.Config) {
	server := NewProxyServer(conf.KrakenApiKey, conf.KrakenPrivateKey)
	server.EnableOtherRequests = conf.EnableOtherRequest
	fmt.Println(http.ListenAndServe(conf.Host+":"+conf.Port, server))
}
