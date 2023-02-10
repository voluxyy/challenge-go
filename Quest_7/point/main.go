package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func PrintR(str string) {
	arrayStr := []rune(str)
	for i := range arrayStr {
		z01.PrintRune(arrayStr[i])
	}
}

func PrintNbr(a int) {
	result := '0'
	if a/10 > 0 {
		PrintNbr(a / 10)
	}
	for i := 0; i < a%10; i++ {
		result++
	}
	z01.PrintRune(result)
}

func setPoint(ptr *point) (int, int) {
	return 42, 21
}

func main() {
	points := &point{}
	x, y := setPoint(points)
	setPoint(points)
	PrintR("x = ")
	PrintNbr(x)
	PrintR(", y = ")
	PrintNbr(y)
	z01.PrintRune('\n')
}
