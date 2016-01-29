package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/AaronO/image-please/bing"
)

type ServerStatus struct {
	Ok bool `json:"ok"`
}

func RunServer(bindTo string) error {
	return http.ListenAndServe(normalizePort(bindTo), handler())
}

func handler() http.Handler {
	r := mux.NewRouter()

	r.Path("/").
		Methods("GET").
		HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		data, _ := json.Marshal(ServerStatus{
			Ok: true,
		})

		rw.Write(data)
	})

	// Handler everything else
	r.Methods("GET").HandlerFunc(queryHandler)

	return r
}

func queryHandler(rw http.ResponseWriter, req *http.Request) {
	word := req.URL.String()[1:]
	if cleanWord, err := url.QueryUnescape(word); err == nil {
		word = cleanWord
	}

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

	// Set cache headers
	cacheSince := time.Now().Format(http.TimeFormat)
	cacheUntil := time.Now().AddDate(60, 0, 0).Format(http.TimeFormat)
	rw.Header().Set("Cache-Control", "max-age:290304000, public")
	rw.Header().Set("Last-Modified", cacheSince)
	rw.Header().Set("Expires", cacheUntil)

	// Output body
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
