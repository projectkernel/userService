package format

import (
	"encoding/json"
	"auth/src/apierror"
)

type JsonFormatter struct {
	errHandler apierror.Handler
}

func NewJsonFormatter(errHandler apierror.Handler) *JsonFormatter{
	return &JsonFormatter{
		errHandler:errHandler,
	}
}

func (format JsonFormatter) Format(val interface{}) (string, error) {
	result, err := json.Marshal(val)
	if err != nil {
		return "", err
	}
	return string(result), nil
}