package proxy

import (
	"net/http"
	"strings"
)

var _msgAcceptableHost []byte = []byte("Acceptable host: " + __krakenApiHost)

var uriBlocked map[string]bool = map[string]bool{
	"/favicon.ico": true,
}

func isValidRequest(w http.ResponseWriter, r *http.Request) bool {
	if r.Host != __krakenApiHost {
		logInvalidHostRequest(r)
		w.WriteHeader(http.StatusMisdirectedRequest)
		w.Write([]byte("Invalid host: " + r.Host + "\n"))
		w.Write(_msgAcceptableHost)
		return false
	}
	if r.Method != "POST" {
		logInvalidMethodError(r)
		return false
	}
	_, isBlocked := uriBlocked[r.URL.Path]
	if isBlocked {
		logBlockedRequest(r)
		w.WriteHeader(http.StatusNotFound)
		return false
	}
	if strings.HasPrefix(r.URL.Path, __krakenPathPrefix) {
		return true
	}
	logIgnoredRequest(r)
	w.WriteHeader(http.StatusForbidden)
	return false
}
