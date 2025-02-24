package dal

import (
	"github.com/XJTU-zxc/MyCloudWeGo/app/frontend/biz/dal/mysql"
	"github.com/XJTU-zxc/MyCloudWeGo/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
