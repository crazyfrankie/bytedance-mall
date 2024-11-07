package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	home "github.com/crazyfrankie/bytedance-mall/app/frontend/hertz_gen/frontend/home"
)

type AboutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAboutService(Context context.Context, RequestContext *app.RequestContext) *AboutService {
	return &AboutService{RequestContext: RequestContext, Context: Context}
}

func (h *AboutService) Run(req *home.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	res := make(map[string]any)

	res["Title"] = "About"

	return res, nil
}
