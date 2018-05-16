package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/haozzzzzzzz/go-lambda/resource/dyndb/table"
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
	counter := table.CounterTable{
		DynamoDBTable: table.DynamoDBTable{
			TableName: "video_buddy_counter_dev",
			Client:    svc,
			Ctx:       aws.BackgroundContext(),
		},
	}

	newNum, err := counter.Incr("hello", "1", 1)
	if nil != err {
		log.Fatal(err)
		return
	}

	fmt.Println(newNum)
}
