package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(UniqueSortedUserIDs([]int64{}))
}

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	sort.Slice(userIDs, func(i, j int) bool {
		return userIDs[i] < userIDs[j]
	})
	userSwapIDs := []int64{}

	for i, el := range userIDs {
		if i != len(userIDs)-1 {
			if el != userIDs[i+1] {
				userSwapIDs = append(userSwapIDs, el)
			}
		} else {
			userSwapIDs = append(userSwapIDs, el)
		}
	}

	return userSwapIDs
}
