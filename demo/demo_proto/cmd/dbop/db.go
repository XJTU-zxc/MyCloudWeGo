package main

import (
	"log"

	"github.com/XJTU-zxc/MyCloudWeGo/demo/demo_proto/biz/dal"
	// "github.com/XJTU-zxc/MyCloudWeGo/demo/demo_proto/biz/dal/mysql"
	// "github.com/XJTU-zxc/MyCloudWeGo/demo/demo_proto/biz/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dal.Init()
	// CURD
	// mysql.DB.Create(&model.User{Email: "zxc@qq.com", Password: "123456",})
	// mysql.DB.Model(&model.User{}).Where("email = ?", "zxc@qq.com").Update("password", "q1q11")
	// var row model.User
	// mysql.DB.Model(&model.User{}).Where("email = ?", "zxc@qq.com").First(&row)
	// fmt.Printf("%+v", row)
	// soft delete
	// mysql.DB.Where("email = ?", "zxc@qq.com").Delete(&model.User{})
	// hard delete
	// mysql.DB.Unscoped().Where("email = ?", "zxc@qq.com").Delete(&model.User{})

}
