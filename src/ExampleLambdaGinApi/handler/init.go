package handler

var mainHandler = ApiGatewayEventHandler

func GetMainHandler() interface{} {
	return mainHandler
}
