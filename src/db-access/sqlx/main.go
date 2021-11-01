package main

import (
	"dbAccessSqlx/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

func main() {
	initDB()
	//get()
	add()
}

func initDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test"
	_db, err := sqlx.Open("mysql", dsn)
	//_db, err := sqlx.Connect("mysql", "user=root dbname=test sslmode=disable")
	if err != nil {
		log.Fatalln(err)
		return
	}

	db = _db
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	return
}

func get() {
	orgCourse := model.OrgCourse{}
	err := db.Get(&orgCourse, "SELECT * FROM `org_course` WHERE id = ?", 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(orgCourse)
}

func add() {
	rows := map[string]interface{}{
		"name": "sqlx测试01",
	}
	result, err := db.NamedExec(`INSERT INTO org_course (name) VALUES (:name)`, rows)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
