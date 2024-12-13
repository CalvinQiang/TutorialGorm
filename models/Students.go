package models

type Students struct {
	StudentId int    `gorm:"column:student_id;AUTO_INCREMENT;primary_key" json:"student_id"` // 学生的ID
	Name      string `gorm:"column:name;NOT NULL" json:"name"`                               // 学生的姓名
}

func (m *Students) TableName() string {
	return "t_teachers"
}
