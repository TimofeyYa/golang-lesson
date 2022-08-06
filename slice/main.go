package main

import "fmt"

func main() {
	fmt.Printf("%d", Remove([]int{1, 2, 3, 4, 5}, 0))
}

func Remove(nums []int, i int) []int {
	if len(nums) < i || i < 0 {
		return nums
	}
	var swapNums []int
	for j, el := range nums {
		if j != i {
			swapNums = append(swapNums, el)
		}
	}
	return swapNums
}
