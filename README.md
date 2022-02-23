# 21Wi-AcademicSystem
2021-2022 学年冬季学期《数据库原理 (1)》教务系统大作业

## 使用方法
### 后端
生成文档
```shell
cd ./backend
swag init
```
编译
```shell
go build
```
运行
```shell
./academic-system
```
> 在windows系统中编译的文件带有 .exe 后缀

查看 swagger 文档：[localhost:8080/swagger/index.html](localhost:8080/swagger/index.html)

####  可能出现的问题

```shell
cannot find type definition: gorm.Model
```
这是由于 swaggo 最新版本不能检测到 GORM 提供的嵌套结构体

解决办法：回退版本

```shell
go get -u github.com/swaggo/swag/cmd/swag@v1.6.7
```
或使用
```shell
swag init --parseDependency --parseInternal --parseDepth=10
```
但是会很慢
## 介绍

### 前端

前端页面采用微信小程序开发

### 后端

后端采用 Golang 编写，引入 GORM 对象关关系映射模型框架负责数据库的 CRUD 工作；引入 Gin WEB 框架来完成 API 的构建。项目可编译成可执行程序，便于迁移。

#### 表设计

表设计 ER 图如下：

![ER](https://markdown-1303167219.cos.ap-shanghai.myqcloud.com/ER.svg)

出于统一性和便捷性考虑，所有表的主键都被设置为各自的 `id` 字段，这样做可以：

- 自动自增
- 便于后端判断记录的唯一性
- 便于前端的数据验证处理

数据库表支持与 json 格式、formData（表单）格式相互转化，便于前后端的数据传输

```go
type Course struct {
	gorm.Model
	Number      string `json:"number" form:"number" gorm:"index" example:"0121"` 
	Name        string `json:"name" form:"name" example:"数据库原理"`                 
	Credit      uint8  `json:"credit" form:"credit" example:"4"`                 
	Department  string `json:"department" form:"department" example:"计算机"`       
	Term        string `json:"term" form:"term" example:"22-冬季学期"`               
	TeacherName string `json:"teacher_name" form:"teacher_name" example:"老师A"`  

	Selections []Selection `json:"selections"`
}
```

#### API

本项目所有的 API 均遵循 RESTful 风格规范：

- [GET] `/course/:id` 获取指定 id 的课程信息
- [POST] `/course` 创建一门新的课程
- [GET] `/student/:id/course&hasScore=true` 查询指定 id 学生所选的所有（有成绩的）课程
- ...

并且使用 swagger ui 提供的自动生成文档服务，构建 API 文档：http://1.15.130.83:8080/swagger/index.html

![image-20220223090203105](https://markdown-1303167219.cos.ap-shanghai.myqcloud.com/image-20220223090203105.png)

接口拥有基本的数据校验功能：

```go
// UpdateSelectionScore 更新课程成绩
func UpdateSelectionScore(id uint, score int) error {
	if score < 1 || score > 100 {
		return errors.New("分数不合法！")
	}

	// ...
}
```

#### SQL 语句例

查询指定 id 学生选择的所有课程信息（包含成绩）

多表链接难以用 ORM 模型实现，故使用原生 SQL 语句完成查询	。

```go
var courseByStuResult []CourseByStuResponse
db.Raw("select distinct c.id, sc.score, c.number, c.name, c.credit, c.department, c.term, 		  c.teacher_name, s.name as student_name from courses as c, selections as sc, students as s where sc.course_id = c.id and sc.student_id = s.id and s.id = ?", studentID).Scan(&courseByStuResult).
```

