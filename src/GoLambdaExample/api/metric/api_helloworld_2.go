package metric

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

var HelloWorldFunc ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod:   "POST",
	RelativePath: "/metric/hello_world",
	HandlerFunc: func(ginContext *gin.Context) {
		fmt.Println("hello world func")

	},
}
