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
		Region: aws.String("ap-south-1"),

		Endpoint: aws.String("http://0.0.0.0:8000"),
	})
	if nil != err {
		log.Fatal(err)
		return
	}

	svc := dynamodb.New(sess)
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("video_buddy_home_dev"),
		Key: map[string]*dynamodb.AttributeValue{
			"home_id": {
				N: aws.String("1"),
			},
		},
	})
	if nil != err {
		log.Fatal(err)
		return
	}

	home := &model.Home{}
	err = dynamodbattribute.UnmarshalMap(result.Item, home)
	if nil != err {
		log.Fatal(err)
		return
	}
	fmt.Println(home)
}
