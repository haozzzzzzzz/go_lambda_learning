package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context) (msg string, err error) {
	msg = "Hello, world!"
	return
}

func main() {
	lambda.Start(Handler)
}
