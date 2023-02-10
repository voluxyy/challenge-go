package piscine

func SplitWhiteSpaces(s string) []string {
	esp := 0
	if s[len(s)-1] == ' ' {
		esp--
	}
	for i := 0; i < len(s); i++ {

		if s[i] == ' ' {
			esp++
			if s[i-1] == ' ' {
				esp--
			}
		}
		if i == len(s)-1 && i != ' ' {
			esp++
		}
	}
	pre := make([]string, len(s))
	result := make([]string, esp)
	for i := 0; i < len(s); i++ {
		pre[i] = string(s[i])
	}
	var i int
	for j := 0; j < len(s); {
		for i = i; i <= len(s); i++ {
			if i == len(s) {
				j = len(s)
				break
			}
			if pre[i] == " " {
				j++
				if pre[i-1] == " " {
					j--
				}
			} else {
				result[j] += pre[i]
			}
		}
	}
	return result
}
