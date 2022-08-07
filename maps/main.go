package main

import (
	"fmt"
)

func main() {
	fmt.Println(MostPopularWord([]string{"Машинка", "Разработка", "Машинка"}))
}

func UniqueUserIDs(userIDs []int64) []int64 {
	userSwapIDs := []int64{}
	uniMap := map[int64]bool{}

	for _, el := range userIDs {
		if !uniMap[el] {
			uniMap[el] = true
			userSwapIDs = append(userSwapIDs, el)
		}
	}

	return userSwapIDs
}

// Обход мапы

func MostPopularWord(words []string) string {
	mapMostPopularCount := map[string]int8{}
	for _, el := range words {
		if mapMostPopularCount[el] == 0 {
			mapMostPopularCount[el] = 1
		} else {
			mapMostPopularCount[el] = mapMostPopularCount[el] + 1
		}
	}
	var popularWord string
	var popularMax int = 0

	for key, value := range mapMostPopularCount {
		if value > int8(popularMax) {
			popularMax = int(value)
			popularWord = key
		}
	}

	return popularWord
}
