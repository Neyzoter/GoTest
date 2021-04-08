package main

import (
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//关闭listener
	defer listener.Close()

	//阻塞等待用户连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//关闭连接
	defer conn.Close()
	//接收用户请求
	bytes := make([]byte, 1024)
	num, err := conn.Read(bytes)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("data = ", string(bytes[:num]))
}
