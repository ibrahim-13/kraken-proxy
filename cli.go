package main

import (
	"kraken-proxy/proxy"
	"kraken-proxy/util"
	"log"
)

func main() {
	config := util.GetConfig()
	log.Println("Starting server at " + config.Host + ":" + config.Port)
	proxy.StartProxyServer(config.Host, config.Port, config.KrakenApiKey, config.KrakenPrivateKey)
}
