package info

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

var HelloFunc ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod:   "GET",
	RelativePath: "/info/hello",
	HandlerFunc: func(ginContext *gin.Context) {
		fmt.Println("hello func")
	},
}
