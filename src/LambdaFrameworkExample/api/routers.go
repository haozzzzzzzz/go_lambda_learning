package api

import (
	"LambdaFrameworkExample/api/metric"

	"github.com/gin-gonic/gin"
)

func BindRouters(engine *gin.Engine) (err error) {
	engine.Handle("GET", "/metric", metric.MetricHandlerFunc.HandlerFunc)
	return
}
