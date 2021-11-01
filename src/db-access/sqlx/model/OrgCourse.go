package model

import "database/sql"

type OrgCourse struct {
	Id uint `db:"id"`
	Name string `db:"name"`
	Uid uint `db:"uid"`
	TypeId uint `db:"type_id"`
	Status uint `db:"status"`
	Ctime uint `db:"ctime"`
	Utime uint `db:"utime"`
	Dtime sql.NullString `db:"dtime"`
}
