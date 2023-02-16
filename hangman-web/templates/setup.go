package Hangman

import (
	"math/rand"
	"time"
)

type HangManData struct {
	Name     string //The Name of The Player
	Word     string // Word composed of '_', ex: H_ll_
	ToFind   string // Final word chosen by the program at the beginning. It is the word to find
	Attempts int    // Number of attempts left
	Stock    string //stock the used letters
}

func Setup() HangManData {
	var Data HangManData
	Data.Attempts = 10
	Data.ToFind = Picker(Data.ToFind)
	var index []int
	for e := 0; e < (len(Data.ToFind)/2)-1; e++ {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(len(Data.ToFind) - 1)
		if len(index) != 0 {
			for i := 0; i < len(index); i++ {
				if r == index[i] {
					r = rand.Intn(len(Data.ToFind))
				}
			}
			index = append(index, r)
		}
		index = append(index, r)
	}
	var ts []rune

	for e := 0; e < len(Data.ToFind); e++ {
		ts = append(ts, '_')
	}
	for i := 0; i < len(index); i++ {
		ts[index[i]] = []rune(Data.ToFind)[index[i]]
	}
	Data.Word = string(ts)
	return Data
}
