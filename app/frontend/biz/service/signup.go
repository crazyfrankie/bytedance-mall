package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"

	auth "frontend/hertz_gen/auth"
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
	// TODO user svc api
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 1)
	session.Save()
	return
}
