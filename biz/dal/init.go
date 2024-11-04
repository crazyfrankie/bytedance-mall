package dal

import (
	"bytedance-mall/biz/dal/mysql"
	"bytedance-mall/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
