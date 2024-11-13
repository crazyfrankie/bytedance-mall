package service

import (
	"context"
	"testing"

	"github.com/crazyfrankie/bytedance-mall/app/user/biz/dal/mysql"
	user "github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestRegister_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()

	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "demo@admin.com",
		Password:        "SADASFASSAFAS",
		PasswordConfirm: "SADASFASSAFAS",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
