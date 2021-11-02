package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Base struct {
	ID uint	`json:"id"`
	Ctime uint `gorm:"autoCreateTime" json:"ctime"`
	Utime uint `gorm:"autoUpdateTime" json:"utime"`
	Dtime gorm.DeletedAt `json:"dtime"`
}

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = _db
}
