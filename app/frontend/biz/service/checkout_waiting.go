package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	checkout "github.com/crazyfrankie/bytedance-mall/app/frontend/hertz_gen/frontend/checkout"
	"github.com/crazyfrankie/bytedance-mall/app/frontend/infra/rpc"
	frontendutils "github.com/crazyfrankie/bytedance-mall/app/frontend/util"
	rpccheckout "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/checkout"
	rpcpayment "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/payment"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	userId := frontendutils.GetUserIDFromCtx(h.Context)
	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    uint32(userId),
		Email:     req.Email,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Address: &rpccheckout.Address{
			Country:       req.Country,
			City:          req.City,
			ZipCode:       req.Zipcode,
			State:         req.Province,
			StreetAddress: req.Street,
		},
		CreditCard: &rpcpayment.CreditCardInfo{
			CreditCardNumber:          req.CardNum,
			CreditCardCvv:             req.Cvv,
			CreditCardExpirationYear:  req.ExpirationYear,
			CreditCardExpirationMonth: req.ExpirationMonth,
		},
	})

	if err != nil {
		return nil, err
	}

	return utils.H{
		"Title":    "waiting",
		"Redirect": "/checkout/result",
	}, nil
}
