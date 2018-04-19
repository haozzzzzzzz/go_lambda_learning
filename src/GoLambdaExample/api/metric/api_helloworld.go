package metric

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

func init() {
	fmt.Println("metric init")
}

var HelloFunc ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod:   "GET",
	RelativePath: "/metric/hello",
	HandlerFunc: func(ginContext *gin.Context) {
		fmt.Println("hello func")
	},
}