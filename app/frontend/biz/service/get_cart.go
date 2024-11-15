package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	cart "github.com/crazyfrankie/bytedance-mall/app/frontend/hertz_gen/frontend/cart"
	"github.com/crazyfrankie/bytedance-mall/app/frontend/infra/rpc"
	frontendUtils "github.com/crazyfrankie/bytedance-mall/app/frontend/util"
	rpccart "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/cart"
	"github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/product"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *cart.Empty) (resp map[string]any, err error) {
	cartResp, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: uint32(frontendUtils.GetUserIDFromCtx(h.Context)),
	})
	if err != nil {
		return nil, err
	}

	var items []map[string]string
	var total float64
	for _, v := range cartResp.Cart.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{
			Id: v.ProductId,
		})
		if err != nil {
			continue
		}

		p := productResp.Product
		items = append(items, map[string]string{
			"Name":        p.Name,
			"Description": p.Description,
			"Price":       strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture":     p.Picture,
			"Quantity":    strconv.Itoa(int(v.Quantity)),
		})
		total += float64(p.Price) * float64(v.Quantity)
	}

	return utils.H{
		"Title": "Cart",
		"Items": items,
		"Total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
