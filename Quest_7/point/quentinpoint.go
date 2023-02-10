package main

import "github.com/01-edu/z01"

func PrintInt(a int) {
	result := '0'
	if a/10 > 0 {
		PrintInt(a / 10)
	}
	for i := 0; i < a%10; i++ {
		result++
	}
	z01.PrintRune(result)
}

func PrintStr(str string) {
	arrayStr := []rune(str)
	for i := range arrayStr {
		z01.PrintRune(arrayStr[i])
	}
}

func setPointQ(ptr int) (int, int) {
	return 42, 21
}

func main() {
	points := 0

	x, y := setPointQ(points)

	PrintStr("x = ")
	PrintInt(x)
	PrintStr(", y = ")
	PrintInt(y)
	z01.PrintRune('\n')
}
