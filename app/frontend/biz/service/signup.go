package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	auth "github.com/crazyfrankie/bytedance-mall/app/frontend/hertz_gen/frontend/auth"
	"github.com/hertz-contrib/sessions"
)

type SignupService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSignupService(Context context.Context, RequestContext *app.RequestContext) *SignupService {
	return &SignupService{RequestContext: RequestContext, Context: Context}
}

func (h *SignupService) Run(req *auth.SignupReq) (resp *auth.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 1)
	err = session.Save()
	if err != nil {
		return nil, err
	}

	return
}
