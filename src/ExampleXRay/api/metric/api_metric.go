package metric

import (
	_ "github.com/haozzzzzzzz/go-lambda/resource/xray"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

var MetricHandlerFunc ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod:   "GET",
	RelativePath: "/api/v1/metric",
	Handle: func(ctx *ginbuilder.Context) (err error) {
		ctx.SuccessReturn(map[string]interface{}{
			"info": "hello, world",
		})
		return
	},
}

var TestPathParams ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod:   "GET",
	RelativePath: "/api/v1/test_path/:key",
	Handle: func(ctx *ginbuilder.Context) (err error) {
		// request path data
		pathData := struct {
			Key string `json:"key" form:"key" binding:"required"`
		}{}
		code, err := ctx.BindPathData(&pathData)
		if err != nil {
			ctx.Errorf(code, "verify  path data failed. \n%s.", err)
			return
		}

		ctx.SuccessReturn(pathData)

		return
	},
}
