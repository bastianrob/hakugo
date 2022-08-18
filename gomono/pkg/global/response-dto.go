package global

type ResponseDTO[T any] struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
