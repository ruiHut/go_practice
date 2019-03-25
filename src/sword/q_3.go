package main

import "fmt"

func main() {
	testArray := []int{1, 2, 3, 4, 3}
	fmt.Println(CompareWithinGroups(testArray))
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

// 拿第i个数字和当前大小的m个数字比较
func CompareWithinGroups(array []int) int {
	if len(array) <= 1 {
		return -1
	}
	var tmp = array[0]
	for index := 0; index < len(array)-1; index++ {
		if array[index] < 0 || array[index] > len(array)-1 {
			return -1
		}
	}

	for index := 0; index < len(array)-1; index++ {
		for array[index] != index {
			if array[index] == array[array[index]] {
				return array[index]
			}
			// swap
			tmp = array[index]
			array[index] = array[tmp]
			array[tmp] = tmp
		}
	}
	return -1
}
