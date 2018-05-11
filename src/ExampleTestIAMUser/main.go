package main

import (
	"ExampleTestIAMUser/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.GetMainHandler())
}
