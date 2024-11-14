package service

import (
	"context"

	"github.com/crazyfrankie/bytedance-mall/app/product/biz/dal/mysql"
	"github.com/crazyfrankie/bytedance-mall/app/product/biz/model"
	product "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	productQuery := model.NewProductQuery(mysql.DB)

	products, err := productQuery.SearchProducts(s.ctx, req.Query)
	if err != nil {
		return nil, err
	}

	var results []*product.Product

	for _, pt := range products {
		results = append(results, &product.Product{
			Id:          uint32(pt.ID),
			Name:        pt.Name,
			Description: pt.Name,
			Picture:     pt.Picture,
			Price:       pt.Price,
		})
	}

	return &product.SearchProductsResp{Results: results}, nil
}
