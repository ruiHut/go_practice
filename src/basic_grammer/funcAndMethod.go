package main

import (
	"errors"
	"fmt"
)

// 函数和方法
//	函数值既可以作为其他函数的参数，也可以作为其结果
//	函数的结果可以是一个结构组 与java单一返回结果不同

// 基本形式 func divide(divident int, divisor int) (int ,error){ //省略代码部分 }    结果列表最好保留名称 即型参

// 方法是函数的一种 他其实就是与某个*数据类型*关联在一起的函数
// 在java中可以自定义对象 给对象创建的叫方法 在此 给类型创造的也同理叫方法

type myInt int

func (i myInt) add(another int) myInt {
	i = i + myInt(another)
	return i
}

// 具体使用 该x其实可以理解为 一个对象为int 的实例
var x = myInt(1)

func main() {
	fmt.Println(test()) // 作为其他函数的结果
	test2(test)         // 作为其他函数的参数
	fmt.Println(divide(1, 1))
	x.add(5)
	fmt.Println(x) // add 方法结束后 并未改变x的实际大小 如需改变则需要将其改为指针类型
}

func test() int { return 1 }

func test2(func() int) {}

// 保留返回结果对函数 省略return 后续代码
// 如果在 return 前没有设置数值 则返回默认大小 int = 0 其他非基础类型 为 nil
func divide(divident int, divisor int) (result int, err error) {
	if divisor == 0 {
		err = errors.New("division by zero")
		return 1, err
	}
	result = divident / divisor
	return
}

// 定义一个 二元操作符对函数类型 让功能实现 让接口定义行为不再是唯一途径
type binaryOperation func(operand1 int, operand2 int) (result int, err error)

func operate(op1 int, op2 int, bop binaryOperation) (reslut int, err error) { // 实现权转移给 operate 函数
	if bop == nil {
		err = errors.New("invalid binary operation function")
		return
	}
	return bop(op1, op2)
}
