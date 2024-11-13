package dal

import (
	"github.com/crazyfrankie/bytedance-mall/app/product/biz/dal/mysql"
	"github.com/crazyfrankie/bytedance-mall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
