package main

import "fmt"

// 切片使用类似于 java的 集合 不需要预先设置长度

var ips = []string{"127.0.0.1", "127.0.0.7"} // 定义一个切片 其与数组不同 切片并不携带长度 因为其长度可以改变
// 元素类型相同 切片的类型也相同	一个切片的类型的零值总为 nil
var array [4]int8 = [4]int8{1, 2, 4, 4}

var ips3 = []string{}

var ips2 = []int32{1, 23, 43}

func main() {
	fmt.Println(ips)
	fmt.Println(ips2)
	fmt.Println(ips2[1])
	fmt.Println(ips2[3:3])
	ips = append(ips, "hello") // 切片新增一个对象 自动从最后一个开始增加 新 旧切片可能指向不同底层数组
	fmt.Println(ips)
	fmt.Println(array)
	fmt.Println(len(ips3))
	fmt.Println(cap(ips)) // cap 指底层容量
	fmt.Println(len(ips)) // len 指当前实际使用

	// 内建函数 快速初始化 切片大小
	ips = make([]string, 100)
	fmt.Println(len(ips)) // len = 100
	fmt.Println(cap(ips)) // cap = 100
	fmt.Println(ips)      // 内容全部 为 nil 意味着会重新产生一个底层数组 且数据全部为nil
}
