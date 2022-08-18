package global

type RequestDTO[T any] struct {
	Data T `json:"data"`
}
