package main

import (
	"ExampleLambdaApi/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

var mainHandler = handler.ApiGatewayEventHandler

func main() {
	lambda.Start(mainHandler)
}
