package forward

type sender interface {
	Send([]byte) error
}
