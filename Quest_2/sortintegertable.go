package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table)-1; i++ {
		for i := 0; i < len(table)-1; i++ {
			for table[i] > table[i+1] {
				a := table[i+1]
				table[i+1] = table[i]
				table[i] = a
			}
		}
	}
}
