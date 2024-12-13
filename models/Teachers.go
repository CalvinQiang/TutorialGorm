package models

type Teachers struct {
	TeacherId   int    `gorm:"column:teacher_id;AUTO_INCREMENT;primary_key" json:"teacher_id"` // 老师的ID
	Name        string `gorm:"column:name;NOT NULL" json:"name"`                               // 老师的姓名
	Email       string `gorm:"column:email" json:"email"`                                      // 邮箱
	PhoneNumber string `gorm:"column:phone_number" json:"phone_number"`                        // 电话号码
	Description string `gorm:"column:description" json:"description"`                          // 详细描述
}

func (m *Teachers) TableName() string {
	return "t_teachers"
}
