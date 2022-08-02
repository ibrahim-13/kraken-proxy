package proxy

import (
	"io"
	"kraken-proxy/kraken"
	"net/http"
	"strings"
)

const (
	__krakenApiHost string = "api.kraken.com"
)

var _msgAcceptableHost []byte = []byte("Acceptable host: " + __krakenApiHost)

var uriBlocked map[string]bool = map[string]bool{
	"/favicon.ico": true,
}

var uriAccepted map[string]bool = map[string]bool{
	"/0/private/Balance":     true,
	"/0/private/AddOrder":    true,
	"/0/private/QueryOrders": true,
}

func (p *ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Host != __krakenApiHost {
		logInvalidHostRequest(r)
		w.WriteHeader(http.StatusMisdirectedRequest)
		w.Write([]byte("Invalid host: " + r.Host + "\n"))
		w.Write(_msgAcceptableHost)
		return
	}
	if r.Method != "POST" {
		logInvalidMethodError(r)
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
		r.ParseForm()
		krakenRequest := getSignedRequest(r, p.KrakenApiKey, p.KrakenPrivateKey)
		client := &http.Client{}
		response, err := client.Do(krakenRequest)
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

func NewProxyServer(apiKey string, privateKey string) *ProxyServer {
	return &ProxyServer{KrakenApiKey: apiKey, KrakenPrivateKey: privateKey}
}

func getSignedRequest(r *http.Request, apiKey string, privateKey string) *http.Request {
	sign := kraken.GetSignature(r.URL.Path, &r.PostForm, privateKey)
	krakenRequest, _ := http.NewRequest("POST", "https://"+__krakenApiHost+r.URL.Path, strings.NewReader(r.PostForm.Encode()))
	krakenRequest.Header["Content-Type"] = []string{"application/x-www-form-urlencoded; charset=utf-8"}
	krakenRequest.Header["API-Key"] = []string{apiKey}
	krakenRequest.Header["API-Sign"] = []string{sign}
	return krakenRequest
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
