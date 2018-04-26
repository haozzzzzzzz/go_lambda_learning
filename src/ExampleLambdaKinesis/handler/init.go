package handler

var mainHandler = BasicExecutionEventHandler

func GetMainHandler() interface{} {
	return mainHandler
}
