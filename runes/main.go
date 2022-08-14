package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(latinLetters("hi baby Lena я рад тебя видеть 23"))
}

func isASCIITEST(s string) bool {
	str := "hello world!"
	runa := []rune("hello world!")
	fmt.Printf("Длинна строки %d, а длинна руны %d\n", len(str), len(runa))
	for i := 0; i < len(str); i++ {
		fmt.Println("Строка: ", str[i], reflect.TypeOf(str[i]))
		fmt.Println("Руна: ", runa[i], reflect.TypeOf(runa[i]))
	}
	return true
}

func isASCII(s string) bool {
	runa := []rune(s)
	return len(s) == len(runa)
}

func latinLetters(s string) string {
	latinString := &strings.Builder{}
	runeStr := []rune(s)

	for _, r := range runeStr {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			latinString.WriteString(string(r))
		}
	}

	return latinString.String()
}
