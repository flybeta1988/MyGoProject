package models

import (
	"database/sql"
	"fmt"
	"reflect"
)

type Model struct {
	Table string
	Fields interface{}
}

func (m *Model) Foo() {
	fmt.Println("Model Foo()")
}

func resetStructFieldsV2(user *User) {
	var fields []interface{}
	v := reflect.ValueOf(user)
	s := v.Elem()
	for i := 0; i < s.NumField(); i ++ {
		//fmt.Println(t.Field(i).Name, v.Field(i).Interface())
		if !s.Field(i).CanSet() {
			continue
		}
		fields = append(fields, v.Field(i))
	}
}

func (m *Model) Get(id int64) ([]interface{}, error)  {

	row := DB.QueryRow("SELECT * FROM `?` WHERE id = ?", m.Table, id)

	fmt.Println(m.Fields)

	var item []interface{}
	fields := &m.Fields

	v := reflect.ValueOf(fields)
	s := v.Elem()
	for i := 0; i < s.NumField(); i ++ {
		if !s.Field(i).CanSet() {
			continue
		}
		item = append(item, v.Field(i))
	}

	if err:= row.Scan(&fields); err != nil {
		//if err:= row.Scan(&user); err != nil {
		if err == sql.ErrNoRows {
			return item, fmt.Errorf("GetById: %d: no such row", id)
		}
		return item, fmt.Errorf("GetById: %d %v", id, err)
	}
	return item, nil
}
