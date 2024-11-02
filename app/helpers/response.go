package helpers

type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		Status:  200,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(status int, message string, data interface{}) APIResponse {
	return APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
