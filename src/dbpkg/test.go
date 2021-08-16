package dbpkg

var Test struct{
	Id int64 `db:"id"`
	Name string `db:"name"`
}

func Insert(name string) int64 {
	ret,_ := DB.Exec("insert into test('name') value (?)", name)
	id,_ := ret.LastInsertId()
	return id
}