package structs

type ErrorMessage struct {
	StatusCode  int    `json:"status_code"`
	Message     string `json:"message"`
}