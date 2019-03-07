package main

import "fmt"

func main() {
	x, y := 5, 19
	fmt.Println(x << 5)
	fmt.Println(y << 9)
	fmt.Println(x<<5 + y<<9) // 160 + 9728 = 9888
}

/*
Package regexp implements a simple library for regular expressions.

The syntax of the regular expressions accepted is:

    regexp:
        concatenation { '|' concatenation }
    concatenation:
        { closure }
    closure:
        term [ '*' | '+' | '?' ]
    term:
        '^'
        '$'
        '.'
        character
        '[' [ '^' ] character-ranges ']'
        '(' regexp ')'
*/
