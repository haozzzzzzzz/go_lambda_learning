package api

import (
	"GoLambdaExample/api/metric"

	"github.com/gin-gonic/gin"
)

func BindRouters(engine *gin.Engine) (err error) {
	engine.Handle("GET", "/metric/hello", metric.HelloFunc.HandlerFunc)
	// TODO add-new-handle
	return
}
