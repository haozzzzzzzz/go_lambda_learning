package main

import (
	"GoLambdaExample/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.ApiGatewayHandler)
}
