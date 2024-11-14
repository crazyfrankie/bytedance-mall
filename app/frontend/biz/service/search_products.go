package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"

	product "github.com/crazyfrankie/bytedance-mall/app/frontend/hertz_gen/frontend/product"
	"github.com/crazyfrankie/bytedance-mall/app/frontend/infra/rpc"
	rpcproduct "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	products, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{
		Query: req.Q,
	})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"Items": products.Results,
		"q":     req.Q,
	}, nil
}
