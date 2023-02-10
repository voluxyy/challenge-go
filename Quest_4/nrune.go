package piscine

func NRune(s string, n int) rune {
	r := []rune(s)
	if n <= len(r) && n > 0 {
		return r[n-1]
	}
	return 0
}
