package logforward

import (
	"fmt"
	"strings"
)

type StrNormal struct {
	StrBase
}

func (s *StrNormal) Decode(prefix, hostname string) (ret map[string]string, err error) {
	ret = make(map[string]string)
	for _, kv := range strings.Split(string(s.D), `&`) {
		tmp := strings.Split(kv, `=`)
		if len(tmp) != 2 {
			return ret, fmt.Errorf("decode error")
		}
		ret[tmp[0]] = tmp[1]
	}
	return
}
