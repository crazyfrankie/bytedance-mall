package dal

import (
	"github.com/crazyfrankie/bytedance-mall/app/payment/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
