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
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段Human
	school string
	loan   float32
}

type Employee struct {
	Human   //匿名字段Human
	company string
	money   float32
}

//Human对象实现Sayhi方法
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human对象实现Sing方法
func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

//Human对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee重载Human的Sayhi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行
}

//Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

//Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}

// 定义interface
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}
