package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/crazyfrankie/bytedance-mall/app/order/biz/dal/mysql"
	"github.com/crazyfrankie/bytedance-mall/app/order/biz/model"
	order "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/order"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	if len(req.OrderItems) == 0 {
		err = kerrors.NewBizStatusError(500001, "item is empty")
		return
	}

	mysql.DB.Transaction(func(tx *gorm.DB) error {
		orderId, _ := uuid.NewRandom()

		o := &model.Order{
			OrderId:      orderId.String(),
			UserId:       uint(req.UserId),
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}
		if req.Address != nil {
			a := req.Address
			o.Consignee.State = a.State
			o.Consignee.City = a.City
			o.Consignee.Country = a.Country
			o.Consignee.StreetAddress = a.StreetAddress
		}
		if err := tx.Create(o).Error; err != nil {
			return err
		}

		var items []model.OrderItem
		for _, v := range req.OrderItems {
			items = append(items, model.OrderItem{
				OrderIdRefer: orderId.String(),
				ProductId:    v.Item.ProductId,
				Quantity:     v.Item.Quantity,
				Cost:         v.Cost,
			})
		}

		if err := tx.Create(items).Error; err != nil {
			return err
		}

		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: orderId.String(),
			},
		}

		return nil
	})

	return
}
