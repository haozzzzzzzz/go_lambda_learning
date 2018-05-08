package main

import (
	"ExampleApiAuthorizer1And2/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.GetMainHandler())
}
