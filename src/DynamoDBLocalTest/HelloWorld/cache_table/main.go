package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	dynamodb2 "github.com/haozzzzzzzz/go-lambda/resource/dynamodb"
	"github.com/haozzzzzzzz/go-lambda/resource/dynamodb/table"
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
	client := table.CacheTable{
		DynamoDBTable: dynamodb2.DynamoDBTable{
			TableName: "video_buddy_cache_dev",
			Client:    svc,
			Ctx:       aws.BackgroundContext(),
		},
	}
	_ = client

	success, err := client.SetNxTTL("hello", "123", time.Minute)
	if nil != err {
		log.Fatal(err)
		return
	}

	fmt.Println(success)
}
