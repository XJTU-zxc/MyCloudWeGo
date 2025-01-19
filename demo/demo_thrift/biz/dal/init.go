package dal

import (
	"github.com/XJTU-zxc/MyCloudWeGo/demo/demo_thrift/biz/dal/mysql"
	"github.com/XJTU-zxc/MyCloudWeGo/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
