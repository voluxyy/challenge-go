package piscine

func Sqrt(nb int) int {
	result := 0
	if nb == 1 {
		return 1
	}
	if nb == 2 {
		return 0
	}
	if nb > 2 {
		for i := 0; i < nb; i++ {
			result = i * i
			if result == nb {
				return i
			}
		}
	} else {
		return 0
	}
	return 0
}
