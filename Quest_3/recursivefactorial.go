package piscine

func RecursiveFactorial(nb int) int {
	if nb == 0 || nb == 1 {
		return 1
	}
	if nb >= 0 && nb < 66 {
		if nb > 1 {
			return nb * RecursiveFactorial(nb-1)
		}
	}
	return 0
}
