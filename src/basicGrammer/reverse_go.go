package main

// 将字符串反转

import (
	"fmt"
)

func main() {
	fmt.Println(Reverse("hello"))
	fmt.Println(MyReverse("hello"))
}

func MyReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)

}

func Reverse(s string) string { // 函数格式 函数名（型参 变量类型） 返回类型 void 可以省略
	r := []rune(s)                                           // 将字符串s转换为ascll编码的int数组 保存到r中  runa 内置别名 相当于 int32
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 { // := 声明变量并赋值 相当于 val i = 1 切其使用范围也为仅在for 循环内
		r[i], r[j] = r[j], r[i] // len（）内置函数 可以用来查看数组或slice的长度
	}
	return string(r)
}

// i, j := 0, len(r) - 1  	完成一次初始化 结果为 i = 0，  j = len（r） - 1
// i, j = i + 1 , j - 1   	完成一次给两个变量赋值的操作 结果为 i + 1, j - 1
