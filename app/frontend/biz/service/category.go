package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	category "github.com/crazyfrankie/bytedance-mall/app/frontend/hertz_gen/frontend/category"
	"github.com/crazyfrankie/bytedance-mall/app/frontend/infra/rpc"
	"github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/product"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	p, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{CategoryName: req.Category})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"Title": "Category",
		"Items": p.Products,
	}, nil
}
