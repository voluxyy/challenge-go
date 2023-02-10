package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for mil := 48; mil <= 57; mil++ {
		for cen := 48; cen <= 57; cen++ {
			for dix := 48; dix <= 57; dix++ {
				for uni := 48; uni <= 57; uni++ {
					if mil*10+cen < dix*10+uni {
						z01.PrintRune(rune(mil))
						z01.PrintRune(rune(cen))
						z01.PrintRune(' ')
						z01.PrintRune(rune(dix))
						z01.PrintRune(rune(uni))
						if !(mil == 57 && cen == 56 && dix == 57 && uni == 57) {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
