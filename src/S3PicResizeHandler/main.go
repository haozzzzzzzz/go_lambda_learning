package main

import (
	"context"

	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, event events.S3Event) (err error) {
	for _, eventRecord := range event.Records {
		log.Print(eventRecord)
	}
	return
}

func main() {
	lambda.Start(Handler)
}
