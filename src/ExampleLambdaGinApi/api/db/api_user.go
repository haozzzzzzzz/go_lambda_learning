package db

import (
	dynamodb2 "github.com/haozzzzzzzz/go-lambda/resource/dynamodb"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

type User struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
}

var AddUser ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod:   "POST",
	RelativePath: "/user/add",
	Handle: func(ctx *ginbuilder.Context) (err error) {
		// 校验post data
		postData := struct {
			Uid  string `json:"uid" yaml:"uid" binding:"required"`
			Name string `json:"name" yaml:"name" binding:"required"`
		}{}
		code, err := ctx.BindPostData(&postData)
		if nil != err {
			ctx.Errorf(code, "bind post data failed. %s", err)
			return
		}

		errorServer := ginbuilder.CodeErrorServer.Clone()
		db, err := dynamodb2.GetDynamodb("ap-south-1") // 这里设置可用区
		if nil != err {
			ctx.Errorf(errorServer, "get dynamodb failed. %s", err)
			return
		}

		userMap, err := dynamodbattribute.MarshalMap(postData)
		if nil != err {
			ctx.Errorf(errorServer, "marshal user data failed. %s", err)
			return
		}

		_, err = db.PutItem(&dynamodb.PutItemInput{
			TableName: aws.String("user"),
			Item:      userMap,
		})
		if nil != err {
			ctx.Errorf(errorServer, "put dynamodb item failed. %s", err)
			return
		}

		// 返回成功
		ctx.Success()

		return
	},
}

var GetUser ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod:   "GET",
	RelativePath: "/user/get",
	Handle: func(ctx *ginbuilder.Context) (err error) {

		// 校验query data
		queryData := struct {
			Uid string `form:"uid" json:"uid" binding:"required"`
		}{}
		code, err := ctx.BindQueryData(&queryData)
		if err != nil {
			ctx.Errorf(code, "bind query data failed. %s", err)
			return
		}

		errServer := ginbuilder.CodeErrorServer.Clone()
		// 从DynamoDB中获取数据
		db, err := dynamodb2.GetDynamodb("ap-south-1") // 这里设置可用区
		if nil != err {
			ctx.Errorf(errServer, "get dynamodb failed. %s", err)
			return
		}

		result, err := db.GetItem(&dynamodb.GetItemInput{
			TableName: aws.String("user"), // 设置表名
			Key: map[string]*dynamodb.AttributeValue{
				"uid": {
					S: aws.String(queryData.Uid), // 设置ID
				},
			},
		})
		if nil != err {
			ctx.Errorf(errServer, "get dynamodb item failed. %s", err)
			return
		}

		person := &User{}
		err = dynamodbattribute.UnmarshalMap(result.Item, person)
		if nil != err {
			ctx.Errorf(errServer, "unmarshal map failed. \n%s.", err)
			return
		}

		// 返回用户
		ctx.SuccessReturn(person)

		return
	},
}
