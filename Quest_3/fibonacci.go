package piscine

func Fibonacci(index int) int {
	result := 0
	if index < 0 {
		return -1
	} else if index == 1 {
		return 1
	} else if index >= 2 {
		result = Fibonacci(index-1) + Fibonacci(index-2)
	}
	return result
}
