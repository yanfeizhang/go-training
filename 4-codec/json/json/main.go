package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Book struct {
	Name  string `json:"name"`
	Pages int    `json:"pages"`
}

func main() {
	// 1. 使用 json.Marshal 编码
	book1 := Book{"深入理解Linux网络", 329}
	bytes1, err := json.Marshal(&book1)
	if err == nil {
		fmt.Println("json.Marshal 编码结果: ", string(bytes1))
	}

	// 2. 使用 json.Unmarshal 解码
	str := `{"name":"深入理解Linux网络","Pages":329}`
	bytes2 := []byte(str)
	var book2 Book // 用来接收解码后的结果
	if json.Unmarshal(bytes2, &book2) == nil {
		fmt.Println("json.Unmarshal 解码结果: ", book2)
	}

	// 3. 使用 json.NewEncoder 编码
	book3 := Book{"深入理解Linux网络", 330}
	bytes3 := new(bytes.Buffer)
	_ = json.NewEncoder(bytes3).Encode(book3)
	if err == nil {
		fmt.Print("json.NewEncoder 编码结果: ", string(bytes3.Bytes()))
	}

	// 4. 使用 json.NewDecoder 解码
	str4 := `{"name":"深入理解Linux网络","Pages":331}`
	var book4 Book
	err = json.NewDecoder(strings.NewReader(str4)).Decode(&book4)
	if err == nil {
		fmt.Println("json.NewDecoder 解码结果: ", book4)
	}
}
