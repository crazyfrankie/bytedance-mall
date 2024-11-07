package home

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/crazyfrankie/bytedance-mall/app/frontend/biz/service"
	"github.com/crazyfrankie/bytedance-mall/app/frontend/biz/utils"
	home "github.com/crazyfrankie/bytedance-mall/app/frontend/hertz_gen/frontend/home"
)

// Home .
// @router / [GET]
func Home(ctx context.Context, c *app.RequestContext) {
	var err error
	var req home.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewHomeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "home", utils.WrapResponse(ctx, c, resp))
}
