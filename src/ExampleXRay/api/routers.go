package api

import (
	"ExampleXRay/api/metric"

	"github.com/gin-gonic/gin"
)

// 注意：BindRouters函数体内不能自定义添加任何声明，由lambda-build compile api命令生成api绑定声明
func BindRouters(engine *gin.Engine) (err error) {
	engine.Handle("GET", "/api/v1/test_path/:key", metric.TestPathParams.GinHandler)
	engine.Handle("GET", "/api/v1/metric", metric.MetricHandlerFunc.GinHandler)
	return
}
