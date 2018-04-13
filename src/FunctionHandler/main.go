package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// 定义事件
type MyEvent struct {
	Name string `json:"name"`
}

// 请求处理器
func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", event.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
