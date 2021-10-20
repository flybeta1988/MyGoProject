package core

import "net/http"

type Response struct {
	http.ResponseWriter
	ErrCode int
	Msg string
}
