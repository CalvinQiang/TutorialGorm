package models

type Courses struct {
	CourseId   int    `gorm:"column:course_id;AUTO_INCREMENT;primary_key" json:"course_id"` // 课程的ID
	CourseName string `gorm:"column:course_name;NOT NULL" json:"course_name"`               // 课程的名称
}

func (m *Courses) TableName() string {
	return "t_teachers"
}
