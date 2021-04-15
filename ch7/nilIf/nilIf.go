package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false
// 当不给buf赋值时，buf作为一个io.Writer接口
// io.Writer的type为*bytes.Buffer
// io.Writer的value为nil
// io.Writer不是nil但是无法通过Write写入数据
func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
	if debug {
		fmt.Println(buf)
	}
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n")) // ERROR
	} else {
		fmt.Printf("Out Nil")
	}
}