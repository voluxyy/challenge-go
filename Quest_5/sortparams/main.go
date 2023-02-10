package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	value := os.Args[1:]
	for j := 0; j < len(value)-1; j++ { // Ce for sert à parcourir toutes les caractères de chaque string
		for i := 0; i < len(value)-1-j; i++ { // Ce for sert à parcourir toutes les strings
			if value[i] > value[i+1] { // On compare chaque caractère à celui dans le string suivant
				a := value[i+1]
				value[i+1] = value[i]
				value[i] = a
			}
		}
	}
	for i := 0; i < len(value); i++ {
		j := 0
		for j = 0; j < len(os.Args[i+1]); j++ {
			z01.PrintRune(rune(value[i][j]))
		}
		z01.PrintRune('\n')
	}
}
