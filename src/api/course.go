package api

import (
	"xnw.com/core"
)

type CourseApi struct {
	BaseApi
}

var Course = new(CourseApi)

func (c *CourseApi) Index(request core.Request) *core.Response {
	id := request.Get("id")
	var data map[string]string
	data = make(map[string]string)
	data["id"] = id
	response := &core.Response{}
	response.ErrCode = 0
	response.Msg = "查询成功"
	response.Data = data
	return response
}

func (c *CourseApi) Detail(request core.Request) *core.Response {
	id := request.Get("id")
	var data map[string]string
	data = make(map[string]string)
	data["id"] = id
	response := &core.Response{}
	response.ErrCode = 0
	response.Msg = "查询成功"
	response.Data = data
	return response
}
