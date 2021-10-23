package core

import (
	"log"
	"net/http"
	"strconv"
)

type Request struct {
	httpRequest *http.Request
}

func (r *Request) Get(key string) string {
	return r.httpRequest.FormValue(key)
}

func (r *Request) GetInt(key string) int {
	id, err := strconv.Atoi(r.Get(key))
	if err != nil {
		log.Printf("core.GetInt key:%s errMsg:%s", key, err.Error())
	}
	return id
}

func (r *Request) GetInt64(key string) int64 {
	id, err := strconv.ParseInt(r.Get(key), 10, 64)
	if err != nil {
		log.Printf("core.GetInt key:%s errMsg:%s", key, err.Error())
	}
	return id
}
