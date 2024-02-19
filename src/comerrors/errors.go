package comerrors

const (
	NoErr      = 0
	ParamErr   = -1
	NetworkErr = -2
)

type Errors struct {
	errno int
	err   string
}

func NewError(errno int, err string) *Errors {
	return &Errors{errno, err}
}

func (e *Errors) Errno() int {
	return e.errno
}

func (e *Errors) Error() string {
	return e.err
}
