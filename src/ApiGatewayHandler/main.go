package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
)

// 请求
type Ask struct {
	Operation string                 `json:"operation"` // 操作
	TableName string                 `json:"tableName"` // 表名
	Payload   map[string]interface{} `json:"payload"`   // 数据负载
}

// 返回
type Ack struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
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

	case "read":
	case "update":
	case "delete":
	case "list":
	case "echo":
	case "ping":
	default:

	}

	return
}

func main() {
	lambda.Start(Handler)
}
