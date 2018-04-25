package metric

import (
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

var MetricHandlerFunc ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod:   "GET",
	RelativePath: "/metric",
	Handle: func(ctx *ginbuilder.Context) (err error) {
		ctx.Success(map[string]interface{}{
			"info": "hello, world",
		})
		return
	},
}
