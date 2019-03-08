package main

// 管道 只能运行一次

import (
	"fmt"
	"os/exec"
)

func main() {
	cmdo := exec.Command("echo", "-n", "my first command come from golong,")
	//if err:= cmdo.Start(); err != nil {
	//	fmt.Println("the command cant startup ", err)
	//	return
	//}

	// 创建一个能获取此命令都输出管道
	stdouto, err := cmdo.StdoutPipe() // StdoutPipe（） 方法表示将管道的输出值赋值给 变量stdouto 类型是 io.readcloser
	if err != nil {
		fmt.Println("cant start is ")
		return
	}

	// 获取输出
	outputO := make([]byte, 30)

	if n, err := stdouto.Read(outputO); err == nil {
		fmt.Println("result is  %s :", outputO[:n])
	} else {
		fmt.Println("cant not open it")
	}
}
