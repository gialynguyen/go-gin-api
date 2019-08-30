package _interface

type SuccessBaseResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ErrorBaseResponse struct {
	Success bool        `json:"success"`
	Error   ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	ErrorCode int    `json:"error_code"`
	ErrorMes  interface{} `json:"error_message"`
}
