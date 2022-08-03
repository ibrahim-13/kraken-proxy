# Kraken-Proxy

>[Kraken REST API Documentation](https://docs.kraken.com/rest/)

A proxy server that handles the authentication mechanism when calling Kraken REST APIs. It will-
  - Add `API-Key` header
  - Generate and add `nonce` in payload
  - Generate and add `API-Sign` header

By default, this proxy only executes requests for Kraken Private APIs. This can be changed to execute all requests from `config.json` file.

## Configuration

**Filename:** `config.json`

**Allowed host:** `api.kraken.com`

**Allowed path prefix:** `/0/private/`

```json
{
    "KrakenApiKey": "",
    "KrakenPrivateKey": "",
    "Host": "",
    "Port": "",
    "ServerCertPath": "",
    "ServerKeyPath": "",
    "EnableSsl": false,
    "EnableOtherRequest": false
}
```

| Field                | Description                               |
| -------------------- | ----------------------------------------- |
| `KrakenApiKey`       | API key for Kraken                        |
| `KrakenPrivateKey`   | Secret Key for Kraken                     |
| `Host`               | Host address for running the proxy server |
| `Port`               | Port for running the proxy server         |
| `ServerCertPath`     | Certificate file location                 |
| `ServerKeyPath`      | Private Key file location                 |
| `EnableSsl`          | Run HTTPS proxy serve r                   |
| `EnableOtherRequest` | Proxy for all destination addresses       |

> When `EnableOtherRequest` is enabled, the proxy will execute all proxy request, but will only add authentication and authorization when the target host is Kraken.

> Running the proxy without the configuration file, the application will panic and emit an example configuration file.

## Logging
### Formatting
```
PREFIX METHOD PATH OTHER_INFO...
```
| Prefix             | Description                                 |
| ------------------ | ------------------------------------------- |
| `BLOCKED`          | Request has been blocked                    |
| `KRAKEN`           | Kraken request                              |
| `OTHER_REQUEST`    | Request other then Kraken API               |
| `INVALID_HOST`     | Host does not match the Kraken API host     |
| `INVALID_METHOD`   | Kraken request made without **POST** method |
| `HTTP_REQUEST_ERR` | Error when making request                   |

## HTTPS Server
### Generate private key (.key)
```sh
# Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out server.key 2048

# Key considerations for algorithm "ECDSA" (X25519 || ≥ secp384r1)
# https://safecurves.cr.yp.to/
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out server.key
```
### Generate public key (.crt)
```sh
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

### Generation of self-signed certificates with one command
```sh
# ECDSA recommendation key ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl req -x509 -nodes -newkey ec:secp384r1 -keyout server.ecdsa.key -out server.ecdsa.crt -days 3650
# openssl req -x509 -nodes -newkey ec:<(openssl ecparam -name secp384r1) -keyout server.ecdsa.key -out server.ecdsa.crt -days 3650
# -pkeyopt ec_paramgen_curve:… / ec:<(openssl ecparam -name …) / -newkey ec:…
ln -sf server.ecdsa.key server.key
ln -sf server.ecdsa.crt server.crt

# RSA recommendation key ≥ 2048-bit
openssl req -x509 -nodes -newkey rsa:2048 -keyout server.rsa.key -out server.rsa.crt -days 3650
ln -sf server.rsa.key server.key
ln -sf server.rsa.crt server.crt
```

## Resources
  - Simple Golang HTTPS/TLS Examples: [denji/golang-tls](https://github.com/denji/golang-tls)
  - A simple HTTP proxy by Golang: [yowu/f7dc34bd4736a65ff28d](https://gist.github.com/yowu/f7dc34bd4736a65ff28d)