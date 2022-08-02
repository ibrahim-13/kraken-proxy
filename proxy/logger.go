package proxy

import (
	"fmt"
	"log"
	"net/http"
)

func logRequest(r *http.Request, param ...any) {
	log.Println(r.Method, r.URL.Path, fmt.Sprint(param...))
}

func logKrakenRequest(r *http.Request) {
	log.SetPrefix(__prefixKraken)
	logRequest(r)
	log.SetPrefix("")
}

func logBlockedRequest(r *http.Request) {
	log.SetPrefix(__prefixBlocked)
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

func logInvalidMethodError(r *http.Request) {
	log.SetPrefix(__prefixHttpRequestError)
	logRequest(r, "Invalid Method: "+r.Method+", Allowed method: POST")
	log.SetPrefix("")
}

func logOtherRequest(r *http.Request) {
	log.SetPrefix(__prefixOtherRequest)
	logRequest(r)
	log.SetPrefix("")
}
