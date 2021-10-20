package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xnw.com/utils"
)

var DB *sql.DB

func init() {
	fmt.Println("models init start ...")
	fmt.Println("DB type: %T", DB)
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	utils.CheckError(err)
	err = db.Ping()
	utils.CheckError(err)
	DB = db
}

func Foo() {
	fmt.Println("models const Foo()")
}