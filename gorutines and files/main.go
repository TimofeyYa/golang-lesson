package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go searchSimpesAndWrite(80000, ch)
	go searchSimpesAndWrite2(80000, ch2)
	go searchSimpesAndWrite(80000, ch3)
	fmt.Printf("Я запустил все работы\n")
	fmt.Printf("Горутины закончили работу с кодами %d %d\n", <-ch, <-ch)

}

func searchSimpesAndWrite(nam int, ch chan int) {
	file, err := os.Create("./simples.txt")
	if err != nil {
		fmt.Println("Тут проблемы с созданием файлов")
	}

	for i := 2; i < nam; i++ {
		count := 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count < 3 {
			stringData := strconv.Itoa(i) + "\n"
			file.WriteString(stringData)
		}
	}
	file.Close()
	ch <- 0
	close(ch)
}

func searchSimpesAndWrite2(nam int, ch chan int) {
	file, err := os.Create("./simples2.txt")
	if err != nil {
		fmt.Println("Тут проблемы с созданием файлов")
	}

	for i := 2; i < nam; i++ {
		count := 0
		for j := 1; j < i; j++ {
			if i%2 == 0 {
				count++
			}
		}
		if count > 3 {
			stringData := strconv.Itoa(i) + "\n"
			file.WriteString(stringData)
		}
	}
	file.Close()
	ch <- 0
	close(ch)
}
