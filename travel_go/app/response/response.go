package response

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
	Code    int         `json:"code"`
}

func RequestBadResponse(message string) Response {
	return Response{
		Success: false,
		Message: message,
		Code:    400,
	}
}

func SeverErrorResponse(message string) Response {
	return Response{
		Success: false,
		Message: message,
		Code:    500,
	}
}

func ForbiddenResponse(message string) Response {
	return Response{
		Success: false,
		Message: message,
		Code:    403,
	}
}
