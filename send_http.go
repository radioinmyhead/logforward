package logforward

import (
	"http"
)

type SendHttpBase struct {
	SendBase
	Url    string
	Mothod string
	Header [][]string
}

type HttpSender interface {
	Set(interface{}) error
	Send() error
	HttpSetUrl(string)
	HttpSetMothod(string)
	HttpSetHeader(string, string)
	HttpSetData()
}

var (
	httpSender = map[string]HttpSender{}
)

func (s *SendHttpBase) HttpSetUrl(u string) {
	s.Url = u
}
func (s *SendHttpBase) HttpSetMothod(m string) {
	s.Mothod(m)
}
func (s *SendHttpBase) HttpSetHeader(k, v string) {
	if s.Header == nil {
		s.Header = make([][]string)
	}
	s.Header = append(s, []string{k, v})
}
func (s *SendHttpBase) HttpSetData() {
	return
}

func (h *SendHttpBase) Send() (err error) {
	var s *HttpSender
	s = h
	client = &http.Client{}
	req, err := http.NewRequest(s.Mothod, url, bytes.NewBuffer(s.D))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	return
}
