package piscine

func IsNumeric(s string) bool {
	r := 0
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 48 && rune(s[i]) <= 57 {
			r++
		}
		if r > len(s)-1 {
			return true
		}
	}
	return false
}
