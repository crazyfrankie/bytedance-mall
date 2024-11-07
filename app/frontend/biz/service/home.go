package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	home "github.com/crazyfrankie/bytedance-mall/app/frontend/hertz_gen/frontend/home"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	resp := make(map[string]any)
	items := []map[string]any{
		{"Name": "毛泽东选集", "Price": "100", "Picture": "/static/image/image.png"},
		{"Name": "Go 语言高级编程", "Price": "25", "Picture": "/static/image/goadvance.png"},
		{"Name": "Go 语言高级编程", "Price": "25", "Picture": "/static/image/goadvance.png"},
		{"Name": "Go 语言高级编程", "Price": "25", "Picture": "/static/image/goadvance.png"},
		{"Name": "Go 语言设计与实现", "Price": "25", "Picture": "/static/image/godesign.png"},
		{"Name": "Go 语言设计与实现", "Price": "25", "Picture": "/static/image/godesign.png"},
		{"Name": "Go 语言设计与实现", "Price": "25", "Picture": "/static/image/godesign.png"},
		{"Name": "Go 语言设计与实现", "Price": "25", "Picture": "/static/image/godesign.png"},
	}

	resp["Title"] = "Hot sales"
	resp["Items"] = items

	return resp, nil
}
