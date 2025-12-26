## httptrace

A tiny cli tool for tracing timestamps for http requests during debugging session and performance tests.

### Install

- For a Go-supporting machine:

Requires Go 1.20+

```bash
go install github.com/maen08/httptrace@latest

```

- For other kinds of machine, download a binary ready for usage

### Usage

```bash

httptrace https://example.com  
httptrace https://httpbin.org/post --method POST --json '{"hello":"world"}'

```

- Output:
```
DNS lookup      : 87.75225ms   [ IP:  104.18.26.120 ]
TCP connect     : 33.85475ms   [ Port:  443 ]
TLS handshake   : 76.363458ms  [ TLS: v1.3, Cipher: TLS_AES_128_GCM_SHA256 ]
Server response : 38.071292ms  [ StatusCode: 200 ]
Total           : 236.810709ms

```
