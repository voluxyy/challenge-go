package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	mod := 0
	var tab []int
	if n == 0 {
		z01.PrintRune(48)
	}
	for n != 0 {
		mod = n % 10
		n /= 10
		tab = append([]int{mod}, tab...)
	}
	var table []int
	table = append(tab)
	for i := 0; i < len(table)-1; i++ {
		for i := 0; i < len(table)-1; i++ {
			for table[i] > table[i+1] {
				a := table[i+1]
				table[i+1] = table[i]
				table[i] = a
			}
		}
	}
	for i := 0; i < len(table); i++ {
		z01.PrintRune(rune(table[i]) + 48)
	}
}
