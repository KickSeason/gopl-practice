package main

import (
	"io"
	"fmt"
	"net/http"
	"strings"
	"os"
)

func main () {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url;
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Get %s err, err: %v\n", url, err)
			os.Exit(1)
		}
		statusCode := resp.StatusCode;
		fmt.Printf("Response status code: %d\n", statusCode)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s err: %v\n", url, err)
			os.Exit(1)
		}
	}
}