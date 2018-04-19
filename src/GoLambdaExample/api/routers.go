package api

import (
	"GoLambdaExample/api/info"
	"GoLambdaExample/api/metric"

	"github.com/gin-gonic/gin"
)

func BindRouters(engine *gin.Engine) (err error) {
	engine.Handle("GET", "/info/hello", info.HelloFunc.HandlerFunc)
	engine.Handle("GET", "/metric/hello", metric.HelloFunc.HandlerFunc)
	engine.Handle("POST", "/metric/hello_world", metric.HelloWorldFunc.HandlerFunc)
	return
}
