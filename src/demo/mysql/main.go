package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Hello Mysql !")

	fmt.Println(os.Args)

	args := os.Args
	actStr := args[1]
	actArr := strings.Split(actStr, "=")
	act := actArr[1]
	fmt.Printf("act dataType is: %T\n", act)
	fmt.Println(act)
	switch {
		case act == "list":
			fmt.Println("this is list act")
		case act == "add":
			fmt.Println("this is add act")
			id := add("aaa")
			fmt.Println("lastInsertID:", id)
	}
	os.Exit(0)

	db := getDB()
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

func getDB() (*sql.DB) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/xnw")
	check(err)
	return db
}

func add(name string) int64 {
	db := getDB()
	ret,_ := db.Exec("insert into test('name') value (?)", name)
	id,_ := ret.LastInsertId()
	return id
}
