package proxy

import (
	"fmt"
	"kraken-proxy/util"
	"net/http"
)

func StartProxyServer(conf *util.Config) {
	server := NewProxyServer(conf.KrakenApiKey, conf.KrakenPrivateKey)
	server.DisableOtherRequests = conf.DisableOtherRequest
	fmt.Println(http.ListenAndServe(conf.Host+":"+conf.Port, server))
}
