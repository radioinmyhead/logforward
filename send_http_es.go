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

	tmp := `{"create":{"_index":"log","_type":"service","_id":"` +
		s.Timestamp + `,` + uid + `"}}` + `\n` + string(s.Data)
	s.Data = []byte(tmp)
}
