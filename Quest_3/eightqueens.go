package piscine

import "github.com/01-edu/z01"

var position = [8]int{}

func EightQueens() {
	solu(0)
}

func Safe(nb_reine, ligne int) bool {
	for i := 0; i < nb_reine; i++ {
		ligne_diff := position[i]

		if ligne_diff == ligne || ligne_diff == ligne-(nb_reine-i) || ligne_diff == ligne+(nb_reine-i) {
			return false
		}
	}
	return true
}

func solu(k int) {
	if k == 8 {
		for i := 0; i < 8; i++ {
			PrintResult(position[i] + 1)
		}
		z01.PrintRune('\n')
	} else {
		for i := 0; i < 8; i++ {
			if Safe(k, i) {
				position[k] = i
				solu(k + 1)
			}
		}
	}
}

func PrintResult(s int) {
	z01.PrintRune(rune(s) + 48)
}
