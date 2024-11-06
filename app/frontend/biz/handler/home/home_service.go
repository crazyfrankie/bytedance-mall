package home

import (
	"context"

	"frontend/biz/service"
	"frontend/biz/utils"
	home "frontend/hertz_gen/home"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
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

	resp := &home.Empty{}
	resp, err = service.NewHomeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "home.html", resp)
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
