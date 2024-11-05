package dal

import (
	"frontend/biz/dal/mysql"
	"frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
