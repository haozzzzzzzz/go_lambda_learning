package handler

import (
	"context"
	"fmt"

	"ExampleApiAuthorizer/constant"

	"github.com/aws/aws-lambda-go/events"
	"github.com/haozzzzzzzz/go-lambda/resource/apigateway"
)

func ApiGatewayAuthorizerEventHandler(ctx context.Context, event events.APIGatewayCustomAuthorizerRequestTypeRequest) (response *events.APIGatewayCustomAuthorizerResponse, err error) {
	fmt.Println(event.Headers)
	response = apigateway.GetAllowAuthorizerResponse(constant.LambdaFunctionName, event.MethodArn)
	return
}
