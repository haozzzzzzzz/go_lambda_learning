package handler

var mainHandler = ApiGatewayProxyEventHandler

func GetMainHandler() interface{} {
	return mainHandler
}
