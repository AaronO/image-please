package bing

import (
	"testing"
)

const METADATA_PAYLOAD = `{ns:"images",k:"5087",mid:"B42442782ED7F811AE4E05EEE1436D767D216B7F",cid:"QUmNwvX9",md5:"41498dc2f5fd553d67e209eab6fb9057",surl:"http://www.reduser.net/forum/showthread.php?68202-i-technology-PLEASE-update/page2",imgurl:"http://ontheoutsidecorner.files.wordpress.com/2011/10/please.jpg",tid:"OIP.M41498dc2f5fd553d67e209eab6fb9057H0",ow:"480",docid:"608046681396218083",oh:"424",tft:"36",dls:"images,5557",fmt:"jpeg",mw:"640",mh:"566",thH:"424",thW:"480"}`

func TestParseMetadata(t *testing.T) {
	meta, err := ParseMetadata(METADATA_PAYLOAD)
	if err != nil {
		t.Error(err)
		return
	}

	// Check values
	if meta.Width != 640 {
		t.Errorf("Width should be 640, instead: %d", meta.Width)
	}
}

func TestJStoJSON(t *testing.T) {
	in := `{a:"adasd",b:"99"}`
	expected := `{"a":"adasd","b":"99"}`
	out := jsToJSON(in)
	if out != expected {
		t.Errorf("Got `%s` but expected `%s`", out, expected)
	}
}
