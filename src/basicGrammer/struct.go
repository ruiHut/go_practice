package main

import "fmt"

// 结构体 既可以关联方法 也可以 内置字段

// 简单实例 针对中文的演示级聊天机器人
type simpleCN struct {
	name string
	talk string
}

// 声明使用方式
var simple1 = simpleCN{"hello", "world"} // 可以省略字段名  顺序需要一致 零值为nil

var simple2 simpleCN

func main() {
	// new 函数分配一个指针， 此处的p的类型为 *person  &对象 代表其指针
	simple3 := new(simpleCN)

	simple2.name = "zhouwang"
	simple2.talk = "hello i am zhouwang!"
	fmt.Println(simple1)
	fmt.Println(&simple2)
	testStruct(simple3)
	fmt.Println(simple3)

}
func testStruct(p *simpleCN) {
	p.name = "zhouwang"
	p.talk = "hello i am good boy"

}
