package main

import (
	"auth/src/data/mongo"
	"auth/src/operation"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"auth/src/pojo"
	"auth/src/data/google"
	"auth/src/controller"
	"auth/src/format"
	"encoding/json"
)

var session, err = mongo.GetWithCredentials(
	[]string{""},
	"",
	"",
	"",
	true)
var db = mongo.NewMongo(session)
var persistent = operation.NewPersistent(db)
var prov = google.NewProvider("", "")
var provider = operation.NewProvider(prov)
var fmt = format.NewJsonFormatter(`{"message": Could not parse JSON}`)
var ctrl = controller.NewProviderPersistence(persistent, provider, fmt)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var auth pojo.AuthCode
	json.Unmarshal([]byte(request.Body), &auth)
	return events.APIGatewayProxyResponse{
		Body: ctrl.SocialSignUp(auth.Code),
		StatusCode: 200,
		}, nil
}

func main() {
	lambda.Start(Handler)

}
