package main

//  channel is a data structure to translate data
// 	channel can use to synchronous operation and communication  double goroutine

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum // put sum to channel c
}

func main() {
	s := []int{1, 2, 3123, 12, 4, 42, 21, -312}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // accept from channel c

	fmt.Println(x, y, x+y)
}
