package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Number      string `json:"number" gorm:"index"`
	Name        string `json:"name"`
	Credit      uint8  `json:"credit"`
	Department  string `json:"department" gorm:"comment:所属院系"`
	TeacherName string `json:"teacher_name"`

	Selection Selection `json:"selection"`
}
