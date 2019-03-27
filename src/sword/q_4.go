package main

import (
	"fmt"
)

//二维数组中的查找

// 对于一个向右递增 向下递增的二维数组查找而言
// 当我们需要解决一个复杂问题时， 一个有效的办法就是从一个具体问题入手， 通过分析简单具体的例子，试图寻找普遍的规律
func main(){
	array := [2][2]int{
 		{1,2},
		{2,3},
	}
	fmt.Println(FindByself(array, 2, 2, 2))
}

func Find(array [2][2]int, rows int, columns int, num int)  bool {
	res := false

	if rows>0 && columns>0 {
		row := 0
		column := columns -1

		for row < rows && column >= 0 {	 // 结束条件
			if array[row][column] == num {
				res = true
				return res
			}else if array[row][column] > num {
				column--
			}
				row++
		}
	}

	return	res
}

func FindByself(array [2][2]int, rows int , columns int, num int) bool {
	res := false

	if rows>0 && columns>0 {
		if num < array[0][0] || num > array[rows-1][columns-1] {
			return false
		}

		row := 0
		column := columns -1

		for row < rows && column >= 0 {
			if array[row][column] == num {
				res = true
				return res
			}
			if array[row][column] > num {
				column--
			}
				rows++
		}
	}

	return res
}