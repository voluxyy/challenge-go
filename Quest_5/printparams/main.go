package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		argument := []rune(os.Args[i])
		if i < len(os.Args) {
			for i := 0; i < len(argument); i++ {
				z01.PrintRune(argument[i])
			}
		}
		z01.PrintRune('\n')
	}
}
