package dbpkg

var Test struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

func Insert(name string) int64 {
	ret, err := DB.Exec("insert into test(`name`) value (?)", name)
	if err != nil {
		panic("insert error:" + err.Error())
	}
	id, _ := ret.LastInsertId()
	return id
}
