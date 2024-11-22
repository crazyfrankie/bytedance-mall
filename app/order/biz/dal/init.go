package dal

import (
	"github.com/crazyfrankie/bytedance-mall/app/order/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
