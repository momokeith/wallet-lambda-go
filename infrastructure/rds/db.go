package rds

import (
	"github.com/jinzhu/gorm"
	"os"
	"fmt"
)

func GetDB() *gorm.DB{
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?", os.Getenv("DB_USERNAME"),os.Getenv("DB_PASSWORD"),os.Getenv("DB_ENDPOINT_ADDRESS"),os.Getenv("DB_NAME"))
	println(args)
	db, err := gorm.Open(os.Getenv("DB_DRIVER"),args)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
