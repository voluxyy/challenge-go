package piscine

func ToLower(s string) string {
	t := []rune(s)
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 65 && rune(s[i]) <= 90 {
			t[i] = t[i] + 32
		}
	}
	return string(t)
}
