package main

import (
	"encoding/json"
	"fmt"
	"xnw.com/core"
	"xnw.com/models"
	"xnw.com/utils"
)

func main() {
	/*defer end()
	fmt.Println("hello golang")
	//addUser()
	user := models.User{}
	users := testGetUserList(user)
	testJson(users)*/
	app := &core.App{}
	app.Routes = append(app.Routes, &core.Route{"/", })
	app.Run()
}

func testJson(users []models.User) {
	//user_json, err := json.Marshal(users)
	user_json, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(user_json))
}

func testGetDetail() {
	user := &models.User{}
	u, err := user.Get(2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(u)
}

func testGetUserList(user models.User) []models.User {
	users, err := user.GetList()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(users)
	return users
}

func addUser()  {
	name := "fly" + utils.Rand()
	user := &models.User{
		Account: name,
		Password: "123456",
		Name: name,
	}
	id, err := user.Add()
	utils.CheckError(err)
	fmt.Println(id)
}

func end() {
	if models.DB != nil {
		err := models.DB.Close()
		if err != nil {
			fmt.Println("db close error:%v", err.Error())
		} else {
			fmt.Println("db closed !")
		}
	}
}
