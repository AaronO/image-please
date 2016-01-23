package bingimage

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/moovweb/gokogiri"
)

// ImageResult holds a preview image URL and the source URL
type ImageResult struct {
	PreviewImage string
	Source       string
}

const base = "http://www.bing.com/images/search?q="

// Search accepts a search string and returns
func Search(search string) (results []ImageResult, err error) {
	page, err := fetchPage(search)
	if err != nil {
		return nil, err
	}
	return parseResult(page)
}

func fetchPage(search string) ([]byte, error) {
	url := base + url.QueryEscape(search)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response code %d", resp.StatusCode)
	}

	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return page, nil
}

func parseResult(html []byte) (results []ImageResult, err error) {
	doc, err := gokogiri.ParseHtml([]byte(html))
	if err != nil {
		return nil, err
	}

	root := doc.Root()
	previews, err := root.Search("//a/div/img")
	if err != nil {
		return nil, err
	}

	var images []ImageResult
	for _, v := range previews {
		previewURL := v.Attr("src")
		images = append(images, ImageResult{PreviewImage: previewURL})
	}

	return images, nil
}
