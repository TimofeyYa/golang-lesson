package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Create("./simples.txt")
	if err != nil {
		fmt.Println("Тут проблемы с созданием файлов")
	}

	for i := 2; i < 80000; i++ {
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

	fmt.Println("Я начал выполнять следующий код")

	file2, err2 := os.Create("./simples2.txt")
	if err2 != nil {
		fmt.Println("Тут проблемы с созданием файлов")
	}

	for i := 2; i < 80000; i++ {
		count := 0
		for j := 1; j < i; j++ {
			if i%2 == 0 {
				count++
			}
		}
		if count > 4 {
			stringData := strconv.Itoa(i) + "\n"
			file2.WriteString(stringData)
		}
	}
	file2.Close()

	searchSimpesAndWrite(80000)
	fmt.Printf("Я запустил все работы\n")
	fmt.Println("Я Наконец закончил работу")
}

func searchSimpesAndWrite(nam int) {
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
}
