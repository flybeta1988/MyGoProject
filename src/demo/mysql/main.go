package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {

	err := initDB()
	check(err)
	defer db.Close()

	fmt.Println(os.Args)

	args := os.Args
	actStr := args[1]
	actArr := strings.Split(actStr, "=")
	act := actArr[1]
	fmt.Printf("act dataType is: %T\n", act)
	switch {
	case act == "list":
		fmt.Println("this is list act")
		list()
	case act == "add":
		fmt.Println("this is add act")
		id := add("aaa")
		fmt.Println("lastInsertID:", id)
	case act == "detail":
		fmt.Println("this is detail act")
		detail(3)
	}
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func initDB() (err error) {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:6306)/xnw")
	check(err)
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

/*func getDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:6306)/xnw")
	check(err)
	return db
}*/

func add(name string) int64 {
	ret, err := db.Exec("insert into test(`name`) value (?)", name)
	if err != nil {
		panic("mysql insert error:" + err.Error())
	}
	id, _ := ret.LastInsertId()
	return id
}

func list() {

	rows, err := db.Query("SELECT id, name FROM test")
	check(err)
	defer rows.Close()

	for rows.Next() {
		var (
			id   int
			name string
		)
		err = rows.Scan(&id, &name)
		check(err)
		fmt.Printf("rows id:%d name:%s\n", id, name)
	}
	err = rows.Err()
	check(err)
}

func detail(id int) {
	stmtOut, err := db.Prepare("SELECT name FROM test WHERE id = ?")
	check(err)
	defer stmtOut.Close()

	var name string
	err = stmtOut.QueryRow(id).Scan(&name)
	check(err)
	fmt.Printf("id:%d name:%s\n", id, name)
}
