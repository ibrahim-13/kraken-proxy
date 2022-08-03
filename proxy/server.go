package proxy

import (
	"fmt"
	"kraken-proxy/util"
	"log"
	"net/http"
)

func StartProxyServer(conf *util.Config) {
	server := NewProxyServer(conf.KrakenApiKey, conf.KrakenPrivateKey)
	server.EnableOtherRequests = conf.EnableOtherRequest
	log.Println("Starting HTTP server at " + conf.Host + ":" + conf.Port)
	err := http.ListenAndServe(conf.Host+":"+conf.Port, server)
	if err != nil {
		fmt.Println(err.Error())
	}
}
