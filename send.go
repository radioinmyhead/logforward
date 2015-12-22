package logforward

type SendBase struct {
	D []byte
}

func (s *SendBase) Send() {
	return
}

func (s *SendBase) Set(d interface{}) (err error) {
	j, err := json.Marshal(d)
	if err != nil {
		return
	}
	s.D = d
}
