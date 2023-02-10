package piscine

func ConcatParams(args []string) string {
	result := ""
	count := 0
	for range args {
		count++
	}
	for i := 0; i < len(args); i++ {
		if count == 1 {
			result += args[i]
		} else {
			result += args[i] + "\n"
			count--
		}
	}
	return result
}
