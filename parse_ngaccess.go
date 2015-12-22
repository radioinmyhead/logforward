package logforward

import (
	"fmt"
	"regexp"
)

type StrNgAccess struct {
	StrBase
}

var (
	ngRe   = regexp.MustCompile(`^([^ ]+) - ([^ ]+) \[([^\]]+)\] "([^ ]+) ([^ ]+) ([^ ]+)" ([^ ]+) ([^ ]+) ([^ ]+) ([^ ]+)"([^"]+)" "([^"]+)"$`)
	ngName = []string{"raw", "remote_addr", "remote_user", "time_local", "method", "path", "http", "status", "body_bytes_sent", "request_time", "request_length", "http_referer", "http_user_agent"}
)

func assignment(ret map[string]string, vs, name []string) {
	l := len(vs)
	if len(name) < l {
		l = len(name)
	}
	for i := 0; i < l; i++ {
		ret[name[i]] = vs[i]
	}
}

func (s *StrNgAccess) Decode() (ret map[string]string, err error) {
	matches := ngRe.FindStringSubmatch(string(s.D))
	if matches == nil {
		return ret, fmt.Errorf("parse nginx log error: %v", err)
	}
	ret = make(map[string]string)
	assignment(ret, matches, ngName)

	return
}
