package forward

import (
	"fmt"
	"sync"
)

type forwarder interface {
	Forward([]byte, map[string]string) error
}

var table map[string]forwarder

func Add(prefix string, fd forwarder) {
	if table == nil {
		table = make(map[string]forwarder)
	}
	table[prefix] = fd
}

func Forward(prefix string, postdata []byte, dic map[string]string) (err error) {
	if v, ok := table[prefix]; ok {
		err = v.Forward(postdata, dic)
		return err
	}
	return fmt.Errorf("prefix is not registered")
}

var sendPool map[string]sender

func Set(name, style string, par ...string) {
	if sendPool == nil {
		sendPool = make(map[string]sender)
	}
	if style == "http" {
		if len(par) != 2 {
			panic("logforward set fail")
		}
		s := NewSendHttp(par...)
		sendPool[name] = s
	}
}
func NewSendHttp(par ...string) (ret SendHttp) {
	return SendHttp{
		Method: par[0],
		Url:    par[1],
	}
}
func GetSender(prefix string) (ret sender) {
	if v, ok := sendPool[prefix]; ok {
		return v
	}
	return
}

type SendNode struct {
	s sender
	d []byte
}

func NewSendNote(s sender, d []byte) SendNode {
	return SendNode{s, d}
}

type SendList []sendNode

func (l *SendList) AddSendList(list ...SendNode) {
	l = append(l, list...)
}
func (sl *SendList) Psend() {
	var wg sync.WaitGroup
	wg.Add(len(sl))
	for _, sn := range sl {
		go func(sn SendNode) {
			defer wg.done()
			s, d := sn.s, sn.d
			s(d)
		}(sn)
	}
	wg.Wait()
}
