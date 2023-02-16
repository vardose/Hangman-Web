package Hangman

func Reveal(word, lettre, ToFind string) string { // fonction qui sert a révéler les lettres du mot qui sont trouvées
	tab := []int{}
	for i := 0; i < len(ToFind); i++ {
		if string(ToFind[i]) == lettre {
			tab = append(tab, i)
		}
	}
	for j := 0; j < len(tab); j++ {
		word = word[:tab[j]] + lettre + word[tab[j]+1:]
	}
	return word
}
