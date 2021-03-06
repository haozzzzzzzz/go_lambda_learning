package home

import (
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

var HandlerFuncName ginbuilder.HandleFunc = ginbuilder.HandleFunc{
	HttpMethod:   "POST",
	RelativePath: "/home/hello_world/:path_id",
	Handle: func(ctx *ginbuilder.Context) (err error) {
		// request path data
		pathData := struct {
			PathId string `json:"path_id" form:"path_id" binding:"required"`
		}{}
		code, err := ctx.BindPathData(&pathData)
		if err != nil {
			ctx.Errorf(code, "verify  path data failed. %s.", err)
			return
		}

		// request query data
		queryData := struct {
			Name string `json:"name" form:"name"`
		}{}
		code, err = ctx.BindQueryData(&queryData)
		if err != nil {
			ctx.Errorf(code, "verify  query data failed. %s.", err)
			return
		}

		// request post data
		postData := struct {
			Address string `json:"address" form:"address"`
		}{}
		code, err = ctx.BindPostData(&postData)
		if err != nil {
			ctx.Errorf(code, "verify  post data failed. %s.", err)
			return
		}

		type Tag struct {
			Name string `json:"name"`
		}

		// response data
		respData := &struct {
			Result  string          `json:"result"`
			Tags    []*Tag          `json:"tags"`
			ATag    *Tag            `json:"a_tag"`
			BTag    Tag             `json:"b_tag"`
			MTag    map[string]*Tag `json:"m_tag"`
			Content struct {
				Title string
			}
			Content2 *struct {
				Detail string
			}
		}{}

		ctx.SuccessReturn(respData)

		return
	},
}
