package proxy

import (
	"net/http"
)

const (
	__krakenApiHost string = "api.kraken.com"
)

var _msgAcceptableHost []byte = []byte("Acceptable host: " + __krakenApiHost)

var uriBlocked map[string]bool = map[string]bool{
	"/favicon.ico": true,
}

var uriAccepted map[string]bool = map[string]bool{
	"/0/private/Balance": true,
}

func createProxyHandler(privateKey string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Host != __krakenApiHost {
			logIgnoredRequest(r)
			w.WriteHeader(http.StatusMisdirectedRequest)
			w.Write([]byte("Invalid host: " + r.Host + "\n"))
			w.Write(_msgAcceptableHost)
			return
		}
		_, isBlocked := uriBlocked[r.RequestURI]
		if isBlocked {
			logBlacklistedRequest(r)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_, isAccepted := uriAccepted[r.RequestURI]
		if isAccepted {
			logWhitelistedRequest(r)
			return
		}
		logIgnoredRequest(r)
	}
}
