package exception

import "fmt"

type Exception struct {
	Message string
	Context error
}

func (e *Exception) Error() string {
	return fmt.Sprintf("%s: %v", e.Context, e.Context)
}

func New(err error, message string) error {
	return &Exception{
		Message: message,
		Context: err,
	}
}
