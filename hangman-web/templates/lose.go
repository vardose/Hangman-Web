package Hangman

func Lose(Data HangManData) bool {
	return Data.Attempts <= 0
}
