package dal

import (
	"bytedance-mall/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
