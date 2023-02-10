package piscine

func Split(s, sep string) []string {
	var texte []string
	taille := len(sep)
	save := 0
	for i := 0; i+taille < len(s); i++ {
		if s[i:taille+i] == sep { // Pour parcourir la string s de la longueur du sep
			texte = append(texte, s[save:i]) // Ajoute au table texte depuis la denier save jusqu'a le début de la suivant
			save = i + taille                // On fait une save de l'index auquel est le dernier sep
		}
		if i+taille == len(s)-1 && s[i:taille+i] != sep { // Si on atteint la fin de la string s et que la fin est différente du sep
			texte = append(texte, s[save:]) // Ajoute au tableau texte depuis le dernier sep jusqu'a la fin
		}
	}
	return texte // Retourne le texte
}
