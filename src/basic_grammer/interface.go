package main

import (
	"fmt"
	"strings"
)

// go 的接口类型用于定义一组行为，其中每个行为都由一个方法声明表示
// 接口类型中的方法声明只有方法签名而没有方法体 都一样
// 例子： 定义 "聊天" 相关的一组行为

type Talk interface {
	// 每个方法声明独占一行
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}

// 接口的实现 只要一个类型中实现了该接口的所有方法 则称为该接口的实现（意味着可以不全部实现 当接口
// 新增或者删除是便会出现遗落问题 ） 该实现方式为非侵入式实现
// 具体实现实例
type myTalk string

var helloTalk = myTalk("hello") // 该类型具体实例

func (talk *myTalk) Hello(userName string) string {
	if strings.EqualFold(userName, "hello") { // 比较两个string 是否相等
		return "hello"
	}
	return "hello world!"
}

func main() {
	fmt.Println(helloTalk.Hello("hello"))
}

// 后续有待补充
