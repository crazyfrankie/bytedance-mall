package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/crazyfrankie/bytedance-mall/app/product/biz/dal/mysql"
	"github.com/crazyfrankie/bytedance-mall/app/product/biz/model"
	product "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required")
	}

	productQuery := model.NewProductQuery(mysql.DB)

	p, err := productQuery.GetByID(s.ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Price:       p.Price,
			Description: p.Description,
			Picture:     p.Picture,
		},
	}, nil
}
