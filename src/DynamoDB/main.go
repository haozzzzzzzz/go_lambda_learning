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
	person := &Person{
		Id:      "1",
		Name:    "hao",
		Address: "ShenZhen",
	}

	av, err := dynamodbattribute.MarshalMap(person)
	if nil != err {
		log.Fatal(err)
		return
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Lambda"),
	}

	_, err = svc.PutItem(input)
	if nil != err {
		log.Fatal(err)
		return
	}

	fmt.Println("Successful")
}
