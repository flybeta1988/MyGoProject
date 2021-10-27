package routes

import (
	"fmt"
	"xnw.com/api"
	"xnw.com/core"
)

func init() {
	fmt.Println("routes v2 init...")
}

func Config(app *core.App) {
	routes := make(map[string]core.ControllerFunc)
	appendRoutes(routes)
	app.RoutesMap = routes
}

func appendRoutes(routes map[string]core.ControllerFunc) {
	routes["/v2/user/detail"] = api.User.Detail
	routes["/v2/course/detail"] = api.Course.Detail
}