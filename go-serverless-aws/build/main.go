package main

import (
	"github.com\91843\go\go-serverless-aws\pkg\handlers"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/lambda"
	"go.mongodb.org/mongo-driver/event"
)

func main() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)})
	if err != nil {
		return
	}

	dynaClient = dynamodb.New(awsSession)
	lambda.Start(handler)
}

const tableName = "LambdaInGoUser"

func handler(req event.APIGatewayProxyRequest) (*events.APIGatewayProxyRequest, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetUser
	}
	case "POST"{
		re
	}
}
