package piscine

func IsPrime(nb int) bool {
	var result bool
	if nb < 2 {
		result = false
	} else if 2 == nb || 3 == nb {
		result = true
	} else {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				result = false
			}
		}
	}
	return result
}
