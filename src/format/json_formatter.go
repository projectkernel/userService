package format

import (
	"encoding/json"
)

type JsonFormatter struct {
	fallback string
}

func NewJsonFormatter(fallback string) *JsonFormatter{
	return &JsonFormatter{
		fallback:fallback,
	}
}

// TODO Provide Fallback ensurance
func (format JsonFormatter) Format(val interface{}) (string) {
	result, err := json.Marshal(val)
	if err != nil {
		return format.fallback
	}
	return string(result)
}