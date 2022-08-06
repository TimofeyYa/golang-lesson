package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Greetings(" т Им о феЙ "))
}

func Greetings(name string) string {
	swapName := name
	swapName = strings.Replace(swapName, " ", "", 10)
	swapName = strings.ToLower(swapName)
	return "Привет, " + strings.Title(swapName)
}
