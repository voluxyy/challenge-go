package piscine

func BasicAtoi(s string) int {
	value := 0                     // Variable value
	j := 0                         // Variable j
	table := []rune(s)             // Tableau qui récupère les valeurs en rune
	for _, number := range table { // Boucle : _ variable qui stocke les indices. number variable qui stocke les nombres
		for i := '0'; i < number; i++ { // Boucle tant que i n'est pas supérieur à number.
			j++ // j prend +1 chaque tour de variable et donc obtient la valeur du chiffre contenu dans le tableau
		}
		value = value*10 + j // value se multiplie lui même par dix et ajoute j
		j = 0                // On remet j à 0 pour qu'elle puisse stocker le prochain chiffre
	}
	return value
}
