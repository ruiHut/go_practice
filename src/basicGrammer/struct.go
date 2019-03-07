package main

import "fmt"

// 结构体 既可以关联方法 也可以 内置字段

// 简单实例 针对中文的演示级聊天机器人
type simpleCN struct {
	name string
	talk string
}

var simple1 = simpleCN{"hello", "world"} // 可以省略字段名  顺序需要一致 零值为nil
func main() {
	fmt.Println(simple1)
}
