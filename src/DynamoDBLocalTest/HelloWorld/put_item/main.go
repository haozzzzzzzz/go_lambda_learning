package main

import (
	"fmt"
	"log"

	"DynamoDBLocalTest/HelloWorld/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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

	home := &model.Home{
		HomeId:    3,
		Title:     "首页3",
		DeviceIds: []string{"8", "a"},
	}

	av, err := dynamodbattribute.MarshalMap(home)
	if nil != err {
		log.Fatal(err)
		return
	}

	svc := dynamodb.New(sess)
	result, err := svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("video_buddy_home_dev"),
		Item:      av,
	})
	if nil != err {
		log.Fatal(err)
		return
	}
	fmt.Println(result)
}
