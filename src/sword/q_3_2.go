package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 1, 3}
	fmt.Println(getDuplication(array))
	fmt.Println(4 << 2)
}

// 分治法 解决  nlogn 空间 1

func getDuplication(array []int) int {
	if len(array) == 0 {
		return -1
	}

	start := 1
	end := len(array) - 1
	for end >= start {
		middle := ((end - start) >> 1) + start // (start + end )/2 有溢出可能性 两数相加有溢出可能性 	>> 为除二的小技巧  >>2 为除4 <<1 为乘2 以此类推
		m := model(array, len(array), start, middle)
		count := countRange(m)
		if end == start {
			if count > 1 {
				return start
			} else {
				break
			}
		}

		if count > (middle - start + 1) {
			end = middle
		} else {
			start = middle + 1
		}
	}

	return -1
}

func countRange(m Duplication) int {

	count := 0
	for i := 0; i < m.length; i++ {
		if m.array[i] >= m.start && m.array[i] <= m.middle {
			count++
		}
	}
	return count
}

type Duplication struct {
	array  []int
	length int
	start  int
	middle int
}

func model(array []int, length int, start int, middle int) Duplication {
	var m = Duplication{
		length: length,
		array:  array,
		start:  start,
		middle: middle,
	}
	return m
}
