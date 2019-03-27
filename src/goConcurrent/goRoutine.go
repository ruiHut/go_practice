package main

import (
	"fmt"
	"runtime"
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
	go passForSleep()
	passForGoshched()
	hello()
	say("hello") // tips 结果 hello 为5 world 小于 等于5  原因为 主线程结束 所有线程也相应销毁。
}

// goroutine 暂停
// 1. sleep  作用：将调用该goroutine 进入Gwaiting状态
func passForSleep() {
	time.Sleep(time.Millisecond)
	fmt.Println("hello world")
}

// 2. runtime.Gosched 作用 ： 暂停当前的G 使其他G有机会执行
func passForGoshched() {
	runtime.Gosched()
	fmt.Println("hello world")
}

func hello() {
	names := []string{"good", "boy", "jim", "home"}
	for _, name := range names {
		go func() {
			fmt.Println(name)
		}()
	}

	for _, name := range names {
		go func(who string) {
			fmt.Println(who)
		}(name)
	}
}
