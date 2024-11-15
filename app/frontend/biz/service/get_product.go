package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"

	product "github.com/crazyfrankie/bytedance-mall/app/frontend/hertz_gen/frontend/product"
	"github.com/crazyfrankie/bytedance-mall/app/frontend/infra/rpc"
	rpcproduct "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/product"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	p, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"Title": p.Product.Name,
		"Item":  p.Product,
	}, nil
}
