package global

type ErrorDTO struct {
	Message    string          `json:"message"`
	Extensions *ErrorExtension `json:"extensions"`
}

type ErrorExtension struct {
	Code  string `json:"code"`
	Field string `json:"field"`
}
