package trace

import (
	"bytes"
	"context"
	"crypto/tls"
	"net/http"
	"net/http/httptrace"
	"strings"
	"time"

)


type Options struct {
	URL    string
	Method string
	JSON   string
}

type Result struct {
	DNSDuration      time.Duration
	TCPDuration      time.Duration
	TLSDuration      time.Duration
	ResponseDuration time.Duration
	TotalDuration    time.Duration

	DNSIP      string
	Port       string
	TLSVersion string
	Cipher     string
	StatusCode int
}

func Run(ctx context.Context, opt Options) (*Result, error) {
	result := &Result{}

	start := time.Now()
	last := start

	var bodyReader *bytes.Reader

	if opt.JSON != "" {
		bodyReader = bytes.NewReader([]byte(opt.JSON))
	} else {
		bodyReader = bytes.NewReader(nil)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		strings.ToUpper(opt.Method),
		opt.URL,
		bodyReader,
	)
	if err != nil {
		return nil, err
	}

	if opt.JSON != "" {
		req.Header.Set("Content-Type", "application/json")
	}

	trace := &httptrace.ClientTrace{
		DNSStart: func(httptrace.DNSStartInfo) {
			last = time.Now()
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			result.DNSDuration = time.Since(last)
			if len(info.Addrs) > 0 {
				result.DNSIP = info.Addrs[0].IP.String()
			}
			last = time.Now()
		},

		ConnectStart: func(network, addr string) {
			last = time.Now()
		},
		ConnectDone: func(network, addr string, err error) {
			result.TCPDuration = time.Since(last)
			result.Port = addr[strings.LastIndex(addr, ":")+1:]
			last = time.Now()
		},

		TLSHandshakeStart: func() {
			last = time.Now()
		},
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			result.TLSDuration = time.Since(last)
			result.TLSVersion = tlsVersion(state.Version)
			result.Cipher = tls.CipherSuiteName(state.CipherSuite)
			last = time.Now()
		},

		GotFirstResponseByte: func() {
			result.ResponseDuration = time.Since(last)
			last = time.Now()
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {return nil, err}
	defer resp.Body.Close()

	result.StatusCode = resp.StatusCode
	result.TotalDuration = time.Since(start)

	return result, nil
}


func tlsVersion(v uint16) string {
	switch v {
	case tls.VersionTLS13:
		return "v1.3"
	case tls.VersionTLS12:
		return "v1.2"
	default:
		return "UNKNOWN"
	}
}
