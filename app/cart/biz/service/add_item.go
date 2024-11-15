package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/crazyfrankie/bytedance-mall/app/cart/biz/dal/mysql"
	"github.com/crazyfrankie/bytedance-mall/app/cart/biz/model"
	"github.com/crazyfrankie/bytedance-mall/app/cart/infra/rpc"
	cart "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/cart"
	"github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/product"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	pt, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}
	if pt == nil || pt.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(400000, "product not found")
	}

	cartQuery := model.NewCartQuery(mysql.DB)

	err = cartQuery.AddItem(s.ctx, &model.Cart{
		UserID:    req.UserId,
		ProductID: req.Item.ProductId,
		Quantity:  req.Item.Quantity,
	})

	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}

	return &cart.AddItemResp{}, nil
}
