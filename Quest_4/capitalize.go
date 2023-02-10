package piscine

func Capitalize(s string) string {
	table := []rune(s)
	for i := 0; i < len(s); i++ {
		if table[i] >= 65 && table[i] <= 90 {
			table[i] = rune(s[i] + 32)
		}
	}
	if table[0] >= 97 && table[0] <= 122 {
		table[0] = table[0] - 32
	}
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 {
			if table[i] >= 32 && table[i] <= 47 || table[i] >= 92 && table[i] <= 96 || table[i] >= 58 && table[i] <= 64 || table[i] >= 91 && table[i] <= 96 || table[i] >= 123 && table[i] <= 126 {
				if table[i+1] >= 97 && table[i+1] <= 122 {
					table[i+1] = table[i+1] - 32
				}
			}
		}
	}
	return string(table)
}
