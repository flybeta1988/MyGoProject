package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Hello Mysql !")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:6306)/xnw")
	check(err)
	defer db.Close()

	stmtOut, err := db.Prepare("SELECT id, name FROM test WHERE id = ?")
	check(err)
	defer stmtOut.Close()

	var id int
	var name string
	err = stmtOut.QueryRow(3).Scan(&id, &name)
	check(err)
	fmt.Printf("id:%d name:%s\n", id, name)

	rows, err := db.Query("SELECT id, name FROM test")
	check(err)
	defer rows.Close()

	for rows.Next() {
		var (
			id   int
			name string
		)
		err = rows.Scan(&id, &name)
		fmt.Printf("rows id:%d name:%s\n", id, name)
	}
	err = rows.Err()
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
