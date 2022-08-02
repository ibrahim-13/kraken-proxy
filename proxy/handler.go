package proxy

import (
	"bytes"
	"io"
	"io/ioutil"
	"kraken-proxy/kraken"
	"net/http"
)

const (
	__krakenApiHost string = "api.kraken.com"
)

var _msgAcceptableHost []byte = []byte("Acceptable host: " + __krakenApiHost)

var uriBlocked map[string]bool = map[string]bool{
	"/favicon.ico": true,
}

var uriAccepted map[string]bool = map[string]bool{
	"/0/private/Balance":  true,
	"/0/private/AddOrder": true,
}

func (p *ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
		r.ParseForm()
		signRequest(r, p.KrakenApiKey, p.KrakenPrivateKey)
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

func NewProxyServer(apiKey string, privateKey string) *ProxyServer {
	return &ProxyServer{KrakenApiKey: apiKey, KrakenPrivateKey: privateKey}
}

func signRequest(r *http.Request, apiKey string, privateKey string) {
	sign := kraken.GetSignature(r.URL.Path, &r.Form, privateKey)
	r.Header["API-Key"] = []string{apiKey}
	r.Header["API-Sign"] = []string{sign}
	buff := bytes.NewBuffer([]byte(r.Form.Encode()))
	r.Body = ioutil.NopCloser(buff)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
