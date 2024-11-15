package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	cart "github.com/crazyfrankie/bytedance-mall/app/frontend/hertz_gen/frontend/cart"
	"github.com/crazyfrankie/bytedance-mall/app/frontend/infra/rpc"
	frontendUtils "github.com/crazyfrankie/bytedance-mall/app/frontend/util"
	rpccart "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/cart"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartReq) (resp *cart.Empty, err error) {
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: uint32(frontendUtils.GetUserIDFromCtx(h.Context)),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  req.ProductNum,
		},
	})
	if err != nil {
		return nil, err
	}

	return
}
