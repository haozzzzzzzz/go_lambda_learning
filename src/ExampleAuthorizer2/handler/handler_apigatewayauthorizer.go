package handler

import (
	"ExampleAuthorizer2/constant"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/haozzzzzzzz/go-lambda/resource/apigateway"
)

func ApiGatewayAuthorizerEventHandler(ctx context.Context, event events.APIGatewayCustomAuthorizerRequestTypeRequest) (response *events.APIGatewayCustomAuthorizerResponse, err error) {
	response = apigateway.GetAllowAuthorizerResponse(constant.LambdaFunctionName, event.MethodArn)
	return
}
