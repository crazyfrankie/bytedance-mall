package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"

	"github.com/crazyfrankie/bytedance-mall/app/frontend/hertz_gen/frontend/user"
	"github.com/crazyfrankie/bytedance-mall/app/frontend/infra/rpc"
	ue "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/user"
)

type SignupService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSignupService(Context context.Context, RequestContext *app.RequestContext) *SignupService {
	return &SignupService{RequestContext: RequestContext, Context: Context}
}

func (h *SignupService) Run(req *user.SignupReq) (resp *user.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userResp, err := rpc.UserClient.Register(h.Context, &ue.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	})
	if err != nil {
		return nil, err
	}

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", userResp.UserId)
	err = session.Save()
	if err != nil {
		return nil, err
	}

	return
}
