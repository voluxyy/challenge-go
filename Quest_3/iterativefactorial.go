package piscine

func IterativeFactorial(nb int) int {
	n := 1
	if nb >= 0 && nb <= 20 {
		for i := 1; i <= nb; i++ {
			n = n * i
		}
		return n
	} else {
		return 0
	}
}
