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
    "EnableOtherRequest": false
}
```

| Field                | Description                               |
| -------------------- | ----------------------------------------- |
| `KrakenApiKey`       | API key for Kraken                        |
| `KrakenPrivateKey`   | Secret Key for Kraken                     |
| `Host`               | Host address for running the proxy server |
| `Port`               | Port for running the proxy server         |
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

## Resources
  - A simple HTTP proxy by Golang: [yowu/f7dc34bd4736a65ff28d](https://gist.github.com/yowu/f7dc34bd4736a65ff28d)