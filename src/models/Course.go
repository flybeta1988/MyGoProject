package models

import (
	"database/sql"
	"fmt"
)

var NewCourse = new(Course)
var CModel = &CourseModel{table: "org_course"}

type Course struct {
	Id int64 `json:"id"`
	Account string `json:"account"`
	Password string `json:"password"`
	Name string `json:"name"`
	Status int `json:"status"`
	CreateAt sql.NullString `json:"create_at"`
	UpdateAt sql.NullString `json:"update_at"`
}

type CourseModel struct {
	Model
	table string
}

func (u *Course) Get(id int64) (User, error)  {
	var user User
	row := DB.QueryRow("SELECT * FROM `user` WHERE id = ?", id)
	if err:= row.Scan(&user.Id, &user.Account, &user.Password, &user.Name, &user.Status, &user.CreateAt, &user.UpdateAt); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("GetById: %d: no such row", id)
		}
		return user, fmt.Errorf("GetById: %d %v", id, err)
	}
	return user, nil
}

func (user *Course) GetList() ([]User, error) {

	rows, err := DB.Query("SELECT * FROM `user` ORDER BY `id` DESC")
	if err != nil {
		return nil, fmt.Errorf("getList err:%v", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err:= rows.Scan(&user.Id, &user.Account, &user.Password, &user.Name, &user.Status, &user.CreateAt, &user.UpdateAt); err != nil {
			return nil, fmt.Errorf("getList err:%v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getList err:%v", err)
	}
	//fmt.Printf("%T\n", users)
	return users, nil
}
