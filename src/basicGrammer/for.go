package main

import "fmt"

func main() {
	slice1 := make([]string, 100)
	fmt.Println(slice1)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = string(i)
	}
	// for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环
	for s := range slice1 {
		fmt.Println(s)
	}
}
