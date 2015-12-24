package forward

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type SendHttp struct {
	Method string
	Url    string
	Header map[string]string
	Data   []byte
}

func (s *SendHttp) httpDo() (ret []byte, err error) {

	client := &http.Client{
		Timeout: time.Second,
	}

	req, err := http.NewRequest(s.Method, s.Url, bytes.NewBuffer(data))
	if err != nil {
		return
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return body, nil
}

func (s *SendHttp) Send(data []byte) (err error) {
	s.Data = data
	_, err = s.httpDo()
	return
}
