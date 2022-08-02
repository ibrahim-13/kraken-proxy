package proxy

import (
	"io"
	"kraken-proxy/kraken"
	"net/http"
	"strings"
)

const (
	__krakenApiHost    string = "api.kraken.com"
	__krakenPathPrefix string = "/0/private/"
)

func (p *ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !isValidRequest(w, r) {
		return
	}
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
