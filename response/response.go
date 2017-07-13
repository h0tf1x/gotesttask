package response

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type DataResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func NewErrorResponse(message string) Response {
	return Response{Success: false, Message: message}
}

func NewSuccessResponse(message string) Response {
	return Response{Success: true, Message: message}
}

func NewDataResponse(data interface{}) DataResponse {
	return DataResponse{Success: true, Data: data}
}
