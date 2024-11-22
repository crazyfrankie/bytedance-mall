package dal

import (
	"github.com/crazyfrankie/bytedance-mall/app/email/biz/dal/mysql"
	"github.com/crazyfrankie/bytedance-mall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
