package proxy

import (
	"bytes"
	"io"
	"io/ioutil"
	"kraken-proxy/kraken"
	"net/http"
	"net/url"
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

func createProxyHandler(apiKey string, privateKey string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Host != __krakenApiHost {
			logInvalidHostRequest(r)
			w.WriteHeader(http.StatusMisdirectedRequest)
			w.Write([]byte("Invalid host: " + r.Host + "\n"))
			w.Write(_msgAcceptableHost)
			return
		}
		_, isBlocked := uriBlocked[r.URL.Path]
		if isBlocked {
			logBlockedRequest(r)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_, isAccepted := uriAccepted[r.URL.Path]
		if isAccepted {
			logAcceptedRequest(r)
			signRequest(r, apiKey, privateKey)
			r.RequestURI = ""
			client := &http.Client{}
			response, err := client.Do(r)
			if err != nil {
				logHttpRequestError(r, err)
				return
			}
			defer response.Body.Close()
			copyHeader(w.Header(), response.Header)
			w.WriteHeader(response.StatusCode)
			io.Copy(w, response.Body)
			return
		}
		logIgnoredRequest(r)
		w.WriteHeader(http.StatusForbidden)
	}
}

func signRequest(r *http.Request, apiKey string, privateKey string) {
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()
	values, _ := url.ParseQuery(string(body))
	sign := kraken.GetSignature(r.URL.Path, &values, privateKey)
	r.Header["API-Key"] = []string{apiKey}
	r.Header["API-Sign"] = []string{sign}
	buff := bytes.NewBuffer([]byte(values.Encode()))
	r.Body = ioutil.NopCloser(buff)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
