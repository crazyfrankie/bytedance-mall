package service

import (
	pbapi "bytedance-mall/kitex_gen/pbapi"
	"context"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(req *pbapi.Request) (resp *pbapi.Response, err error) {
	// Finish your business logic.

	return &pbapi.Response{Message: req.Message}, nil
}
