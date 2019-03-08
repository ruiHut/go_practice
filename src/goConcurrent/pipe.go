package main

// 管道 只能运行一次

import (
	"bytes"
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

	// ps aux | grep go

	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "apipe")

	var outputBuf1 bytes.Buffer
	cmd1.Stdout = &outputBuf1
	if err := cmd1.Wait(); err != nil { //	cmd1 的wait方法的调用会一直阻塞， 直到cmd1完全运行结束
		fmt.Println("Error: the first command can not be startup %s\n", err)
		return
	}

	if err := cmd2.Wait(); err != nil {
		fmt.Println("Error: couldn wait for the second command: %s\n", err)
		return
	}
	fmt.Println(outputBuf1.Bytes())

}
