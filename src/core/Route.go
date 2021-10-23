package core

type Route struct {
	Path           string
	ControllerFunc ControllerFunc
}

func (route *Route) Get() {

}