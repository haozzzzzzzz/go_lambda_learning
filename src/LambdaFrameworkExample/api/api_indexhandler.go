package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

var IndexHandler ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod:   "GET",
	RelativePath: "/",
	HandlerFunc: func(ginContext *gin.Context) {
		fmt.Println("/Users/hao/Documents/Projects/Github/go_lambda_learning/src/LambdaFrameworkExample/api/api_indexhandler.go")
	},
}