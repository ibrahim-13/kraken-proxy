package main

import (
	"kraken-proxy/proxy"
	"kraken-proxy/util"
	"log"
)

func main() {
	log.Println("Starting server at localhost:9000")
	proxy.StartProxyServer("localhost", "9000", util.GetConfig().KrakenPrivateKey)
}
