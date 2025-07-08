package api

type ErrorResponse struct {
	Error Error `json:error`
}

type Error struct {
	Code    string `json:code`
	Message string `json:message`
}

func NewErrorResponse(code string, message string) ErrorResponse {
	return ErrorResponse{
		Error: Error{
			Code:    code,
			Message: message,
		},
	}
}

func NewErrorResponse404() ErrorResponse {
	return NewErrorResponse("not_found", "Resource not found")
}
