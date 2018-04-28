package metric

import (
	_ "github.com/haozzzzzzzz/go-lambda/resource/xray"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

var MetricHandlerFunc ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod:   "GET",
	RelativePath: "/metric",
	Handle: func(ctx *ginbuilder.Context) (err error) {
		ctx.SuccessReturn(map[string]interface{}{
			"info": "hello, world",
		})
		return
	},
}
