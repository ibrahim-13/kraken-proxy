package proxy

import (
	"fmt"
	"kraken-proxy/util"
	"net/http"
)

func StartProxyServer(conf *util.Config) {
	server := NewProxyServer(conf.KrakenApiKey, conf.KrakenPrivateKey)
	server.EnableOtherRequests = conf.EnableOtherRequest
	var err error
	if conf.EnableSsl {
		if !util.IsFileExist(conf.ServerCertPath) {
			panic("Certificate file not found:" + conf.ServerCertPath)
		}
		if !util.IsFileExist(conf.ServerKeyPath) {
			panic("Key file not found:" + conf.ServerKeyPath)
		}
		err = http.ListenAndServeTLS(conf.Host+":"+conf.Port, conf.ServerCertPath, conf.ServerKeyPath, server)
	} else {
		err = http.ListenAndServe(conf.Host+":"+conf.Port, server)
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}
