package model

// CourseByStuResult 学生视角下的课程信息
type CourseByStuResult struct {
	ID          uint   `json:"id"`
	Score       int    `json:"score"`        // 成绩
	Number      string `json:"number"`       // 课号
	Name        string `json:"name"`         // 课名
	Credit      uint8  `json:"credit"`       // 学分
	Department  string `json:"department"`   // 所属院系
	Term        string `json:"term"`         // 学期
	TeacherName string `json:"teacher_name"` // 老师名
	StudentName string `json:"student_name"` // 学生名
}
