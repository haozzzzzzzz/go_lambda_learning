package main

import (
	"ExampleSAMLocal/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.GetMainHandler())
}
