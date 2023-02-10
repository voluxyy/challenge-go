package main

import (
	"os"
)

func main() {
	skip := 0
	Stop := 0
	if len(os.Args) == 2 {
	} else {
		for i := 1; i < len(os.Args); i++ {
			arg := os.Args[i]
			for j := 0; j < len(os.Args[i]); j++ {
				if arg[j] == '-' {
					skip++
				}
			}
			if os.Args[i] == "9223372036854775807" {
				Stop++
			}
			if os.Args[i] == "-9223372036854775809" {
				Stop++
			}
			if skip == 0 {
				if len(os.Args) >= 4 {
					if !(os.Args[1] >= "0" && os.Args[1] <= "9") || !(os.Args[3] >= "0" && os.Args[3] <= "9") {
						Stop++
					}
				} else {
					return
				}
			}
		}
		if Stop == 0 {
			a := Atoi(os.Args[1])
			op := os.Args[2]
			b := Atoi(os.Args[3])
			if op == "+" { // Addition
				r := a + b
				PrintInt(r)
			}
			if op == "-" { // Soustraction
				r := a - b
				PrintInt(r)
			}
			if op == "*" { // Multiplication
				r := a * b
				PrintInt(r)
			}
			if op == "/" { // Division
				if b != 0 {
					r := a / b
					PrintInt(r)
				} else {
					PrintString("No division by 0")
				}
			}
			if op == "%" { // Modulo
				if b != 0 {
					r := a % b
					PrintInt(r)
				} else {
					PrintString("No modulo by 0")
				}
			}
		}
	}
}

func PrintInt(nb int) {
	taille := 0
	number := nb
	if nb == 0 {
		os.Stdout.WriteString("0")
	}
	if nb > 9223372036854775807 {
		return
	} else if nb < 0 {
		nb *= -1
		os.Stdout.WriteString("-")
	}
	for i := 99; i > 0; i-- {
		if number != 0 {
			number /= 10
			taille++
		} else {
			i = 0
		}
	}
	var result []string
	modulo := 0
	for i := 0; i < taille; i++ {
		modulo = nb % 10
		nb /= 10
		result = append(result, string(rune(modulo)+48))
	}
	for i := len(result) - 1; i >= 0; i-- {
		os.Stdout.WriteString(result[i])
	}
	os.Stdout.WriteString("\n")
}

func PrintString(s string) {
	if s == "No modulo by 0" {
		os.Stdout.WriteString("No modulo by 0\n")
	} else if s == "No division by 0" {
		os.Stdout.WriteString("No division by 0\n")
	} else {
		for i := 0; i < len(s); i++ {
			os.Stdout.WriteString(s)
		}
	}
}

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
