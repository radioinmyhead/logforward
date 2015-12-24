package forward

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type LogNgxWb struct {
}

func (l *LogNgxWb) Parse(data []byte, v interface{}) (err error) {

	// check
	t := reflect.TypeOf(v)
	tmp := make(map[string]string)
	if !t.ConvertibleTo(reflect.TypeOf(&tmp)) {
		return fmt.Errorf("type error")
	}

	// parse
	var js struct {
		Host string
		Code map[string]interface{}
	}
	if err := json.Unmarshal(data, &js); err != nil {
		return err
	}
	tmp["host"] = js.Host
	for k, v := range js.Code {
		tmp[k] = fmt.Sprintf("%v", v)
	}

	// set
	reflect.ValueOf(v).Elem().Set(reflect.ValueOf(tmp))

	return nil
}
