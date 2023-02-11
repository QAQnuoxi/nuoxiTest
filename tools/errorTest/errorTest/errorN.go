package errorTest

func New(text string) error1 {
	return &errorString1{text}
}

// errorString is a trivial implementation of error.
type errorString1 struct {
	s string
}

func (e *errorString1) Error1() string {
	return e.s
}
