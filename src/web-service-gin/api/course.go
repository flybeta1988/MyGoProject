package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-service-gin/model"
	"xnw.com/core"
)

func GetCourses(c *gin.Context) {

	response := core.Response{}
	response.Msg = "查询成功"
	response.ErrCode = 0

	keyword := c.Param("keyword")

	courses := model.GetAllCourseList(keyword)

	var data map[string]interface{}
	data = make(map[string]interface{}, 100)
	data["total"] = len(courses)
	data["courses"] = courses
	response.Data = data

	c.IndentedJSON(http.StatusOK, response)
}
