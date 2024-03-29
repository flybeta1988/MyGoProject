package dbpkg

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var DbErr error

func init() {
	DB, DbErr = sql.Open("mysql", "root:root@tcp(127.0.0.1:6306)/xnw")
	//defer DB.Close()

	if DbErr != nil {
		panic("数据库配置错误：" + DbErr.Error())
	}

	// 最大连接数
	DB.SetMaxOpenConns(100)
	// 闲置连接数
	DB.SetMaxIdleConns(20)
	// 最大连接周期
	DB.SetConnMaxLifetime(100 * time.Second)

	if DbErr = DB.Ping(); nil != DbErr {
		panic("数据库连接失败: " + DbErr.Error())
	}
}

func Foo() {
	fmt.Println("dbpkg Foo")
}
