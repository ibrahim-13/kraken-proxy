package proxy

import (
	"fmt"
	"log"
	"net/http"
)

const (
	__prefixBlocked          string = "BLOCKED         "
	__prefixAccepted         string = "ACCEPTED        "
	__prefixIgnored          string = "IGNORED         "
	__prefixInvalidHost      string = "INVALID_HOST    "
	__prefixHttpRequestError string = "HTTP_REQUEST_ERR"
)

func logRequest(r *http.Request, param ...any) {
	log.Println(r.Method, r.URL.Path, fmt.Sprint(param...))
}

func logAcceptedRequest(r *http.Request) {
	log.SetPrefix(__prefixAccepted)
	logRequest(r)
	log.SetPrefix("")
}

func logBlockedRequest(r *http.Request) {
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

func logHttpRequestError(r *http.Request, err error) {
	log.SetPrefix(__prefixHttpRequestError)
	logRequest(r, err)
	log.SetPrefix("")
}
