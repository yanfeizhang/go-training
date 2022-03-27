package main

import (
	"os"
	"text/template"
)

type Service struct {
	Namespace string
	App       string
	Func      string
}

func main() {
	//解析模板
	//这里不使用new()，使用默认文件名来命名
	tmpl, err := template.ParseFiles("1-basic/template/hello.tpl")
	if err != nil {
		panic(err)
	}

	//渲染模板
	sweaters := Service{"TestNamespace", "MyApp", "Func1"}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}
