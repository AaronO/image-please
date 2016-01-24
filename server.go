package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/AaronO/image-please/bing"
)

func RunServer(bindTo string) error {
	return http.ListenAndServe(normalizePort(bindTo), http.HandlerFunc(handler))
}

func handler(rw http.ResponseWriter, req *http.Request) {
	word := req.URL.String()[1:]

	// Search
	results, err := bing.Search(word)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}

	// Get first element
	url := results[0].URL

	// Stream
	resp, err := http.Get(url)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}
	io.Copy(rw, resp.Body)
}

// Normalize port string to an "addr"
// as expected by ListenAndServe
func normalizePort(port string) string {
	if strings.Contains(port, ":") {
		return port
	}
	return ":" + port
}
