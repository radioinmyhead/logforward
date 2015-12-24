package forward

import (
	"strings"

	"github.com/satori/go.uuid"
)

type SendHttpEs struct {
	SendHttp
	Timestamp string
}

func (s *SendHttpEs) SetTimestamp(t string) {

	s.Timestamp = t
	s.Timestamp = strings.TrimSuffix(s.Timestamp, "000")

	uid := uuid.NewV4().String()

	s.Data = `{"create":{"_index":"log","_type":"service","_id":"` +
		s.Timestamp + `,` + uid + `"}}` + `\n` + s.Data
}
