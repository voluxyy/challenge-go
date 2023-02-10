package piscine

func PrintNbrBase(nbr int, base string) {
	if len(base)-1 >= 2 {
		for i := 0; i < len(base); i++ {
			if base[i] != base[i+1] {
				if base[i] != 43 || base[i] != 45 {
				} else {
					return
				}
			} else {
				return
			}
		}
	} else {
		return
	}
}
