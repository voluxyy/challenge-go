package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	table := make([]int, max-min)
	for i := 0; i+min < max; i++ {
		t := min + i
		table[i] = t
	}
	return table
}
