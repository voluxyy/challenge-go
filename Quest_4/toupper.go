package piscine

func ToUpper(s string) string {
	t := []rune(s)
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 97 && rune(s[i]) <= 122 {
			t[i] = t[i] - 32
		}
	}
	return string(t)
}
