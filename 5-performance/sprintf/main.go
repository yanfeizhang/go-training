package main

import (
	"fmt"
	"strconv"
)

func sprintfConvert(roomid int64) string {
	for i := 0; i < 1000; i++ {
		fmt.Sprintf("%d", roomid)
	}
	return fmt.Sprintf("%d", roomid)
}

func strconvConvert(roomid int64) string {
	for i := 0; i < 1000; i++ {
		strconv.FormatInt(roomid, 10)
	}
	return strconv.FormatInt(roomid, 10)
}

func main() {
	fmt.Println()
}
