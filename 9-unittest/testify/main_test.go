package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//Add 函数的单元测试
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"case1", 1, 2, 3},
		{"case2", 2, 2, 4},
		{"case3", 2, 3, 5},
	}

	for _, tt := range tests {
		assert.Equal(t, Add(tt.a, tt.b), tt.want, tt.name+"they should be equal")
	}

	assert.NotEqual(t, Add(2, 3), 456, "they should not be equal")
}

//作者
type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//testify的Json assert使用
func TestJson(t *testing.T) {
	a := &Author{
		ID:   1,
		Name: "Yanfei Zhang",
	}

	j, _ := json.Marshal(a)
	assert.JSONEq(
		t,
		`{"id":1, "name":"Yanfei Zhang"}`,
		string(j),
	)
}

type MyTestSuit struct {
	suite.Suite
	testCount uint32
}

func (s *MyTestSuit) SetupSuite() {
	fmt.Println("SetupSuite")
}

func (s *MyTestSuit) TearDownSuite() {
	fmt.Println("TearDownSuite")
}

func (s *MyTestSuit) SetupTest() {
	fmt.Printf("SetupTest test count:%d\n", s.testCount)
}

func (s *MyTestSuit) TearDownTest() {
	s.testCount++
	fmt.Printf("TearDownTest test count:%d\n", s.testCount)
}

func (s *MyTestSuit) BeforeTest(suiteName, testName string) {
	fmt.Printf("BeforeTest suite:%s test:%s\n", suiteName, testName)
}

func (s *MyTestSuit) AfterTest(suiteName, testName string) {
	fmt.Printf("AfterTest suite:%s test:%s\n", suiteName, testName)
}

func (s *MyTestSuit) TestExample() {
	fmt.Println("TestExample")
}

func TestExample(t *testing.T) {
	suite.Run(t, new(MyTestSuit))
	suite.Run(t, new(MyTestSuit))
}
