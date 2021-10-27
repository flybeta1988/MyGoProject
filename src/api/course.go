package api

import (
	"xnw.com/core"
)

var Course = new(CourseApi)

type CourseApi struct {
	BaseApi
}

func (c *CourseApi) Detail(request core.Request) *core.Response {
	id := request.GetInt64("id")

	var data map[string]interface{}
	data = make(map[string]interface{})
	data["id"] = id

	response := &core.Response{}
	response.ErrCode = 0
	response.Msg = "查询成功"
	response.Data = data

	return response
}
