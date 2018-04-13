package main

import "github.com/aws/aws-lambda-go/lambda"

func hello() (string, error) {
	return "Hello, lambda", nil
}

func main() {
	// make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
