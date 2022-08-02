package proxy

import (
	"io"
	"kraken-proxy/kraken"
	"net/http"
	"strings"
)

func (p *ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if isKrakenRequest(w, r) {
		logKrakenRequest(r)
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
	} else if !p.DisableOtherRequests {
		logOtherRequest(r)
		client := &http.Client{}
		delHopHeaders(r.Header)
		r.RequestURI = ""
		response, err := client.Do(r)
		if err != nil {
			logHttpRequestError(r, err)
			return
		}
		defer response.Body.Close()
		copyHeader(w.Header(), response.Header)
		w.WriteHeader(response.StatusCode)
		io.Copy(w, response.Body)
	} else {
		logBlockedRequest(r)
		w.WriteHeader(http.StatusForbidden)
		w.Write(_msgRequestBlocked)
	}
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
