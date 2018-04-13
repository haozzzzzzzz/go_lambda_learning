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

var counter int

// 请求处理器
func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	counter++
	return fmt.Sprintf("Hello %s!. num: %d.", event.Name, counter), nil
}

func main() {
	lambda.Start(HandleRequest)
}
