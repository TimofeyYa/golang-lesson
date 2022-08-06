package main

import "fmt"

type Numbers struct {
	num1 int
	num2 int
}

type NumbersInterface interface {
	sum() int
}

func (n Numbers) sum() int {
	return n.num1 + n.num2
}

func main() {
	var i NumbersInterface
	numbers := Numbers{5, 4}
	i = numbers
	fmt.Printf("%d", i.sum())
}
