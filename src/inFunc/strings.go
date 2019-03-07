package main

import (
	"fmt"
	"strings"
)

const testString  = "hello world"
func main() {
	testlist := strings.Split(testString , " ")
	fmt.Println(testlist)
}
