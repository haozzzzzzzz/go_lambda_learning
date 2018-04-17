package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/sirupsen/logrus"
)

/**
此程序需要在DynamoDB中建立一张Lambda表，包含字段Id string, name string
*/

// 请求
type Ask struct {
	Operation string `json:"operation"` // 操作
	TableName string `json:"tableName"` // 表名
	Payload   struct {
		Id   string `json:"Id"`
		Name string `json:"name"`
	} `json:"payload"` // 数据负载
}

// 返回
type Ack struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Ask     *Ask
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	var (
		ask = new(Ask)
		ack = new(Ack)
	)

	defer func() {
		ack.Ask = ask

		bytesAck, errMarshal := json.Marshal(ack)
		if errMarshal != nil {
			err = errMarshal
			logrus.Warnf("marshal ack failed. %s", err)
		}

		response.Body = string(bytesAck)
		if err != nil {
			response.StatusCode = 500
		} else {
			response.StatusCode = 200
		}

	}()

	ack.Code = 0
	ack.Message = "请求成功"

	err = json.Unmarshal([]byte(request.Body), ask)
	if nil != err {
		logrus.Warnf("unmarshal request body failed. %s", err)
		return
	}

	switch ask.Operation {
	case "create":
		err = DynamoDBPutItem(&Person{
			Id:   ask.Payload.Id,
			Name: ask.Payload.Name,
		})
		if nil != err {
			logrus.Warnf("put dynamo db item failed. %s", err)
			return
		}

		ack.Data = "上传成功"

	case "read":
		person, errGet := DynamoDBGetItem(ask.Payload.Id, ask.Payload.Name)
		if nil != errGet {
			err = errGet
			logrus.Warnf("get dynamo db item failed. %s", err)
			return
		}
		ack.Data = person

	default:

	}

	return
}

type Person struct {
	Id   string `json:"Id"`
	Name string `json:"name"`
}

func GetDynamoDB() (db *dynamodb.DynamoDB, err error) {
	ses, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		logrus.Warnf("new aws session failed. %", err)
		return
	}

	db = dynamodb.New(ses)
	return
}

func DynamoDBPutItem(person *Person) (err error) {
	attribute, err := dynamodbattribute.MarshalMap(person)
	if nil != err {
		logrus.Warnf("marshal map failed. %s", err)
		return
	}

	input := &dynamodb.PutItemInput{
		Item:      attribute,
		TableName: aws.String("Lambda"),
	}

	db, err := GetDynamoDB()
	if nil != err {
		logrus.Warnf("get dynamo db failed. %s", err)
		return
	}

	_, err = db.PutItem(input)
	if nil != err {
		logrus.Warnf("put item failed. %s", err)
		return
	}

	fmt.Println("Successfully put item")

	return
}

func DynamoDBGetItem(id string, name string) (person *Person, err error) {
	db, err := GetDynamoDB()
	if nil != err {
		logrus.Warnf("get dynamo db failed. %s", err)
		return
	}

	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Lambda"),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(id),
			},
			"name": {
				S: aws.String(name),
			},
		},
	})
	if nil != err {
		logrus.Warnf("get item failed. %s", err)
		return
	}

	person = &Person{}
	err = dynamodbattribute.UnmarshalMap(result.Item, person)
	if nil != err {
		logrus.Warnf("unmarshal map failed. %s", err)
		return
	}

	fmt.Println("Successfully get item")

	return
}

func main() {
	lambda.Start(Handler)
}
