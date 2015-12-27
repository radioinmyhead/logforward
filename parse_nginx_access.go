package forward

import (
	"fmt"
	"reflect"
	"regexp"
	"time"
)

var (
	ngxRe     = regexp.MustCompile(`^([^ ]+) - ([^ ]+) \[([^\]]+)\] "([^ ]+) ([^ ]+) ([^ ]+)" ([^ ]+) ([^ ]+) ([^ ]+) ([^ ]+) "([^"]+)" "([^"]+)"$`)
	ngxName   = []string{"raw", "ip", "remote_user", "timestamp", "method", "path", "proto", "status", "body_bytes_sent", "request-time", "request_length", "http-referer", "user-agent"}
	ngxLayout = "02/Jan/2006:15:04:05 -0700"
)

type LogNgxAccess struct {
}

func ngxGetTs(data map[string]string, field string) (err error) {
	if v, ok := data[field]; ok {
		t, err := time.Parse(ngxLayout, v)
		if err != nil {
			return err
		}
		data[field] = fmt.Sprintf("%v000", t.Unix())
	}
	return nil
}

func ngxAssignment(ret map[string]string, values, names []string) {
	if ret == nil {
		ret = make(map[string]string)
	}
	size := len(values)
	if len(names) < size {
		size = len(names)
	}
	for i := 0; i < size; i++ {
		ret[names[i]] = values[i]
	}
}

func (l *LogNgxAccess) Parse(data []byte, v interface{}) (err error) {
	// check
	t := reflect.TypeOf(v)
	tmp := make(map[string]string)
	if !t.ConvertibleTo(reflect.TypeOf(&tmp)) {
		return fmt.Errorf("type error")
	}
	// parse
	values, names := []string{}, []string{}
	matchs := ngxRe.FindStringSubmatch(string(data))
	if matchs == nil {
		return fmt.Errorf("parse nginx log error: %v", err)
	}
	values = append(values, matchs...)
	names = append(names, ngxName...)

	ngxAssignment(tmp, values, names)
	ngxGetTs(tmp, "timestamp")

	// set
	reflect.ValueOf(v).Elem().Set(reflect.ValueOf(tmp))

	return nil
}
