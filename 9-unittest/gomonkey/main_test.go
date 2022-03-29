package main

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"

	"github.com/stretchr/testify/assert"
)

//GoMonkey 之 mock 函数测试
func TestMockFunc(t *testing.T) {
	//测试用例
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"case1", 1, 2, 7},
		{"case2", 2, 2, 12},
		{"case3", 2, 3, 13},
	}

	//函数 Mock
	patches := gomonkey.ApplyFunc(GetFromSql, func() int {
		return 5
	})
	defer patches.Reset()

	for _, tt := range tests {
		assert.Equal(t, Compute(tt.a, tt.b), tt.want, tt.name+"they should be equal")
	}
}

//GoMonkey 之 mock 方法测试
func TestMockMethod(t *testing.T) {
	//测试用例
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"case1", 1, 2, 7},
		{"case2", 2, 2, 12},
		{"case3", 2, 3, 13},
	}

	//方法 Mock
	var s *Student
	patches := gomonkey.ApplyMethod(reflect.TypeOf(s), "GetFromSql", func(s *Student, n int) int {
		return 5
	})
	defer patches.Reset()

	s = &Student{}
	for _, tt := range tests {
		assert.Equal(t, Compute(tt.a, tt.b), tt.want, tt.name+"they should be equal")
	}
}

//GoMonkey 之 mock 全局变量
func TestMockGlobalVar(t *testing.T) {
	patches := gomonkey.ApplyGlobalVar(&Version, 12)
	defer patches.Reset()
	assert.Equal(t, Version, 12, "they should be equal")
}

//GoMonkey 之 mock 函数序列测试
//记住要避免函数内联
//go test -v -gcflags=all=-l
func TestMockFuncSeries(t *testing.T) {
	//模拟函数前三次调用输出
	outputs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{1}}, // 模拟函数的第1次输出
		{Values: gomonkey.Params{2}}, // 模拟函数的第2次输出
		{Values: gomonkey.Params{3}}, // 模拟函数的第3次输出
	}

	patches := gomonkey.ApplyFuncSeq(GetFromSql, outputs)
	defer patches.Reset()

	assert.Equal(t, 3, Compute(1, 2), "they should be equal")
	assert.Equal(t, 5, Compute(1, 3), "they should be equal")
	assert.Equal(t, 5, Compute(1, 2), "they should be equal")
}
