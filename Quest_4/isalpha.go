package piscine

func IsAlpha(s string) bool {
	r := 0
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 97 && rune(s[i]) <= 122 || rune(s[i]) >= 65 && rune(s[i]) <= 90 || rune(s[i]) >= 48 && rune(s[i]) <= 57 {
			r++
		}
		if rune(s[i]) == ' ' {
			r++
		}
		if r > len(s)-1 {
			return true
		}
	}
	return false
}
