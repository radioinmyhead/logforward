package forward

import (
	"fmt"
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
