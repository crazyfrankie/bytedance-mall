package dal

import (
	"github.com/crazyfrankie/bytedance-mall/app/product/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
