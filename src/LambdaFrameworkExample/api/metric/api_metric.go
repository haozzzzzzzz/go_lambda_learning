package metric

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

var MetricHandlerFunc ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod:   "GET",
	RelativePath: "/metric",
	HandlerFunc: func(ginContext *gin.Context) {
		fmt.Println("/Users/hao/Documents/Projects/Github/go_lambda_learning/src/LambdaFrameworkExample/api/metric/api_metric.go")
	},
}
