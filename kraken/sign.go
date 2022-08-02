package kraken

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"net/url"
	"strconv"
	"time"
)

const (
	__formKeyNonce string = "nonce"
)

func GetSignature(url_path string, payload *url.Values, privateKey string) string {

	sha := sha256.New()
	nonce := strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
	if payload.Has(__formKeyNonce) {
		payload.Del(__formKeyNonce)
	}
	payload.Add(__formKeyNonce, nonce)
	sha.Write([]byte(nonce + payload.Encode()))
	shasum := sha.Sum(nil)

	b64DecodedPrivateKey, _ := base64.StdEncoding.DecodeString(privateKey)
	mac := hmac.New(sha512.New, b64DecodedPrivateKey)
	mac.Write(append([]byte(url_path), shasum...))
	macsum := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(macsum)
}
