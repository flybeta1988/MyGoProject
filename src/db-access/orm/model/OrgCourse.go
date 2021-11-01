package model

type OrgCourse struct {
	Base
	Name string
	Uid uint
	TypeId uint
	Status uint `gorm:"default:1"`
}

func (OrgCourse) TableName() string {
	return "org_course"
}
