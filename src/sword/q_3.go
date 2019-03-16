package main

func main() {

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
