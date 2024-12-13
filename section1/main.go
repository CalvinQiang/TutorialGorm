package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"models"
)

func main() {

	dsn := "root:cd635750e8d693e6@tcp(127.0.0.1:3306)/db_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("打开数据失败，错误内容：%#v", err)
	}

	teacher := models.Teachers{
		Name:        "测试",
		Email:       "ceshiCeshi",
		PhoneNumber: "13300000000",
		Description: "测试毕业于XXX",
	}
	result := db.Select("Name").Create(&teacher)
	if result.Error != nil {
		log.Fatalf("创建数据表错误，错误内容：%#v", result.Error)
	}
	log.Printf("插入成功%#v\n", result)
}
