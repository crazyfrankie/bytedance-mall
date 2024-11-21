package utils

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/crazyfrankie/bytedance-mall/app/frontend/infra/rpc"
	frontendUtils "github.com/crazyfrankie/bytedance-mall/app/frontend/util"
	"github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/cart"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WrapResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	userId := frontendUtils.GetUserIDFromCtx(ctx)
	content["user_id"] = ctx.Value(frontendUtils.SessionUserId)

	if userId > 0 {
		cartResp, err := rpc.CartClient.GetCart(ctx, &cart.GetCartReq{
			UserId: uint32(frontendUtils.GetUserIDFromCtx(ctx)),
		})
		if err == nil && cartResp != nil {
			content["cart_num"] = len(cartResp.Cart.Items)
		}
	}

	return content
}
