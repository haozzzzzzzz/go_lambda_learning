package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	dynamodb2 "github.com/haozzzzzzzz/go-lambda/resource/dynamodb"
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
	client := dynamodb2.DynamoDBTable{
		TableName: "video_buddy_counter_dev",
		Client:    svc,
	}
	_ = client

	newNum, err := client.IncrCounter(map[string]*dynamodb.AttributeValue{
		"home_id": {
			N: aws.String("1"),
		},
	}, "counter", 1)
	if nil != err {
		log.Fatal(err)
		return
	}

	fmt.Println(newNum)
}
