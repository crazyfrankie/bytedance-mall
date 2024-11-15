package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"

	"github.com/crazyfrankie/bytedance-mall/app/frontend/conf"
	frontendUtils "github.com/crazyfrankie/bytedance-mall/app/frontend/util"
	"github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client
	CartClient    cartservice.Client
	once          sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
	})
}

func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	CartClient, err = cartservice.NewClient("cart", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}
