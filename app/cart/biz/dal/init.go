package dal

import (
	"github.com/crazyfrankie/bytedance-mall/app/cart/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
