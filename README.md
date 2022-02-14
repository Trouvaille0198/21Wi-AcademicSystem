# 21Wi-AcademicSystem
2021-2022 学年冬季学期《数据库原理 (1)》教务系统大作业

查看swagger: localhost:8080/swagger/index.html
生成文档
```shell
swag init
```
运行
```shell
go run ./main.go
```

## swag 的版本问题
```shell
cannot find type definition: gorm.Model
```
解决办法：
```shell
go get -u github.com/swaggo/swag/cmd/swag@v1.6.7
```
或使用
```shell
swag init --parseDependency --parseInternal --parseDepth=10
```
但是很慢