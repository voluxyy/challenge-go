package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	result := 0
	if nbr < "0" {
	} else {
		if baseFrom == "01" {
			if baseTo == "0123456789" {
				binaryToBase10(nbr)
			}
		}
	}
	return string(result)
}

func binaryToBase10(s string) []string {
	tab := make([]string, len(s))
	result := make([]string, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		tab[i] = string(s[i])
	}
	base2 := 1
	for i := 0; i < len(tab); i++ {
		if tab[i] == "1" {
			base2 *= 2
		} else {
			result[i] = "0"
			base2 *= 2
		}
	}
	return result
}

// Solution :
//	ln := 0
//	ln2 := 0
//	ln3 := 0
//	result := ""
//	Nbr_baseFrom_baseTo := map[rune]int{}
//	for c := range nbr {
//		ln = c + 1
//	}
//	for i, c := range baseFrom {
//		Nbr_baseFrom_baseTo[c] = i
//		ln2 = i + 1
//	}
//	for c := range baseTo {
//		ln3 = c + 1
//	}
//	pw := 1
//	cnt := 0
//	for i := ln - 1; i >= 0; i-- {
//		cnt = cnt + Nbr_baseFrom_baseTo[[]rune(nbr)[i]]*pw
//		pw *= ln2
//	}
//	for cnt != 0 {
//		result = string(baseTo[cnt%ln3]) + result
//		cnt /= ln3
//	}
//	return result
//	}
