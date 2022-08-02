package proxy

// https://gist.github.com/yowu/f7dc34bd4736a65ff28d

import (
	"net/http"
	"strings"
)

func isKrakenRequest(w http.ResponseWriter, r *http.Request) bool {
	if strings.HasPrefix(r.URL.Path, __krakenPathPrefix) {
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
		_, isBlocked := _uriBlocked[r.URL.Path]
		if isBlocked {
			logBlockedRequest(r)
			w.WriteHeader(http.StatusNotFound)
			return false
		}
	}
	return false
}
