package main

import "fmt"

func main() {
	fmt.Println(MergeNumberLists([]int{1, 2, 3}, []int{4}, []int{5, 6, 7, 8}, []int{9, 10}))
}

func MergeNumberLists(numberLists ...[]int) []int {
	joinSlise := []int{}
	for _, el := range numberLists {
		joinSlise = append(joinSlise, el...)
	}

	return joinSlise
}
