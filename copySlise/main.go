package main

import "fmt"

func main() {
	fmt.Println(IntsCopy([]int{1, 2, 3}, 2))
}

func IntsCopy(src []int, maxLen int) []int {
	if maxLen < 0 {
		maxLen = 0
	}
	if maxLen > len(src) {
		maxLen = len(src)
	}
	arrCp := make([]int, maxLen)
	copy(arrCp, src)
	return arrCp
}
