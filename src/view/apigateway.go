package view

import (
	"auth/src/format"
	"github.com/aws/aws-lambda-go/events"
	"auth/src/apierror"
)

type Lambda struct {
	formatter  format.Formatter
	errHandler apierror.Handler
}

func NewLambda(formatter format.Formatter, errHandler apierror.Handler) *Lambda {
	return &Lambda{
		formatter:  formatter,
		errHandler: errHandler,
	}
}

func (aws Lambda) Transform(value interface{}, err error) (events.APIGatewayProxyResponse, error) {
	var result interface{}
	if err != nil {
		result = aws.errHandler.Handle(err)
	}
	str, err := aws.formatter.Format(result)
	if err != nil {
		// Formatter Error
		str, err = aws.formatter.Format(aws.errHandler.Handle(err))
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
			}, nil
		}
	}
	return events.APIGatewayProxyResponse {
		Body: str,
	}, nil

}
