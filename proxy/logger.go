package proxy

import (
	"log"
	"net/http"
)

const (
	__prefixBlocked     string = "BLOCKED     "
	__prefixAccepted    string = "ACCEPTED    "
	__prefixIgnored     string = "IGNORED     "
	__prefixInvalidHost string = "INVALID_HOST"
)

func logRequest(r *http.Request) {
	log.Println(r.Method, r.RequestURI)
}

func logWhitelistedRequest(r *http.Request) {
	log.SetPrefix(__prefixAccepted)
	logRequest(r)
	log.SetPrefix("")
}

func logBlacklistedRequest(r *http.Request) {
	log.SetPrefix(__prefixBlocked)
	logRequest(r)
	log.SetPrefix("")
}

func logIgnoredRequest(r *http.Request) {
	log.SetPrefix(__prefixIgnored)
	logRequest(r)
	log.SetPrefix("")
}

func logInvalidHostRequest(r *http.Request) {
	log.SetPrefix(__prefixInvalidHost)
	logRequest(r)
	log.SetPrefix("")
}
