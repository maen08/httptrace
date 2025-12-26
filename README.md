## httptrace

A tiny cli tool for tracing timestamps for http requests during debugging session and performance tests.

### Why this ?
Scenario: assume your service takes too long to respond on a request (eg. 5 seconds) and you dont know at what stage does 
your request takes longer eg. at DNS resolution (to troubleshoot your DNS server), is it TLS handshaking stage (to fix your gateway server),
is it the connection itself to the server? So to avoid a guess-work hence `httptrace`

Originally inspired by the Golang `httptrace` stdlib itself - Read more: https://pkg.go.dev/net/http/httptrace

### Install

- For a Golang-supporting machine:

Requires Go 1.20+

```bash
go install github.com/maen08/httptrace@latest

```

- For other kinds of machine, download a binary ready for usage [here](https://github.com/maen08/httptrace/releases/tag/v0.1)

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
