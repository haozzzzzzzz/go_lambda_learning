package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
)

// 定义事件
type MyEvent struct {
	Name string `json:"name"`
}

// 请求处理器
func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	byteConfig, err := ioutil.ReadFile("./config.yaml")
	if nil != err {
		logrus.Errorf("read config.yaml failed. \n%s.", err)
		return "error", nil
	}

	event.Name = string(byteConfig)

	return fmt.Sprintf("Hello %s!.", event.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
