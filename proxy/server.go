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

func StartProxyServerTLS(conf *util.Config) {
	server := NewProxyServer(conf.KrakenApiKey, conf.KrakenPrivateKey)
	server.EnableOtherRequests = conf.EnableOtherRequest
	if !util.IsFileExist(conf.ServerCertPath) {
		panic("Certificate file not found:" + conf.ServerCertPath)
	}
	if !util.IsFileExist(conf.ServerKeyPath) {
		panic("Key file not found:" + conf.ServerKeyPath)
	}
	log.Println("Starting HTTPS server at " + conf.Host + ":" + conf.Port)
	err := http.ListenAndServeTLS(conf.Host+":"+conf.Port, conf.ServerCertPath, conf.ServerKeyPath, server)
	if err != nil {
		fmt.Println(err.Error())
	}
}
