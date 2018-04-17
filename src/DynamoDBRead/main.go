package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Person struct {
	Id      string `json:"Id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if nil != err {
		log.Fatal(err)
		return
	}

	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Lambda"),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String("1"),
			},
			"name": {
				S: aws.String("hao"),
			},
		},
	})
	if nil != err {
		log.Fatal(err)
		return
	}

	newPerson := Person{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &newPerson)
	if nil != err {
		log.Fatal(err)
		return
	}

	fmt.Println(newPerson)
}
