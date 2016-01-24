package bing

import (
	"encoding/json"
	"regexp"
	"strconv"
)

func ParseMetadata(str string) (*imageMetadata, error) {
	meta := imageMetadata{}
	if err := json.Unmarshal([]byte(jsToJSON(str)), &meta); err != nil {
		return nil, err
	}
	return &meta, nil
}

// Converts a JS like structure to JSON (adds quotes around keys)
var _jsToJSON = regexp.MustCompile(`[^"](\w+):\"`)

func jsToJSON(str string) string {
	return _jsToJSON.ReplaceAllStringFunc(str, func(match string) string {
		head := match[:1]
		rest := match[1 : len(match)-2]
		return head + `"` + rest + `":"`
	})
}

// A useful type to parse json strings as ints
type stringInt int

func (s *stringInt) UnmarshalJSON(data []byte) error {
	str := string(data[1 : len(data)-1])
	if x, err := strconv.Atoi(str); err == nil {
		*s = stringInt(x)
	} else {
		*s = stringInt(0)
	}

	return nil
}

type imageMetadata struct {
	MD5       string    `json:"md5"`
	Namespace string    `json:"ns"`
	Format    string    `json:"fmt"`
	ImageUrl  string    `json:"imgurl"`
	SourceUrl string    `json:"surl"`
	Tid       string    `json:"tid"`
	Height    stringInt `json:"mh"`
	Width     stringInt `json:"mw"`
}
