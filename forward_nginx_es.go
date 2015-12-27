package forward

import (
	"encoding/json"
)

type ForNgxEs struct {
	LogNgxAccess
}

func (f *ForNgxEs) Forward(data []byte, dic map[string]string) (err error) {
	var ret map[string]string
	if err = f.Parse(data, &ret); err != nil {
		return
	}
	for k, v := range dic {
		ret[k] = v
	}
	sendhttp := GetSender("es").(*SendHttp)
	sendes := &SendHttpEs{*sendhttp, ret["timestamp"]}
	d, _ := json.Marshal(ret)
	sendes.Send(d)
	return
}
