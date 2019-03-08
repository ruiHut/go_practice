package main

import (
	"fmt"
)

//  方法 ：带有接收者的函数

// 定义一个矩型对象
type Rectangle struct {
	width, height float64
}

// area 不是作为Rectangle的方法实现的 仅仅是一个方法
func area(r Rectangle) float64 {
	return r.width * r.height
}

// method 定义常量 全包可用
const (
	WHITH = iota // iota == 0
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte // uint8 别称

type Box struct {
	width, height, depth float64
	color                Color // 可以看作为byte的别名
}

type BoxList []Box // a slice of boxes

// 方法名前设置 type 即看作为给对象设置的方法 是与对象绑定的
func (b Box) Volume() float64 {
	return b.height * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

// 寻找最大的Coler
func (bl BoxList) BiggerstColer() Color {
	v := 0.00
	k := Color(WHITH)

	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

// 全部都设为黑色
func (bl BoxList) PainrItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

// toString

func (c Color) String() string {
	strings := []string{"w", "y", "r", "b"}
	return strings[c]
}
func main() {
	fmt.Println(byte(1))
	r1 := Rectangle{12, 2}
	r2 := Rectangle{2, 4}
	fmt.Println("area of r1 is ", area(r1))
	fmt.Println("area of r2 is ", area(r2))
}
