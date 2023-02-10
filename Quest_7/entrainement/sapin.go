package main

import (
	"fmt"
	"os"
)

func Atoi(s string) int {
	m := 0
	l := 0
	b := 0
	value := 0
	n := len(s)
	table := []rune(s)
	sign := 1
	if n == 1 {
		if table[0] >= '0' && table[0] <= '9' {
			return int(table[0]) - '0'
		} else {
			return 0
		}
	}
	for i := 0; i < n; i++ {
		if table[len(s)-1] == 43 || table[len(s)-1] == 45 {
			return 0
		}
		if table[i] == 45 { // Condition si y'a un signe "-"
			sign = -1
			m = 1
			i++
		}
		if table[i] == 43 { // Condition si y'a un signe "+"
			sign = 1
			l = 1
			i++
		}
		if m == 1 && l == 1 {
			return 0
		}
		if table[i] >= '0' && table[i] <= '9' { // Condition si la valeur dans le tableau à l'indice i est comprise entre 0 et 9
			b += 1
		} else {
			b += 1000
		}
		if b > 1000 {
			return 0
		} else {
			value *= 10
			value += int(table[i]) - '0'
		}
	}
	return value * sign
}

func main() {
	arg := Atoi(os.Args[1]) // Nombre de ligne
	nbs_esp := arg          // Moitié
	nbs_star := 1

	// Sapin
	for i := 0; i < arg; i++ { // Nombre de lignes
		for j := 0; j < nbs_esp; j++ { // Nombre d'espace
			fmt.Print(" ")
		}
		for k := 0; k < nbs_star; k++ { // Nombre d'étoile
			fmt.Printf("*")
		}
		fmt.Printf("\n")
		nbs_star += 2
		nbs_esp -= 1
	}
	// Tronc
	nbs_star -= 2
	nbs_esp += 1
	nbt_t := nbs_star / 3
	nbt_esp := nbs_star / 3

	for l := 0; l < nbt_esp+1; l++ { // Nombre d'espace
		fmt.Print(" ")
	}
	for m := 0; m < nbt_t; m++ { // Nombre de tronc
		fmt.Printf("|")
	}
	fmt.Printf("\n")
}
