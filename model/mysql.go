package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/structe/go-meeting/config"
)

var DB *gorm.DB

func InitMysql() (*gorm.DB, error) {
	//fmt.Println(config.MysqlPort)
	//fmt.Println(config.Sqlurl)
	db, err := gorm.Open("mysql", config.Sqlurl)
	if err == nil {
		DB = db
		db.AutoMigrate(&User{}, &Todo{})
		return db, err
	}
	return nil, err
}
