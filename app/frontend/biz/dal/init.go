package dal

import (
	"github.com/crazyfrankie/bytedance-mall/app/frontend/biz/dal/mysql"
	"github.com/crazyfrankie/bytedance-mall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
