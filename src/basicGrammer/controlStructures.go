package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 	go语言流程控制特点如下
// 		1. 没有while和do-while 但是for更加灵活
//		2. switch语句灵活多变， 还可以用于类型判断
//		3. if语句和switch语句都包含一条初始化子语句
//		4. defer语句可以使我们更加方便地执行异常捕获和资源回收任务
//		5. select 语句也用于多分支选择，但只与通道配合使用
//		6. go语言用于异步启用goroutime并执行指定函数

// 简单实例 "聊天机器人"
func main() {
	// 准备从标准输入读取数据
	inputReader := bufio.NewReader(os.Stdin)
	//fmt.Println(inputReader)
	fmt.Println("please input you name")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("an error occurred : %s\n", err)
		os.Exit(1) // 系统退出
	} else {
		//
		name := input[:len(input)-1] // 等价与从o开始切
		fmt.Printf("hello ,%s ! what can i do for you?\n ", name)
	}
	for { // for {} 类似与 while(true)
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("an error occurred : %s", err)
			continue
		}
		input = input[0 : len(input)-1]
		input = strings.ToLower(input) // 全部转为小写
		switch input {
		case "":
			continue
		case "nothing", "bey":
			fmt.Printf("bey!")
			os.Exit(1)
		default:
			fmt.Printf("Sorry, i dont know!\n")
		}
	}
}
