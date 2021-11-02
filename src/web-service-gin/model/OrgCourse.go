package model

type OrgCourse struct {
	Base
	Name string `json:"name"`
	Uid uint `json:"uid"`
	TypeId uint `json:"type_id"`
	Status uint `gorm:"default:1"`
}

func (OrgCourse) TableName() string {
	return "org_course"
}

func GetAllCourseList(keyword string) []OrgCourse {
	courses := []OrgCourse{}
	if "" != keyword {
		db.Where("name LIKE ?", keyword + "%").Find(&courses)
	} else {
		db.Find(&courses)
	}
	return courses
}