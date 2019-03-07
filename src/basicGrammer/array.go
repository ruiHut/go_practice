package main

import "fmt"

var ipv4 [4]uint8 = [4]uint8{192, 23, 23} // 如果不设置初始大小 则为0 且其长度不可变 对于需要规划内存时，数组类型非常有用 如果超过将编译期报错

func main() {
	fmt.Println(ipv4)
	//if len(ipv4) == cap(ipv4) {		// 对于数组而言 len 和 cap 恒相等
	//	fmt.Println("is equal ", cap(ipv4), "is array`s len")
	//}
	// 获取数组的数值
	fmt.Println(ipv4[1])   // 索引表达式获取
	fmt.Println(ipv4[2:3]) // 切片表达式获取数组数值
	// 切片表达式同样适用于数组 且其为左开右闭区间 如果左大于右编译时期报错
}
