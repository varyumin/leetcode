package main

import "fmt"

func singleNumber(nums []int) int {
	hash := make(map[int]int)
	for _, v := range nums {
		val, ok := hash[v]
		if ok {
			hash[v] = val + 1
		} else {
			hash[v] = 0
		}
	}
	for k, _ := range hash {
		if hash[k] == 0 {
			return k
		}
	}
	return 0
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 2, 3, 3}))
}
