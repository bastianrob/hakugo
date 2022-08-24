package exception

import "fmt"

const (
	CodeInputMalformed  = "INPUT_MALFORMED"
	CodeNotFound        = "NOT_FOUND"
	CodeValidationError = "VALIDATION_ERROR"
	CodeUnexpectedError = "UNEXPECTED_ERROR"
)

type Exception struct {
	Message string
	Context error
	Code    string
}

func (e *Exception) Error() string {
	return fmt.Sprintf("%s: %v", e.Context, e.Context)
}

func New(err error, message, code string) error {
	return &Exception{
		Context: err,
		Code:    code,
		Message: message,
	}
}

func IsException(err error) (*Exception, bool) {
	if err == nil {
		return nil, false
	}

	exc, ok := err.(*Exception)
	return exc, ok
}
