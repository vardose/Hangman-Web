package Hangman

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

func Picker(ToFind string) string { // fonction qui sert a choisir un mot au hasard dans le fichier text passé en parametres
	tab := []string{}
	rand.Seed(time.Now().UnixNano())
	f, err := os.Open("templates/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		tab = append(tab, scanner.Text())
	}
	r := rand.Intn(len(tab))
	for i := 0; i <= r; i++ {
		if i == r {
			ToFind = tab[i]
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ToFind
}
