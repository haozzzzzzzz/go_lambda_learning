package main

import (
	"ExampleAuthorizer2/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.GetMainHandler())
}
