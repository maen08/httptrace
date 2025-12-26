package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/maen08/httptrace/trace"
)


func Execute() {
	method := flag.String("method", "GET", "HTTP method (GET, POST, PUT, etc)")
	jsonBody := flag.String("json", "", "JSON body to send with the request")

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("default with GET-method usage: httptrace <url>")
		fmt.Println("with other http method usage: httptrace --method POST --json '{...}' <url>")
		os.Exit(1)
	}

	url := flag.Arg(0)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := trace.Run(ctx, trace.Options{
		URL:    url,
		Method: *method,
		JSON:   *jsonBody,
	})
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	printResult(result)
}


func printResult(r *trace.Result) {
	fmt.Printf("DNS lookup      : %v   [ IP:  %s ]\n",
		r.DNSDuration, r.DNSIP)

	fmt.Printf("TCP connect     : %v   [ Port:  %s ]\n",
		r.TCPDuration, r.Port)

	if r.TLSDuration > 0 {
		fmt.Printf("TLS handshake   : %v  [ TLS: %s, Cipher: %s ]\n",
			r.TLSDuration, r.TLSVersion, r.Cipher)
	}

	fmt.Printf("Server response : %v  [ StatusCode: %d ]\n",
		r.ResponseDuration, r.StatusCode)

	fmt.Printf("Total           : %v\n", r.TotalDuration)
}
