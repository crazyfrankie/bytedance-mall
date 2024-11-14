package service

import (
	"context"

	"github.com/crazyfrankie/bytedance-mall/app/product/biz/dal/mysql"
	"github.com/crazyfrankie/bytedance-mall/app/product/biz/model"
	product "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	categoryQuery := model.NewCategoryQuery(mysql.DB)

	cgs, err := categoryQuery.GetProductsByCategoryName(s.ctx, req.CategoryName)
	if err != nil {
		return nil, err
	}
	
	resp = &product.ListProductsResp{}

	for _, ct := range cgs {
		for _, pt := range ct.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(pt.ID),
				Name:        pt.Name,
				Price:       pt.Price,
				Picture:     pt.Picture,
				Description: pt.Description,
			})
		}
	}
	return resp, nil
}
