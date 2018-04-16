package main

import (
	"context"

	"fmt"

	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, event events.S3Event) (msg string, err error) {
	for _, eventRecord := range event.Records {
		log.Print(eventRecord)
	}

	msg = fmt.Sprintf("%+v", event)
	return
}

func main() {
	lambda.Start(Handler)
}
