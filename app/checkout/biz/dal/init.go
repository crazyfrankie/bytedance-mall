package dal

import (
	"github.com/crazyfrankie/bytedance-mall/app/checkout/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
