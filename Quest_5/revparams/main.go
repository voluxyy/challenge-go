package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		argument := []rune(os.Args[i])
		if i > 0 {
			for i := 0; i < len(argument); i++ {
				z01.PrintRune(argument[i])
			}
		}
		z01.PrintRune('\n')
	}
}
