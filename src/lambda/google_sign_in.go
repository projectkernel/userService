package lambda

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"auth/src/signin/google"
	"encoding/json"
	"auth/src/pojo"
	"errors"
	"auth/src/signin"
)

type Result struct {
	User *pojo.User `json:"user"`
	Token string `json:"token"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var auth pojo.AuthCode
	controller := signin.NewModel(google.New("", ""))
	json.Unmarshal([]byte(request.Body), &auth)
	user, token, err := controller.SignIn(auth.Code)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 404}, err
	}
	result, err := json.Marshal(Result { User:user, Token:token })
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, errors.New("could not create a valid response")
	}
	return events.APIGatewayProxyResponse{
		Body: string(result),
		StatusCode: 200,
		}, nil
}

func main() {
	lambda.Start(Handler)
}
