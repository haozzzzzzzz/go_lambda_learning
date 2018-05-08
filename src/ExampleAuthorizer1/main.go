package main

import (
	"ExampleAuthorizer1/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.GetMainHandler())
}
