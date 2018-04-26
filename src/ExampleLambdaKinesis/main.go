package main

import (
	"ExampleLambdaKinesis/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

var mainHandler = handler.BasicExecutionEventHandler

func main() {
	lambda.Start(handler.GetMainHandler())
}
