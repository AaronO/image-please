package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/AaronO/image-please/bing"
)

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

func main() {
	fmt.Println("Listening..")
	http.ListenAndServe(":9999", http.HandlerFunc(handler))
}
