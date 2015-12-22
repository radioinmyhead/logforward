package logforward

import (
	"encoding/json"
)

type StrJson struct {
	StrBase
}

func (s *StrJson) Decode() (ret map[string]string, err error) {
	err = json.Unmarshal(s.D, &ret)
	return
}
