package core

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"xnw.com/utils"
)

type App struct {
	id string
	db *sql.DB
	Routes []Route
}

func init() {
	fmt.Println("core's init() start ...")
}

func (app *App) InitDB() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:6306)/xnw")
	utils.CheckError(err)
	err = db.Ping()
	utils.CheckError(err)
	app.db = db
}

func (app *App) initRoutes() {
	for _, route := range app.Routes {
		http.HandleFunc(route.path, safeHandler(route.httpHandlerFunc))
	}
}

func (app *App) Run() {
	app.initRoutes()
	err := http.ListenAndServe(":7070", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	//@todo 未完全理解
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				// 或者输出自定义的50x错误页面
				//w.WriteHeader(http.StatusInternalServerError)
				//renderHtml(w, "error", err)
				log.Println("WARN: panic in %v. -%v", fn, err)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

func (app *App) End() {
	defer app.db.Close()
}

func Test() {
	fmt.Println("App's Test func!")
}
