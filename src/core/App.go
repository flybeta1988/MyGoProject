package core

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime/debug"
	"xnw.com/utils"
)

type App struct {
	id string
	db *sql.DB
	Routes []Route
	RoutesMap map[string]ControllerFunc
}

type ControllerFunc func(Request) *Response

func init() {
	fmt.Println("core's init() start ...")
}

func (app *App) Run() {
	app.initRoutes()
	err := http.ListenAndServe(":7070", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func (app *App) InitDB() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	utils.CheckError(err)
	err = db.Ping()
	utils.CheckError(err)
	app.db = db
}

func (app *App) initRoutes() {
	for path, cfunc := range app.RoutesMap {
		http.HandleFunc(path, safeHandlerV2(cfunc))
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

func safeHandlerV2(fn ControllerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := Request{r}
		response := fn(request)
		data, err := json.Marshal(response)
		if err != nil {

		}
		io.WriteString(w, string(data))
	}
}

func (app *App) End() {
	defer app.db.Close()
}
