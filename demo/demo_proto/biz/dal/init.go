package dal

import (
	"github.com/XJTU-zxc/MyCloudWeGo/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
