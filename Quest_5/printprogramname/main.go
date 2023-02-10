package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := []rune(os.Args[0])
	for i := range argument {
		if i > 1 {
			z01.PrintRune(argument[i])
		}
	}
	z01.PrintRune('\n')
}
