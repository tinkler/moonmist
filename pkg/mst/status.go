package mst

type Status interface {
	error
	Code() int
	Msg() string
}

type HttpStatus struct {
	code int
	msg  string
}

func (s *HttpStatus) Error() string {
	return s.msg
}

func (s *HttpStatus) Code() int {
	return s.code
}

func (s *HttpStatus) Msg() string {
	return s.msg
}
