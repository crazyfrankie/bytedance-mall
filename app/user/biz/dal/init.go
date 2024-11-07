package dal

import (
	"github.com/crazyfrankie/bytedance-mall/app/user/biz/dal/mysql"
	"github.com/crazyfrankie/bytedance-mall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
