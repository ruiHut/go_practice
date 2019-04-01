package main

import (
	"fmt"
	"sync"
)

func Countoff(n int) {
	fmt.Println("i am : ", n)
}

func main() {
	wg := &sync.WaitGroup{}
	if nil == wg {
		fmt.Println("what the fuck!")
	} else {
		wg.Add(10)
		for i := 0; i < 10; i++ {
			go func(n int) {
				Countoff(n)
				wg.Done()
			}(i)
		}
		wg.Wait()
	}
}
