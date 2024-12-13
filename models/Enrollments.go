package models

import "time"

type Enrollments struct {
	EnrollmentId   int       `gorm:"column:enrollment_id;AUTO_INCREMENT;primary_key" json:"enrollment_id"` // 选课记录的ID
	StudentId      int       `gorm:"column:student_id;NOT NULL" json:"student_id"`                         // 学生的ID
	CourseId       int       `gorm:"column:course_id;NOT NULL" json:"course_id"`                           // 课程的ID
	TeacherId      int       `gorm:"column:teacher_id;NOT NULL" json:"teacher_id"`                         // 老师的ID
	EnrollmentDate time.Time `gorm:"column:enrollment_date" json:"enrollment_date"`                        // 选课日期
}

func (m *Enrollments) TableName() string {
	return "t_teachers"
}
