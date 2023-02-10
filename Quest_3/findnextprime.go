package piscine

func FindNextPrime(nb int) int {
	result := nb
	if nb <= 1 {
		return 2
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return FindNextPrime(nb + 1)
		}
	}
	return result
}
