package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//关闭连接
	defer conn.Close()

	//读取连接上的数据
	var buf [1024]byte
	len, err := conn.Read(buf[:])
	if len == 0 || err != nil {
		fmt.Println("accept failed, err:%v\n", err)
	}
	fmt.Println("accept data:%v\n", string(buf[0:len]))

	//发送数据
	_, err = conn.Write([]byte("I am server!"))
	if err != nil {
		fmt.Println("send data failed, err:%v\n", err)
	}
	fmt.Println("send data success: I am server!")

}

func main() {
	//构造一个listener
	listener, _ := net.Listen("tcp", "127.0.0.1:9008")
	defer listener.Close()

	for {
		//接收请求
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err:%v\n", err)
			continue
		}

		go process(conn)
	}
}
