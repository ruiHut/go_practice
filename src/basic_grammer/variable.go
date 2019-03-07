package main

import "fmt"

// 语言变量
// 声明一般使用 var 关键字

// 第一种 指定类型 声明若不赋值 则使用默认值
var a string

// 第二种，根据值自行判定变量类型。
var b = "abc"

func main() {
	// 第三种，省略var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误。。 且该方法声明的变量必须被使用 否则报错
	// 且只能声明在函数内
	c := "c"
	fmt.Println("string value is", a)
	fmt.Println("b value is", b)
	fmt.Println("c value is", c)
}
