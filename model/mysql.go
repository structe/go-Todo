package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitMysql() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		DB = db
		db.AutoMigrate(&User{})
		return db, err
	}
	return nil, err
}
