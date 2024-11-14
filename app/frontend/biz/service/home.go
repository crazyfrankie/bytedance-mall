package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	home "github.com/crazyfrankie/bytedance-mall/app/frontend/hertz_gen/frontend/home"
	"github.com/crazyfrankie/bytedance-mall/app/frontend/infra/rpc"
	"github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/product"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (map[string]any, error) {
	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{CategoryName: "books"})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"Title": "Hot Sale",
		"Items": products.Products,
	}, nil
}
