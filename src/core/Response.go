package core

import "net/http"

type Response struct {
	_ http.ResponseWriter
	ErrCode int	`json:"errcode"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}
