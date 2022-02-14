package model

import (
	"errors"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Number      string `json:"number" gorm:"index" example:"0121"` // 课号
	Name        string `json:"name" example:"数据库原理"`               // 课名
	Credit      uint8  `json:"credit" example:"4"`                 // 学分
	Department  string `json:"department" example:"计算机"`           // 所属院系
	Term        string `json:"term" example:"22-冬季学期"`             // 学期
	TeacherName string `json:"teacher_name" example:"老师A"`         // 教师姓名

	Selections []Selection `json:"selections"`
}

// CreateCoursesExample 创建课程样例
func CreateCoursesExample() (courses []Course) {
	courses = []Course{
		{Number: "0121", Name: "数据库原理", Credit: 4, Department: "计算机", Term: "22-春季学期", TeacherName: "老师A"},
		{Number: "0830", Name: "数据结构", Credit: 5, Department: "计算机", Term: "21-冬季学期", TeacherName: "老师B"},
		{Number: "0451", Name: "大学语文", Credit: 2, Department: "计算机", Term: "21-秋季学期", TeacherName: "老师C"},
		{Number: "0279", Name: "算法设计与分析", Credit: 2, Department: "计算机", Term: "22-春季学期", TeacherName: "老师D"},
		{Number: "0022", Name: "操作系统", Credit: 4, Department: "计算机", Term: "22-春季学期", TeacherName: "老师E"}}

	db.Model(&Course{}).Create(&courses)
	return
}

// GetAllCourses 获取所有课程信息
func GetAllCourses() ([]*Course, error) {
	var courses []*Course
	err := db.Model(&Course{}).Find(&courses).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return courses, nil
}

// GetCourseByID 获取指定id的课程信息
func GetCourseByID(id int) (*Course, error) {
	var course Course
	err := db.Model(&Course{}).Where("id = ?", id).First(&course).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &course, nil
}

// GetCourseByAttrs 通用方法 根据条件获取所有课程信息
func GetCourseByAttrs(attrs interface{}) (*[]Course, error) {
	var courses []Course
	err := db.Model(&Course{}).Where(attrs).Find(&courses).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &courses, nil
}

// GetCoursesByStudent 获取指定学生的所有课程
func GetCoursesByStudent(studentID int) (*[]CourseByStuResult, error) {
	var courseByStuResult []CourseByStuResult
	err := db.Raw("select distinct c.id, sc.score, c.number, c.name, c.credit, c.department, c.term, "+
		"c.teacher_name, s.name as student_name "+
		"from courses as c, selections as sc, students as s "+
		"where sc.course_id = c.id and sc.student_id = s.id "+
		"and s.id = ?", studentID).Scan(&courseByStuResult).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &courseByStuResult, nil
}

// GetCoursesByStudentWithScore 获取指定学生的所有有成绩的课程
func GetCoursesByStudentWithScore(studentID int) (*[]CourseByStuResult, error) {
	var courseByStuResult []CourseByStuResult
	err := db.Raw("select distinct c.id, sc.score, c.number, c.name, c.credit, c.department, c.term, "+
		"c.teacher_name, s.name as student_name "+
		"from courses as c, selections as sc, students as s "+
		"where sc.course_id = c.id and sc.student_id = s.id "+
		"and s.id = ? and sc.score <> -1", studentID).Scan(&courseByStuResult).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &courseByStuResult, nil
}

// GetCoursesByStudentWithoutScore 获取指定学生的所有无成绩的课程
func GetCoursesByStudentWithoutScore(studentID int) (*[]CourseByStuResult, error) {
	var courseByStuResult []CourseByStuResult
	err := db.Raw("select distinct c.id, sc.score, c.number, c.name, c.credit, c.department, c.term, "+
		"c.teacher_name, s.name as student_name "+
		"from courses as c, selections as sc, students as s "+
		"where sc.course_id = c.id and sc.student_id = s.id "+
		"and s.id = ? and sc.score = -1", studentID).Scan(&courseByStuResult).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &courseByStuResult, nil
}

// UpdateCourse 通用方法 修改指定id的课程信息
func UpdateCourse(id int, data map[string]interface{}) (err error) {
	err = db.Model(&Course{}).Where("id = ?", id).Updates(data).Error
	return
}

// CreateCourse 创建课程实例
func CreateCourse(data map[string]interface{}) (*Course, error) {
	course := Course{
		Number:      data["number"].(string),
		Name:        data["name"].(string),
		Credit:      data["credit"].(uint8),
		Department:  data["department"].(string),
		Term:        data["term"].(string),
		TeacherName: data["teacher_name"].(string),
	}
	err := db.Create(&course).Error
	if err != nil {
		return &Course{}, err
	}
	return &course, nil
}

// DeleteCourse 删除指定id课程
func DeleteCourse(id int) (err error) {
	err = db.Model(&Course{}).Where("id = ?", id).Delete(&Course{}).Error
	return
}
