package piscine

func BasicAtoi2(s string) int {
	value := 0
	n := len(s)
	table := []rune(s)
	for i := 0; i < n; i++ {
		if table[i] < '0' || table[i] > '9' {
			return 0
		} else {
			value *= 10
			value += int(table[i]) - '0'
		}
	}
	return value
}
