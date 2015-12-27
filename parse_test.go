package forward

import (
	"testing"
)

func TestLogNgxAccess(t *testing.T) {
	p := &LogNgxAccess{}
	var tmp map[string]string
	err := p.Parse([]byte(`182.92.153.124 - - [21/Dec/2015:15:39:37 +0800] "POST /faceid/v1/verify HTTP/1.0" 200 173 0.496 27005 "-" "-"`), &tmp)
	if err != nil {
		t.Error(err)
	}
}

func TestLogNgxWb(t *testing.T) {
	p := &LogNgxWb{}
	var tmp map[string]string
	err := p.Parse([]byte(`{
		"host": "test.com",
		"code": {
			"a": 1,
			"b": 2,
			"c": "ok",
			"d": true
		}
	}`), &tmp)
	if err != nil {
		t.Error("wb parse fail:", err)
	}
}
