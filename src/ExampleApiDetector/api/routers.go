package api

import (
	"ExampleApiDetector/api/metric"
	"ExampleApiDetector/api/user"

	"github.com/gin-gonic/gin"
)

// 注意：BindRouters函数体内不能自定义添加任何声明，由lambda-build compile api命令生成api绑定声明
func BindRouters(engine *gin.Engine) (err error) {
	engine.Handle("GET", "/metric", metric.MetricHandlerFunc.HandlerFunc)
	engine.Handle("GET", "/user/read", user.DynamoDBRead.HandlerFunc)
	return
}
