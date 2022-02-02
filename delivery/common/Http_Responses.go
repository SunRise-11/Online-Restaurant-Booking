package common

type DefaultResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewSuccessOperationResponse() DefaultResponse {
	return DefaultResponse{
		200,
		"Successful Operation",
	}
}

func NewInternalServerErrorResponse() DefaultResponse {
	return DefaultResponse{
		500,
		"Internal Server Error",
	}
}

func NewNotFoundResponse() DefaultResponse {
	return DefaultResponse{
		404,
		"Not Found",
	}
}

func NewBadRequestResponse() DefaultResponse {
	return DefaultResponse{
		400,
		"Bad Request",
	}
}

func NewConflictResponse() DefaultResponse {
	return DefaultResponse{
		409,
		"Data Has Been Modified",
	}
}

func NewStatusNotAcceptable() DefaultResponse {
	return DefaultResponse{
		406,
		"Not Accepted",
	}
}
