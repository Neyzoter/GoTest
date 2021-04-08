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

	for ;; {
		//阻塞等待用户连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	//关闭连接
	defer conn.Close()
	//接收用户请求
	bytes := make([]byte, 1024)
	remoteAddr := conn.RemoteAddr()
	// 查询并处理buffer中的消息
	// conn.Read是阻塞的？应该是
	// 如果断开连接，则num会是0  一定会是0吗？
	//              err会是EOF 一定会是EOF吗？
	var num int
	var err error
	for num, err = conn.Read(bytes); num > 0; num, err = conn.Read(bytes) {
		if err != nil {
			fmt.Printf("[From %v] err : %v\n", remoteAddr, err)
			return
		}
		fmt.Printf("[From %v] %d data = %s\n", remoteAddr, num, string(bytes[:num]))
	}
	// 如果客户端断开连接，则会推出上面的for循环
	fmt.Printf("[From %v] Conn Closed\n", remoteAddr)
	fmt.Printf("          num : %d ; err : %v", num, err)
}