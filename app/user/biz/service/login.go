package service

import (
	"context"
	"errors"

	"github.com/crazyfrankie/bytedance-mall/app/user/biz/dal/mysql"
	"github.com/crazyfrankie/bytedance-mall/app/user/biz/model"
	user "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	u, err := model.FindByEmail(mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("email or password error")
	}

	resp = &user.LoginResp{
		UserId: int32(u.ID),
	}

	return resp, nil
}
