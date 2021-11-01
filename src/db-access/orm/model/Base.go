package model

import "gorm.io/gorm"

type Base struct {
	ID uint
	Ctime uint `gorm:"autoCreateTime"`
	Utime uint `gorm:"autoUpdateTime"`
	Dtime gorm.DeletedAt
}
