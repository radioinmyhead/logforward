package logforward

type StrBase struct {
	D []byte
}

func (s *StrBase) Decode() (ret map[string]string, err error) {
	return
}
func (s *StrBase) Set(d []byte) {
	s.D = d
}
