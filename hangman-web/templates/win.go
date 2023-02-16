package Hangman

import "strings"

func Win(Data HangManData) bool {
	return !strings.Contains(Data.Word, "_")
}
