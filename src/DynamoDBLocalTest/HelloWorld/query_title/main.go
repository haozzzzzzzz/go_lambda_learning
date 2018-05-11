package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("ap-south-1"),
		Endpoint: aws.String("http://0.0.0.0:8000"),
	})
	if nil != err {
		log.Fatal(err)
		return
	}

	svc := dynamodb.New(sess)
	output, err := svc.Query(&dynamodb.QueryInput{
		TableName: aws.String("video_buddy_home_dev"),
	})
	if nil != err {
		log.Fatal(err)
		return
	}

	fmt.Println(output)
}
