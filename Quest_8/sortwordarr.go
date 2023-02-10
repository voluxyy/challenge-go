package piscine

func SortWordArr(a []string) {
	for j := 0; j < len(a)-1; j++ { // Ce for sert à parcourir toutes les caractères de chaque string
		for i := 0; i < len(a)-1-j; i++ { // Ce for sert à parcourir toutes les strings
			if a[i] > a[i+1] { // On compare chaque caractère à celui dans le string suivant
				b := a[i+1]
				a[i+1] = a[i]
				a[i] = b
			}
		}
	}
}
