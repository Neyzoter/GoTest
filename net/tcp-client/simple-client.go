package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const hostname = "aliyun"

// 一个简单的tcp客户端
// 可以实现将输入的数据发送到远程服务器
// 在断网情况下也可以保存数据，在重新连接后重新发送数据
func main() {
	// 连接服务器
	conn,err := net.Dial("tcp",hostname + ":8000")
	if err != nil {
		fmt.Println("[ERR] Connect to TCP server failed ,err:",err)
	}
	if err != nil {
		fmt.Println("[ERR] Set Write Deadline err : ", err)
	}
	// 读取命令行输入
	inputReader := bufio.NewReader(os.Stdin)
	buffer := []byte{}
	// step1. 阻塞读取数据
	// step2. 将数据保存在slice
	// step3. 检查conn是否为nil（如果断网，则conn可能为nil）
	// step3.1 conn不是nil，发送数据
	//         检查发送数据结果，如果没有成功发送，则重新连接服务器
	// step3.2 conn是nil，直接重新连接服务器
	for ;; {
		input,err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("[ERR] Read from console failed,err:",err)
			continue
		}

		// 读取到字符"Q"退出
		str := strings.TrimSpace(input)
		if str == "Q"{
			break
		}
		buffer = append(buffer, []byte(input)...)
		// 响应服务端信息
		if conn != nil {
			_,err = conn.Write(buffer)
			if err != nil{
				fmt.Println("[ERR] Write failed, err : ",err)
				fmt.Println("[INFO] Data Saved")
				// 连接服务器
				conn,err = net.Dial("tcp",hostname + ":8000")
				if err != nil {
					fmt.Println("[ERR] Connect to TCP server failed ,err:",err)
				}
			} else { // 清空
				buffer = []byte{}
			}
		} else {
			fmt.Println("conn is nil")
			// 连接服务器
			conn,err = net.Dial("tcp",hostname + ":8000")
			if err != nil {
				fmt.Println("[ERR] Connect to TCP server failed ,err:",err)
			}
		}
	}
}
