package main

import (
	"fmt"
	"log"

	"DynamoDBLocalTest/HelloWorld/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	dynamodb2 "github.com/haozzzzzzzz/go-lambda/resource/dynamodb"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("ap-south-1"),
		Endpoint: aws.String("http://192.168.0.36:8000"),
	})
	if nil != err {
		log.Fatal(err)
		return
	}

	svc := dynamodb.New(sess)
	client := dynamodb2.DynamoDBTable{
		TableName: "video_buddy_home_dev",
		Client:    svc,
	}
	var records []model.Home
	err = client.Query(&dynamodb.QueryInput{
		KeyConditionExpression: aws.String("home_id = :home_id"),
		FilterExpression:       aws.String("contains(device_ids, :device_id)"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":home_id": {
				N: aws.String("1"),
			},
			":device_id": {
				S: aws.String("a"),
			},
		},
	}, &records)
	if nil != err {
		log.Fatal(err)
		return
	}

	for _, rec := range records {
		fmt.Println(rec)
	}
}
