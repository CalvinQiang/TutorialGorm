package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"models"
)

func main() {

	dsn := "root:cd635750e8d693e6@tcp(127.0.0.1:3306)/db_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	teacher := models.Teachers{}
	tx := db.First(&teacher)
	log.Println(tx.RowsAffected)
	log.Printf("%#v\n", teacher)

}
