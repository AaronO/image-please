package bing

import (
	"encoding/json"
	"strconv"
)

func ParseMetadata(str string) (*imageMetadata, error) {
	meta := imageMetadata{}
	if err := json.Unmarshal([]byte(str), &meta); err != nil {
		return nil, err
	}
	return &meta, nil
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
