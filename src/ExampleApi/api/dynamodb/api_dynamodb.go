package dynamodb

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type User struct {
	Id   string `json:"id"`
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

func DynamoDBGetItem(id string) (person *User, err error) {
	db, err := GetDynamoDB()
	if nil != err {
		logrus.Warnf("get dynamo db failed. %s", err)
		return
	}

	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("user"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if nil != err {
		logrus.Warnf("get item failed. %s", err)
		return
	}

	person = &User{}
	err = dynamodbattribute.UnmarshalMap(result.Item, person)
	if nil != err {
		logrus.Warnf("unmarshal map failed. %s", err)
		return
	}

	fmt.Println("Successfully get item")

	return
}

var ReadDynamoDB ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod:   "GET",
	RelativePath: "/dynamodb/read",
	HandlerFunc: func(ginContext *gin.Context) {
		var (
			msg gin.H
			err error
		)

		defer func() {
			if err != nil {
				ginContext.JSON(200, gin.H{
					"err": err.Error(),
				})
			} else {
				ginContext.JSON(200, msg)
			}
		}()

		id, ok := ginContext.GetQuery("id")
		if !ok {
			err = errors.New("no id specified.")
			return
		}

		user, err := DynamoDBGetItem(id)
		if nil != err {
			err = errors.New(fmt.Sprintf("get item failed. %s", err))
			return
		}

		msg["user"] = user
	},
}
