package core

import "net/http"

type Request struct {
	http.ResponseWriter
}

func (r *Request) Get()  {

}
