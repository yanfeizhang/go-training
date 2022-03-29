package main

//写单元测试用的 compute 函数
func Compute(a int, b int) int {
	return a*GetFromSql() + b
}

//写单元测试用的 GetFromSql 函数
func GetFromSql() int {
	return 5
}

//Student类，用来承载单测方法
type Student struct {
	score int
}

//写单元测试用的 Compute 方法
func (s *Student) Compute(a int, b int) int {
	return a*GetFromSql() + b
}

//写单元测试用的 GetFromSql 方法
func (s *Student) GetFromSql(n int) int {
	return n - n + 5
}

//全局变量
var Version int
