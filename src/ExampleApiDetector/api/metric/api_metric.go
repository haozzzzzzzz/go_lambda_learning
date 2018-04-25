package metric

import (
	"github.com/gin-gonic/gin"
	_ "github.com/haozzzzzzzz/go-lambda/resource/dynamodb"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

var MetricHandlerFunc ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod:   "GET",
	RelativePath: "/metric",
	HandlerFunc: func(ginContext *gin.Context) {
		ginContext.JSON(200, gin.H{
			"message": "hello, lambda api with gin style.",
		})
	},
}
