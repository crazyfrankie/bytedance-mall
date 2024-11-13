package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"

	"github.com/crazyfrankie/bytedance-mall/app/frontend/conf"
	frontendUtils "github.com/crazyfrankie/bytedance-mall/app/frontend/util"
	"github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient userservice.Client

	once sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}
