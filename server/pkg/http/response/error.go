package response

type ErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Message string `json:"message"`
}

func NewErrorRes(err error) ErrorResponse {
	return ErrorResponse{
		Error: Error{
			Message: err.Error(),
		},
	}
}
