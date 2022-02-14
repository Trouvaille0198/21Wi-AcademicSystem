package model

import "gorm.io/gorm"

// Selection 记录学生选课情况
// 三种情况: 未选 已选 已修
type Selection struct {
	gorm.Model
	StudentID uint `gorm:"student_id"`
	CourseID  uint `gorm:"course_id"`
	Score     int  `json:"score" gorm:"default:-1" example:"75"` // 分数, -1表示未评分
}

// CreateSelectionsExample 创建课程关联实例
func CreateSelectionsExample() (selections []Selection) {
	selections = []Selection{
		{StudentID: 1, CourseID: 2, Score: 96},
		{StudentID: 1, CourseID: 3, Score: -1},
		{StudentID: 1, CourseID: 4, Score: 75},
		{StudentID: 2, CourseID: 1, Score: 84},
		{StudentID: 2, CourseID: 2, Score: 68},
		{StudentID: 2, CourseID: 3, Score: -1},
		{StudentID: 2, CourseID: 5, Score: -1},
		{StudentID: 3, CourseID: 2, Score: 75},
		{StudentID: 5, CourseID: 5, Score: -1},
	}

	db.Model(&Selection{}).Create(&selections)
	return
}

// CreateSelection 创建选课记录
func CreateSelection(studentID, CourseID uint, Score int) (err error) {
	selection := Selection{
		StudentID: studentID,
		CourseID:  CourseID,
		Score:     Score,
	}

	err = db.Model(&Selection{}).Create(&selection).Error
	return
}

// UpdateSelection 更新选课记录
func UpdateSelection(id int, data map[string]interface{}) (err error) {
	err = db.Model(&Selection{}).Where("id = ?", id).Updates(data).Error
	return
}

// DeleteSelection 删除指定id选课记录
func DeleteSelection(id int) (err error) {
	err = db.Model(&Selection{}).Where("id = ?", id).Delete(&Selection{}).Error
	return
}

// GetSelectionsByStudentID 获取指定学生的选课记录
func GetSelectionsByStudentID(studentID int) ([]*Selection, error) {
	var selections []*Selection
	err := db.Model(&Selection{}).Where("student_id = ?", studentID).Find(&selections).Error
	if err != nil {
		return []*Selection{}, err
	}
	return selections, nil
}
