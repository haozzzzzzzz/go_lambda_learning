package main

import (
	"LambdaFrameworkExample/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.ApiGatewayHandler)
}
