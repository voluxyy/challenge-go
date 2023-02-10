package piscine

func IsUpper(s string) bool {
	r := 0
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 65 && rune(s[i]) <= 90 {
			r++
		}
		if r > len(s)-1 {
			return true
		}
	}
	return false
}
