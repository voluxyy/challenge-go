package piscine

func IsPrintable(s string) bool {
	r := 0
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 32 && rune(s[i]) <= 127 {
			r++
		}
		if r > len(s)-1 {
			return true
		}
	}
	return false
}
