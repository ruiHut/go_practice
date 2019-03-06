package main

import "fmt"

//	go 中字典的类型是散列表（hash table）的一个实现 官方称为map 类似于切片类型 map也是一个引用类型
//		散列表是一个实现了关联数组的数据结构（映射关系）同一个索引在一个map中唯一

//	简单声明
var ipSwitches = map[string]bool{"false": false, "true": true} // 其中key 为 string value 为 bool

func main() {
	// 添加和修改 索引表达式
	ipSwitches["s"] = false
	ipSwitches["r"] = true
	ipSwitches["s"] = true
	// 根据 key查找value
	fmt.Println(ipSwitches["s"])
	// foreach 遍历map
	// 删除 根据内置函数
	delete(ipSwitches, "s")
}
