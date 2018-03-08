package main

import (
	"auth/src/data/mongo"
	"auth/src/pipeline"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"auth/src/pojo"
	"auth/src/data/google"
	"auth/src/controller"
	"auth/src/format"
	"encoding/json"
	"auth/src/view"
	"auth/src/apierror"
)

var session, err = mongo.GetWithCredentials(
	[]string{""},
	"",
	"",
	"",
	true)
var db = mongo.NewMongo(session)
var googleProv = google.NewProvider("", "")
var provPip = pipeline.NewProvider(googleProv)
var ctrl = controller.NewProviderPersistence(provPip, db)
var formatter = format.NewJsonFormatter(&apierror.SimpleHandler{})
var aws = view.NewLambda(formatter, &apierror.SimpleHandler{})

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Could not open a session with database
	if err != nil {
		return aws.Transform(nil, err)
	}
	var auth pojo.AuthCode
	json.Unmarshal([]byte(request.Body), &auth)
	return aws.Transform(ctrl.SocialSignUp(auth.Code))
}

func main() {
	lambda.Start(Handler)
}
