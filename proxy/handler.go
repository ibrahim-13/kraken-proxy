package proxy

import (
	"log"
	"net/http"
)

var uriBlacklist map[string]bool = map[string]bool{
	"/favicon.ico": true,
}

func createProxyHandler(privateKey string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, isBlacklisted := uriBlacklist[r.RequestURI]
		if isBlacklisted {
			logBlockedRequest(r)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		logRequest(r)
	}
}

func logRequest(r *http.Request) {
	log.Println(r.Method, r.RequestURI)
}

func logBlockedRequest(r *http.Request) {
	log.Println(r.Method, r.RequestURI, "BLOCKED")
}
