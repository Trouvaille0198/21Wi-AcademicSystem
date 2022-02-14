package model

import (
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	Number     string `json:"number" gorm:"unique;index"`
	Name       string `json:"name"`
	Department string `json:"department" gorm:"comment:所属院系"`
	Password   string `json:"password" gorm:"default:123"`

	Courses []Selection `json:"courses"`
}

// CreateTeachersExample 创建老师样例
func CreateTeachersExample() (teachers []Teacher) {
	teachers = []Teacher{
		{Number: "0196", Name: "老师A", Department: "计算机"},
		{Number: "0197", Name: "老师B", Department: "社会学"},
		{Number: "0198", Name: "老师C", Department: "医学"},
		{Number: "0199", Name: "老师D", Department: "通信"},
		{Number: "0200", Name: "老师E", Department: "计算机"}}

	db.Model(&Teacher{}).Create(&teachers)
	return
}

// GetAllTeachers 获取所有老师信息
func GetAllTeachers() (*[]Teacher, error) {
	var teachers []Teacher
	err := db.Model(&Teacher{}).Find(&teachers).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &teachers, nil
}

// GetTeacherByID 获取指定id的老师信息
func GetTeacherByID(id int) (*Teacher, error) {
	var teacher Teacher
	err := db.Model(&Teacher{}).Where("id = ?", id).First(&teacher).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &teacher, nil
}

// GetTeacherByAttrs 通用方法 根据条件获取所有老师
//信息
func GetTeacherByAttrs(attrs interface{}) (*[]Teacher, error) {
	var teachers []Teacher
	err := db.Model(&Teacher{}).Where(attrs).Find(&teachers).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &teachers, nil
}

// UpdateTeacher 通用方法 修改指定id的老师信息
func UpdateTeacher(id int, data map[string]interface{}) error {
	err := db.Model(&Teacher{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateTeacher 创建老师实例
func CreateTeacher(data map[string]interface{}) (*Teacher, error) {
	teacher := Teacher{
		Number:     data["number"].(string),
		Name:       data["name"].(string),
		Department: data["department"].(string),
		Password:   data["password"].(string),
	}
	err := db.Create(&teacher).Error
	if err != nil {
		return &Teacher{}, err
	}
	return &teacher, nil
}

// DeleteTeacher 删除指定id老师
func DeleteTeacher(id int) error {
	err := db.Model(&Teacher{}).Where("id = ?", id).Delete(&Teacher{}).Error
	if err != nil {
		return err
	}
	return nil
}
