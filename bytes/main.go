package main

func main() {
	println(shiftASCII("Zorro", 1))
}

func nextASCII(b byte) byte {
	return b + 1
}

func prevASCII(b byte) byte {
	return b - 1
}

// Обход строки
func shiftASCII(s string, step int) string {
	if step == 0 {
		return s
	}

	sSwap := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		sSwap[i] = byte(int(s[i]) + step)
	}

	return string(sSwap)
}
