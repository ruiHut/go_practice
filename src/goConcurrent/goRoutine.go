package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello") // tips 结果 hello 为5 world 小于 等于5  原因为 主线程结束 所有线程也相应销毁。
}
