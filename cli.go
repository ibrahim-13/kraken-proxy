package main

import (
	"kraken-proxy/proxy"
	"kraken-proxy/util"
)

func main() {
	config := util.GetConfig()
	if config.EnableSsl {
		proxy.StartProxyServerTLS(&config)
	} else {
		proxy.StartProxyServer(&config)
	}
}
