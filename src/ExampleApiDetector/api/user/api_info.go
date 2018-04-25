package user

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
	dynamodb2 "github.com/haozzzzzzzz/go-lambda/resource/dynamodb"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
	"github.com/sirupsen/logrus"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var DynamoDBRead ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod:   "GET",
	RelativePath: "/user/read",
	HandlerFunc: func(ginContext *gin.Context) {
		var err error
		response := gin.H{}

		defer func() {
			if err != nil {
				ginContext.JSON(200, gin.H{
					"error": err.Error(),
				})
			} else {
				ginContext.JSON(200, response)
			}
		}()

		id, ok := ginContext.GetQuery("id")
		if !ok {
			err = errors.New("require id.")
			return
		}

		db, err := dynamodb2.GetDynamodb("us-east-1")
		if nil != err {
			logrus.Errorf("get dynamodb failed. \n%s.", err)
			return
		}

		result, err := db.GetItem(&dynamodb.GetItemInput{
			TableName: aws.String("user"),
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					N: aws.String(id),
				},
			},
		})

		if nil != err {
			logrus.Errorf("get item failed. \n%s.", err)
			return
		}

		person := &User{}
		err = dynamodbattribute.UnmarshalMap(result.Item, person)
		if nil != err {
			logrus.Errorf("unmarshal map failed. \n%s.", err)
			return
		}

		response["user"] = person
	},
}
