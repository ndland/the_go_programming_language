package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	var newUrl string

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(os.Args[0], "http://") {
			newUrl = "http://" + url
		} else {
			newUrl = url
		}
		resp, err := http.Get(newUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Response Status: %s\n", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
