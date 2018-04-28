package main

import (
	"ExampleXRay/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

var mainHandler = handler.ApiGatewayEventHandler

func main() {
	lambda.Start(handler.GetMainHandler())
}
