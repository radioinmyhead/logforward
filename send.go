package forward

import (
	"sync"
)

type sender interface {
	Send([]byte) error
	//GetUrl()
}

type SendNode struct {
	s sender
	d []byte
}

func NewSendNote(s sender, d []byte) SendNode {
	return SendNode{s, d}
}

type SendList []SendNode

func (l *SendList) AddSendList(list ...SendNode) {
	*l = append(*l, list...)
}
func (sl *SendList) Psend() {
	var wg sync.WaitGroup
	wg.Add(len(*sl))
	for _, sn := range *sl {
		go func(sn SendNode) {
			defer wg.Done()
			s, d := sn.s, sn.d
			s.Send(d)
		}(sn)
	}
	wg.Wait()
}

//var sendPool map[string]sender
type csend struct {
	t string
	p []string
}

var sendPool map[string]csend

//func Set(name, style string, par ...string) {
//	if sendPool == nil {
//		sendPool = make(map[string]sender)
//	}
//	if style == "http" {
//		if len(par) != 2 {
//			panic("logforward set fail")
//		}
//		s := NewSendHttp(par...)
//		sendPool[name] = s
//	}
//}
func Set(name, Type string, par ...string) {
	if sendPool == nil {
		sendPool = make(map[string]csend)
	}
	sendPool[name] = csend{Type, par}
}

//func GetSender(prefix string) (ret sender) {
//	if v, ok := sendPool[prefix]; ok {
//		return v
//	}
//	return
//}
func GetSender(prefix string) (ret interface{}) {
	if v, ok := sendPool[prefix]; ok {
		switch v.t {
		case "http":
			return NewSendHttp(v.p...)
		}
	}
	return nil
}

// new

func NewSendHttp(par ...string) (ret *SendHttp) {
	return &SendHttp{
		Method: par[0],
		Url:    par[1],
	}
}
