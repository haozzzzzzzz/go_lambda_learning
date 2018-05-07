package handler

var mainHandler = ApiGatewayAuthorizerEventHandler

func GetMainHandler() interface{} {
	return mainHandler
}
