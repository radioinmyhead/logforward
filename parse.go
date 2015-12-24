package forward

type parser interface {
	Parse([]byte, interface{}) error
}
