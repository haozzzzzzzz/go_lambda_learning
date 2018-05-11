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

	svc := dynamodb.New(sess)
	output, err := svc.Query(&dynamodb.QueryInput{
		TableName:              aws.String("video_buddy_home_dev"),
		KeyConditionExpression: aws.String("home_id = :home_id"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":home_id": {
				N: aws.String("1"),
			},
		},
	})
	if nil != err {
		log.Fatal(err)
		return
	}

	var resultRecords []*model.Home
	err = dynamodbattribute.UnmarshalListOfMaps(output.Items, &resultRecords)
	if nil != err {
		log.Fatal(err)
		return
	}

	for _, rec := range resultRecords {
		fmt.Println(rec)
	}
}
