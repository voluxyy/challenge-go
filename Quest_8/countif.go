package piscine

func CountIf(f func(string) bool, tab []string) int {
	nb := 0
	for i := 0; i < len(tab); i++ {
		if f(tab[i]) == true {
			nb++
		}
	}
	return nb
}
