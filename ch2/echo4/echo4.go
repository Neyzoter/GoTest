package main

import (
	"../pack"
	"flag"
	"fmt"
	"strings"
)

// 命令行参数 -n
// 是否换行
var n = flag.Bool("n", false, "omit trailing newline")
// 命令行参数 -s
var sep = flag.String("s", " ", "separator")
func main() {
	// 更新flag参数
	flag.Parse()
	// 将Args和分隔符合并
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	fmt.Println(pack.Trans("123"))
}
