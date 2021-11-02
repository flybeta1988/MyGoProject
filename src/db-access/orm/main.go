package main

import (
	"dbAccessOrm/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var db *gorm.DB

func main() {
	initDB()
	test2()
	//update()
}

func initDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = _db
}

func update() {
	var orgCourse model.OrgCourse
	db.First(&orgCourse, 5)
	orgCourse.Name = fmt.Sprintf("测试%d", orgCourse.ID)
	db.Save(&orgCourse)
	fmt.Println(orgCourse)
}

func test2() {
	db.Create(&model.OrgCourse{Name: "测试01"})

	var orgCoruse model.OrgCourse
	db.First(&orgCoruse, 1)
	fmt.Println(orgCoruse)
}

func test1()  {
	// Migrate the schema
	//db.AutoMigrate(&Product{})

	// Create
	//db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	db.First(&product, 1) // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	fmt.Println(product)
}
