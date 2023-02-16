package Hangman

import (
	"strings"
)

func Jeu(Data HangManData, r string) HangManData {
	for Data.Attempts > 0 {
		lettre := r
		if strings.Contains(Data.Stock, lettre) {
			break
		} else if len(lettre) != 1 && lettre != Data.ToFind {
			Data.Attempts -= 2
			break
		} else if len(lettre) == len(Data.ToFind) && strings.Compare(lettre, Data.ToFind) == 0 { //found the word
			Data.Word = Data.ToFind
			break
		} else if strings.Contains(Data.ToFind, lettre) {
			Data.Word = Reveal(Data.Word, lettre, Data.ToFind)
			Data.Stock += lettre + " "
			break
		} else {
			Data.Attempts--
			Data.Stock += lettre + " "
			break
		}
	}
	return Data
}
