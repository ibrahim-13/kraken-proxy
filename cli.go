package main

import (
	"kraken-proxy/proxy"
	"kraken-proxy/util"
)

func main() {
	config := util.GetConfig()
	proxy.StartProxyServer(&config)
}
