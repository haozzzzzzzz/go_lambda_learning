package metric

import (
	"fmt"

	dynamodb2 "github.com/haozzzzzzzz/go-lambda/resource/dynamodb"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
	"github.com/sirupsen/logrus"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var MetricHandlerFunc ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod:   "GET",
	RelativePath: "/metric",
	Handle: func(ctx *ginbuilder.Context) (resp interface{}, err error) {
		defer func() {
			if err != nil {
				resp = &ginbuilder.ResponseBase{
					ReturnCode: 1,
					Message:    err.Error(),
				}
			}
		}()

		queryData := struct {
			Id string `form:"id" binding:"required"`
		}{}
		err = ctx.BindQueryData(&queryData)
		if err != nil {
			return
		}

		response := &struct {
			ginbuilder.ResponseBase
			User *User `json:"user"`
		}{}
		resp = response

		response.ReturnCode = 0
		response.Message = "normal"

		db, err := dynamodb2.GetDynamodb("us-east-1")
		if nil != err {
			logrus.Errorf("get dynamodb failed. \n%s.", err)
			return
		}

		result, err := db.GetItem(&dynamodb.GetItemInput{
			TableName: aws.String("user"),
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					N: aws.String(queryData.Id),
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

		fmt.Println(person)

		response.User = person

		return
	},
}
