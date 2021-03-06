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
		Region: aws.String("ap-south-1"),

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
	home := &model.Home{}
	err = client.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"home_id": {
				N: aws.String("1"),
			},
		},
	}, home)
	if nil != err {
		log.Fatal(err)
		return
	}
	fmt.Println(home)
}
