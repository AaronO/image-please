package bing

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/moovweb/gokogiri"
)

// ImageResult holds a preview image URL and the source URL
type ImageResult struct {
	URL     string
	Preview string
	Source  string

	Format string
	Width  int
	Height int
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
	// Required to get correct HTML
	req.Header.Set("User-Agent", "Mozilla/5.0 Chrome")

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

	// Get image tags
	imagesTags, err := doc.Search(`//div[@class='dg_u']/a`)
	if err != nil {
		return nil, err
	}

	// Filter and parse
	var images []ImageResult
	for _, tag := range imagesTags {
		if meta, err := ParseMetadata(tag.Attr("m")); err == nil {
			images = append(images, metaToResult(meta))
		}
	}

	// No results
	if images == nil {
		return nil, fmt.Errorf("No results")
	}

	return images, nil
}

func metaToResult(meta *imageMetadata) ImageResult {
	return ImageResult{
		URL:    meta.ImageUrl,
		Width:  int(meta.Width),
		Height: int(meta.Height),
		Format: meta.Format,
	}
}
