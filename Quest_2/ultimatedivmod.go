package piscine

func UltimateDivMod(a *int, b *int) {
	var z int
	z = *a / *b
	*b = *a % *b
	*a = z

	if *b > 0 && *b < 1 {
		*b = 1
	} else if *b > -1 && *b < 0 {
		*b = -1
	} else {
	}
}
