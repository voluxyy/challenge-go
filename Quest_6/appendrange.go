package piscine

func AppendRange(min, max int) []int {
	table := []int{}
	if min >= max {
		return nil
	} else {
		for i := 0; i+min < max; i++ {
			t := min + i
			table = append(table, t)
		}
	}
	return table
}
