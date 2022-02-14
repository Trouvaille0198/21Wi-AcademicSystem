package model

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Number     string `json:"number" gorm:"unique;index" example:"0198"`
	Name       string `json:"name" example:"王二"`
	Sex        string `json:"sex" example:"男"`
	Age        uint8  `json:"age" example:"21"`
	Department string `json:"department" example:"计算机"` // 所属院系
	Password   string `json:"password" gorm:"default:123" example:"123"`

	Selection []Selection `json:"selection"` // 选课情况
}

// CreateStudentsExample 创建学生样例
func CreateStudentsExample() (students []Student) {
	students = []Student{
		{Number: "0196", Name: "学生A", Sex: "男", Age: 21, Department: "计算机"},
		{Number: "0197", Name: "学生B", Sex: "女", Age: 22, Department: "社会学"},
		{Number: "0198", Name: "学生C", Sex: "男", Age: 20, Department: "医学"},
		{Number: "0199", Name: "学生D", Sex: "女", Age: 21, Department: "计算机"},
		{Number: "0200", Name: "学生E", Sex: "男", Age: 21, Department: "数学"}}

	db.Model(&Student{}).Create(&students)
	return
}

// GetAllStudents 获取所有学生信息
func GetAllStudents() (*[]Student, error) {
	var students []Student
	err := db.Model(&Student{}).Find(&students).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &students, nil
}

// GetStudentByID 获取指定id的学生信息
func GetStudentByID(id int) (*Student, error) {
	var student Student
	err := db.Model(&Student{}).Where("id = ?", id).First(&student).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &student, nil
}

// GetStudentByAttrs 通用方法 根据条件获取所有学生信息
func GetStudentByAttrs(attrs interface{}) (*[]Student, error) {
	var students []Student
	err := db.Model(&Student{}).Where(attrs).Find(&students).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &students, nil
}

// UpdateStudent 通用方法 修改指定id的学生信息
func UpdateStudent(id int, data map[string]interface{}) error {
	err := db.Model(&Student{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateStudent 创建学生实例
func CreateStudent(data map[string]interface{}) (*Student, error) {
	student := Student{
		Number:     data["number"].(string),
		Name:       data["name"].(string),
		Sex:        data["sex"].(string),
		Age:        data["age"].(uint8),
		Department: data["department"].(string),
		Password:   data["password"].(string),
	}
	err := db.Create(&student).Error
	if err != nil {
		return &Student{}, err
	}
	return &student, nil
}

// DeleteStudent 删除指定id学生
func DeleteStudent(id int) error {
	err := db.Model(&Student{}).Where("id = ?", id).Delete(&Student{}).Error
	if err != nil {
		return err
	}
	return nil
}
