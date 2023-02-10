package piscine

func AlphaCount(s string) int {
	nb := 0
	r := []rune(s)
	for i := 0; i < len(s); i++ {
		if r[i] >= 97 && r[i] <= 122 || r[i] >= 65 && r[i] <= 90 {
			nb++
		}
	}
	return nb
}
