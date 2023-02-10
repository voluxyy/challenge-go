package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	for !(nbr == 1 || nbr == 2) {
		nbr -= 2
	}
	if nbr == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	n := len(os.Args) - 1
	if isEven(n) == false {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
