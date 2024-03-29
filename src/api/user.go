package api

import (
	"xnw.com/core"
	"xnw.com/models"
)

type UserApi struct {
	BaseApi
}

var User = new(UserApi)

func (c *UserApi) Detail(request core.Request) *core.Response {
	id := request.GetInt64("id")

	var data map[string]interface{}
	data = make(map[string]interface{})
	data["id"] = id

	user, _ := models.NewUserModel.Get(id)
	data["user"] = user

	response := &core.Response{}
	response.ErrCode = 0
	response.Msg = "查询成功"
	response.Data = data
	return response
}
