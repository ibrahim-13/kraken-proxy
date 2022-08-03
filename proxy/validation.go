package proxy

import (
	"errors"
	"net/http"
	"strings"
)

func isKrakenRequest(w http.ResponseWriter, r *http.Request) (bool, error) {
	if strings.HasPrefix(r.URL.Path, __krakenPathPrefix) {
		if r.Host != __krakenApiHost {
			logInvalidHostRequest(r)
			w.WriteHeader(http.StatusMisdirectedRequest)
			w.Write([]byte("Invalid host: " + r.Host + "\n"))
			w.Write(_msgAcceptableHost)
			return false, errors.New("invalid host")
		}
		if r.Method != "POST" {
			logInvalidMethodError(r)
			return false, errors.New("invalid method")
		}
		_, isBlocked := _uriBlocked[r.URL.Path]
		if isBlocked {
			logBlockedRequest(r)
			w.WriteHeader(http.StatusNotFound)
			return false, errors.New("Blocked")
		}
		return true, nil
	}
	return false, nil
}
