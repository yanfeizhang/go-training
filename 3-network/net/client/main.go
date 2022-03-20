package main

import (
	"fmt"
	"net"
)

func main() {
	//与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:9008")
	if err != nil {
		fmt.Printf("conn server failed, err:%v\n", err)
		return
	}
	defer conn.Close()

	//给服务器端发送数据
	_, err = conn.Write([]byte("I am a client\n"))
	if err != nil {
		fmt.Printf("send data ti server failed, err:%v\n", err)
		return
	}

	//接收服务器端返回
	var buf [1024]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Printf("read failed:%v\n", err)
		return
	}
	fmt.Printf("收到服务端回复:%v\n", string(buf[:n]))

}
