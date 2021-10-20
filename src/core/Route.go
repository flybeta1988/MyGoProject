package core

import "net/http"

type Route struct {
	path string
	httpHandlerFunc http.HandlerFunc
}

func (route *Route) Get() {

}