package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Number     string `json:"number" gorm:"unique;index"`
	Name       string `json:"name"`
	Sex        string `json:"sex"`
	Age        uint8  `json:"age"`
	Department string `json:"department" gorm:"comment:所属院系"`
	Password   string `json:"password" gorm:"default:123"`

	Selection Selection `json:"selection"`
}

// CreateStudentsExample 创建学生样例
func CreateStudentsExample() (students []Student) {
	students = []Student{
		{Number: "0196", Name: "学生A", Sex: "男", Age: 21, Department: "计算机"},
		{Number: "0197", Name: "学生B", Sex: "女", Age: 21, Department: "社会学"},
		{Number: "0198", Name: "学生C", Sex: "男", Age: 21, Department: "医学"},
		{Number: "0199", Name: "学生D", Sex: "女", Age: 21, Department: "计算机"},
		{Number: "0200", Name: "学生E", Sex: "男", Age: 21, Department: "数学"}}

	db.Model(&Student{}).Create(&students)
	return
}
