package kraken

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"net/url"
)

func GetSignature(url_path string, values url.Values, privateKey string) string {

	sha := sha256.New()
	sha.Write([]byte(values.Get("nonce") + values.Encode()))
	shasum := sha.Sum(nil)

	b64DecodedPrivateKey, _ := base64.StdEncoding.DecodeString(privateKey)
	mac := hmac.New(sha512.New, b64DecodedPrivateKey)
	mac.Write(append([]byte(url_path), shasum...))
	macsum := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(macsum)
}
