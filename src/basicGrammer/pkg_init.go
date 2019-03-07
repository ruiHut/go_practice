package main

import (
	"fmt"
	"runtime"
)

func init() { // 代码包的初始化函数 会在main函数执行前被执行 info函数在变量自定义初始化之后执行  例如m输出的结果
	fmt.Printf("Map: %v\n", m)
	info = "hello， world！"
	fmt.Println(info)
	info = fmt.Sprintf("OS: %s, Arch: %s", runtime.GOOS, runtime.GOARCH) //	Sprintf : 将次格式化输出且输出结果赋值给变量
}

var m = map[int]string{1: "2", 2: "b", 3: "c"} // go 中 定义一个 map集合
//  var m = map[int]string{1: "a", 2: "b"}
var info string // 定义一个未被初始化的string 变量 可由第二句输出看出

func main() {
	fmt.Println(info)
}
