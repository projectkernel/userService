package pojo

type ApiError struct {
    Message string `json:"message"`
	Code int `json:"code"`
}

func (err ApiError) Error() string {
	return err.Message
}