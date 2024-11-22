package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/crazyfrankie/bytedance-mall/app/checkout/infra/mq"
	"github.com/crazyfrankie/bytedance-mall/app/checkout/infra/rpc"
	"github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/cart"
	checkout "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/checkout"
	"github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/email"
	"github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/order"
	"github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/payment"
	"github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/product"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// 1. 先从购物车获取商品
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Cart.Items == nil {
		return nil, kerrors.NewBizStatusError(5004001, "cart is empty")
	}

	var (
		total float32
		oi    []*order.OrderItem
	)

	// 2. 计算商品总价
	for _, item := range cartResult.Cart.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: item.ProductId})
		if resultErr != nil {
			return nil, resultErr
		}

		if productResp.Product == nil {
			continue
		}

		p := productResp.Product.Price

		cost := p * float32(item.Quantity)
		total += cost

		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: item.ProductId,
				Quantity:  item.Quantity,
			},
			Cost: cost,
		})
	}

	// 3. 创建订单
	var orderId string

	// u, _ := uuid.NewRandom()
	// orderId = u.String()
	zipCode, _ := strconv.Atoi(req.Address.ZipCode)
	orderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId: req.UserId,
		Email:  req.Email,
		Address: &order.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			Country:       req.Address.Country,
			State:         req.Address.State,
			ZipCode:       int32(zipCode),
		},
		OrderItems: oi,
	})
	if err != nil {
		return nil, kerrors.NewBizStatusError(5004002, err.Error())
	}

	if orderResp != nil && orderResp.Order != nil {
		orderId = orderResp.Order.OrderId
	}

	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
		},
	}

	// 4. 清空购物车
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err.Error())
	}

	// 5. 付款
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}

	//6. 添加邮箱生产者
	data, _ := proto.Marshal(&email.EmailReq{
		From:        "123@qq.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "You have just created an order in GoShop",
		Content:     "You have just created an order in GoShop",
	})

	msg := &nats.Msg{Subject: "email", Data: data}

	_ = mq.Nc.PublishMsg(msg)

	klog.Info(paymentResult)

	return &checkout.CheckoutResp{OrderId: orderId, TransactionId: paymentResult.TransactionId}, nil
}
