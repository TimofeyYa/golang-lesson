package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%d", SafeWrite(nums, 10, 10))

}

func SafeWrite(nums [5]int, i, val int) [5]int {
	if len(nums) > i && i > -1 {
		nums[i] = val
	}
	return nums
}
