package main

import (
	"LambdaApiExample/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

var mainHandler = handler.BasicExecutionEventHandler

func main() {
	lambda.Start(mainHandler)
}
