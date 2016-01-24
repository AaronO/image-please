package bing

import (
	"fmt"
	"testing"
)

func TestFetchHTML(t *testing.T) {
	results, err := Search("birthday cake")
	if err != nil {
		t.Errorf("error searching: %s", err)
		return
	}
	fmt.Printf("got back %s", results[0].URL)
}
