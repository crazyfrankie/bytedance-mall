package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/crazyfrankie/bytedance-mall/app/cart/biz/dal/mysql"
	"github.com/crazyfrankie/bytedance-mall/app/cart/biz/model"
	cart "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	cartQuery := model.NewCartQuery(mysql.DB)

	carts, err := cartQuery.GetCart(s.ctx, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50002, err.Error())
	}

	var items []*cart.CartItem
	for _, ct := range carts {
		items = append(items, &cart.CartItem{
			ProductId: ct.ProductID,
			Quantity:  ct.Quantity,
		})
	}

	return &cart.GetCartResp{Cart: &cart.Cart{UserId: req.UserId, Items: items}}, nil
}
