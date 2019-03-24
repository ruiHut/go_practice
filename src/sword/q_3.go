package main

import "fmt"

func main() {
	testArray := []int32{1, 2, 2}
	fmt.Println(DuplicateNumberByHashMap(testArray))
}

// 找出数组中重复的数组 o（n） 空间o（n）
func DuplicateNumber(array []int32) int32 {
	res := make([]int32, len(array))
	for _, a := range array {
		res[a] = res[a] + 1
	}

	for _, a := range res {
		if res[a] > 1 {
			return res[a]
		}
	}

	return -1
}

// by hash
func DuplicateNumberByHashMap(array []int32) int32 {
	var res2 = map[int32]int{}
	for _, vars := range array {
		i := res2[vars]
		if i == 0 {
			res2[vars] = 1
		} else {
			return vars
		}
	}
	return -1
}
