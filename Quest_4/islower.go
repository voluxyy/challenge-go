package piscine

func IsLower(s string) bool {
	r := 0
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 97 && rune(s[i]) <= 122 {
			r++
		}
		if r > len(s)-1 {
			return true
		}
	}
	return false
}
