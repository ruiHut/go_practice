package main

import "fmt"

// defer 延迟 用于延迟调用指定的函数
// 只能出现在函数的内部 由defer关键字以及针对某个函数的调用表达式组成
// 简单实例
func outerFunc() { // 针对与此例中 outer 称为外围函数 调用者即main 称为调用函数 位置不限 数量不限
	defer fmt.Println("this is defer func1")
	defer fmt.Println("this is defer func2")
	fmt.Println("this is after func")
}

// 具体规则 由先到后	本体顺序函数  》 延迟函数 》 return 》 调用函数
// 两大特点
func printNumbers() {
	for i := 0; i < 5; i++ {
		defer func(n int) { // 	如果要延迟函数要调用外部变量， 就应该通过参数传入 因为执行时可能该外部变量已经改变
			fmt.Printf("%d", n) // 4  3  2  1  0	当i 没有传入结果为 5 5 5 5 5
		}(i * 2) // 为什么不是  0 1 2 3 4 5 cause：同一个外围函数内有多个延迟函数调用顺序，会与其所属当defer语句的执行顺序完全相反
		// 当遇到时会押入同一个栈中 所以取出来时顺序相反 切如果有参数传入 会保存该参数当前数值
	}
}

func main() {
	printNumbers()
	outerFunc()
}
