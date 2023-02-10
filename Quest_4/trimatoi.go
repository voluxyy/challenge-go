package piscine

func TrimAtoi(s string) int {
	m := 0
	l := 0
	value := 0
	n := len(s)
	table := []rune(s)
	sign := 1
	if n == 1 {
		if table[0] >= '0' && table[0] <= '9' {
			return int(table[0]) - '0'
		} else {
			return 0
		}
	}
	for i := 0; i < n; i++ {
		if table[i] == 45 { // Condition si y'a un signe "-"
			if table[i-1] >= '0' && table[i-1] <= '9' {
				sign = 1
			} else {
				sign = -1
				m = 1
				i++
			}
		}
		if table[i] == 43 { // Condition si y'a un signe "+"
			sign = 1
			l = 1
			i++
		}
		if m == 1 && l == 1 {
			return 0
		}
		if table[i] >= '0' && table[i] <= '9' { // Condition si la valeur dans le tableau à l'indice i est comprise entre 0 et 9
			value *= 10
			value += int(table[i]) - '0'
		}
	}
	return value * sign
}
