package models

import "gorm.io/gorm"

// Selection 记录学生选课情况
// 三种情况: 未选 已选 已修
type Selection struct {
	gorm.Model
	StudentID int   `gorm:"student_id"`
	CourseID  int   `gorm:"course_id"`
	Score     int64 `json:"score" gorm:"default:-1;comment:分数, -1表示未评分"`
}
