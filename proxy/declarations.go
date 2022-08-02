package proxy

import "net/http"

const (
	__krakenApiHost    string = "api.kraken.com"
	__krakenPathPrefix string = "/0/private/"
)

const (
	__prefixBlocked          string = "BLOCKED          "
	__prefixKraken           string = "KRAKEN           "
	__prefixOtherRequest     string = "OTHER_REQUEST    "
	__prefixInvalidHost      string = "INVALID_HOST     "
	__prefixInvalidMethod    string = "INVALID_METHOD   "
	__prefixHttpRequestError string = "HTTP_REQUEST_ERR "
)

// Hop-by-hop headers. These are removed when sent to the backend.
// http://www.w3.org/Protocols/rfc2616/rfc2616-sec13.html
var hopHeaders = []string{
	"Connection",
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te", // canonicalized version of "TE"
	"Trailers",
	"Transfer-Encoding",
	"Upgrade",
}

var _msgAcceptableHost []byte = []byte("Acceptable host: " + __krakenApiHost)
var _msgRequestBlocked []byte = []byte("Request blocked by proxy")

var _uriBlocked map[string]bool = map[string]bool{
	"/favicon.ico": true,
}

type ProxyServer struct {
	http.Handler
	EnableOtherRequests bool
	KrakenApiKey        string
	KrakenPrivateKey    string
}
