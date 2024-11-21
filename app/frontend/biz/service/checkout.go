package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	checkout "github.com/crazyfrankie/bytedance-mall/app/frontend/hertz_gen/frontend/checkout"
	"github.com/crazyfrankie/bytedance-mall/app/frontend/infra/rpc"
	frontendutils "github.com/crazyfrankie/bytedance-mall/app/frontend/util"
	rpccart "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/product"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *checkout.Empty) (resp map[string]any, err error) {
	var items []map[string]any
	userId := frontendutils.GetUserIDFromCtx(h.Context)

	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}

	var total float32

	for _, v := range carts.Cart.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: v.ProductId})
		if err != nil {
			return nil, err
		}

		if productResp.Product == nil {
			continue
		}

		p := productResp.Product
		items = append(items, map[string]any{
			"Name":     p.Name,
			"Price":    strconv.FormatFloat((float64(p.Price)), 'f', 2, 64),
			"Picture":  p.Picture,
			"Quantity": strconv.Itoa(int(v.Quantity)),
		})
		total += float32(v.Quantity) * p.Price
	}

	return utils.H{
		"Title": "Checkout",
		"Items": items,
		"Total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
