package models

import "time"

type Scores struct {
	ScoreId   int       `gorm:"column:score_id;AUTO_INCREMENT;primary_key" json:"score_id"` // 分数记录的ID
	StudentId int       `gorm:"column:student_id;NOT NULL" json:"student_id"`               // 学生的ID
	CourseId  int       `gorm:"column:course_id;NOT NULL" json:"course_id"`                 // 课程的ID
	Score     string    `gorm:"column:score;NOT NULL" json:"score"`                         // 分数
	ScoreDate time.Time `gorm:"column:score_date" json:"score_date"`                        // 分数录入日期
}

func (m *Scores) TableName() string {
	return "t_teachers"
}
