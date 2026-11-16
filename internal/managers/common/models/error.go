package models

// ErrorResponse — единый формат ответа API при ошибке
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
