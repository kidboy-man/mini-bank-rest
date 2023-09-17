package schemas

type CustomError struct {
	Code       int    `json:"code"`
	HTTPStatus int    `json:"http_status"`
	Message    string `json:"message"`
}

func (ce *CustomError) Error() string {
	return ce.Message
}
