package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var result bool
	if a[0] == 0 {
		result = true
	}
	for i := 1; i < len(a); i++ {
		if i < len(a)-1 {
			if a[i] >= a[i-1] && a[i] <= a[i+1] {
				result = true
			} else if a[i] <= a[i-1] && a[i] >= a[i+1] {
				result = true
			} else {
				return false
			}
		}
	}
	return result
}
