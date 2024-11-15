package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/crazyfrankie/bytedance-mall/app/cart/biz/dal/mysql"
	"github.com/crazyfrankie/bytedance-mall/app/cart/biz/model"
	cart "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/cart"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	cartQuery := model.NewCartQuery(mysql.DB)

	err = cartQuery.EmptyCart(s.ctx, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}

	return &cart.EmptyCartResp{}, nil
}
