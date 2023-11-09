package main

import "fmt"

func removeDuplicates(nums []int) int {
	length := len(nums)

	if length <= 1 {
		return length
	}
	counter := 0

	for i := 1; i < length; i++ {
		if nums[counter] != nums[i] {
			counter++
			nums[counter] = nums[i]
		}

	}

	return counter + 1
}

func main() {
	nums := []int{1, 1, 2}
	expectedNums := []int{1, 2}
	k := removeDuplicates(nums)

	if k == len(expectedNums) {
		fmt.Println("OK!")
	}
	for i := 0; i < k; i++ {
		if nums[i] == expectedNums[i] {
			fmt.Println("OK!")
		}
	}
}
